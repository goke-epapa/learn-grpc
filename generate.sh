#!/bin/bash

protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
protoc sumapi/sumapipb/sum.proto --go_out=plugins=grpc:.