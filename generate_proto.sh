#!/bin/bash

protoc -I=shared/proto --go_out=shared/proto --go_opt=paths=source_relative \
shared/proto/protoevents/order_created.proto