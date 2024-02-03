package main

const PlayerComponentID = "MyComponent"
type PlayerComponent struct {
	fireCooldown float32
}

const SizeComponentID = "Size"
type SizeComponent struct {
	w, h float32
}

const PositionComponentID = "Position"
type PositionComponent struct {
	x, y float32
}

const TimerComponentID = "Timer"
type TimerComponent struct {
	timer float32
}

const SpeedComponentID = "Speed"
type SpeedComponent struct {
	speed float32
}

const EnemyComponentID = "Enemy"
type EnemyComponent struct {}

const BulletComponentID = "Bullet"
type BulletComponent struct {}
