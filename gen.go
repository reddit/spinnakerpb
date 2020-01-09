//TODO(eac): listing the filenames here to avoid globbing is unwieldy. probably should switch to make at some point.
//go:generate protoc --gogofast_out=. artifact.proto notification.proto pipeline.proto stage.proto trigger.proto
package spinnakerpb
