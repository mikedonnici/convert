# Convert package

Provides functions for handling units of measurement and conversions between them.

## Usage examples

Convert gallons per acre to litres per hectare, using compound units in the form `gal1ac-1`.

```go
v, _ := convert.ValueFromTo(10, "gal1ac-1", "l1ha-1")
fmt.Println(v) // 93.5396
```

Convert gallons per acre to litres per hectare, using compound units in the form `gal/ac`.

```go
v, _ := convert.ValueFromTo(10, "gal/ac", "l/ha")
fmt.Println(v) // 93.5396
```

Try to convert a mass to a volume
    
```go
v, err := convert.ValueFromTo(10, "kg", "l")
fmt.Println(err) // cannot convert from kg to l
```





