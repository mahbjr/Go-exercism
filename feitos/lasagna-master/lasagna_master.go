package lasagna

// PreparationTime calculates the total preparation time based on the number of layers
// and an average preparation time per layer.
// If averagePreparationTime is 0, a default value of 2 minutes per layer is used.
func PreparationTime(layers []string, averagePreparationTime int) int {
    if averagePreparationTime == 0 {
        averagePreparationTime = 2
    }
    return len(layers) * averagePreparationTime
}

// Quantities calculates the amounts of noodles and sauce needed based on the layers.
// Each noodle layer requires 50 grams of noodles.
// Each sauce layer requires 0.2 liters of sauce.
func Quantities(layers []string) (int, float64) {
    noodles := 0
    sauce := 0.0
    
    for _, layer := range layers {
        if layer == "noodles" {
            noodles += 50
        } else if layer == "sauce" {
            sauce += 0.2
        }
    }
    
    return noodles, sauce
}

// AddSecretIngredient takes your friend's list of ingredients and your list, then
// adds your friend's secret ingredient (the last item in their list) to your list
// (which has a ? as placeholder for the secret ingredient).
func AddSecretIngredient(friendsList, myList []string) {
    secretIngredient := friendsList[len(friendsList)-1]
    myList[len(myList)-1] = secretIngredient
}

// ScaleRecipe scales the amounts in a recipe by the requested number of portions.
// The original recipe is for 2 portions.
func ScaleRecipe(amounts []float64, portions int) []float64 {
    if len(amounts) == 0 {
        return []float64{}
    }
    
    scaledAmounts := make([]float64, len(amounts))
    scaleFactor := float64(portions) / 2.0
    
    for i, amount := range amounts {
        scaledAmounts[i] = amount * scaleFactor
    }
    
    return scaledAmounts
}
