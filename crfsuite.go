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

type Feature struct {
	Key   string
	Value float32
}

type FeatureExtractor func(items []string, position int) []Feature

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

func (d *Dictionary) ToString(id int) string {
	s := C.DictionaryToString(d.Original, C.int(id))
	defer C.DictionaryFree(d.Original, s)
	return C.GoString(s)
}

type Model struct {
	Labels     Dictionary
	Attributes Dictionary
	Original   *C.struct_tag_crfsuite_model
}

func (m *Model) getAttributes() Dictionary {
	dictionary := C.GetModelAttributes(m.Original)
	return Dictionary{Original: dictionary}
}

func (m *Model) getLabels() Dictionary {
	dictionary := C.GetModelLabels(m.Original)
	return Dictionary{Original: dictionary}
}

func (m *Model) GetTagger() Tagger {
	tagger := C.GetModelTagger(m.Original)
	return Tagger{
		Labels:     &m.Labels,
		Attributes: &m.Attributes,
		Original:   tagger,
	}
}

func NewModelFromFile(path string) Model {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	model := C.NewModelFromFile(cPath)
	m := Model{Original: model}
	m.Labels = m.getLabels()
	m.Attributes = m.getAttributes()
	return m
}

type Tagger struct {
	Labels     *Dictionary
	Attributes *Dictionary
	Original   *C.struct_tag_crfsuite_tagger
}

func (t *Tagger) Set(inst Instance) {
	C.SetTaggerInstance(t.Original, inst.Original)
}

func (t *Tagger) Tag(items []string, extractor FeatureExtractor) []int {
	inst := NewInstance()
	defer C.InstanceFinish(inst.Original)
	for i := 0; i < len(items); i++ {
		item := NewItem()
		features := extractor(items, i)
		for i := 0; i < len(features); i++ {
			feature := features[i]
			id := t.Attributes.ToID(feature.Key)
			if id > 0 { // TODO: skip unknown features
				attribute := NewAttribute(id, feature.Value)
				item.AddAttribute(attribute)
			}
		}
		inst.AddItem(item, 0) // TODO:
	}
	if !inst.Empty() {
		t.Set(inst)
		length := inst.Length()
		array := C.TaggerDecode(t.Original, C.int(length))
		slice := (*[1 << 30]C.int)(unsafe.Pointer(array))[:length:length] // slice that points to original data
		labels := make([]int, length)                                     // brand new slice, that can be tracked by go gc
		for i := 0; i < length; i++ {
			labels[i] = int(slice[i])
		}
		defer C.free(unsafe.Pointer(array))
		return labels
	} else {
		return []int{}
	}
}

func (t *Tagger) IDsToLabels(ids []int) []string {
	labels := make([]string, len(ids))
	for i := 0; i < len(ids); i++ {
		labels[i] = t.Labels.ToString(ids[i])
	}
	return labels
}

type Attribute struct {
	Original unsafe.Pointer
}

func NewAttribute(id int, value float32) Attribute {
	attribute := C.NewAttribute(C.int(id), C.float(value))
	return Attribute{Original: unsafe.Pointer(&attribute)}
}

type Item struct {
	Original unsafe.Pointer
}

func (i *Item) AddAttribute(attr Attribute) {
	C.AppendAttributeToItem(i.Original, attr.Original)
}

func NewItem() Item {
	item := C.NewItem()
	return Item{Original: unsafe.Pointer(&item)}
}

type Instance struct {
	Original unsafe.Pointer
}

func (i *Instance) Empty() bool {
	status := int(C.EmptyInstance(i.Original))
	if status != 0 {
		return true
	} else {
		return false
	}
}

func (i *Instance) Length() int {
	return int(C.InstanceLength(i.Original))
}

func (i *Instance) AddItem(item Item, label_id int) {
	C.AddItemToInstance(i.Original, item.Original, C.int(label_id))
}

func NewInstance() Instance {
	inst := C.NewInstance()
	return Instance{Original: unsafe.Pointer(&inst)}
}
