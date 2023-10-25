package abstract_factory

import (
	"testing"
)

func TestCarFactory(t *testing.T) {

	carFactory, err := BuildFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	luxuryCarVehicle, err := carFactory.Build(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}

	_, ok := luxuryCarVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion failed.")
	}
}

func TestMotoFactory(t *testing.T) {

	motorbikeFactory, err := BuildFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	cruiseMotorbike, err := motorbikeFactory.Build(CruiseMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}

	_, ok := cruiseMotorbike.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion failed.")
	}
}
