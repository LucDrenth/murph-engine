// Demonstrate how to delete an entity from the world
package main

import (
	"fmt"

	"github.com/lucdrenth/murph_engine/src/ecs"
)

type NPC struct{ ecs.Component }

func main() {
	world := ecs.NewWorld()
	entity, _ := ecs.Spawn(&world, NPC{})

	fmt.Printf("Before deleting: %d entity in the world\n", world.CountEntities())
	ecs.Delete(&world, entity)
	fmt.Printf("After deleting: %d entities in the world\n", world.CountEntities())
}
