package main

import (
	"github.com/TheLazyLemur/tinyecs/ecs"

	rl "github.com/gen2brain/raylib-go/raylib"

	lo "github.com/samber/lo"
)

type BulletCollisionSystem struct{}

func (b BulletCollisionSystem) Update(world *ecs.World, dt float32, entity []ecs.Entity) {
	bullets := lo.FilterMap(entity, func(x ecs.Entity, _ int) (ecs.Entity, bool) {
		if x.Components[BulletComponentID] != nil {
			return x, true
		}

		return x, false
	})

	enemies := lo.FilterMap(entity, func(x ecs.Entity, _ int) (ecs.Entity, bool) {
		if x.Components[EnemyComponentID] != nil {
			return x, true
		}

		return x, false
	})

	for _, bullet := range bullets {
		bulletSize := ecs.GetComponent[*SizeComponent](bullet, SizeComponentID)
		bulletPos := ecs.GetComponent[*PositionComponent](bullet, PositionComponentID)
		bulletRec := rl.NewRectangle(float32(bulletPos.x), float32(bulletPos.y), float32(bulletSize.w), float32(bulletSize.h))

		for _, enemy := range enemies {
			enemySize := ecs.GetComponent[*SizeComponent](enemy, SizeComponentID)
			enemyPos := ecs.GetComponent[*PositionComponent](enemy, PositionComponentID)
			enemyRec := rl.NewRectangle(float32(enemyPos.x), float32(enemyPos.y), float32(enemySize.w), float32(enemySize.h))

			if rl.CheckCollisionRecs(bulletRec, enemyRec) {
				world.RemoveEntity(bullet.ID)
				world.RemoveEntity(enemy.ID)
			}
		}
	}
}

func (b BulletCollisionSystem) GetFilters() (ecs.FilterType, []string) {
	return ecs.EntityTag, []string{EnemyTag, BulletTag}
}
