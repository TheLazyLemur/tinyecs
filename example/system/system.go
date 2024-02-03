package system

import (
	"github.com/TheLazyLemur/tinyecs/ecs"
	comp "github.com/TheLazyLemur/tinyecs/example/component"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlayerSystem struct{
	BaseSystem *ecs.BaseSystem
}

func (s *PlayerSystem) Update(entites []*ecs.Entity, dt float32) {
	for _, ent := range entites {
		s.updateOne(ent, dt)
	}
}

func (s *PlayerSystem) updateOne(ent *ecs.Entity, dt float32) {
	pos := ent.Components[comp.PositionComponentName].(*comp.PositionComponent)
	vel := ent.Components[comp.VelocityComponentName].(*comp.VelocityComponent)
	speed := ent.Components[comp.SpeedComponentName].(*comp.SpeedComponent)

	if rl.IsKeyDown(rl.KeyD) {
		vel.X = speed.Speed * dt
	}

	if rl.IsKeyDown(rl.KeyA) {
		vel.X = -speed.Speed * dt
	}

	if rl.IsKeyDown(rl.KeyW) {
		vel.Y = -speed.Speed * dt
	}

	if rl.IsKeyDown(rl.KeyS) {
		vel.Y = speed.Speed * dt
	}

	pos.X += vel.X
	pos.Y += vel.Y

	rl.DrawRectangleV(rl.Vector2{X: float32(pos.X), Y: float32(pos.Y)}, rl.Vector2{X: 10, Y: 10}, rl.Red)
}

func(s *PlayerSystem) SetBaseSystem(_ *ecs.BaseSystem) () {}

func (s *PlayerSystem) GetFilter() (ecs.FilterType, []string) {
	return ecs.All, []string{
		comp.PositionComponentName,
		comp.VelocityComponentName,
		comp.SpeedComponentName,
	}
}
