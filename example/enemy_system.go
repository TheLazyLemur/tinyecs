package main

import (
	"github.com/TheLazyLemur/tinyecs/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type EnemySystem struct{}

func (b EnemySystem) Update(world *ecs.World, dt float32, entity []ecs.Entity) {
	for _, e := range entity {
		pc := e.Components[PositionComponentID].(*PositionComponent)
		sz := e.Components[SizeComponentID].(*SizeComponent)
		sc := e.Components[SpeedComponentID].(*SpeedComponent)

		pc.x -= sc.speed * rl.GetFrameTime()

		rl.DrawRectangle(int32(pc.x), int32(pc.y), int32(sz.w), int32(sz.h), rl.Red)
	}
}

func (b EnemySystem) GetFilters() (ecs.FilterType, []string) {
	return ecs.EntityTag, []string{EnemyTag}
}
