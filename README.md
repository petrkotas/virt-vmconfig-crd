Persistent Virtual Machine kubevirt add-on

The persistent virtual machine (pvm) add-on adds a functionality to kubevirt
to store a non running VMs in kubernetes.

# Installing the add-on to vagrant deployment

The command

```bash
make vagrant-deploy
```

executed on a machine kubevirt is deployed from deploys the add-on via
kubevirt's kubectl proxy.

# Usage of PVMs

## Creating new PVM

Create a new PVM definition and push it to the cluster via kubectl. An example
definition is provided in example.yaml. That definition can be pushed simply by
calling

```bash
cluster/kubectl.sh --core create -f example.yaml
```

## Listing PVMs

```bash
cluster/kubectl.sh --core get pvms
```

## Deleting PVM

```bash
cluster/kubectl.sh --core delete -f example.yaml
```

## Starting a new PVM

Due to the nature of CRD, the VM can be directly constructed from the PVM as
long as the format of example.yaml is followed.

```bash
./cluster/kubectl.sh --core get pvm testvm -o yaml | sed 's/PersistentVirtualMachine/VirtualMachine/' | ./cluster/kubectl.sh create -f -
```
