package builder

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeat() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type ManofacturingDirector struct {
	builder BuildProcess
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

func (f *ManofacturingDirector) Construct() {
	f.builder.SetSeat().SetStructure().SetWheels()
}

func (f *ManofacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeat() BuildProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) SetSeat() BuildProcess {
	b.v.Seats = 1
	return b
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bike"
	return b
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}
