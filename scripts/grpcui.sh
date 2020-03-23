#!/usr/bin/env bash
if [ -n "$1" ]; then
  grpcui -import-path ../proto -proto "../proto/$1" -port 5000 -plaintext localhost:50051
else
  echo "You need to provide the relative path from ./proto to the proto file"
fi
