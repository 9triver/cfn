#!/bin/bash

PROTOC_VERSION=26.1
PROTOC_URL="https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip"
PROTOC="./protoc/bin/protoc"
export PATH="$PATH:$HOME/go/bin"

if [ ! -f $PROTOC ]; then
  echo "Downloading protoc"
  wget -O protoc.zip $PROTOC_URL
  unzip protoc.zip -d protoc
  rm protoc.zip
fi

ACTOR_SRC=$(go list -f {{.Dir}} github.com/asynkron/protoactor-go/actor)
ACTOR_PROTO=$ACTOR_SRC/actor.proto

PROTOC="$PROTOC -I $ACTOR_SRC -I ./messages"
PROTO_SRC="./messages/*.proto ./messages/executor/*.proto ./messages/dag/*.proto ./messages/controller/*.proto"
if [ ! -d ./ts ]; then
  echo "Creating output directory for TypeScript: ./ts"
  mkdir -p ./ts
fi

if [ ! -d ./python ]; then
  echo "Creating output directory for Python: ./python"
  mkdir -p ./python
fi

echo "Generating protobuf files for Go"
$PROTOC --go_out=./go --go_opt=paths=source_relative --go-grpc_out=./go --go-grpc_opt=paths=source_relative $PROTO_SRC

echo "Generating protobuf files for TypeScript"
$PROTOC --ts_proto_out=./ts --plugin=./node_modules/.bin/protoc-gen-ts_proto --ts_proto_opt=esModuleInterop=true $PROTO_SRC $ACTOR_PROTO

PY_OUTPUTS="./python ../clients/py/actorc/protos"

echo "Generating protobuf files for Python"
for PY_OUTPUT in $PY_OUTPUTS; do
  if [ ! -d $PY_OUTPUT ]; then
    echo "Creating output directory for Python: $PY_OUTPUT"
    mkdir -p $PY_OUTPUT
  fi
  $PROTOC --python_out=pyi_out:$PY_OUTPUT $PROTO_SRC $ACTOR_PROTO
done
