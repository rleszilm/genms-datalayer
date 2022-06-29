package datalayer

import "google.golang.org/protobuf/compiler/protogen"

// Generator is a struct that reads and sources data that should be used for file generation.
type Generator struct {
	Plugin *protogen.Plugin
}

// NewGenerator returns a new Generator.
func NewGenerator() {

}
