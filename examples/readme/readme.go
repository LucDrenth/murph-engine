// The example in the readme. It is added here to catch compiler errors
package main

import (
	"fmt"

	"github.com/lucdrenth/murphy/src/ecs"
)

func main() {
	world := ecs.NewWorld()
	fmt.Printf("hello %T\n", world)
}
