KUBEVIRT="$GOPATH/src/kubevirt.io/kubevirt/"
KUBECTL="$KUBEVIRT/cluster/kubectl.sh"

# We need to grab absolute path to all manifests since we'll have to pushd
# later on.
MANIFESTS=$(find $(pwd)/manifests -name *.yaml)

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
for i in $MANIFESTS; do
    $KUBECTL create -f $i
done
popd
