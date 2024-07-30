//go:build !noasm || !appengine
// +build !noasm !appengine

// Code generated by asm2asm, DO NOT EDIT.

package base64x

import (
	"github.com/bytedance/sonic/loader"
)

const (
	_entry__b64decode = 1328
	_entry__b64encode = 256
)

const (
	_stack__b64decode = 152
	_stack__b64encode = 40
)

const (
	_size__b64decode = 17616
	_size__b64encode = 864
)

var (
	_pcsp__b64decode = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{10, 32},
		{12, 40},
		{13, 48},
		{17560, 152},
		{17564, 48},
		{17565, 40},
		{17567, 32},
		{17569, 24},
		{17571, 16},
		{17573, 8},
		{17577, 0},
		{17608, 152},
	}
	_pcsp__b64encode = [][2]uint32{
		{1, 0},
		{4, 8},
		{6, 16},
		{8, 24},
		{10, 32},
		{852, 40},
		{853, 32},
		{855, 24},
		{857, 16},
		{859, 8},
		{864, 0},
	}
)

var funcs = []loader.CFunc{
	{"__native_entry__", 0, 67, 0, nil},
	{"_b64decode", _entry__b64decode, _size__b64decode, _stack__b64decode, _pcsp__b64decode},
	{"_b64encode", _entry__b64encode, _size__b64encode, _stack__b64encode, _pcsp__b64encode},
}
