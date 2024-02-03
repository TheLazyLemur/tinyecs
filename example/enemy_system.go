package main

import (
	"github.com/TheLazyLemur/tinyecs/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)


type EnemySystem struct{}

func (b EnemySystem) Update(world *ecs.World, entity []ecs.Entity) {
	for _, e := range entity {
		pc := e.Components[PositionComponentID].(*PositionComponent)
		sz := e.Components[SizeComponentID].(*SizeComponent)


		rl.DrawRectangle(int32(pc.x), int32(pc.y), int32(sz.w), int32(sz.h), rl.Red)
	}
}

func (b EnemySystem) GetFilters() (ecs.FilterType, []string) {
	return ecs.All, []string{PositionComponentID, SizeComponentID, EnemyComponentID}
}
