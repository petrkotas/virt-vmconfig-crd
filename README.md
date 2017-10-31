# Persistent Virtual Machine Kubernetes add-on

This add-on provides the functionality to store VM configuration in the
Kubernetes cluster. It is designed to provide common grounds for building 
high level Virtual Machine management tools on top of the 
[Kubevirt](http://kubevirt.io).

That means, this add-on is designed to work in tanden with the 
[Kubevirt](http://kubevirt.io). 

The functionality present in this add-on is implemented as a [Kubernetes Custom
resource definition (CRD)](https://kubernetes.io/docs/concepts/api-extension/custom-resources/).
Because of thisa approach, api heavylifting is done by Kubernetes and does not
have to be re-implemented.

Moreover, this this add-on is also designed to provide golang language binding
for building custom resource client. These bindings are autogenereted, therefore
developers are not supposed to modify them.

## Get started
Before you can do anything with this add-on, you have to get a local copy. 
To use the codes further as a library, use the golang tool.
```bash
go get github.com/petrkotas/virt-vmconfig-crd.git
```
With this method you get functioning codes stored in your local `GOPATH`.

If this is not desired, just clone the repository whereever you desire.


## Install
Instalation is simple. Just run 
```bash
kubectl create -f manifests/persistentvm-resource.yaml
```
give it few seconds and you are ready to go.

To verify everything went OK get existing CRD from the Kubernetes cluster.
```bash
kubectl get crd
```
If `persistentvirtualmachines.kubevirt.io` is present, you are
good to go.

### Legacy
Besides the direct kubernetes installation. In case you are using the original
Kubevirt vagrant deployment present in its github repository. Use following
command.

```bash
make vagrant-deploy
```
This deploys the Kubevirt to the virtual machine running the Kubernetes
with the Kubevirt.


## Usage of PVMs
Once the new PVM CRD has been added, you can start using it.
You can:

## Create new PVM
Before you can register any new PVM in Kubernetes cluster, you have to define it
first. New vm definition is done the same way as any other Kubernetes object,
with the YAML file.
```yaml
apiVersion: kubevirt.io/v1alpha1
kind: PersistentVirtualMachine
metadata:
  name: testvm
spec:
  domain:
    devices:
      consoles:
      - type: neuralink
    memory:
      unit: MB
      value: 64
    os:
      type:
        os: AwesomeBrainOS
    type: qemu
```
Once you have the spicification ready, registering to Kubernetes cluster is as
simple as
```bash
kubectl create -f my-pvm-example.yaml
```

### Legacy
If you are using the original Kubevirt vagrant environment, you can use this 
command
```bash
cluster/kubectl.sh --core create -f example.yaml
```

## List PVMs
To get the list of existing PVMs in your Kubernetes cluster use the command
```bash
kubectl get pvms
```

### Legacy
For the original Kubevirt vagrant environment use
```bash
cluster/kubectl.sh --core get pvms
```

## Delete PVM
To delete existing PVM from you Kubernetes cluster use the command
```bash
kubectl delete pvm <your-PVM-name>
```
or if you have a YAML specification of the PVM, use
```bash
kubectl delete -f your-PVM-name.yaml
```

### Legacy
For the original Kubevirt vagrant environment use
```bash
cluster/kubectl.sh --core delete -f example.yaml
```

## Use as a library
To use the generated bindings in your own code, just import the package in the
go file and start using it.
```go
import pvmcrd github.com/petrkotas/virt-vmconfig-crd/...

//TODO: Add proper example
```
