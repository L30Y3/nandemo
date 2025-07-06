#!/bin/bash

# Exit on first error
set -e

# Check for input
if [ -z "$1" ]; then
  echo "Usage: ./generate_proto.sh <proto_file_name.proto>"
  exit 1
fi

PROTO_NAME="$1"
PROTO_DIR="shared/proto"
PROTO_PATH="$PROTO_DIR/protoevents/$PROTO_NAME"

# Run protoc
protoc -I="$PROTO_DIR" --go_out="$PROTO_DIR" --go_opt=paths=source_relative "$PROTO_PATH"

echo "✅ Generated Go code for $PROTO_NAME"