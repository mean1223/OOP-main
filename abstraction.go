package main

import "fmt"

// ประกาศ Interface สำหรับรถ
type Car interface {
	Start()
	Stop()
	Accelerate(speed int)
	Turn(direction string)
}

// Struct สำหรับรถประเภทต่างๆ
type ElectricCar struct {
	Brand      string
	Model      string
	Color      string
	Year       int
	BatteryPct int
}

type GasCar struct {
	Brand     string
	Model     string
	Color     string
	Year      int
	FuelLevel int
}

// Implement Method ของ ElectricCar
func (e ElectricCar) Start() {
	fmt.Println("Electric car starts: turning on the electrical system")
}

func (e ElectricCar) Stop() {
	fmt.Println("Electric car stops safely")
}

func (e ElectricCar) Accelerate(speed int) {
	fmt.Printf("Electric car accelerates to %d km/h\n", speed)
}

func (e ElectricCar) Turn(direction string) {
	fmt.Printf("Electric car turns %s\n", direction)
}

// Implement Method ของ GasCar
func (g GasCar) Start() {
	fmt.Println("Gas car starts with the key ignition")
}

func (g GasCar) Stop() {
	fmt.Println("Gas car stops safely")
}

func (g GasCar) Accelerate(speed int) {
	fmt.Printf("Gas car accelerates to %d km/h\n", speed)
}

func (g GasCar) Turn(direction string) {
	fmt.Printf("Gas car turns %s\n", direction)
}

func main() {
	electricCar := ElectricCar{
		Brand:      "Tesla",
		Model:      "Model 3",
		Color:      "Red",
		Year:       2023,
		BatteryPct: 80,
	}

	gasCar := GasCar{
		Brand:     "Toyota",
		Model:     "Corolla",
		Color:     "Blue",
		Year:      2020,
		FuelLevel: 60,
	}

	// เรียกใช้ Method ของรถแต่ละประเภท
	electricCar.Start()
	electricCar.Accelerate(100)
	electricCar.Turn("left")
	electricCar.Stop()

	gasCar.Start()
	gasCar.Accelerate(120)
	gasCar.Turn("right")
	gasCar.Stop()
}
