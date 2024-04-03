## vSphere Permissions

### Compiling 

Use golang >=1.22.

```shell
go build -o vperm main.go
```

### Running 

Set the required variables for connecting on vSphere, since the output is fairly long in some cases save on a file.

```shell
VSPHERE_HOST="192.168.0.1" \
VSPHERE_PASSWORD="randompassword" \
VSPHERE_USERNAME="administrator@vsphere.local" \
go run main.go > permissions.log
```

### Sections explanations

* Fetching Roles attached to user 

The format is [role id] - Name and a list of permissions

```
	[-1] - Admin 
		 * Alarm.Acknowledge 
		 * Alarm.Create 
		 * Alarm.Delete 
		 * Alarm.DisableActions 
		 * Alarm.Edit 
		 * Alarm.SetStatus 
		 * Alarm.ToggleEnableOnEntity 
```


* Fetching permissions for user

Format is [perm.id] entity and principal

```
	[-1] 	 Folder:group-d1 	 VSPHERE.LOCAL\vpxd-extension-5888a232-cd28-458e-a604-bd8bb041e13c 	 {} 
	[1003] 	 Folder:group-d1 	 VSPHERE.LOCAL\vsphere-webclient-5888a232-cd28-458e-a604-bd8bb041e13c 	 {} 
	[-1] 	 Folder:group-d1 	 VSPHERE.LOCAL\vpxd-5888a232-cd28-458e-a604-bd8bb041e13c 	 {} 
	[-1] 	 Folder:group-d1 	 VSPHERE.LOCAL\Administrator 	 {} 
	[-1] 	 Folder:group-d1 	 VSPHERE.LOCAL\Administrator 	 {} 
	[-7] 	 Folder:group-d1 	 VSPHERE.LOCAL\TrustedAdmins 	 {} 
	[11] 	 Folder:group-d1 	 VSPHERE.LOCAL\AutoUpdate 	 {} 
	[-1] 	 Folder:group-d1 	 VSPHERE.LOCAL\Administrators 	 {} 
	[1003] 	 Folder:group-d1 	 VSPHERE.LOCAL\vSphereClientSolutionUsers 	 {} 
	[1002] 	 Folder:group-d1 	 VSPHERE.LOCAL\SyncUsers 	 {} 
	[963545493] 	 Folder:group-d1 	 VSPHERE.LOCAL\vStatsGroup 	 {} 
```

* Fetching privileges for DataStores in teh system

Format is [datastore id] and a list of privileges for the ds

```
	[Datastore:datastore-40] 
		 * System.Anonymous 
		 * System.View 
		 * System.Read 
		 * TrustedAdmin.ConfigureTokenConversionPolicy 
		 * TrustedAdmin.ManageKMSTrust 
		 * TrustedAdmin.ReadKMSTrust 
		 * TrustedAdmin.ReadStsInfo 
		 * TrustedAdmin.ConfigureHostCertificates 
		 * TrustedAdmin.ConfigureHostMetadata 
		 * TrustedAdmin.ManageAttestingSSO 
		 * TrustedAdmin.ReadAttestingSSO 
		 * TrustedAdmin.RetrieveHostMetadata 
		 * TrustedAdmin.RetrieveTPMHostCertificates 
		 * TrustedAdmin.ManageTrustedHosts 
```
