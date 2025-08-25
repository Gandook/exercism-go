package thefarm

import (
    "errors"
    "fmt"
)

func DivideFood (calculator FodderCalculator, cowNum int) (float64, error) {
    amount, fodderErr := calculator.FodderAmount(cowNum)
    if fodderErr != nil {
        return 0.0, fodderErr
    }

    factor, fatteningErr := calculator.FatteningFactor()
	if fatteningErr != nil {
        return 0.0, fatteningErr
    }

    return amount * factor / float64(cowNum), nil
}

func ValidateInputAndDivideFood (calculator FodderCalculator, cowNum int) (float64, error) {
    if cowNum > 0 {
        return DivideFood(calculator, cowNum)
    } else {
        return 0.0, errors.New("invalid number of cows")
    }
}

type InvalidCowsError struct {
    cowNum int
    message string
}

func (icErr *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", icErr.cowNum, icErr.message)
}

func ValidateNumberOfCows(cowNum int) error {
    if cowNum < 0 {
        return &InvalidCowsError {
            cowNum: cowNum,
            message: "there are no negative cows",
        }
    } else if cowNum == 0 {
        return &InvalidCowsError {
            cowNum: cowNum,
            message: "no cows don't need food",
        }
    } else {
        return nil
    }
}
