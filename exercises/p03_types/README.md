# Exercise 3: (Intro to) Structs, Interfaces, Errors, and Testing

## Language Notes

**Pointers vs. Values (For Python/Java Developers)**
In Python and Java, objects are automatically passed by "reference". In Go, everything is passed by **value** (a photocopy) by default.
To modify original data, you must use a **pointer**:

- Use `&` to get the pointer (address) of a value: `myDrink := &Drink{}`
- Use `*` to declare that a type is a pointer: `func (d *Drink) Prepare(...) error`
  _Best Practice:_ If a struct has any methods that require a pointer receiver to mutate data, **all** methods on that struct should use pointer receivers for consistency.

**Implicit Interfaces**
Interfaces define behavior contracts. A type satisfies an interface automatically simply by implementing its methods.

```go
type Pricer interface {
    Price() float64
}
// Employee satisfies Pricer!
func (e *Employee) Price() float64 { return 15.0 }
```

**Custom Errors**
A custom error is just a Struct that implements the built-in `error` interface (which requires an `Error() string` method).

```go
type MachineError struct { Code int }
func (e *MachineError) Error() string { return "machine down" }
```

You can use `errors.As` to extract a custom error and check its fields:

```go
var mErr *MachineError
if errors.As(err, &mErr) {
    fmt.Println("Code was:", mErr.Code)
}
```

## Assignment

Open `main.go` and implement the following engine components:

1. **Custom Error `OutOfStockError`**
   - Define the `OutOfStockError` struct with fields: `Item` (string), `Needed` (int), and `Available` (int).
   - Implement the `Error() string` method to return: `"out of stock: need <Needed> of <Item> but have <Available>"`.

2. **`MenuItem` Interface**
   - Requires two methods: `Price() float64` and `Prepare(inventory map[string]int) error`.

3. **`Drink` Struct & Methods (Pointer Receivers)**
   - `Drink` has fields: `Name` (string), `BasePrice` (float64), `MilkType` (string), and `MilkOunces` (int).
   - `Price()`: Returns `BasePrice`.
   - `Prepare(inventory map[string]int) error`:
     - If `MilkType` is `"None"`, return `nil` (success).
     - Otherwise, check if `inventory[MilkType]` has at least `MilkOunces`.
     - If it is short, return a pointer to an `OutOfStockError`.
     - If available, deduct `MilkOunces` from the inventory and return `nil`.

4. **`Pastry` Struct & Methods (Value Receivers)**
   - `Pastry` has fields: `Name` (string), and `BasePrice` (float64).
   - `Price()`: Returns `BasePrice`.
   - `Prepare(inventory map[string]int) error`:
     - Pastries require exactly `1` unit of `Name`.
     - Check if `inventory[Name]` is at least `1`.
     - If it is short, return a pointer to an `OutOfStockError`.
     - If available, deduct `1` and return `nil`.

5. **`Order` Struct & `Fulfill` Method**
   - `Order` has fields: `TicketID` (string) and `Items` ([]MenuItem).
   - `Fulfill(inventory map[string]int) (float64, error)`:
     - Iterate through `Items`. For each item, add its `Price()` to a running total, and call `Prepare(inventory)`.
     - **Important:** If `Prepare` returns an error, immediately return `0.0` and that error (abort the rest of the order). It is expected that items successfully prepared before the failure will have already deducted their ingredients from inventory.
     - If all items succeed, return the total price and `nil`.

## Part 2: Writing Tests

Open `main_test.go`. We have provided the basic success tests.
**Your assignment is to write `TestOrder_OutOfStock`.**

Verify that:

- `Fulfill` returns a non-nil error when an item is out of stock.
- The error is specifically of type `*OutOfStockError` (using `errors.As`).
- The fields `Item`, `Needed`, and `Available` on the error precisely match the failure condition.

## Testing

```bash
go test -v ./...
```
