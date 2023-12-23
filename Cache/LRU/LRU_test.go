package Cache

import (
	"testing"
)

var lrucache = New[int, string](2)
var lrucache2 = New[int, int](5)

func TestNew(t *testing.T) {
	cap := lrucache.GetCapacity()
	if cap != 2 {
		t.Errorf("FAILED,error in getCapacity expected 2 got %v", cap)
	} else {
		t.Logf("successfull expected 2 got %v", cap)
	}
}

func TestSet(t *testing.T) {
	lrucache.Set(4, "four")
	getValue, err := lrucache.Get(4)
	if err != nil {
		t.Errorf("Failed,error in GetValue method %v", err)
	} else if getValue != "four" {
		t.Errorf("Failed,error in GetValue method got %v expected four", getValue)
	} else {
		t.Logf("Successfull expected four got %v", getValue)
	}

	lrucache.Set(5, "five")
	lrucache.Set(6, "six")
	lrucache.Set(7, "seven")
	lrucache.Set(8, "eight")
	lrucache.Set(8, "eight")
	getValue1, err := lrucache.Get(4)
	if err != nil && getValue1 == "" {
		t.Logf("Successfull,error in GetValue method caught %v", err)
	} else {
		t.Errorf("Failed, Didn't delete LRU element, got %v", getValue1)
	}

	//multiple insertions simeltenously
	for i := 0; i < 5; i++ {
		go lrucache2.Set(i, i+i)
		go lrucache.Get(i)
	}
}

func Test_getInsertNode(t *testing.T) {
	insertNode := getInsertNode[int, string](3, "three")
	if insertNode.key == 3 && insertNode.val == "three" {
		t.Logf("Successfull expected struct with key: 3 and value: four got %v", insertNode)
	} else {
		t.Errorf("Failed, Got wrong node to insert in getInsert,expected struct with key: 3 and value: four got %v", insertNode)
	}
}
