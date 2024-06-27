// simple for Loop
for i := 0; i < 5; i++ {
    fmt.Println(i)
}


//while loop
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}


// infine loop
i := 0
for {
    fmt.Println(i)
    i++
    if i == 5 {
        break
    }
}


// for each loop
numbers := []int{1, 2, 3, 4, 5}
for index, num := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, num)
}

// select  ----------------------------------------------------------------------------------

func main() {
    fruit := "apple"

    switch fruit {
    case "banana":
        fmt.Println("It's a banana")
    case "apple":
        fmt.Println("It's an apple")
    default:
        fmt.Println("Unknown fruit")
    }
}

// select with or
    num := 3

    switch num {
    case 1, 2:
        fmt.Println("One or two")
    case 3, 4:
        fmt.Println("Three or four")
    default:
        fmt.Println("Other number")
    }


// select with condition
age := 20

switch {
case age < 18:
	fmt.Println("Underage")
case age >= 18 && age < 65:
	fmt.Println("Adult")
default:
	fmt.Println("Senior")
}

// select type
func getType(t interface{}) {
    switch t := t.(type) {
    case int:
        fmt.Println("Integer")
    case string:
        fmt.Println("String")
    default:
        fmt.Printf("Unknown type %T\n", t)
    }
}
func main() {
    getType(42)         // prints: Integer
    getType("Hello")    // prints: String
    getType(3.14)       // prints: Unknown type float64
}

//conditional----------------------------------------------------
num := 10
if num < 0 {
	fmt.Println("Number is negative")
} else if num > 0 {
	fmt.Println("Number is positive")
} else {
	fmt.Println("Number is zero")
}


