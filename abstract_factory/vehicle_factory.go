package abstract_factory

type VehicleFactory interface {
	Build(v int) (Vehicle, error)
}
