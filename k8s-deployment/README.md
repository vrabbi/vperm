# Build an image (Optional)
Using pack CLI run the following from the root of the directory:
```bash
pack build vperm --path . --builder paketobuildpacks/builder-jammy-tiny
```
# Push the image to your registry (Optional)
```bash
docker tag vperm:latest ghcr.io/vrabbi/vperm:1.0.0
docker push ghcr.io/vrabbi/vperm:1.0.0
```

# Use YTT to update the image reference and credentials (Required)
* First update the values.yaml file in this folder with your environment details.
* Generate the final manifest with YTT
  ```bash
  ytt -f manifest.ytt.yaml -f values.yaml > manifest.yaml
  ```
# Deploy to your cluster (Required)
```bash
kubectl apply -f manifest.yaml
kubectl wait --for=condition=complete job/vperm -n vperm --timeout 5m
```

# Retrieve the output (Required)
You can output to STDOUT
```bash
kubectl logs jobs/vperm -n vperm
```  
You can also export to a file
```bash
kubectl logs jobs/vperm -n vperm > permissions.log
```

# Cleanup (Optional)
```bash
kubectl delete ns vperm
```
