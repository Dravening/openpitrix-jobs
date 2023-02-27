#!/bin/bash

set -e

GV="application:v1alpha1 cluster:v1alpha1"

rm -rf ./pkg/client
./hack/generate_group.sh "client,lister,informer" d3os.io/openpitrix-jobs/pkg/client d3os.io/openpitrix-jobs/pkg/apis "$GV" --output-base=./  -h "$PWD/hack/boilerplate.go.txt"
mv d3os.io/openpitrix-jobs/pkg/client ./pkg/
rm -rf ./d3os.io
