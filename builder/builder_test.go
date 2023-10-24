package builder

import "testing"

func TestBuilderPatternCar(t *testing.T) {
	manofacturingDirector := ManofacturingDirector{}

	carBuilder := &CarBuilder{}
	manofacturingDirector.SetBuilder(carBuilder)
	manofacturingDirector.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Error("Wheels error")
	}

	if car.Seats != 5 {
		t.Error("Seat error")
	}

	if car.Structure != "Car" {
		t.Error("No car")
	}
}

func TestBuilderPatternBike(t *testing.T) {
	manofacturingDirector := ManofacturingDirector{}

	bikeBuilder := &BikeBuilder{}
	manofacturingDirector.SetBuilder(bikeBuilder)
	manofacturingDirector.Construct()

	bike := bikeBuilder.GetVehicle()

	if bike.Wheels != 2 {
		t.Error("Wheels error")
	}

	if bike.Seats != 1 {
		t.Error("Seat error")
	}

	if bike.Structure != "Bike" {
		t.Error("No bike")
	}
}
