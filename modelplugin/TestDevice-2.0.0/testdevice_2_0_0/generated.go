// Copyright 2019-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Package testdevice_2_0_0 is a generated package which contains definitions
of structs which represent a YANG schema. The generated schema can be
compressed by a series of transformations (compression was false
in this case).

This package was generated by /home/scondon/go/src/github.com/openconfig/ygot/ygen/commongen.go
using the following YANG input files:
	- test1@2019-06-10.yang
Imported modules were sourced from:
	- yang/...
*/
package testdevice_2_0_0

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/ytypes"
)

// Binary is a type that is used for fields that have a YANG type of
// binary. It is used such that binary fields can be distinguished from
// leaf-lists of uint8s (which are mapped to []uint8, equivalent to
// []byte in reflection).
type Binary []byte

// YANGEmpty is a type that is used for fields that have a YANG type of
// empty. It is used such that empty fields can be distinguished from boolean fields
// in the generated code.
type YANGEmpty bool

var (
	SchemaTree map[string]*yang.Entry
)

func init() {
	var err error
	if SchemaTree, err = UnzipSchema(); err != nil {
		panic("schema error: " +  err.Error())
	}
}

// Schema returns the details of the generated schema.
func Schema() (*ytypes.Schema, error) {
	uzp, err := UnzipSchema()
	if err != nil {
		return nil, fmt.Errorf("cannot unzip schema, %v", err)
	}

	return &ytypes.Schema{
		Root: &Device{},
		SchemaTree: uzp,
		Unmarshal: Unmarshal,
	}, nil
}

// UnzipSchema unzips the zipped schema and returns a map of yang.Entry nodes,
// keyed by the name of the struct that the yang.Entry describes the schema for.
func UnzipSchema() (map[string]*yang.Entry, error) {
	var schemaTree map[string]*yang.Entry
	var err error
	if schemaTree, err = ygot.GzipToSchema(ySchema); err != nil {
		return nil, fmt.Errorf("could not unzip the schema; %v", err)
	}
	return schemaTree, nil
}

// Unmarshal unmarshals data, which must be RFC7951 JSON format, into
// destStruct, which must be non-nil and the correct GoStruct type. It returns
// an error if the destStruct is not found in the schema or the data cannot be
// unmarshaled. The supplied options (opts) are used to control the behaviour
// of the unmarshal function - for example, determining whether errors are
// thrown for unknown fields in the input JSON.
func Unmarshal(data []byte, destStruct ygot.GoStruct, opts ...ytypes.UnmarshalOpt) error {
	tn := reflect.TypeOf(destStruct).Elem().Name()
	schema, ok := SchemaTree[tn]
	if !ok {
		return fmt.Errorf("could not find schema for type %s", tn )
	}
	var jsonTree interface{}
	if err := json.Unmarshal([]byte(data), &jsonTree); err != nil {
		return err
	}
	return ytypes.Unmarshal(schema, destStruct, jsonTree, opts...)
}

// Device represents the /device YANG schema element.
type Device struct {
	Cont1A	*Test1_Cont1A	`path:"cont1a" module:"test1"`
	Cont1BState	*Test1_Cont1BState	`path:"cont1b-state" module:"test1"`
	LeafAtTopLevel	*string	`path:"leafAtTopLevel" module:"test1"`
}

// IsYANGGoStruct ensures that Device implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Device) IsYANGGoStruct() {}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Device) Validate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Device"], t, opts...); err != nil {
		return err
	}
	return nil
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Device) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }


// Test1_Cont1A represents the /test1/cont1a YANG schema element.
type Test1_Cont1A struct {
	Cont2A	*Test1_Cont1A_Cont2A	`path:"cont2a" module:"test1"`
	Leaf1A	*string	`path:"leaf1a" module:"test1"`
	List2A	map[string]*Test1_Cont1A_List2A	`path:"list2a" module:"test1"`
}

// IsYANGGoStruct ensures that Test1_Cont1A implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Test1_Cont1A) IsYANGGoStruct() {}

// NewList2A creates a new entry in the List2A list of the
// Test1_Cont1A struct. The keys of the list are populated from the input
// arguments.
func (t *Test1_Cont1A) NewList2A(Name string) (*Test1_Cont1A_List2A, error){

	// Initialise the list within the receiver struct if it has not already been
	// created.
	if t.List2A == nil {
		t.List2A = make(map[string]*Test1_Cont1A_List2A)
	}

	key := Name

	// Ensure that this key has not already been used in the
	// list. Keyed YANG lists do not allow duplicate keys to
	// be created.
	if _, ok := t.List2A[key]; ok {
		return nil, fmt.Errorf("duplicate key %v for list List2A", key)
	}

	t.List2A[key] = &Test1_Cont1A_List2A{
		Name: &Name,
	}

	return t.List2A[key], nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Test1_Cont1A) Validate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Test1_Cont1A"], t, opts...); err != nil {
		return err
	}
	return nil
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Test1_Cont1A) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }


// Test1_Cont1A_Cont2A represents the /test1/cont1a/cont2a YANG schema element.
type Test1_Cont1A_Cont2A struct {
	Leaf2A	*uint8	`path:"leaf2a" module:"test1"`
	Leaf2B	*float64	`path:"leaf2b" module:"test1"`
	Leaf2C	*string	`path:"leaf2c" module:"test1"`
	Leaf2D	*bool	`path:"leaf2d" module:"test1"`
}

// IsYANGGoStruct ensures that Test1_Cont1A_Cont2A implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Test1_Cont1A_Cont2A) IsYANGGoStruct() {}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Test1_Cont1A_Cont2A) Validate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Test1_Cont1A_Cont2A"], t, opts...); err != nil {
		return err
	}
	return nil
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Test1_Cont1A_Cont2A) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }


// Test1_Cont1A_List2A represents the /test1/cont1a/list2a YANG schema element.
type Test1_Cont1A_List2A struct {
	Name	*string	`path:"name" module:"test1"`
	RxPower	*uint16	`path:"rx-power" module:"test1"`
	TxPower	*uint16	`path:"tx-power" module:"test1"`
}

// IsYANGGoStruct ensures that Test1_Cont1A_List2A implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Test1_Cont1A_List2A) IsYANGGoStruct() {}

// ΛListKeyMap returns the keys of the Test1_Cont1A_List2A struct, which is a YANG list entry.
func (t *Test1_Cont1A_List2A) ΛListKeyMap() (map[string]interface{}, error) {
	if t.Name == nil {
		return nil, fmt.Errorf("nil value for key Name")
	}

	return map[string]interface{}{
		"name": *t.Name,
	}, nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Test1_Cont1A_List2A) Validate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Test1_Cont1A_List2A"], t, opts...); err != nil {
		return err
	}
	return nil
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Test1_Cont1A_List2A) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }


// Test1_Cont1BState represents the /test1/cont1b-state YANG schema element.
type Test1_Cont1BState struct {
	Leaf2D	*uint16	`path:"leaf2d" module:"test1"`
	List2B	map[uint8]*Test1_Cont1BState_List2B	`path:"list2b" module:"test1"`
}

// IsYANGGoStruct ensures that Test1_Cont1BState implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Test1_Cont1BState) IsYANGGoStruct() {}

// NewList2B creates a new entry in the List2B list of the
// Test1_Cont1BState struct. The keys of the list are populated from the input
// arguments.
func (t *Test1_Cont1BState) NewList2B(Index uint8) (*Test1_Cont1BState_List2B, error){

	// Initialise the list within the receiver struct if it has not already been
	// created.
	if t.List2B == nil {
		t.List2B = make(map[uint8]*Test1_Cont1BState_List2B)
	}

	key := Index

	// Ensure that this key has not already been used in the
	// list. Keyed YANG lists do not allow duplicate keys to
	// be created.
	if _, ok := t.List2B[key]; ok {
		return nil, fmt.Errorf("duplicate key %v for list List2B", key)
	}

	t.List2B[key] = &Test1_Cont1BState_List2B{
		Index: &Index,
	}

	return t.List2B[key], nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Test1_Cont1BState) Validate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Test1_Cont1BState"], t, opts...); err != nil {
		return err
	}
	return nil
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Test1_Cont1BState) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }


// Test1_Cont1BState_List2B represents the /test1/cont1b-state/list2b YANG schema element.
type Test1_Cont1BState_List2B struct {
	Index	*uint8	`path:"index" module:"test1"`
	Leaf3C	*string	`path:"leaf3c" module:"test1"`
}

// IsYANGGoStruct ensures that Test1_Cont1BState_List2B implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Test1_Cont1BState_List2B) IsYANGGoStruct() {}

// ΛListKeyMap returns the keys of the Test1_Cont1BState_List2B struct, which is a YANG list entry.
func (t *Test1_Cont1BState_List2B) ΛListKeyMap() (map[string]interface{}, error) {
	if t.Index == nil {
		return nil, fmt.Errorf("nil value for key Index")
	}

	return map[string]interface{}{
		"index": *t.Index,
	}, nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Test1_Cont1BState_List2B) Validate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Test1_Cont1BState_List2B"], t, opts...); err != nil {
		return err
	}
	return nil
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Test1_Cont1BState_List2B) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }



var (
	// ySchema is a byte slice contain a gzip compressed representation of the
	// YANG schema from which the Go code was generated. When uncompressed the
	// contents of the byte slice is a JSON document containing an object, keyed
	// on the name of the generated struct, and containing the JSON marshalled
	// contents of a goyang yang.Entry struct, which defines the schema for the
	// fields within the struct.
	ySchema = []byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x5c, 0x4b, 0x6f, 0xe3, 0x36,
		0x10, 0xbe, 0xfb, 0x57, 0x08, 0x3a, 0x16, 0x31, 0x22, 0xc9, 0x71, 0x90, 0xfa, 0xe6, 0xac, 0xbb,
		0x28, 0xb0, 0x49, 0xbb, 0xd8, 0x0d, 0x7a, 0x68, 0x60, 0x14, 0xb4, 0x44, 0x7b, 0x89, 0xca, 0x94,
		0x21, 0x51, 0xbb, 0x36, 0x02, 0xff, 0xf7, 0x42, 0x0f, 0x3f, 0xf4, 0x22, 0x87, 0x8a, 0x9f, 0xdd,
		0x39, 0x05, 0x91, 0x86, 0x22, 0xe7, 0xa1, 0x6f, 0x46, 0xdf, 0x90, 0x7e, 0xeb, 0x18, 0x86, 0x61,
		0x98, 0x7f, 0x90, 0x39, 0x35, 0x07, 0x86, 0xe9, 0xd1, 0xef, 0xcc, 0xa5, 0xe6, 0x4d, 0x76, 0xf5,
		0x13, 0xe3, 0x9e, 0x39, 0x30, 0xec, 0xfc, 0xdf, 0x0f, 0x01, 0x9f, 0xb2, 0x99, 0x39, 0x30, 0xac,
		0xfc, 0xc2, 0x88, 0x85, 0xe6, 0xc0, 0xc8, 0x1e, 0x91, 0x5e, 0x70, 0x03, 0x2e, 0x6c, 0x52, 0xb8,
		0x56, 0x78, 0x7c, 0x7e, 0xff, 0xa6, 0x78, 0xb7, 0x38, 0xcd, 0xf6, 0x72, 0x79, 0xba, 0xed, 0x8d,
		0xcf, 0x21, 0x9d, 0xb2, 0x65, 0x65, 0x96, 0xc2, 0x4c, 0x82, 0x46, 0xc2, 0x2e, 0x4d, 0x94, 0x0a,
		0x7c, 0x0d, 0xe2, 0xd0, 0xa5, 0xb5, 0x83, 0xb3, 0xc5, 0xd0, 0xd5, 0x8f, 0x20, 0x4c, 0xd6, 0x63,
		0x2e, 0xb2, 0x79, 0x6e, 0xea, 0x05, 0x7f, 0x27, 0xd1, 0x30, 0x9c, 0xc5, 0x73, 0xca, 0x85, 0x39,
		0x30, 0x44, 0x18, 0xd3, 0x06, 0xc1, 0x3d, 0xa9, 0x7c, 0x59, 0x15, 0xb9, 0x75, 0xe1, 0xca, 0xba,
		0xa4, 0x6f, 0xd9, 0xcc, 0x05, 0x73, 0x3b, 0xa4, 0x59, 0x97, 0x7d, 0xb3, 0x3b, 0xa4, 0x49, 0x91,
		0x7a, 0xf3, 0x2b, 0xdd, 0x00, 0x71, 0x07, 0xd8, 0x2d, 0x50, 0xf7, 0x68, 0xbb, 0x49, 0xdb, 0x5d,
		0x3a, 0x6e, 0xab, 0x77, 0x5f, 0x83, 0x1b, 0x95, 0xee, 0xdc, 0x0a, 0xf8, 0x94, 0x4c, 0x25, 0x6e,
		0xad, 0x18, 0x35, 0x97, 0x57, 0xe8, 0x33, 0xa2, 0x53, 0x12, 0xfb, 0xa9, 0x3a, 0x8e, 0x4a, 0x36,
		0x0f, 0x09, 0x4b, 0x21, 0xa6, 0x0a, 0x0d, 0x9d, 0x10, 0xd1, 0x0e, 0x15, 0xdd, 0x90, 0x69, 0x1d,
		0x3a, 0xad, 0x43, 0xa8, 0x4d, 0x28, 0xc9, 0x43, 0x4a, 0x11, 0x5a, 0xdb, 0x09, 0x5f, 0x56, 0x0b,
		0xaa, 0x67, 0xed, 0x98, 0x71, 0xf1, 0x00, 0xb1, 0x76, 0x1e, 0x1a, 0x7d, 0x80, 0xe8, 0x17, 0xc2,
		0x67, 0xc9, 0xc3, 0x5f, 0x41, 0x56, 0x82, 0x79, 0x2f, 0x7d, 0xf0, 0x33, 0xe3, 0x60, 0x77, 0x6b,
		0xc6, 0x74, 0x65, 0xd8, 0x5f, 0xc4, 0x8f, 0x69, 0x33, 0x3c, 0x36, 0x8e, 0xfb, 0x18, 0x12, 0x57,
		0xb0, 0x80, 0x8f, 0xd8, 0x8c, 0x89, 0x28, 0x99, 0x18, 0x3c, 0x7e, 0x7d, 0xa3, 0x61, 0x0a, 0xb2,
		0x3c, 0xb9, 0x29, 0x7a, 0x27, 0x34, 0x45, 0xe7, 0x80, 0x06, 0xbb, 0xf4, 0x08, 0xc3, 0x10, 0xdb,
		0xd9, 0xe2, 0xf2, 0x62, 0x4c, 0x29, 0x35, 0xee, 0xb4, 0x1b, 0x2f, 0xf1, 0x45, 0x96, 0xda, 0x27,
		0x9a, 0xa5, 0xc0, 0xe4, 0xd0, 0xe9, 0xdd, 0xc6, 0xf4, 0x7e, 0xe5, 0xe9, 0xdd, 0xa3, 0x2e, 0x9b,
		0x13, 0xff, 0xfe, 0x4e, 0x23, 0xc5, 0xdb, 0x0e, 0x40, 0xb6, 0xf2, 0xde, 0xf5, 0xb0, 0x30, 0xd0,
		0xb5, 0xd8, 0xff, 0x06, 0xb5, 0x1d, 0xcb, 0xb2, 0x4e, 0x68, 0x8d, 0x4b, 0xc7, 0x6d, 0x57, 0x13,
		0xb7, 0xdd, 0x43, 0xe3, 0xb6, 0x83, 0xb8, 0x7d, 0xe5, 0xb8, 0x1d, 0x89, 0x90, 0xf1, 0x99, 0x0e,
		0x68, 0x3f, 0x1c, 0x2b, 0x9a, 0x3d, 0xcd, 0x68, 0xf6, 0x30, 0x9a, 0x31, 0x9a, 0x8b, 0xd6, 0x9e,
		0x04, 0x81, 0x4f, 0x09, 0xd7, 0x09, 0x67, 0xbb, 0x6d, 0x38, 0x6b, 0x51, 0x72, 0x43, 0xce, 0x03,
		0x41, 0x92, 0x3c, 0x24, 0x67, 0xe6, 0x22, 0xf7, 0x1b, 0x9d, 0x93, 0x05, 0x11, 0xdf, 0x12, 0x75,
		0x6e, 0x53, 0x33, 0xdf, 0x66, 0xa4, 0xf6, 0xad, 0x94, 0x64, 0xcd, 0x46, 0x8b, 0x30, 0x76, 0x05,
		0xcf, 0x8d, 0xf1, 0x92, 0x0c, 0xfe, 0xe7, 0x43, 0x32, 0x78, 0x98, 0xfe, 0x71, 0x86, 0xf5, 0x1e,
		0xab, 0xaa, 0x52, 0xa3, 0x46, 0xfa, 0xce, 0xd9, 0x00, 0x2e, 0x38, 0x97, 0x93, 0x73, 0xc1, 0x16,
		0x72, 0xc1, 0xa7, 0xe1, 0x82, 0x95, 0xef, 0x10, 0x3c, 0x13, 0xec, 0x32, 0x80, 0x44, 0xe6, 0x89,
		0xf2, 0x59, 0x1a, 0xbc, 0xf2, 0xd2, 0x1b, 0xf0, 0x4e, 0xeb, 0x94, 0xda, 0x9a, 0x75, 0xe5, 0xb6,
		0x9e, 0xec, 0x03, 0xe5, 0x5b, 0x7c, 0xfc, 0x03, 0x4a, 0x69, 0xad, 0x12, 0xba, 0xad, 0x8a, 0xb6,
		0x75, 0x44, 0x1d, 0x5b, 0x62, 0xe7, 0xf8, 0x3d, 0x30, 0xc4, 0x22, 0x50, 0x4b, 0x2a, 0x97, 0xc3,
		0x96, 0xd4, 0x75, 0xb4, 0xa4, 0xf2, 0xa4, 0x05, 0xac, 0xff, 0x52, 0x69, 0x6c, 0x31, 0x61, 0xf5,
		0xf7, 0xde, 0x6f, 0x19, 0x80, 0x2c, 0x30, 0xa3, 0x5d, 0x09, 0x99, 0x74, 0x87, 0x2d, 0x80, 0x8d,
		0x29, 0x1e, 0xb0, 0x03, 0x90, 0xa8, 0x15, 0x2e, 0xbb, 0x8b, 0xe0, 0x07, 0x0d, 0xe1, 0xe8, 0xbb,
		0x1d, 0x81, 0x08, 0x8c, 0x08, 0x6c, 0x94, 0x9b, 0xfc, 0xf6, 0xbd, 0x06, 0x02, 0xdf, 0xff, 0x6c,
		0x64, 0xbe, 0x63, 0x21, 0x00, 0x6f, 0xdb, 0xfc, 0x16, 0x22, 0x70, 0xa2, 0x96, 0xd0, 0x46, 0x60,
		0x81, 0x08, 0x8c, 0x08, 0x8c, 0x08, 0xdc, 0x0a, 0x75, 0x70, 0x13, 0xcc, 0x79, 0x92, 0xd1, 0xb9,
		0x01, 0x58, 0x8b, 0xaf, 0xf8, 0x44, 0x57, 0x0a, 0x9e, 0xc1, 0x7c, 0x62, 0x91, 0x18, 0x0a, 0xa1,
		0xe0, 0x35, 0x9e, 0x19, 0xff, 0xcd, 0xa7, 0x09, 0x68, 0x24, 0xe6, 0xe2, 0xb1, 0xef, 0x4b, 0x12,
		0xc1, 0x33, 0x59, 0xee, 0x09, 0x03, 0x73, 0x81, 0x6a, 0xc7, 0x85, 0x0e, 0xae, 0xee, 0x63, 0xea,
		0x9c, 0x2c, 0xbb, 0x74, 0xb3, 0x1a, 0x00, 0x48, 0xb4, 0x41, 0xd5, 0x02, 0xa2, 0xde, 0x99, 0x47,
		0xc8, 0xad, 0x7f, 0x86, 0x1e, 0x0d, 0xa9, 0xf7, 0xb8, 0xca, 0xad, 0x7f, 0xf2, 0xae, 0x8d, 0x94,
		0x87, 0x34, 0x14, 0x5d, 0x9b, 0x24, 0xc6, 0x34, 0xba, 0x36, 0xd2, 0x4d, 0xfe, 0x0a, 0x65, 0x64,
		0x4a, 0xd4, 0x1d, 0x73, 0x68, 0x5e, 0x75, 0x71, 0xb9, 0xbb, 0x45, 0xed, 0x2d, 0x28, 0x3b, 0xa7,
		0x31, 0xe9, 0x46, 0x82, 0x08, 0xaa, 0x38, 0xcd, 0xb1, 0x91, 0xd2, 0x3c, 0xd3, 0xe1, 0xe0, 0x99,
		0x8e, 0x7a, 0x65, 0x15, 0xbd, 0x76, 0x58, 0x8f, 0x1d, 0xdc, 0xc7, 0x73, 0x90, 0x40, 0x3f, 0x48,
		0x42, 0x82, 0xf7, 0xf1, 0x94, 0x15, 0x20, 0xa0, 0xf2, 0x03, 0x56, 0x7c, 0x17, 0xd2, 0xc5, 0xb3,
		0x2d, 0xf0, 0xb6, 0xb0, 0x6b, 0xee, 0xe4, 0x39, 0x47, 0x56, 0xf3, 0x5c, 0xcd, 0xbc, 0x09, 0xb0,
		0x99, 0x37, 0x79, 0x6f, 0x33, 0x0f, 0xb1, 0xe8, 0x30, 0x58, 0xa4, 0x6c, 0xe6, 0x31, 0xee, 0xd1,
		0x25, 0xbc, 0x82, 0xcd, 0xc4, 0x91, 0xca, 0x40, 0x2a, 0xc3, 0xc0, 0x13, 0x63, 0x1a, 0x09, 0x01,
		0xa9, 0xe4, 0x5d, 0x6e, 0xec, 0xf7, 0x91, 0x4b, 0xde, 0x54, 0xf7, 0x3d, 0xcd, 0x7d, 0xe1, 0x3d,
		0xdc, 0x17, 0x8e, 0xe0, 0x5b, 0xb2, 0x36, 0xee, 0xa5, 0x40, 0x26, 0x19, 0x99, 0xe4, 0xc3, 0x31,
		0xc9, 0xb2, 0x1a, 0xf7, 0xf8, 0x54, 0xb2, 0x42, 0xf8, 0xcc, 0x24, 0x69, 0xce, 0xf0, 0xdd, 0x4a,
		0xbf, 0xf2, 0xe4, 0xa4, 0xe3, 0xe3, 0xd7, 0xe4, 0x09, 0x19, 0x5f, 0xfa, 0x78, 0x6e, 0xbe, 0xb4,
		0x9e, 0xb1, 0x84, 0x28, 0x00, 0xa1, 0x4e, 0x93, 0x8c, 0x3d, 0x14, 0x2f, 0xc1, 0xe2, 0x89, 0x7e,
		0xa7, 0x7e, 0x33, 0x79, 0x5a, 0x92, 0xab, 0xa7, 0x4f, 0x2d, 0xfc, 0x49, 0x9c, 0x7a, 0xef, 0x37,
		0x66, 0x4c, 0x75, 0x86, 0x94, 0x65, 0x44, 0xf3, 0x33, 0x11, 0x82, 0x86, 0xbc, 0x31, 0x05, 0x9a,
		0xaf, 0xc3, 0xee, 0xdf, 0xe3, 0xb7, 0xde, 0xba, 0xfb, 0x6a, 0x75, 0x7f, 0x1d, 0xff, 0x52, 0x5d,
		0xf7, 0xb8, 0x29, 0x46, 0x3a, 0x7b, 0x7a, 0x34, 0x45, 0xaf, 0xc9, 0xa2, 0x8f, 0xe4, 0x5f, 0xfa,
		0x25, 0x08, 0xaa, 0x06, 0x2c, 0x47, 0xb4, 0xb9, 0x7f, 0xab, 0x10, 0xb7, 0xa3, 0xec, 0x47, 0x9c,
		0xb2, 0x09, 0x3b, 0xeb, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc6,
		0x93, 0x63, 0x50, 0xe3, 0x49, 0x00, 0x00,
	}
)


// ΛEnumTypes is a map, keyed by a YANG schema path, of the enumerated types that
// correspond with the leaf. The type is represented as a reflect.Type. The naming
// of the map ensures that there are no clashes with valid YANG identifiers.
var ΛEnumTypes = map[string][]reflect.Type{
}

