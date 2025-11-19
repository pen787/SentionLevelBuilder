package objects

import (
	"log"

	"github.com/pen787/SentionLevelBuilder/emath"
)

type IDObjects map[uint32]Object
type Layers struct {
	IDGen *emath.IDGenerator
	Data  [255]IDObjects
}

func (l *Layers) AddObject(layer uint8, obj Object) bool {
	if layer > 255 || layer < 0 {
		panic("Out of layer bound to add object!")
	}

	if l.Data[layer] == nil {
		l.Data[layer] = make(IDObjects)
	}
	obj.SetID(l.IDGen.Next())
	obj.SetLayer(layer)

	objs := l.Data[layer]
	if _, ok := objs[obj.GetID()]; ok {
		return false
	}
	l.Data[layer][obj.GetID()] = obj
	err := obj.Ready()
	if err != nil {
		log.Fatalln(err)
	}

	return true
}

func (l *Layers) RemoveObject(layer uint8, obj Object) bool {
	if layer > 255 || layer < 0 {
		panic("Out of layer bound to remove object!")
	}

	if l.Data[layer] == nil {
		return false
	}

	objs := l.Data[layer]
	for key, v := range objs {
		if v == obj {
			objs[key] = nil
			obj.SetLayer(255)
			err := obj.Exit()
			if err != nil {
				log.Fatalln(err)
			}
			return true
		}
	}

	return false
}

func (l *Layers) RemoveObjectById(layer uint8, id uint32) bool {
	if layer > 255 || layer < 0 {
		panic("Out of layer bound to remove by id object!")
	}

	if l.Data[layer] == nil {
		return false
	}

	objs := l.Data[layer]
	if _, ok := objs[id]; ok {
		objs[id].SetLayer(255)
		objs[id] = nil
	}

	return true
}

func MakeLayer(idgen *emath.IDGenerator) *Layers {
	return &Layers{
		IDGen: idgen,
		Data:  [255]IDObjects{},
	}
}
