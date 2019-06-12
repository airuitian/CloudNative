package map_test

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Log(len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Log(len(m2))
	m3 := make(map[int]int, 10)
	t.Log(len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	v, ok := m1[3]

	if ok {
		t.Log(v)
	} else {
		t.Log("not existing")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	if mySet[1] {
		t.Log("existing")
	} else {
		t.Log("not existing")
	}
}
