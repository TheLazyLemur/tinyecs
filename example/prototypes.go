package main

import "github.com/TheLazyLemur/tinyecs/ecs"

const PlayerTag = "Player"
func NewPlayer() ecs.Entity {
	return ecs.NewEntity(PlayerTag,
		map[string]ecs.Component{
			PlayerComponentID: &PlayerComponent{},
			SizeComponentID: &SizeComponent{
				w: 10,
				h: 10,
			},
			PositionComponentID: &PositionComponent{
				x: 0,
				y: 0,
			},
			SpeedComponentID: &SpeedComponent{
				speed: 300,
			},
		},
	)
}

const EnemyTag = "Enemy"
func NewEnemy() ecs.Entity {
	return ecs.NewEntity(EnemyTag,
		map[string]ecs.Component{
			EnemyComponentID: &EnemyComponent{},
			SizeComponentID: &SizeComponent{
				w: 100,
				h: 100,
			},
			PositionComponentID: &PositionComponent{
				x: 500,
				y: 300,
			},
			SpeedComponentID: &SpeedComponent{
				speed: 100,
			},
		},
	)
}

const BulletTag = "Bullet"
func NewBullet(x, y float32) ecs.Entity {
	return ecs.NewEntity(BulletTag,
		map[string]ecs.Component{
			BulletComponentID: &BulletComponent{},
			SizeComponentID: &SizeComponent{
				w: 10,
				h: 10,
			},
			PositionComponentID: &PositionComponent{
				x: x,
				y: y,
			},
			TimerComponentID: &TimerComponent{
				timer: 3,
			},
			SpeedComponentID: &SpeedComponent{
				speed: 800,
			},
		},
	)
}
