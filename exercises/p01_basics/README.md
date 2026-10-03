# Practice 1: Go Basics

This first practice problem covers basic Go syntax: variable declaration, constants, control flow (if/else, switch, for), and multiple return values.

## Getting Started

Before writing any code, copy this exercise from the `template/` directory into your personal student folder so you don't edit the shared starter files.
You can copy the folder using the GoLand visual interface or you can ask Andrew or Calvin about the specific terminal commands if you are curious.

## Language Notes

**Variables & Constants**

```go
var age int = 20
name := "Alex" // Type inferred, usable inside functions only
const BasePrice float64 = 3.50
```

**Control Flow (If & Switch)**
Notice that `if` and `switch` statements don't use parentheses around conditions, and `switch` cases don't require `break`.

```go
if shots > 2 {
    fmt.Println("Highly caffeinated")
}

switch customer {
case "BYOM":
    fmt.Println("Bring Your Own Mug discount applied")
default:
    fmt.Println("Standard pricing")
}
```

**For Loops & Strings**
Go only has one loop keyword: `for`. You can iterate using standard C-style syntax. You can also access individual characters in a string by index (they are treated as byte values).

```go
// Standard loop
for i := 0; i < len(word); i++ {
    // You can compare bytes directly!
    if word[i] >= 'a' && word[i] <= 'z' {
        fmt.Println("Found a lowercase letter!")
    }
}
```

**String Formatting**
To combine strings and variables, use string concatenation (`+`) or `fmt.Sprintf`.

```go
greeting := "Hello, " + name + "!"
receipt := fmt.Sprintf("Total is $%.2f", 4.50)
```

**Multiple Return Values**
Go functions often return more than one value, which is very common for returning a result alongside a status flag or error.

```go
func ApplyTax(price float64) (float64, bool) {
    if price > 0 {
        return price * 1.08, true
    }
    return 0.0, false
}
```

## Assignment

You are writing the core pricing logic for the Chaus Point of Sale (POS) system. Open `main.go` and complete the functions marked with `// TODO`.

1. `GreetCustomer(name string) string`
   - Return a welcome greeting formatted exactly as: `"Welcome to Chaus, <name>!"`

2. `CalculateDrinkPrice(basePrice float64, extraShots int, altMilk bool) float64`
   - Calculate the total price of a drink (like a Cup of Ambition or a London Fog).
   - Start with `basePrice`.
   - Add `ShotPrice` ($0.50) for every extra shot.
   - Add `AltMilkPrice` ($0.75) if `altMilk` is true.
   - If `basePrice` is 0 or less, return `0.0`.

3. `ApplyDiscount(total float64, discountType string) (float64, bool)`
   - Use a `switch` statement on `discountType`.
   - `"BYOM"` (Bring Your Own Mug): return the total minus $0.25, and `true`.
   - `"Student"`: return 10% off (`total * 0.90`), and `true`.
   - `"Faculty"`: return 5% off (`total * 0.95`), and `true`.
   - For any unrecognized discount type, return the original `total` and `false`.

4. `ValidatePromoCode(code string) bool`
   - Check if a promotional discount code is valid.
   - The code length must be between 5 and 10 characters, inclusive. Use `len(code)`.
   - It must contain at least one numeric digit (`'0'` through `'9'`).
   - It must contain at least one uppercase letter (`'A'` through `'Z'`).
   - Use a `for` loop to inspect each character byte (`code[i]`).

## Testing

Run the test suite to check your work:

```bash
go test -v ./...
```
