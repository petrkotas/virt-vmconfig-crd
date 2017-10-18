#!/bin/bash

KUBEVIRT_PATH=$GOPATH/src/kubevirt.io/kubevirt/ $GOPATH/src/kubevirt.io/kubevirt/cluster/kubectl.sh "$@"
