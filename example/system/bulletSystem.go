package system

import (
	comp "github.com/TheLazyLemur/tinyecs/example/component"
	proto "github.com/TheLazyLemur/tinyecs/example/entities"
	"github.com/TheLazyLemur/tinyecs/ecs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type BulletSystem struct{
	baseSystem *ecs.BaseSystem
}

func (s *BulletSystem) Update(entites []*ecs.Entity, dt float32) {
	for _, ent := range entites {
		s.updateOne(ent, entites, dt)
	}
}

func (s *BulletSystem) updateOne(ent *ecs.Entity, entities []*ecs.Entity, dt float32) {
	_, isBullet := ent.Components[comp.LinearSpeedComponentName]
	if isBullet {
		pos := ent.Components[comp.PositionComponentName].(*comp.PositionComponent)
		speed := ent.Components[comp.LinearSpeedComponentName].(*comp.LinearSpeedComponent)

		pos.X += speed.Speed * dt
		if pos.X > 800 {
			s.baseSystem.World.RemoveEntity(ent)
		}

		rl.DrawRectangleV(rl.Vector2{X: float32(pos.X), Y: float32(pos.Y)}, rl.Vector2{X: 10, Y: 10}, rl.Red)

		for _, e := range entities {
			_, isEnemy := e.Components[comp.PlayerFinderComponentName]
			if isEnemy {
				enemyPos, _ := e.Components[comp.PositionComponentName].(*comp.PositionComponent)
				bulRec := rl.Rectangle{
					X:      float32(pos.X),
					Y:      float32(pos.Y),
					Width:  10,
					Height: 10,
				}

				enemyRec := rl.Rectangle{
					X:      float32(enemyPos.X),
					Y:      float32(enemyPos.Y),
					Width:  100,
					Height: 100,
				}

				if rl.CheckCollisionRecs(bulRec, enemyRec) {
					s.baseSystem.World.RemoveEntity(e)
					s.baseSystem.World.RemoveEntity(ent)
					s.baseSystem.World.AddEntity(proto.Drop(pos.X, pos.Y))
				}
			}
		}
	}
}

func (s *BulletSystem) SetBaseSystem(bs *ecs.BaseSystem) {
	s.baseSystem = bs
}

func (s *BulletSystem) GetFilter() (ecs.FilterType, []string) {
	return ecs.One, []string{
		comp.LinearSpeedComponentName,
		comp.PositionComponentName,
		comp.PlayerFinderComponentName,
	}
}
