package client

import "google.golang.org/protobuf/types/known/structpb"

type TaskInputs structpb.Struct

type TaskOutputs structpb.Struct

type TaskFunction interface {
	Apply(*TaskInputs) (*TaskOutputs, error)
}
