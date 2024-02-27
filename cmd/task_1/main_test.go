package main

import (
	"reflect"
	"testing"
)

func TestAddTopic(t *testing.T) {
	api := myApi{
		storage: make(map[uint64]topic),
	}

	for i := 0; i < 5; i++ {
		ok := api.AddTopic(topic{
			id: uint64(i),
		})

		if !ok {
			t.Error("error at adding topic")
		}
	}

	for i := 0; i < 5; i++ {
		_, ok := api.storage[uint64(i)]
		if !ok {
			t.Error("topic not in storage")
		}
	}

	for i := 0; i < 5; i++ {
		ok := api.AddTopic(topic{
			id: uint64(i),
		})

		if ok {
			t.Error("no error recieved")
		}
	}
}

func TestTGetopic(t *testing.T) {
	api := myApi{
		storage: make(map[uint64]topic),
	}

	for i := 0; i < 5; i++ {
		api.storage[uint64(i)] = topic{
			id: uint64(i),
		}
	}

	for i := 0; i < 5; i++ {
		recievedTopic := api.GetTopic(uint64(i))

		if !reflect.DeepEqual(recievedTopic, topic{
			id: uint64(i),
		}) {
			t.Error("error at getting topic")
		}
	}

	for i := 5; i < 7; i++ {
		recievedTopic := api.GetTopic(uint64(i))

		if reflect.DeepEqual(recievedTopic, topic{
			id: uint64(i),
		}) {
			t.Error("no error recieved")
		}
	}
}

func TestRemoveTopic(t *testing.T) {
	api := myApi{
		storage: make(map[uint64]topic),
	}

	for i := 0; i < 5; i++ {
		api.storage[uint64(i)] = topic{
			id: uint64(i),
		}
	}

	for i := 0; i < 5; i++ {
		ok := api.RemoveTopic(uint64(i))

		if !ok {
			t.Error("error at removing topic")
		}
	}

	for i := 0; i < 5; i++ {
		_, ok := api.storage[uint64(i)]
		if ok {
			t.Error("topic in storage")
		}
	}

	for i := 0; i < 5; i++ {
		ok := api.RemoveTopic(uint64(i))

		if ok {
			t.Error("no error recieved")
		}
	}
}
