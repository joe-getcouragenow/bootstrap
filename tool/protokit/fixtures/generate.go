package main

//go:generate protoc --descriptor_set_out=fileset.pb --include_imports --include_source_info -I. ./booking.proto ./todo.proto ./extend.proto
