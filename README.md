# scaler-operator


## Description
Can be used to scale deployments

## Getting Started
You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Running this locally

-   ```
    make manifests
    ```

-   ```
    kubectl apply -f config/crd/bases/api.rakshitgondwal.online_scalers.yaml
    ```

- Run a `nginx` deployment on your cluster.

    ```
    kubectl create deployment nginx --image=nginx
    ```
**Run the following command to check for the number of replicas running right now**(Should be 1)
```
kubectl get pods
```

**Run the following commands to scale up the replicas**
-   ```
    make run
    ```
- To be done in other terminal
    ```
    kubectl apply -f config/samples/api_v1alpha1_scaler.yaml
    ```

**You can now check the number of replicas using the following command**(Should be 5)
```
kubectl get pods
```

## License

Copyright 2023 Rakshit Gondwal.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

