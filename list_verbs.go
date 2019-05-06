package main

import (
	"./file"
	"strings"
	"fmt"
	"math/rand"
    "time"
)

func main (){

	rand.Seed(time.Now().UnixNano())

	var opcion int

	listEnglish := file.ReadAndGetFile("english")
	listSpanish := file.ReadAndGetFile("spanish")

	arrayListEnglish := strings.Split(listEnglish, ",")
	arrayListSpanish := strings.Split(listSpanish, ",")
	lengthArray := len(arrayListEnglish)

	fmt.Print("De Ingles-Español(1) o Español-Ingles(2) ?: ")
	fmt.Scanf("%d \n", &opcion)

	switch opcion {
	case 1:
		startGame(lengthArray, arrayListEnglish, arrayListSpanish)
	case 2:
		startGame(lengthArray, arrayListSpanish, arrayListEnglish)
	default:
		fmt.Println("La opción ingresada no existe")	
	}

}

func startGame (lengthArray int, firstArrayList []string, secondArrayList []string){
	var opcion int

	fmt.Println("\nComenzemos.....\n")

	for index, value := range firstArrayList {
		opcion = 0

		fmt.Printf("\nCual es la traducción de %s ? \n", value)

		getPositionsRandoms(index, lengthArray)

		if (lengthArray-2) > index {

			fmt.Printf("1) %s \n", secondArrayList[index])
			fmt.Printf("2) %s \n", secondArrayList[index+1])
			fmt.Printf("3) %s \n", secondArrayList[index+2])		
			
		}else{
			fmt.Printf("1) %s \n", secondArrayList[index])
			fmt.Printf("2) %s \n", secondArrayList[lengthArray - index])
			fmt.Printf("3) %s \n", secondArrayList[lengthArray - (index+1)])			
		}

		fmt.Print("Respuesta:")
		fmt.Scanf("%d \n", &opcion)

	}
}


func getPositionsRandoms (initIndex int, maxIndex int ) {
	var firstRandomPosition, secondRandomPosition int
	var positionVerbs [3] int

	firstRandomPosition = rand.Intn(maxIndex - initIndex)

	for (firstRandomPosition == initIndex) {
		firstRandomPosition = rand.Intn(maxIndex - initIndex)
	}

	secondRandomPosition = rand.Intn(maxIndex - initIndex)

	for (firstRandomPosition == secondRandomPosition) || (secondRandomPosition == initIndex) {
		secondRandomPosition = rand.Intn(maxIndex - secondRandomPosition)
	}

	positionVerbs[0], positionVerbs[1], positionVerbs[2] = firstRandomPosition, secondRandomPosition, initIndex

}

func printArray (array []string) {
	for index, value := range array {
		fmt.Printf("Index: %d, Value: %s \n", index, value)
	}
	fmt.Println("")
}

func disOrder (array []int) []int{
	positioRandom := rand.Intn(0 - (len(array)-1))

	/*
		number_random, value_temporal := 1, 0
		value_temporal = array[i]
		array[i] = array[number_random]


	*/

}