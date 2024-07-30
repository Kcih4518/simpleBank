//go:build !noasm || !appengine
// +build !noasm !appengine

// Code generated by asm2asm, DO NOT EDIT.

package sse

import (
	"github.com/bytedance/sonic/loader"
)

const (
	_entry__html_escape = 64
)

const (
	_stack__html_escape = 64
)

const (
	_size__html_escape = 1296
)

var (
	_pcsp__html_escape = [][2]uint32{
		{0x1, 0},
		{0x6, 8},
		{0x8, 16},
		{0xa, 24},
		{0xc, 32},
		{0xd, 40},
		{0x11, 48},
		{0x505, 64},
		{0x506, 48},
		{0x508, 40},
		{0x50a, 32},
		{0x50c, 24},
		{0x50e, 16},
		{0x50f, 8},
		{0x510, 0},
	}
)

var _cfunc_html_escape = []loader.CFunc{
	{"_html_escape_entry", 0, _entry__html_escape, 0, nil},
	{"_html_escape", _entry__html_escape, _size__html_escape, _stack__html_escape, _pcsp__html_escape},
}
