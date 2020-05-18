package main

import "fmt"

func fuelGuage(fuel int){
  fmt.Println("Fuel Left:", fuel)
}

func calculateFuel(planet string) int{
  var fuel int
  switch planet {
  case "Venus":
    fuel = 300000
  case "Mercury":
    fuel = 500000
  case "Mars":
    fuel = 700000
  default:
    fuel = 0
  }
  return fuel
}

func greenPlanet(planet string){
  fmt.Println("Welcome to", planet)
}


func cantFly(){
  fmt.Println("We do not have the available fuel to fly there.")
}


func flyToPlanet(planet string, fuel int) int{
  var fuelRemaining int
  var fuelCost int
  fuelRemaining = fuel
  fuelCost = calculateFuel(planet)
  if (fuelRemaining >= fuelCost) {
    greenPlanet(planet)
    fuelRemaining -= fuelCost
  } else {
    cantFly()
  }
  return fuelRemaining
}


func main() {  
  fuel := 1000000
  planetChoice := "Mars"
  fuel = flyToPlanet(planetChoice, fuel)
  // And then liftoff!
  fuelGuage(fuel)
  
}
