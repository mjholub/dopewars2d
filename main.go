package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type Weapon struct {
	name                         string
	Price                        int
	Damage, Accuracy, FiringRate int
	Melee                        bool
	MeleeOnly                    bool
	MeleeDmg                     int
	Throwable                    bool
	ThrowingDamage               int
	ThrowingAccuracy             int
	MaxStock                     int
	Default                      bool
}

var knuckle = Weapon{"Knuckle", 0, 3, 100, 1, true, true, 1, false, 0, 0, 1, true}
var knife = Weapon{"Knife", 100, 10, 100, 1, false, true, 0, true, 5, 50, 5, false}
var baseballBat = Weapon{"Baseball Bat", 200, 20, 100, 1, false, true, 0, false, 0, 0, 1, false}
var machete = Weapon{"Machete", 300, 30, 100, 1, false, true, 0, true, 15, 30, 2, false}
var pistol = Weapon{"Pistol", 1200, 10, 80, 2, false, false, 0, false, 0, 0, 1, false}
var SMG = Weapon{"SMG", 3000, 20, 50, 3, false, false, 0, false, 0, 0, 1, false}
var shotgun = Weapon{"Shotgun", 2000, 30, 60, 1, false, false, 0, false, 0, 0, 1, false}
var machineGun = Weapon{"Machine Gun", 5000, 40, 30, 4, true, false, 10, false, 0, 0, 1, false}
var handgrenade = Weapon{"Handgrenade", 800, 50, 100, 1, false, false, 0, true, 30, 80, 8, false}

type Drug struct {
	Name        string
	Price       int
	Stock       int
	RaiseWanted int
}

var weed = Drug{"Weed", 50, 0, 2}
var cocaine = Drug{"Cocaine", 300, 0, 4}
var heroin = Drug{"Heroin", 200, 0, 6}
var acid = Drug{"Acid", 40, 0, 0}
var ketamine = Drug{"Ketamine", 100, 0, 1}
var amphetamine = Drug{"Amphetamine", 60, 0, 3}
var meth = Drug{"Meth", 150, 0, 5}
var morphine = Drug{"Morphine", 80, 0, 5}
var shrooms = Drug{"Shrooms", 30, 0, 1}

type weaponUnits struct {
	Name  string
	Count int
}

var Player struct {
	Name                    string
	Health                  int
	Reputation, WantedLevel int
	cash, debt              int
	drugs                   [8]Drug
	weaponsAvailable        [8]Weapon
	weaponUnits             weaponUnits
	weapons                 [4]Weapon
	CurrentDistrict         districtProperties
}
type districtProperties struct {
	name string
	neighbour_a []district
	neighbour_b []district
	drugsAvailable [5]Drug
	hospital bool
	bank bool
	loanShark bool
}
type district interface {
	Name()                      string
	neighbour_a()               []district
	neighbour_b()               []district
	drugsAvailable()            [5]Drug
	hospital()					bool
	bank()						bool
	loanShark() 				bool
	starting()                  bool
	ID()						int
	Properties()				districtProperties
}

type dist struct {
	properties districtProperties
	ID int
}
//create a manhattan struct. The drugs can be the same for each, they should be updated upon drugs_available() call
var manhattan = dist{ districtProperties{"Manhattan", nil, nil, [5]Drug{weed, cocaine, heroin, meth, ketamine}, false, false, false}, 0}
var brooklyn = dist{ districtProperties{"Brooklyn", nil, nil, [5]Drug{weed, cocaine, heroin, meth, ketamine}, false, false, false}, 1}
var queens = dist{ districtProperties{"Queens", nil, nil, [5]Drug{weed, cocaine, heroin, meth, ketamine}, false, false, false}, 2}
var bronx = dist{ districtProperties{"Bronx", nil, nil, [5]Drug{weed, cocaine, heroin, meth, ketamine}, false, false, false}, 3}
var statenIsland = dist{ districtProperties{"Staten Island", nil, nil, [5]Drug{weed, cocaine, heroin, meth, ketamine}, false, false, false}, 4}
//create a manhattan array

//var manhattan = district{"Manhattan", nil, nil, [5]Drug{weed, cocaine, heroin, meth, ketamine}, true, true, false, false}
//var brooklyn = district{"Brooklyn", nil, nil, [5]Drug{amphetamine, meth, morphine, shrooms, heroin}, false, true, false, false}
//var queens = district{"Queens", nil, nil, [5]Drug{weed, cocaine, heroin, acid, amphetamine}, true, false, false, false}
//var statenIsland = district{"Staten Island", nil, nil, [5]Drug{weed, amphetamine, shrooms, acid, ketamine}, false, true, false, false}
//var bronx = district{"Bronx", nil, nil, [5]Drug{meth, morphine, heroin, shrooms, acid}, true, false, true, false}

func drugs_available() {
	alldrugs := [8]Drug{weed, cocaine, heroin, acid, ketamine, amphetamine, meth, morphine}
	//write 5 random drugs to player's current district. Load upon entering district
	for i := 0; i < 5; i++ {
		Player.CurrentDistrict.drugsAvailable[i] = alldrugs[rand.Intn(8)]
	}
}

func main() {
	fmt.Println("Welcome to the city of New York.")
	err := os.Remove("save.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	Player.cash = 10000
	Player.debt = 15000
	fmt.Println("Welcome to Dope Wars!")
	fmt.Println("What is your name?")
	fmt.Scanln(&Player.Name)
	fmt.Println("Welcome to the world of Dope Wars, " + Player.Name + "!")
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
	fmt.Println("You are a small time drug dealer in the city of New York.\n After failing one job after another, you have decided to start a small business.")
	fmt.Println("You have a small amount of cash, but you need to make a lot of money.\nAfter one of your drug deals went down, you were left with a debt.")
	fmt.Println("You have $" + strconv.Itoa(Player.debt) + " to pay off.")
	fmt.Println("Press h for help. Press q to quit.")
	var key string
	fmt.Scanln("%s", &key)
	if key == "h" {
		fmt.Println("h - help")
		fmt.Println("q - quit")
		fmt.Println("i - Player")
		fmt.Println("d - district info and the available drugs. Press d again to see the drugs in stock.\n to buy a drug, type the drug number and press enter.")
		fmt.Println("t - travel to a district. Press t again to see the districts you can travel to.\n to travel to a district, type the district number and press enter.")
		fmt.Println("w - weapon info and the available weapons. Press w again to see the weapons in stock.\n to buy a weapon, type the weapon number and press enter.")
		fmt.Println("a - current weapon info. Press a again to see the current weapon stats. Press s to sell the current weapon.")
		fmt.Println("f - fight the opponent. For throwable weapons, press j to throw the weapon. Note you will lose the weapon if you do not deal a critical hit\n or if it's a handgrenade.")
		fmt.Println("s - sell the drugs. Type the drug number and press enter.")
		fmt.Println("o - make a payment or withdraw/borrow money from the bank or loan shark. Type the amount and press enter.")
		fmt.Println("r - run away. You might lose some cash or drugs and the wanted level will go down.")
		fmt.Println("b - bribe the law enforcement. You will lose some cash and the wanted level will go down.")
		fmt.Println("g - visit the bank or loan shark.")
		fmt.Println("u - visit the hospital.")
		fmt.Println("Press enter to continue.")
		fmt.Scanln()
	} else if key == "q" {
		os.Exit(0)
	}
}

func reputation() {
	switch {
	case Player.Reputation > 0 && Player.Reputation < 10:
		Player.weaponsAvailable = [8]Weapon{knife, baseballBat, knuckle, knuckle, knuckle, knuckle, knuckle, knuckle}
		//a chance of 20% to multiply the price of up to 2 drugs in the Player by 1.5
		if rand.Intn(100) < 30 {
			for i := 0; i < len(Player.drugs); i++ {
				if Player.drugs[i].Price > 0 {
					Player.drugs[i].Price = int(float64(Player.drugs[i].Price) * 1.5)
				}
			}
		}
	case Player.Reputation > 10 && Player.Reputation < 25:
		Player.weaponsAvailable = [8]Weapon{knife, baseballBat, machete, pistol, knuckle, knuckle, knuckle, knuckle}
		//a chance of 40% to multiply the price of up to 3 drugs in the Player by 1.5
		if rand.Intn(100) < 40 {
			for i := 0; i < len(Player.drugs); i++ {
				if Player.drugs[i].Price > 0 {
					Player.drugs[i].Price = int(float64(Player.drugs[i].Price) * 1.5)
				}
			}
		}
	case Player.Reputation > 25 && Player.Reputation < 50:
		Player.weaponsAvailable = [8]Weapon{knife, baseballBat, machete, pistol, SMG, shotgun, knuckle, knuckle}
		//a chance of 60% to multiply the price of up to 4 drugs in the Player by 1.75
		if rand.Intn(100) < 60 {
			for i := 0; i < len(Player.drugs); i++ {
				if Player.drugs[i].Price > 0 {
					Player.drugs[i].Price = int(float64(Player.drugs[i].Price) * 1.75)
				}
			}
		}
	case Player.Reputation > 50:
		Player.weaponsAvailable = [8]Weapon{knife, baseballBat, machete, pistol, SMG, shotgun, machineGun, handgrenade}
		//a chance of 80% to multiply the price of up to 5 drugs in the Player by 2
		if rand.Intn(100) < 80 {
			for i := 0; i < len(Player.drugs); i++ {
				if Player.drugs[i].Price > 0 {
					Player.drugs[i].Price = int(float64(Player.drugs[i].Price) * 2)
				}
			}
		}
	}
}

func buyDrug() {
	fmt.Println("You have $" + strconv.Itoa(Player.cash) + " to spend.")
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
	fmt.Println("What drug would you like to buy?")
	//prints the drugs in the current district
	//if the drug is not available, it will not be printed
	//get the current district

	fmt.Println(district.drugsAvailable)
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
	fmt.Println("How many would you like to buy?")
	fmt.Scanln(&Player.drugs[0].Stock)
	fmt.Println("You have $" + strconv.Itoa(Player.cash) + " to spend.")
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
	fmt.Println("You have bought " + strconv.Itoa(Player.drugs[0].Stock) + " " + Player.drugs[0].Name + ".")
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
}

// sellDrug is a function that allows the Player to sell drugs. Each sale will increase the Player's reputation, but also increase the wanted level, multiplied by the amount of d sold.
func sellDrug() {
	fmt.Println("You have " + strconv.Itoa(Player.drugs[0].Stock) + " " + Player.drugs[0].Name + " to sell.")
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
	// print the numbered list of drugs in the Player with their current stock and price per unit
	for i := 0; i < len(Player.drugs); i++ {
		if Player.drugs[i].Stock > 0 {
			fmt.Println(strconv.Itoa(i+1) + ". " + Player.drugs[i].Name + " - " + strconv.Itoa(Player.drugs[i].Stock) + " units - $" + strconv.Itoa(Player.drugs[i].Price) + " per unit")
		}
	}
	fmt.Println("Which drug would you like to sell?.  Please type the number and press enter.")
	fmt.Scanln(&Player.drugs[0].Name)
	fmt.Println("How many would you like to sell?")
	var unitsSell int
	fmt.Scanln("%d", &unitsSell)

	if unitsSell > Player.drugs[0].Stock {
		fmt.Println("You don't have that many units to sell.")
		fmt.Println("Press enter to continue.")
		fmt.Scanln()
	} else {
		Player.drugs[0].Stock -= unitsSell
		Player.cash += unitsSell * Player.drugs[0].Price
		Player.WantedLevel += Player.drugs[0].RaiseWanted * unitsSell
		fmt.Println("You have sold " + strconv.Itoa(unitsSell) + " " + Player.drugs[0].Name + ".")
		fmt.Println("You have" + strconv.Itoa(Player.drugs[0].Stock) + " " + Player.drugs[0].Name + " left.")
		fmt.Println("Your current cash is $" + strconv.Itoa(Player.cash) + ".")
		fmt.Println("Your reputation has increased to " + strconv.Itoa(Player.Reputation) + ".")
		fmt.Println("Your wanted level has increased to " + strconv.Itoa(Player.WantedLevel) + ".")
		fmt.Println("Press enter to continue.")
		fmt.Scanln()
	}
	//If the Player has a reputation lower than 25, the reputation will increase by 4 for each 4 units sold.
	if Player.Reputation < 25 {
		Player.Reputation += 4 * (unitsSell / 4)
	} else {
		//If the Player has a reputation higher than 25 and lower than 50, the reputation will increase by 3 for each 5 units sold.
		if Player.Reputation > 25 && Player.Reputation < 50 {
			Player.Reputation += 3 * (unitsSell / 5)
		} else {
			//If the Player has a reputation higher than 50, the reputation will increase by 2 for each 6 units sold.
			if Player.Reputation > 50 {
				Player.Reputation += 2 * (unitsSell / 6)
			}
		}
	}

	fmt.Println("Press enter to continue.")
	fmt.Scanln()
}

func min (a, b int) int {
	if a < b {
		return a
	}
	return b
}

func buyWeapon() {
	Player.weapons = [4]Weapon{knuckle, knuckle, knuckle, knuckle}

	fmt.Println("You have $" + strconv.Itoa(Player.cash) + " to spend.")
	fmt.Println("Press enter to continue.")
	fmt.Scanln()
	fmt.Println("What weapon would you like to buy?")
	//print a numbered list of weapons available to the Player, based on their reputation, writable to a weaponChoice variable
	var weaponChoice int
	var maxObtainable int
	for i := 0; i < len(Player.weaponsAvailable); i++ {
		if Player.weaponsAvailable[i].Price > 0 {
			weaponChoice = append(weaponChoice, strconv.Itoa(i+1)+". "+Player.weaponsAvailable[i].Name+" - $"+strconv.Itoa(Player.weaponsAvailable[i].Price)+" per unit")
		}
	}
	//prompt the Player to select the number of the weapon to buy, using the weapon's number in the list as the index
	fmt.Println("Type the number of the weapon you would like to buy and press enter.")
	fmt.Scanln(&weaponChoice)
	//if weapon.MaxStock > 1, prompt the Player to provide the quantity. Read the weapon quantity owned.
	//If the Player owns at least 1 unit of a weapon, subtract the quantity owned and set the maxObtainable variable.
	minObtainable := 1
	//max obtainable is the minimum of the max stock and the Player's cash modulo divided by the weapon's price
	maxObtainable = min(Player.weaponsAvailable[weaponChoice].MaxStock, Player.cash/Player.weaponsAvailable[weaponChoice].Price)
	fmt.Println("Please provide the quantity you wish to purchase (%d - %d):", minObtainable, maxObtainable)
	var weaponQuantity int
	fmt.Scanln(&weaponQuantity)
	switch {
	case weaponQuantity < minObtainable:
		//terminate the function if the Player tries to buy less than 1 unit
		return
	case weaponQuantity > maxObtainable:
		fmt.Println("You cannot afford or carry that many.")
		fmt.Println("Press space to continue. To abort purchase, press c.")
		var abort string
		fmt.Scanln(&abort)
		if abort == "c" {
			return
		}
	default:
		//if the Player has enough cash, subtract the cost of the weapon from the Player's cash and add the weapon to the Player's inventory
		if Player.cash >= weaponQuantity*Player.weaponsAvailable[weaponChoice].Price {
			Player.cash -= weaponQuantity * Player.weaponsAvailable[weaponChoice].Price
			Player.weapons[weaponChoice].Stock += weaponQuantity
			fmt.Println("You have purchased " + strconv.Itoa(weaponQuantity) + " " + Player.weaponsAvailable[weaponChoice].Name + ".")
			fmt.Println("You have $" + strconv.Itoa(Player.cash) + " left.")
			fmt.Println("Press enter to continue.")
			fmt.Scanln()
		} else {
			fmt.Println("You cannot afford that many.")
			fmt.Println("Press enter to continue.")
			fmt.Scanln()
		}

	}
	//charge the Player the price of the weapon
	Player.cash -= Player.weaponsAvailable[weaponChoice-1].Price
	//add the weapon to the Player's Player
	Player.weapons = append(Player.weapons, Player.weaponsAvailable[weaponChoice-1])
}

func travel() {
	//update neighbour_a and neighbour_b in districtProperties for each district
	manhattan.properties.neighbour_a = brooklyn
	manhattan.properties.neighbour_b = queens


	currentDistrict := Player.CurrentDistrict
	//read the t keypress
	//the Player can travel to neighbour_a or neighbour_b
	fmt.Println("Where would you like to travel to? Type 1 or 2 and press enter.")
	var travelChoice int
	fmt.Println("1. " + currentDistrict.properties.neighbour_a.name)
	fmt.Println("2. " + currentDistrict.properties.neighbour_b.Name)
	fmt.Scanln("%s", &travelChoice)
	//if the Player selects 1, travel to neighbour_a
	if travelChoice == 1 {
		Player.CurrentDistrict = currentDistrict.properties.neighbour_a()[0]
	} else {
		//if the Player selects 2, travel to neighbour_b
		Player.CurrentDistrict = currentDistrict.neighbour_b()[0]
	}
	fmt.Println("You have arrived at " + Player.District.Name + ".")
}
