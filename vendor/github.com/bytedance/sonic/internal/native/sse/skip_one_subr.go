//go:build !noasm || !appengine
// +build !noasm !appengine

// Code generated by asm2asm, DO NOT EDIT.

package sse

import (
	"github.com/bytedance/sonic/loader"
)

const (
	_entry__skip_one = 176
)

const (
	_stack__skip_one = 160
)

const (
	_size__skip_one = 10428
)

var (
	_pcsp__skip_one = [][2]uint32{
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

var _cfunc_skip_one = []loader.CFunc{
	{"_skip_one_entry", 0, _entry__skip_one, 0, nil},
	{"_skip_one", _entry__skip_one, _size__skip_one, _stack__skip_one, _pcsp__skip_one},
}