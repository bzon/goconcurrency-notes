# My Concurrency Notes


## Normal Loop

Let us start with a normal loop by running __foo__ and __bar__ functions in a synchronous manner and nothing fancy. Our goal here is how we can improve the program's performance by using go routines and measuring execution time of the program which will be discussed in the proceeding topics.


```go
package main

import (
	"fmt"
	"time"
)

func main() {
	foo()
	bar()
}

func foo() {
	for i := 0; i < 30; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 30; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Bar:", i)
	}
}
```

## Using `go` routines in functions

Let's add the __go__ statement when calling __foo__ and __bar__. By default in the nature of __Go__, running the code below will just exit the program immediately.

We will add a __WaitGroup__ in the next topic to make the __main__ function wait for __foo__ and __bar__ to complete, and gives us a standard output and measure the execution time.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	go foo()
	go bar()
}

func foo() {
	for i := 0; i < 30; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 30; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Bar:", i)
	}
}
```

## Using `sync.WaitGroup` to capture standard output

With a __WaitGroup__ we will now be able to see the printed outputs for each __foo__ and __bar__ calls. Try running the first code and the last code below and you should notice a significant increase in the execution time. The difference maker is because __foo__ and __bar__ concurrently executed at the same time (parallelism).

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 30; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 30; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Bar:", i)
	}
	wg.Done()
}
```
