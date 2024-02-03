package system

import (
	"github.com/TheLazyLemur/tinyecs/ecs"
	"github.com/TheLazyLemur/tinyecs/example/component"
	"github.com/TheLazyLemur/tinyecs/example/entities"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type DropSystem struct {
	BaseSystem *ecs.BaseSystem
}

func (s *DropSystem) Update(entites []*ecs.Entity, dt float32) {
	for _, ent := range entites {
		s.updateOne(ent, dt)
	}
}

func (s *DropSystem) updateOne(ent *ecs.Entity, dt float32) {
	rv := rl.GetRandomValue(0, 100)
	if rv > 70 {
		pos := ent.Components[component.PositionComponentName].(*component.PositionComponent)
		s.BaseSystem.World.AddEntity(entities.Pickup(pos.X, pos.Y))
	}

	s.BaseSystem.World.RemoveEntity(ent)
}

func (s *DropSystem) SetBaseSystem(bs *ecs.BaseSystem) {
	s.BaseSystem = bs
}

func (s *DropSystem) GetFilter() (ecs.FilterType, []string) {
	return ecs.All, []string{
		component.DropComponentName,
	}
}
