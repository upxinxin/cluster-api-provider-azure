# Deploy cluster on Public MEC

- **Feature status:** Experimental
- **Feature gate:** EdgeZone=true

## Overview

Cluster API Provider Azure (CAPZ) experimentally supports deploying clusters on [Azure Public MEC](https://azure.microsoft.com/en-us/solutions/public-multi-access-edge-compute-mec). Before you begin, you may need to prepare an azure subscription which has access to Public MEC.

To deploy a cluster on Public MEC, provide extendedlocation info through environment variable, and use the [Edgezone flavor](https://raw.githubusercontent.com/kubernetes-sigs/cluster-api-provider-azure/main/templates/cluster-template-edgezone.yaml).

## Example: Deploy cluster on Public MEC by `clusterctl`

A clusterctl flavor exists to deploy cluster on Public MEC. This flavor requires the following environment variables to be set before executing clusterctl.

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
export the following variables in you current shell.
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

Public MEC enabled clusters also require the following feature flags set as environment variables:

```bash
export EXP_EDGEZONE=true
```

Create a local kind cluster to run the managemenet cluster components:

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



