Persistent Virtual Machine KubeVirt add-on

The persistent virtual machine (pvm) add-on adds a functionality to KubeVirt
to store a non running VMs in kubernetes.

# Overview

The persistence mechanism is based on the idea to use a separate resource to
hold persisted VMs.
The API is the same as for `VirtualMachines` just the `kind` is different.

# Installation

In order to enable persistent VMs, you need to define the new resource type in
the cluster:

```bash
$ kubectl create -f manifests/persistentvm-resource.yaml
customresourcedefinition "persistentvirtualmachines.kubevirt.io" created
```

## To the KubeVirt vagrant deployment

The developer setup needs a slightly different flow:

```bash
make vagrant-deploy
```

executed on a machine KubeVirt is deployed from deploys the add-on via
KubeVirt's kubectl proxy.

# Usage of PVMs

## Creating new PVM

Create a new PVM definition and push it to the cluster via kubectl. An example
definition is provided in `example.yaml`. That definition can be pushed simply
by calling

> **Note:** In order to use PVMs with the KubeVirt developer setup, you need to
> replace `kubectl` with `cluster/kubectl.sh --core` int he following examples.

```bash
$ kubectl create -f example.yaml
```

## Listing PVMs

```bash
$ kubectl get pvms
```

## Deleting PVM

```bash
$ kubectl delete -f example.yaml
# OR
$ kubectl delete pvm testvm
```

## Starting a new PVM

Due to the nature of CRD, the VM can be directly constructed from the PVM as
long as the format of `example.yaml` is followed.

```bash
$ kubectl get pvm testvm -o json \
  | jq ".kind = 'VirtualMachine" \
  | kubectl create -f -
```

This will fetch the PVM `testvm` and just change the `kind` (using `jq`) to
`VirtualMachine` in order to run it.
