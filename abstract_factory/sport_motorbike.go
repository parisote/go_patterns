package abstract_factory

type SuperMotorbike struct{}

func (s *SuperMotorbike) NumWheels() int {
	return 2
}

func (s *SuperMotorbike) NumSeats() int {
	return 1
}

func (s *SuperMotorbike) GetMotorbikeType() int {
	return SuperMotorbikeType
}
