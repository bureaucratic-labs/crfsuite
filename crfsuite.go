package crfsuite

/*
#cgo LDFLAGS: -lcrfsuite
#include <stdlib.h>
#include "crfsuite.h"
*/
import "C"

import (
	"unsafe"
)

type Dictionary struct {
	Original *C.struct_tag_crfsuite_dictionary
}

// Obtain the number of strings in the dictionary.
func (d *Dictionary) Length() int {
	return int(C.DictionaryLength(d.Original))
}

// Assign and obtain the integer ID for the string.
func (d *Dictionary) Get(key string) int {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	return int(C.DictionaryGet(d.Original, cKey))
}

// Obtain the integer ID for the string.
func (d *Dictionary) ToID(key string) int {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	return int(C.DictionaryToID(d.Original, cKey))
}

type Model struct {
	Original *C.struct_tag_crfsuite_model
}

func (m *Model) GetAttributes() Dictionary {
	dictionary := C.GetModelAttributes(m.Original)
	return Dictionary{Original: dictionary}
}

func (m *Model) GetLabels() Dictionary {
	dictionary := C.GetModelLabels(m.Original)
	return Dictionary{Original: dictionary}
}

func NewModelFromFile(path string) Model {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	model := C.NewModelFromFile(cPath)
	return Model{Original: model}
}
