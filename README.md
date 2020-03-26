# go-sort

Sorting algorithms implemented in Go for educational purpose.

Create a file with one number on each row.

```go
// create file of numbers
path := "./numbers.txt"
if err := utils.CreateNumbers(path, 50000); err != nil {
    log.Fatalln(err)
}
```

Then create a slice of items that implements the `Comparable` interface.

```go
// read file of numbers
numbers, err := utils.ReadNumbersFromFile(path)
if err != nil {
    log.Fatalln(err)
}
```

Finally sort the slice.

## Bubble sort example

```go
fmt.Println("Starting bubblesort")
bubblesort.Sort(numbers)
```

## Insertion sort example

```go
fmt.Println("Starting insertion sort")
insertionsort.Sort(numbers)
```

---

## Check for errors after the sorting algorithm

```go
fmt.Println("Done, now checking for errors")

err = utils.CheckSorted(numbers)
if err != nil {
    log.Fatalln(err)
}

fmt.Println("Numbers are sorted")
```