# Custom Virtual Networks

## Pre-existing vnet and subnets

To deploy a cluster using a pre-existing vnet, modify the `AzureCluster` spec to include the name and resource group of the existing vnet as follows, as well as the control plane and node subnets as follows:

```yaml
apiVersion: infrastructure.cluster.x-k8s.io/v1beta1
kind: AzureCluster
metadata:
  name: cluster-byo-vnet
  namespace: default
spec:
  location: southcentralus
  networkSpec:
    vnet:
      resourceGroup: custom-vnet
      name: my-vnet
    subnets:
      - name: my-control-plane-subnet
        role: control-plane
        securityGroup:
          name: my-control-plane-nsg
      - name: my-node-subnet
        role: node
        routeTable:
          name: my-node-routetable
        securityGroup:
          name: my-node-nsg
  resourceGroup: cluster-byo-vnet
  ```

When providing a vnet, it is required to also provide the two subnets that should be used for control planes and nodes.

If providing an existing vnet and subnets with existing network security groups, make sure that the control plane security group allows inbound to port 6443, as port 6443 is used by kubeadm to bootstrap the control planes. Alternatively, you can [provide a custom control plane endpoint](https://github.com/kubernetes-sigs/cluster-api-bootstrap-provider-kubeadm#kubeadmconfig-objects) in the `KubeadmConfig` spec.

The pre-existing vnet can be in the same resource group or a different resource group in the same subscription as the target cluster. When deleting the `AzureCluster`, the vnet and resource group will only be deleted if they are "managed" by capz, ie. they were created during cluster deployment. Pre-existing vnets and resource groups will *not* be deleted.

## Virtual Network Peering

Alternatively, pre-existing vnets can be peered with a cluster's newly created vnets by specifying each vnet by name and resource group.

```yaml
apiVersion: infrastructure.cluster.x-k8s.io/v1beta1
kind: AzureCluster
metadata:
  name: cluster-vnet-peering
  namespace: default
spec:
  location: southcentralus
  networkSpec:
    vnet:
      name: my-vnet
      cidrBlocks: 
        - 10.255.0.0/16
      peerings:
      - resourceGroup: vnet-peering-rg
        remoteVnetName: existing-vnet-1
      - resourceGroup: vnet-peering-rg
        remoteVnetName: existing-vnet-2
    subnets:
      - name: my-subnet-cp
        role: control-plane
        cidrBlocks: 
          - 10.255.0.0/24
      - name: my-subnet-node
        role: node
        cidrBlocks: 
          - 10.255.1.0/24
  resourceGroup: cluster-vnet-peering
  ```

Currently, only virtual networks on the same subscription can be peered. Also, note that when creating workload clusters with internal load balancers, the management cluster must be in the same VNet or a peered VNet. See [here](https://capz.sigs.k8s.io/topics/api-server-endpoint.html#warning) for more details.

## Custom Network Spec

It is also possible to customize the vnet to be created without providing an already existing vnet. To do so, simply modify the `AzureCluster` `NetworkSpec` as desired. Here is an illustrative example of a cluster with a customized vnet address space (CIDR) and customized subnets:

```yaml
apiVersion: infrastructure.cluster.x-k8s.io/v1beta1
kind: AzureCluster
metadata:
  name: cluster-example
  namespace: default
spec:
  location: southcentralus
  networkSpec:
    vnet:
      name: my-vnet
      cidrBlocks: 
        - 10.0.0.0/16
    subnets:
      - name: my-subnet-cp
        role: control-plane
        cidrBlocks: 
          - 10.0.1.0/24
      - name: my-subnet-node
        role: node
        cidrBlocks: 
          - 10.0.2.0/24
  resourceGroup: cluster-example
```

If no CIDR block is provided, `10.0.0.0/8` will be used by default, with default internal LB private IP `10.0.0.100`.

### Custom Security Rules

<aside class="note">

<h1> Note </h1>

Security Rules were previously known as `ingressRule` in v1alpha3.

</aside>

Security rules can also be customized as part of the subnet specification in a custom network spec.
Note that ingress rules for the Kubernetes API Server port (default 6443) and SSH (22) are automatically added to the controlplane subnet only if security rules aren't specified.
It is the responsibility of the user to supply those rules themselves if using custom rules.

Here is an illustrative example of customizing rules that builds on the one above by adding an egress rule to the control plane nodes:

```yaml
apiVersion: infrastructure.cluster.x-k8s.io/v1beta1
kind: AzureCluster
metadata:
  name: cluster-example
  namespace: default
spec:
  location: southcentralus
  networkSpec:
    vnet:
      name: my-vnet
      cidrBlocks: 
        - 10.0.0.0/16
    subnets:
      - name: my-subnet-cp
        role: control-plane
        cidrBlocks: 
          - 10.0.1.0/24
        securityGroup:
          name: my-subnet-cp-nsg
          securityRules:
            - name: "allow_ssh"
              description: "allow SSH"
              direction: "Inbound"
              priority: 2200
              protocol: "*"
              destination: "*"
              destinationPorts: "22"
              source: "*"
              sourcePorts: "*"
            - name: "allow_apiserver"
              description: "Allow K8s API Server"
              direction: "Inbound"
              priority: 2201
              protocol: "*"
              destination: "*"
              destinationPorts: "6443"
              source: "*"
              sourcePorts: "*"
            - name: "allow_port_50000"
              description: "allow port 50000"
              direction: "Outbound"
              priority: 2202
              protocol: "Tcp"
              destination: "*"
              destinationPorts: "50000"
              source: "*"
              sourcePorts: "*"
      - name: my-subnet-node
        role: node
        cidrBlocks: 
          - 10.0.2.0/24
  resourceGroup: cluster-example
```

### Virtual Network service endpoints

Sometimes it's desirable to use [Virtual Network service endpoints](https://docs.microsoft.com/en-us/azure/virtual-network/virtual-network-service-endpoints-overview) to establish secure and direct connectivity to Azure services from your subnet(s). Service Endpoints are configured on a per-subnet basis. Vnets managed by either `AzureCluster` or `AzureManagedControlPlane` can have `serviceEndpoints` optionally set on each subnet.

```yaml
apiVersion: infrastructure.cluster.x-k8s.io/v1beta1
kind: AzureCluster
metadata:
  name: cluster-example
  namespace: default
spec:
  location: southcentralus
  networkSpec:
    vnet:
      name: my-vnet
      cidrBlocks:
        - 10.0.0.0/16
    subnets:
      - name: my-subnet-cp
        role: control-plane
        cidrBlocks:
          - 10.0.1.0/24
        serviceEndpoints:
          - service: Microsoft.AzureActiveDirectory
            locations: ["*"]
      - name: my-subnet-node
        role: node
        cidrBlocks:
          - 10.0.2.0/24
        serviceEndpoints:
          - service: Microsoft.AzureActiveDirectory
            locations: ["*"]
          - service: Microsoft.Storage
            locations: ["southcentralus"]
  resourceGroup: cluster-example
```

### Custom subnets

Sometimes it's desirable to use different subnets for different node pools.
Several subnets can be specified in the `networkSpec` to be later referenced by name from other CR's like `AzureMachine` or `AzureMachinePool`.
When more than one `node` subnet is specified, the `subnetName` field in those other CR's becomes mandatory because the controllers wouldn't know which subnet to use.

The subnet used for the control plane must use the role `control-plane` while the subnets for the worker nodes must use the role `node`.


```yaml
---
apiVersion: infrastructure.cluster.x-k8s.io/v1beta1
kind: AzureCluster
metadata:
  name: cluster-example
  namespace: default
spec:
  location: southcentralus
  networkSpec:
    subnets:
    - name: control-plane-subnet
      role: control-plane
    - name: subnet-mp-1
      role: node
    - name: subnet-mp-2
      role: node
    vnet:
      name: my-vnet
      cidrBlocks:
        - 10.0.0.0/16
  resourceGroup: cluster-example
---
apiVersion: infrastructure.cluster.x-k8s.io/v1beta1
kind: AzureMachinePool
metadata:
  name: mp1
  namespace: default
spec:
  location: southcentralus
  strategy:
    rollingUpdate:
      deletePolicy: Oldest
      maxSurge: 25%
      maxUnavailable: 1
    type: RollingUpdate
  template:
    osDisk:
      diskSizeGB: 30
      managedDisk:
        storageAccountType: Premium_LRS
      osType: Linux
    sshPublicKey: ${YOUR_SSH_PUB_KEY}
    subnetName: subnet-mp-1
    vmSize: Standard_D2s_v3
---
apiVersion: infrastructure.cluster.x-k8s.io/v1beta1
kind: AzureMachinePool
metadata:
  name: mp2
  namespace: default
spec:
  location: southcentralus
  strategy:
    rollingUpdate:
      deletePolicy: Oldest
      maxSurge: 25%
      maxUnavailable: 1
    type: RollingUpdate
  template:
    osDisk:
      diskSizeGB: 30
      managedDisk:
        storageAccountType: Premium_LRS
      osType: Linux
    sshPublicKey: ${YOUR_SSH_PUB_KEY}
    subnetName: subnet-mp-2
    vmSize: Standard_D2s_v3
```

If you don't specify any `node` subnets, one subnet with role `node` will be created and added to the `networkSpec` definition.
