package system

import (
	"github.com/TheLazyLemur/tinyecs/ecs"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type UiSystem struct {
	BaseSystem *ecs.BaseSystem
}

func (s *UiSystem) Update(entites []*ecs.Entity, dt float32) {
	button := gui.Button(rl.NewRectangle(50, 150, 100, 40), "Start Game")
	if button {
		s.BaseSystem.World.CallSceneSwap("game")
	}
}

func (s *UiSystem) SetBaseSystem(bs *ecs.BaseSystem) {
	s.BaseSystem = bs
}

func (s *UiSystem) GetFilter() (ecs.FilterType, []string) {
	return ecs.All, []string{}
}
