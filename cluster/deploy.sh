#!/bin/bash

KUBECTL="cluster/kubectl.sh"
MANIFESTS=$(ls manifests/*.yaml);

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
