package component

const PositionComponentName = "Position"
type PositionComponent struct {
	X float32
	Y float32
}

const VelocityComponentName = "Velocity"
type VelocityComponent struct {
	X float32
	Y float32
}

const SpeedComponentName = "Speed"
type SpeedComponent struct {
	Speed float32
}

const PlayerFinderComponentName = "PlayerFinder"
type PlayerFinderComponent struct {
	PlayerPosition *PositionComponent
}

const ShootingComponentName = "Shooting"
type ShootingComponent struct {
	Cooldown float32
}

const LinearSpeedComponentName = "LinearSpeed"
type LinearSpeedComponent struct {
	Speed float32
}

const DropComponentName = "Drop"
type DropComponent struct {
}

const PickupComponentName = "Pickup"
type PickupComponent struct {
	X float32
	Y float32
}
