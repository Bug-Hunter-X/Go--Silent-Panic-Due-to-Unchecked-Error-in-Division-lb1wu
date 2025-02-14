func myFunc(a, b int) (int, error) {
    if a == 0 {
        return 0, errors.New("a cannot be zero")
    }
    if b == 0 {
        return 0, errors.New("b cannot be zero")
    }
    return a / b, nil
}

func main() {
    result, err := myFunc(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
    result, err = myFunc(10, 2) //added a correct case to check
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
} 