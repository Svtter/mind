package model_test

import (
	"testing"

	model "github.com/svtter/mind/internal"
)

func TestAddMinds(t *testing.T) {
	mind := &model.Mind{
		Title:   "test",
		Content: "test2",
	}

	if mind == nil {
		t.Fatal("the mind not create.")
	} else {
		t.Log(mind)
	}
}
