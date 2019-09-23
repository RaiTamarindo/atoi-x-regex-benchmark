# Atoi x Regex for integer validation

This project has benchmarks to measure performance metrics to answer the following question:

> Which is the best method for validating an string as an integer number, `strconv.Atoi` or `regexp.MatchString`?

### How this benchmark was made

There was two functions for validation: `validateWithAtoi` and `validateWithRegex`.

Both functions receives a string and returns a boolean that is the validation result.

For validating with regex, the compiling time was not included to measures.

### How to start

`go test -bench=. -benchmem`

### Results on my machine

| Validation Method | Throughput (ns/op) |
|---|--:|--:|--:|
| Atoi | 10.1 |
| Regex | 222 |
