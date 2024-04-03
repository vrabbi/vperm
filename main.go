package main

import (
	"context"
	"fmt"
	"log"
	"maps"
	"os"
	"sync"

	cnsvsphere "sigs.k8s.io/vsphere-csi-driver/v3/pkg/common/cns-lib/vsphere"

	"github.com/fatih/color"
	"github.com/vmware/govmomi/object"
	vim25types "github.com/vmware/govmomi/vim25/types"
)


var(
	DsPriv = "Datastore.FileManagement"
	SysReadPriv = "System.Read"

    title = color.New(color.FgBlue).Add(color.Bold)
    content = color.New(color.FgCyan)

    vc *cnsvsphere.VirtualCenter 
)

func init() {
    vc = &cnsvsphere.VirtualCenter{
        Config: &cnsvsphere.VirtualCenterConfig {
            Username: os.Getenv("VSPHERE_USERNAME"),
            Password: os.Getenv("VSPHERE_PASSWORD"),
            Host: os.Getenv("VSPHERE_HOST"),
            Port: 443,
            Insecure: true,
            Scheme: "https",
        },
        ClientMutex: &sync.Mutex{},
    }
}

func main() {
    ctx := context.Background()
    title.Printf("-- Connecting to %v\n", vc.Config.Host)

    cli, err := vc.NewClient(ctx) // New VC Client based on previous input.
    fatalError(err)
    datastores, err := GetDatastores(ctx, vc) // Fetch all datastores from the cluster.
    fatalError(err)

	// Get authorization manager object from VC
    mgr := object.NewAuthorizationManager(cli.Client)

    title.Printf("\n-- Fetching Roles attached to user %v\n", vc.Config.Username)
    listRoles(ctx, mgr)
    title.Printf("\n-- Fetching permissions for %v\n", vc.Config.Username)
    listPermissions(ctx, mgr)

    entities := []vim25types.ManagedObjectReference{}
    for _, ds := range datastores { entities = append(entities, ds.Reference()) }
    title.Printf("\n-- Fetching privileges for DS %v\n", vc.Config.Username)
    listPrivilegesOnDSs(ctx, entities, vc.Config.Username, mgr)

    fmt.Println("------------------------------")
}

func listPrivilegesOnDSs(ctx context.Context, entities []vim25types.ManagedObjectReference, username string, mgr *object.AuthorizationManager) {
    privileges, err := mgr.FetchUserPrivilegeOnEntities(ctx, entities, username)
    fatalError(err)

    for _, priv := range privileges {
        content.Printf("\n\t[%v] \n", priv.Entity)
        for i := 0; i <= len(priv.Privileges)-1; i++ {
            fmt.Printf("\t\t * %s \n", priv.Privileges[i])
        }
    }
}

// listRoles fetch roles binded to the user
func listRoles(ctx context.Context, mgr *object.AuthorizationManager) {
    roles, err := mgr.RoleList(ctx)
    fatalError(err)

    content.Add(color.Bold)

    for _, role := range roles {
        content.Printf("\n\t[%v] - %s \n", role.RoleId, role.Name)
        for i := 0; i <= len(role.Privilege)-1; i++ {
            fmt.Printf("\t\t * %s \n", role.Privilege[i])
        }
    }
}

// listPermissions fetches all permissions for the user
func listPermissions(ctx context.Context, mgr *object.AuthorizationManager) {
    permissions, err := mgr.RetrieveAllPermissions(ctx)
    fatalError(err)

    content.Add(color.Bold)

    for _, perm := range permissions {
        content.Printf("\t[%v] \t", perm.RoleId)
        content.Printf(" %s \t %s \t %v \n", perm.Entity, perm.Principal, perm.DynamicData)
    }
}

func GetDatastores(ctx context.Context, vc *cnsvsphere.VirtualCenter) (map[string]*cnsvsphere.DatastoreInfo, error) {
    dcList, err := vc.GetDatacenters(ctx)
    if err != nil { return nil, err }

    dsURLInfoMap := map[string]*cnsvsphere.DatastoreInfo{}
    for _, dc := range dcList {
        tmpURL, err := dc.GetAllDatastores(ctx)
        if err != nil {
            return nil, err
        }
        maps.Copy(dsURLInfoMap, tmpURL)
    }
    return dsURLInfoMap, nil
}

func fatalError(err error) {
    if err != nil {
        log.Fatal(err)
    }
}


