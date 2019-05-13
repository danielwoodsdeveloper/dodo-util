package dodo

import (
	"testing"
)

var id string
var aid string

func TestStore(t *testing.T) {
	st := Settings{"http://localhost:6060", "", "", "", ""}
	d, err := NewDodoConnection(st)
	if err != nil {
		panic(err)
	}

	doc := map[string]interface{}{"id": "12345"}
	id, err = d.Store(doc)
	if err != nil {
		panic(err)
	}
}

func TestGet(t *testing.T) {
	st := Settings{"http://localhost:6060", "", "", "", ""}
	d, err := NewDodoConnection(st)
	if err != nil {
		panic(err)
	}

	r, err := d.Get(id)
	if err != nil {
		panic(err)
	}

	if r["id"] != "12345" {
		t.Errorf("Stored document contains incorrect data, got: %s, want: %s.", r["id"], "12345")
	}
}

func TestModify(t *testing.T) {
	st := Settings{"http://localhost:6060", "", "", "", ""}
	d, err := NewDodoConnection(st)
	if err != nil {
		panic(err)
	}

	doc := map[string]interface{}{"id": "67890"}
	err = d.Modify(id, doc)
	if err != nil {
		panic(err)
	}
}

func TestGetAgain(t *testing.T) {
	st := Settings{"http://localhost:6060", "", "", "", ""}
	d, err := NewDodoConnection(st)
	if err != nil {
		panic(err)
	}

	r, err := d.Get(id)
	if err != nil {
		panic(err)
	}

	if r["id"] != "67890" {
		t.Errorf("Stored document contains incorrect data, got: %s, want: %s.", r["id"], "67890")
	}
}

func TestStoreAnother(t *testing.T) {
	st := Settings{"http://localhost:6060", "", "", "", ""}
	d, err := NewDodoConnection(st)
	if err != nil {
		panic(err)
	}

	doc := map[string]interface{}{"id": "12345"}
	aid, err = d.Store(doc)
	if err != nil {
		panic(err)
	}
}

func TestGetAll(t *testing.T) {
	st := Settings{"http://localhost:6060", "", "", "", ""}
	d, err := NewDodoConnection(st)
	if err != nil {
		panic(err)
	}

	r, err := d.GetAll()
	if err != nil {
		panic(err)
	}

	if len(r) < 2 {
		t.Errorf("Not enough documents found in the store")
	}
}

func TestDelete(t *testing.T) {
	st := Settings{"http://localhost:6060", "", "", "", ""}
	d, err := NewDodoConnection(st)
	if err != nil {
		panic(err)
	}

	err = d.Delete(id)
	if err != nil {
		panic(err)
	}

	err = d.Delete(aid)
	if err != nil {
		panic(err)
	}
}

func TestGetMissing(t *testing.T) {
	st := Settings{"http://localhost:6060", "", "", "", ""}
	d, err := NewDodoConnection(st)
	if err != nil {
		panic(err)
	}

	r, err := d.Get(id)
	if err != nil {
		panic(err)
	}

	if len(r) != 0 {
		t.Errorf("Found a document, but this is supposed to be deleted")
	}

	r, err = d.Get(aid)
	if err != nil {
		panic(err)
	}

	if len(r) != 0 {
		t.Errorf("Found a document, but this is supposed to be deleted")
	}
}