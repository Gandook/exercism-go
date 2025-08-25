package lasagna

func PreparationTime(layers []string, timePerLayer int) int {
    if timePerLayer == 0 {
        return len(layers) * 2
    } else {
        return len(layers) * timePerLayer
    }
}

func Quantities(layers []string) (int, float64) {
    noodleGram := 0
    var sauceLiter float64 = 0.0
    for _, layer := range layers {
        if layer == "sauce" {
            sauceLiter += 0.2
        } else if layer == "noodles" {
            noodleGram += 50
        }
    }
    return noodleGram, sauceLiter
}

func AddSecretIngredient(friendsList, myList []string) {
    myList = append(myList[:len(myList) - 1], friendsList[len(friendsList) - 1])
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
    ratio := float64(portions) / 2.0
    result := make([]float64, len(quantities))
    copy(result, quantities)
    for i, _ := range result {
        result[i] *= ratio
    }
    return result
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
