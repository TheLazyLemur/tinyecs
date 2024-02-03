package main

import (
	"github.com/TheLazyLemur/tinyecs/ecs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlayerSystem struct {
}

func (m PlayerSystem) Update(world *ecs.World, entity []ecs.Entity) {
	p := entity[0]
	playerComponent := p.Components[PlayerComponentID].(*PlayerComponent)
	playerComponent.fireCooldown -= rl.GetFrameTime()

	size := p.Components[SizeComponentID].(*SizeComponent)
	pos := p.Components[PositionComponentID].(*PositionComponent)
	sc := p.Components[SpeedComponentID].(*SpeedComponent)

	if rl.IsKeyDown(rl.KeyD) {
		pos.x += sc.speed * rl.GetFrameTime()
	}

	if rl.IsKeyDown(rl.KeyA) {
		pos.x -= sc.speed * rl.GetFrameTime()
	}

	if rl.IsKeyDown(rl.KeyW) {
		pos.y -= sc.speed * rl.GetFrameTime()
	}

	if rl.IsKeyDown(rl.KeyS) {
		pos.y += sc.speed * rl.GetFrameTime()
	}

	rl.DrawRectangleV(rl.Vector2{X: pos.x, Y: pos.y}, rl.Vector2{X: size.w, Y: size.h}, rl.Red)

	if playerComponent.fireCooldown <= 0 && rl.IsKeyDown(rl.KeySpace) {
		playerComponent.fireCooldown = 0
		world.AddEntity(NewBullet(pos.x, pos.y))
	}
}

func (m PlayerSystem) GetFilters() (ecs.FilterType, []string) {
	return ecs.All, []string{PlayerComponentID, SizeComponentID, PositionComponentID, SpeedComponentID}
}
