# useful links

[ Multi dimensional example ](https://www.tutorialspoint.com/go/go_multi_dimensional_arrays.html)

# manipulate slices

```
    var s []int                   // a nil slice
    s1 := []string{"foo", "bar"}
    s2 := make([]int, 2)          // same as []int{0, 0}
    s3 := make([]int, 2, 4)       // same as new([4]int)[:2]
    fmt.Println(len(s3), cap(s3)) // 2 4
    var copyS = []i               // a nil slice
    copy(copyS,s)                 // same the  s slice
```

# manipulate string

```
    greetings :=  []string{"Hello","world!"}
    fmt.Println(strings.Join(greetings, " "))
```

# simple for example

```
    sum := 0
    for i := 1; i < 5; i++ {
        sum += i
    }
    fmt.Println(sum) // 10 (1+2+3+4)
```

# "foreach" examples

```
    strings := []string{"hello", "world"}
    // i is index
    // s is value
    for i, s := range strings {
        fmt.Println(i, s)
    }
```

# while

```
    n := 1
    for n < 5 {
        n *= 2
    }
    fmt.Println(n) // 8 (1*2*2*2)
```

# while TRUE

```
    sum := 0
    for {
        sum++ // repeated forever
    }
```

# map or dict

```
    map [ keyType ] ValueType
    m := map[string]int{
        "one":   1,
        "two":   2,
        "three": 3,
    }
    for key, value := range m {
        fmt.Println(key, value)
    }
```
---
```
    countryCapitalMap := map[string] string 
    {
        "France":"Paris",
        "Italy":"Rome",
        "Japan":"Tokyo"
    }
    value  , finded
    capital, ok    := countryCapitalMap["United States"]

    /* if ok is true, entry is present otherwise entry is absent*/
    if(ok){
        fmt.Println("Capital of United States is", capital)  
    } else {
        fmt.Println("Capital of United States is not present") 
    }
```
