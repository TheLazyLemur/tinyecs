package entities

import (
	comp "github.com/TheLazyLemur/tinyecs/example/component"
	"github.com/TheLazyLemur/tinyecs/ecs"
)

func Player() *ecs.Entity {
	return &ecs.Entity{
		Components: map[string]ecs.Component{
			comp.PositionComponentName: &comp.PositionComponent{
				X: 0,
				Y: 0,
			},
			comp.VelocityComponentName: &comp.VelocityComponent{
				X: 0,
				Y: 0,
			},
			comp.SpeedComponentName: &comp.SpeedComponent{
				Speed: 300,
			},
			comp.ShootingComponentName: &comp.ShootingComponent{
				Cooldown: 0.5,
			},
		},
	}
}

func Enemy(x, y float32) *ecs.Entity {
	return &ecs.Entity{
		Components: map[string]ecs.Component{
			comp.PositionComponentName: &comp.PositionComponent{
				X: x,
				Y: y,
			},
			comp.PlayerFinderComponentName: &comp.PlayerFinderComponent{},
		},
	}
}

func Bullet(x, y float32) *ecs.Entity {
	return &ecs.Entity{
		Components: map[string]ecs.Component{
			comp.PositionComponentName: &comp.PositionComponent{
				X: x,
				Y: y,
			},
			comp.LinearSpeedComponentName: &comp.LinearSpeedComponent{
				Speed: 600,
			},
		},
	}
}

func Drop(x, y float32) *ecs.Entity {
	return &ecs.Entity{
		Components: map[string]ecs.Component{
			comp.PositionComponentName: &comp.PositionComponent{
				X: x,
				Y: y,
			},
			comp.DropComponentName: &comp.DropComponent{},
		},
	}
}

func Pickup(x, y float32) *ecs.Entity {
	return &ecs.Entity{
		Components: map[string]ecs.Component{
			comp.PickupComponentName: &comp.PickupComponent{
				X: x,
				Y: y,
			},
		},
	}
}
