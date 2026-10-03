# Practice 2: Collections & Strings

## Language Notes

**Arrays vs. Slices**
In Go, an **Array** has a fixed size determined at compile time. A **Slice** is a dynamically-sized view into an underlying array. You will almost always use slices in Go.

```go
// Array (fixed size of 7)
var weeklyData [7]float64

// Slice (dynamic size)
var orders []string
orders = append(orders, "Cup of Ambition") // adding to a slice
```

**Maps**
Maps are key-value stores (similar to Python dictionaries). You must initialize them using `make()` before inserting data, otherwise your program will panic.

```go
counts := make(map[string]int)
counts["London Fog"]++ // In Go, missing map keys default to 0, so this safely increments to 1

// Checking if a key exists
val, exists := counts["Green Tea Freeze"]
```

**Range Iteration**
The `range` keyword is the standard way to iterate over collections.

```go
// Iterating over a slice
for index, item := range orders {
    fmt.Println(index, item)
}

// Iterating over a map
for key, value := range counts {
    fmt.Println(key, value)
}
```

_(Tip: If you don't need the index or key, you must use the blank identifier `_` to ignore it: `for _, item := range orders`)_

**Strings Package**
The standard library `strings` package is extremely useful for parsing raw data.

```go
raw := "Cup of Ambition, extra foam, whole milk"
parts := strings.Split(raw, ",") // ["Cup of Ambition", " extra foam", " whole milk"]

clean := strings.TrimSpace(" extra foam ") // "extra foam"
isOnline := strings.HasPrefix("ONLINE: Order 1", "ONLINE:") // true
```

## Assignment

You are building the end-of-day analytics and log parser for the Chaus POS system. Open `main.go` and implement the following functions.

1. `CalculateWeeklyRevenue(sales [7]float64) (float64, float64)`
   - Given a fixed array of 7 transaction totals representing Monday-Sunday, calculate the total revenue and the average daily revenue.
   - Return `(total, average)`.

2. `FilterOnlineOrders(logs []string) []string`
   - You are given a slice of raw order logs.
   - Filter the logs to keep _only_ the strings that start with the exact prefix `"ONLINE:"`.
   - Return a new slice containing only these online orders.
   - _(Hint: Use `strings.HasPrefix` and `append()`)_.

3. `ParseModifiers(ticket string) []string`
   - You are given a single ticket string, formatted as `"Drink Name, Modifier 1, Modifier 2"`.
   - Use `strings.Split` to separate the string by commas.
   - The first element is the drink itself. You should extract only the modifiers (everything after the first element).
   - Use `strings.TrimSpace` to clean up any leading/trailing spaces on each modifier.
   - If a modifier is completely empty after trimming, do not include it.
   - Return a slice of the cleaned modifiers. If there are no modifiers, return an empty slice or `nil`.

4. `TallyDrinks(drinks []string) map[string]int`
   - You are given a slice of drink names (e.g., `["London Fog", "Cup of Ambition", "London Fog"]`).
   - Create a `map[string]int` and tally how many times each drink appears in the slice.
   - Return the map.

## Testing

Run the test suite to check your work:

```bash
go test -v ./...
```
