package main

import "fmt"

// creaet a struct type of vehicle
type vehicle struct {
	doors int
	color string
}

// creaet a struct type of truct, embedded with vehicle
type truck struct {
	vehicle
	fourWheel bool
}

// creaet a struct type of sedan, embedded with vehicle
type sedan struct {
	vehicle
	luxury bool
}

func main() {

	// create type vehicle for truck 1
	vehicle_truck_1 := vehicle{
		doors: 4,
		color: "skyblue",
	}

	// create type vehicle for truck 2
	vehicle_truck_2 := vehicle{
		doors: 4,
		color: "grey",
	}

	// create truct 1, embedded with vehicle_truck_1
	truck_1 := truck{
		vehicle:   vehicle_truck_1,
		fourWheel: true,
	}

	// create truct 2, embedded with vehicle_truck_2
	truck_2 := truck{
		vehicle:   vehicle_truck_2,
		fourWheel: false,
	}
	fmt.Println("trucks: ")
	fmt.Println(truck_1)
	fmt.Println(truck_1.vehicle.doors, truck_1.vehicle.color, truck_1.fourWheel)
	fmt.Println(truck_2)
	fmt.Println(truck_2.vehicle.doors, truck_2.vehicle.color, truck_2.fourWheel)

	// create type vehicle for sedan 1
	vehicle_sedan_1 := vehicle{
		doors: 4,
		color: "black",
	}
	// create type vehicle for sedan 2
	vehicle_sedan_2 := vehicle{
		doors: 4,
		color: "red",
	}

	// create sedan 1, embedded with vehicle_sedan_1
	sedan_1 := sedan{
		vehicle: vehicle_sedan_1,
		luxury:  true,
	}
	// create sedan 1, embedded with vehicle_sedan_1
	sedan_2 := sedan{
		vehicle: vehicle_sedan_2,
		luxury:  true,
	}
	fmt.Println("sedans: ")
	fmt.Println(sedan_1)
	fmt.Println(sedan_1.vehicle.doors, sedan_1.vehicle.color, sedan_1.luxury)
	fmt.Println(sedan_2)
	fmt.Println(sedan_2.vehicle.doors, sedan_2.vehicle.color, sedan_2.luxury)

}
