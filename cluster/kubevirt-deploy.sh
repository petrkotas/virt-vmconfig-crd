KUBEVIRT="$GOPATH/src/kubevirt.io/kubevirt/"
KUBECTL="$KUBEVIRT/cluster/kubectl.sh"

pushd $KUBEVIRT
# This is a pattern taken from kubevirt:
# first, we delete the manifests that were previously deployed,
for i in $MANIFESTS; do
    $KUBECTL delete -f $i --grace-period 0 2>/dev/null || :
done

# wait a bit,
sleep 2

# and then deploy the new manifests.
echo "Deploying ..."
$KUBECTL create -f manifests/persisentvm-resource.yaml
popd
