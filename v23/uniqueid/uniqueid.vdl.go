// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: uniqueid

package uniqueid

import (
	"v.io/v23/vdl"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// An Id is a likely globally unique identifier.
type Id [16]byte

func (Id) VDLReflect(struct {
	Name string `vdl:"v.io/v23/uniqueid.Id"`
}) {
}

func (x Id) VDLIsZero() bool {
	return x == Id{}
}

func (x Id) VDLWrite(enc vdl.Encoder) error {
	if err := enc.WriteValueBytes(__VDLType_array_1, x[:]); err != nil {
		return err
	}
	return nil
}

func (x *Id) VDLRead(dec vdl.Decoder) error {
	bytes := x[:]
	if err := dec.ReadValueBytes(16, &bytes); err != nil {
		return err
	}
	return nil
}

// Hold type definitions in package-level variables, for better performance.
var (
	__VDLType_array_1 *vdl.Type
)

var __VDLInitCalled bool

// __VDLInit performs vdl initialization.  It is safe to call multiple times.
// If you have an init ordering issue, just insert the following line verbatim
// into your source files in this package, right after the "package foo" clause:
//
//    var _ = __VDLInit()
//
// The purpose of this function is to ensure that vdl initialization occurs in
// the right order, and very early in the init sequence.  In particular, vdl
// registration and package variable initialization needs to occur before
// functions like vdl.TypeOf will work properly.
//
// This function returns a dummy value, so that it can be used to initialize the
// first var in the file, to take advantage of Go's defined init order.
func __VDLInit() struct{} {
	if __VDLInitCalled {
		return struct{}{}
	}
	__VDLInitCalled = true

	// Register types.
	vdl.Register((*Id)(nil))

	// Initialize type definitions.
	__VDLType_array_1 = vdl.TypeOf((*Id)(nil))

	return struct{}{}
}