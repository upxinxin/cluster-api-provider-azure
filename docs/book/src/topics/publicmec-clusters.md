# Deploy cluster on Public MEC

- **Feature status:** Experimental
- **Feature gate:** EdgeZone=true

## Overview

Cluster API Provider Azure (CAPZ) has experimental support for deploying clusters on [Azure Public MEC](https://azure.microsoft.com/en-us/solutions/public-multi-access-edge-compute-mec). Before you begin, you need an Azure subscription which has access to Public MEC.

To deploy a cluster on Public MEC, provide extended location info through environment variables and use the "edgezone" flavor.

## Example: Deploy cluster on Public MEC by `clusterctl`

The clusterctl "edgezone" flavor exists to deploy clusters on Public MEC. This flavor requires the following environment variables to be set before executing `clusterctl`.

```bash
# Kubernetes values
export CLUSTER_NAME="my-cluster"
export WORKER_MACHINE_COUNT=2
export CONTROL_PLANE_MACHINE_COUNT=1
export KUBERNETES_VERSION="v1.25.0"

# Azure values
export AZURE_LOCATION="eastus2euap"
export AZURE_EXTENDEDLOCATION_TYPE="EdgeZone"
export AZURE_EXTENDEDLOCATION_NAME="microsoftrrdclab3"
export AZURE_RESOURCE_GROUP="${CLUSTER_NAME}"
```

Create a new service principal and save to local file:
```bash
az ad sp create-for-rbac --role Contributor --scopes="/subscriptions/${AZURE_SUBSCRIPTION_ID}" --sdk-auth > sp.json
```
Export the following variables to you current shell.
```bash
export AZURE_SUBSCRIPTION_ID="$(cat sp.json | jq -r .subscriptionId | tr -d '\n')"
export AZURE_CLIENT_SECRET="$(cat sp.json | jq -r .clientSecret | tr -d '\n')"
export AZURE_CLIENT_ID="$(cat sp.json | jq -r .clientId | tr -d '\n')"
export AZURE_CONTROL_PLANE_MACHINE_TYPE="Standard_D2s_v3"
export AZURE_NODE_MACHINE_TYPE="Standard_D2s_v3"
export AZURE_CLUSTER_IDENTITY_SECRET_NAME="cluster-identity-secret"
export AZURE_CLUSTER_IDENTITY_SECRET_NAMESPACE="default"
export CLUSTER_IDENTITY_NAME="cluster-identity"
```

Public MEC-enabled clusters also require the following feature flags set as environment variables:

```bash
export EXP_EDGEZONE=true
```

Create a local kind cluster to run the management cluster components:

```bash
kind create cluster
```

Create an identity secret on the management cluster:

```bash
kubectl create secret generic "${AZURE_CLUSTER_IDENTITY_SECRET_NAME}" --from-literal=clientSecret="${AZURE_CLIENT_SECRET}"
```

Execute clusterctl to template the resources, then apply to your kind management cluster.

```bash
clusterctl init --infrastructure azure
clusterctl generate cluster ${CLUSTER_NAME} --kubernetes-version ${KUBERNETES_VERSION} --flavor edgezone > edgezone-cluster.yaml

# assumes an existing management cluster
kubectl apply -f edgezone-cluster.yaml
```
## Known issues
After the deployment above, the cluster will look like this:
```bash
NAME                                                                      READY  SEVERITY  REASON                       SINCE  MESSAGE
Cluster/private-external-8620                                             False  Warning   ScalingUp                    18m    Scaling up control plane to 3 replicas (actual 1)
├─ClusterInfrastructure - AzureCluster/private-external-8620              True                                          18m
├─ControlPlane - KubeadmControlPlane/private-external-8620-control-plane  False  Warning   ScalingUp                    18m    Scaling up control plane to 3 replicas (actual 1)
│ └─Machine/private-external-8620-control-plane-9rp2g                     True                                          15m
└─Workers
  └─MachineDeployment/private-external-8620-md-0                          False  Warning   WaitingForAvailableMachines  21m    Minimum availability requires 2 replicas, current 0 available
    └─2 Machines...                                                       True                                          13m    See private-external-8620-md-0-7cbcd647f-9vqs8, private-external-8620-md-0-7cbcd647f-hjg8n
```
To fix the "False" READY status, [Azure cloud provider components](https://github.com/kubernetes-sigs/cloud-provider-azure/tree/master/helm/cloud-provider-azure) need to be installed by Helm.
First, get the kubeconfig of the cluster:
```bash
kubectl get secrets ${CLUSTER_NAME}-kubeconfig -o json | jq -r .data.value | base64 --decode > ./kubeconfig
```
Then
```bash
helm install --repo https://raw.githubusercontent.com/kubernetes-sigs/cloud-provider-azure/master/helm/repo cloud-provider-azure --generate-name --set infra.clusterName=${CLUSTER_NAME} --kubeconfig=./kubeconfig
```

After a while, the cluster should look like this:
```bash
NAME                                                                      READY  SEVERITY  REASON  SINCE  MESSAGE
Cluster/private-external-8620                                             True                     6m38s
├─ClusterInfrastructure - AzureCluster/private-external-8620              True                     46m
├─ControlPlane - KubeadmControlPlane/private-external-8620-control-plane  True                     6m38s
│ └─3 Machines...                                                         True                     7m47s  See private-external-8620-control-plane-6lb57, private-external-8620-control-plane-79mls, ...
└─Workers
  └─MachineDeployment/private-external-8620-md-0                          True                     10m
    └─2 Machines...                                                       True                     41m    See private-external-8620-md-0-7cbcd647f-9vqs8, private-external-8620-md-0-7cbcd647f-hjg8n
```




