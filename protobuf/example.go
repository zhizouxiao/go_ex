package main

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	example "github.com/zhizouxiao/go_ex/protobuf/proto"
	"log"
)

type Test_OptionalGroup struct {
	RequiredField string
}

type Test struct {
	Label         string
	Type          int32
	Reps          []int64
	Optionalgroup *Test_OptionalGroup
}

func main() {
	test := &example.Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Reps:  []int64{1, 2, 3},
		Optionalgroup: &example.Test_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	log.Print("len proto:", len(data), string(data))
	newTest := &example.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if test.GetLabel() != newTest.GetLabel() {
		log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	}

	//json序列化，要比proto大的多
	testForJson := &Test{
		Label: "hello",
		Type:  17,
		Reps:  []int64{1, 2, 3},
		Optionalgroup: &Test_OptionalGroup{
			RequiredField: "good bye",
		},
	}
	data, err = json.Marshal(testForJson)
	log.Print("len proto:", len(data), string(data))
	// etc.
}
