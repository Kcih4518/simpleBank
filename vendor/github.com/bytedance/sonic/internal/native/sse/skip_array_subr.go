//go:build !noasm || !appengine
// +build !noasm !appengine

// Code generated by asm2asm, DO NOT EDIT.

package sse

import (
	"github.com/bytedance/sonic/loader"
)

const (
	_entry__skip_array = 176
)

const (
	_stack__skip_array = 160
)

const (
	_size__skip_array = 10428
)

var (
	_pcsp__skip_array = [][2]uint32{
		{0x1, 0},
		{0x6, 8},
		{0x8, 16},
		{0xa, 24},
		{0xc, 32},
		{0xd, 40},
		{0x11, 48},
		{0x26b7, 160},
		{0x26b8, 48},
		{0x26ba, 40},
		{0x26bc, 32},
		{0x26be, 24},
		{0x26c0, 16},
		{0x26c1, 8},
		{0x26c2, 0},
		{0x28bc, 160},
	}
)

var _cfunc_skip_array = []loader.CFunc{
	{"_skip_array_entry", 0, _entry__skip_array, 0, nil},
	{"_skip_array", _entry__skip_array, _size__skip_array, _stack__skip_array, _pcsp__skip_array},
}
