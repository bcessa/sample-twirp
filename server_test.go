package main

import (
	"github.com/bcessa/sample-twirp/proto"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/struct"
	"log"
	"testing"
)

func toStruct(el proto.Message) (*structpb.Struct, error) {
	m := jsonpb.Marshaler{}
	js, _ := m.MarshalToString(el)
	s := &structpb.Struct{}
	err := jsonpb.UnmarshalString(js, s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func TestServer(t *testing.T) {
	m := jsonpb.Marshaler{}
	t.Run("Simple", func(t *testing.T) {
		el := &sample.Item{
			Name: "custom",
			Time: ptypes.TimestampNow(),
		}
		pb, _ := proto.Marshal(el)
		pj, _ := m.MarshalToString(el)
		el2 := &sample.Item{}
		proto.Unmarshal(pb, el2)
		log.Printf("message name: %s", proto.MessageName(el))
		log.Printf("encoded: %x", pb)
		log.Printf("json: %s", pj)
		log.Printf("decoded: %+v", el2)
	})

	t.Run("Extensions", func(t *testing.T) {
		// Get extension
		ab := &sample.AddressBook{}
		ab.Contacts = append(ab.Contacts, &sample.Contact{
			Name: "Rick",
			LastName: "Sanchez",
			Email: "rick@c137.cit",
		})
		ab.Contacts = append(ab.Contacts, &sample.Contact{
			Name: "Morty",
			LastName: "Smith",
			Email: "morty@c137.cit",
		})
		s, _ := toStruct(ab)
		ex := &sample.Extension{
			Id: "sample.contacts",
			Version: "0.1.0",
			Data: s,
		}

		// Build record with extensions
		el := &sample.Item{
			Name: "record",
			Time: ptypes.TimestampNow(),
			Extensions: []*sample.Extension{ex},
		}

		pb, err := proto.Marshal(el)
		if err != nil {
			t.Error("failed to encode element", err)
		}
		pj, err := m.MarshalToString(el)
		if err != nil {
			t.Error("failed to encode as JSON", err)
		}
		el2 := &sample.Item{}
		err = proto.Unmarshal(pb, el2)
		if err != nil {
			t.Error("failed to decode", err)
		}

		log.Printf("message name: %s", proto.MessageName(el))
		log.Printf("encoded: %x", pb)
		log.Printf("json: %s", pj)
		log.Printf("decoded: %+v", el2)
	})
}