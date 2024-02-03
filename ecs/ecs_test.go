package ecs

import "testing"

type testSystem struct {
	baseSystem *BaseSystem
	state      int
	numEnts    int
}

func (s *testSystem) Update(e []*Entity, dt float32) {
	s.state = 1
	s.numEnts = len(e)
}

func (s *testSystem) GetFilter() (FilterType, []string) { return All, []string{"hello"} }
func (s *testSystem) SetBaseSystem(bs *BaseSystem)      { s.baseSystem = bs }

func TestAddSystem(t *testing.T) {
	w := World{}
	s := &testSystem{}

	w.AddEntity(&Entity{
		Components: map[string]Component{
			"hello": s,
		},
	})

	w.AddEntity(&Entity{
		Components: map[string]Component{
			"world": s,
		},
	})

	w.AddSystem(s)
	w.Update(0)

	if s.baseSystem.World != &w {
		t.Error("World not set on system")
	}

	if s.numEnts == 2 {
		t.Error("Unexpected number of entities")
	}

	if s.state != 1 {
		t.Error("System not updated")
	}
}

func TestAddEntity(t *testing.T) {
	w := World{}
	w.AddEntity(&Entity{})
	w.Update(0)

	if len(w.entities) != 1 {
		t.Error("Entity not added")
	}
}

func TestUpdate(t *testing.T) {
	w := World{}
	s := &testSystem{}

	w.AddSystem(s)

	w.Update(0)

	if s.state != 1 {
		t.Error("System not updated")
	}
}
