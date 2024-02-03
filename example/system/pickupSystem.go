package system

import (
	"github.com/TheLazyLemur/tinyecs/ecs"
	"github.com/TheLazyLemur/tinyecs/example/component"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PickupSystem struct {
	BaseSystem *ecs.BaseSystem
}

func (s *PickupSystem) Update(entites []*ecs.Entity, dt float32) {
	for _, ent := range entites {
		s.updateOne(ent, dt)
	}

}

func (s *PickupSystem) updateOne(ent *ecs.Entity, dt float32) {
	pickup := ent.Components[component.PickupComponentName].(*component.PickupComponent)
	rl.DrawRectangleV(rl.Vector2{X: float32(pickup.X), Y: float32(pickup.Y)}, rl.Vector2{X: 10, Y: 10}, rl.Red)
}

func (s *PickupSystem) SetBaseSystem(bs *ecs.BaseSystem) {
	s.BaseSystem = bs
}

func (s *PickupSystem) GetFilter() (ecs.FilterType, []string) {
	return ecs.All, []string{
		component.PickupComponentName,
	}
}
