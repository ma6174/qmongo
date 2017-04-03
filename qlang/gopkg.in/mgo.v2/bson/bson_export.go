package bson

import (
	"encoding/binary"
	"io"

	"gopkg.in/mgo.v2/bson"

	qlang "qlang.io/spec"
)

func ReadOne(r io.Reader) (b []byte, err error) {
	var docLen int32
	err = binary.Read(r, binary.LittleEndian, &docLen)
	if err != nil {
		return
	}
	b = make([]byte, int(docLen))
	binary.LittleEndian.PutUint32(b, uint32(docLen))
	_, err = io.ReadFull(r, b[4:])
	return
}

func Unmarshal(in interface{}) (out interface{}, err error) {
	switch tin := in.(type) {
	case []byte:
		err = bson.Unmarshal(tin, &out)
		return
	case string:
		err = bson.Unmarshal([]byte(tin), &out)
		return
	case io.Reader:
		one, err := ReadOne(tin)
		if err != nil {
			return nil, err
		}
		err = bson.Unmarshal(one, &out)
		return out, err
	default:
		panic("unexpect in")
	}
}

var Exports = map[string]interface{}{
	"_name": "gopkg.in/mgo.v2/bson",

	// var
	"MaxKey":    bson.MaxKey,
	"MinKey":    bson.MinKey,
	"SetZero":   bson.SetZero,
	"Undefined": bson.Undefined,

	// func
	"now":                 bson.Now,
	"readOne":             ReadOne,
	"marshal":             bson.Marshal,
	"unmarshal":           Unmarshal,
	"objectIdHex":         bson.ObjectIdHex,
	"newObjectId":         bson.NewObjectId,
	"isObjectIdHex":       bson.IsObjectIdHex,
	"newObjectIdWithTime": bson.NewObjectIdWithTime,

	"M":              qlang.StructOf((*bson.M)(nil)),
	"D":              qlang.StructOf((*bson.D)(nil)),
	"Binary":         qlang.StructOf((*bson.Binary)(nil)),
	"DocElem":        qlang.StructOf((*bson.DocElem)(nil)),
	"ObjectId":       qlang.StructOf((*bson.ObjectId)(nil)),
	"MongoTimestamp": qlang.StructOf((*bson.MongoTimestamp)(nil)),
}
