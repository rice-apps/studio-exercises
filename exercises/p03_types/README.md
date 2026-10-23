# Exercise 3: Structs, Interfaces, Errors, and Testing

You are building an order fulfillment engine for a cafe across multiple Go files. You will practice custom errors, interface polymorphism, enums, and testing.

## Core Concepts

Go approaches object-orientated programming (OOP), memory, and polymorphism with very different design decisions than Python or Java.

---

### 1. Value Types vs. Reference Types (Pointers)

In **Python** and **Java**, object variables are **reference types**:
- A variable does not contain the object's data directly; it holds a reference pointing to the object.
- When you assign `b = a` or pass an object into a function, both variables point to the exact same object in memory:

```python
# python

class Pastry:
    def __init__(self, name, base_price):
        self.name = name
        self.base_price = base_price

a = Pastry("Croissant", 3.00)
b = a
b.base_price = 5.00

print(a.base_price)  # Prints 5.0 — 'a' was mutated because 'b' references the same object!
```

In **Go**, structs are **value types** by default:
- A variable of type `Pastry` holds the actual struct data directly.
- Assigning `b := a` or passing a struct into a function creates an independent **copy** of the entire struct:

```go
// golang
type Pastry struct {
    Name      string
    BasePrice float64
}

a := Pastry{Name: "Croissant", BasePrice: 3.00}
b := a
b.BasePrice = 5.00

fmt.Println(a.BasePrice) // Prints 3.0 — 'a' was NOT mutated because 'b' is a separate copy!
```

- If a method is defined with a **value receiver**:
  ```go
  // golang
  
  func (p Pastry) Price() float64
  ```
  Go passes a copy of `Pastry` into the method. Mutating any field inside the method only changes the local copy, leaving the original struct untouched.

To share and mutate an existing struct across functions or variables, you explicitly use **pointers**:
- **`&` (Address-of operator):** Obtains a pointer (memory address) to an existing struct.
- By using `&`, you get reference behavior similar to Python:

```go
// golang

a := Pastry{Name: "Croissant", BasePrice: 3.00}
b := &a // 'b' is a pointer (*Pastry) storing the memory address of 'a'
b.BasePrice = 5.00

fmt.Println(a.BasePrice) // Prints 5.0 — 'a' WAS mutated because 'b' points directly to 'a'!
```

In methods, you use the pointer type `*` to declare a **pointer receiver**:
- **`*` (Pointer receiver):** Declares that a method receives the pointer to the original struct rather than a copy:
  ```go
  // golang
  
  func (p *Pastry) SetPrice(newPrice float64) {
      p.BasePrice = newPrice // Modifies the caller's original struct!
  }
  ```

> **Design Rule:** 
> - Use a pointer receiver (`*MyType`) if the method mutates the struct or if the struct is large.
> - Use a value receiver (`MyType`) if the method is read-only and the struct is small.

---

### 2. OOP & Polymorphism in Go

In traditional OOP languages (Java, C++, Python):
- Polymorphism relies on **class hierarchies and inheritance**.
- Conformance is **nominal**: a class must explicitly declare that it inherits from or implements a named interface:

```java
// java

public interface Notifier {
    void send(String message);
}

public class EmailNotifier implements Notifier { // Must explicitly state 'implements Notifier'
    private String address;

    public EmailNotifier(String address) {
        this.address = address;
    }

    public void send(String message) {
        System.out.println("Sending email to " + this.address + ": " + message);
    }
}

// Usage:
Notifier n = new EmailNotifier("alice@example.com");
```

Go intentionally omits classes, inheritance, and the `implements` keyword entirely:
- **No Class Hierarchies:** Go favors composition over inheritance. Structs are simple data containers, and methods are functions bound to a specific receiver type.
- **Structural Polymorphism (Implicit Interfaces):** An interface is purely a behavior contract consisting of method signatures:
  ```go
  // golang
  
  type Notifier interface {
      Send(message string) error
  }
  ```
- If a type declares methods matching the interface's names and signatures, it **automatically satisfies that interface**:

```go
// golang

// Defining a matching Send method automatically satisfies Notifier.
type EmailNotifier struct {
    Address string
}

func (e *EmailNotifier) Send(message string) error {
    fmt.Printf("Sending email to %s: %s\n", e.Address, message)
    return nil
}

// Usage:
var n Notifier = &EmailNotifier{Address: "alice@example.com"}
```

#### Why This Matters
Polymorphism allows collections to process heterogeneous types through a single interface without knowing their concrete implementations:
```go
// golang

// A slice of Notifier can hold any type that satisfies the interface:
alerts := []Notifier{
    &EmailNotifier{Address: "alice@example.com"},
    &SMSNotifier{PhoneNumber: "555-0199"},
}

for _, alert := range alerts {
    alert.Send("System update available") // Dispatches polymorphically
}
```
The caller does not care whether an alert sends an email, an SMS, or a push notification—it only relies on the behavioral contract defined by `Notifier`.

#### Receiver Types & Interface Satisfaction
In Go, the receiver type determines what satisfies an interface:
- If a method has a **pointer receiver** (`func (e *EmailNotifier) Send(...)`), then **`*EmailNotifier`** (the pointer) satisfies `Notifier`, but `EmailNotifier` (the value) does not. That is why instances with pointer receiver methods are passed as pointers (`&EmailNotifier{...}`).
- If all methods have **value receivers**, then both the value and the pointer satisfy the interface.

---

### 3. Enums, Errors, & Packages Quick Reference

- **Enums:** Go does not have a dedicated `enum` keyword. Instead, use typed string constants (`type MilkType string`) when values map to dictionary/map keys, or `iota` (`type DrinkSize int`) for sequential integer constants.
- **Custom Errors:** Any struct implementing the `Error() string` method satisfies the standard `error` interface. In tests, use `errors.As(err, &target)` to inspect fields of a custom error.
- **Multi-File Packages:** All files in this directory declare `package main`. Types and functions declared in one file are visible across all other files in the same directory without import statements. Run your application with `go run .`.

---

## Business Rules & Specifications

### 1. Stock Shortage Reporting (`errors.go`)
- When inventory is insufficient, the system reports an `OutOfStockError` recording the item name, amount needed, and amount currently available.
- Its error message must follow the format: `"out of stock: need <needed> of <item> but have <available>"`.

### 2. Drinks (`drink.go`)
- **Pricing:** Large drinks (`SizeLarge`) have a \$1.00 surcharge added to their base price. Small drinks (`SizeSmall`) cost their base price.
- **Preparation:** Deducts the drink's `MilkOunces` from the inventory map. Drinks with no milk (`MilkNone`) require no ingredients and always succeed. If inventory cannot cover the required ounces, return an `OutOfStockError`.
- _Verify with:_ `go test -v -run TestDrink ./...`

### 3. Pastries (`pastry.go`)
- **Pricing:** Pastries always sell for their flat base price.
- **Preparation:** Deducts 1 unit of that pastry from inventory. If fewer than 1 unit is available, return an `OutOfStockError`.
- _Verify with:_ `go test -v -run TestPastry ./...`

### 4. Order Fulfillment (`order.go`)
- An order fulfills each of its items in sequence:
  - Accumulates the total price across all items using the `MenuItem` interface.
  - Prepares each item against the shared inventory map.
  - **Failure behavior:** If any item fails to prepare, fulfillment aborts immediately and returns a total of \$0.00 and the error. (Any items prepared before the failure remain deducted from inventory).
  - If all items succeed, returns the total price and `nil`.

### 5. Shortage Test (`order_test.go`)
- Implement `TestOrder_OutOfStock` to verify that attempting to fulfill an order with insufficient inventory:
  1. Returns a non-nil error.
  2. The error is specifically an `*OutOfStockError` (verified via `errors.As`).
  3. The error's `Item`, `Needed`, and `Available` fields accurately reflect the shortage.
- _Verify everything:_ `go test -v ./...`
- _Run demo:_ `go run .`
