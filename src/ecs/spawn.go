package ecs

import (
	"fmt"

	"github.com/lucdrenth/murph/engine/src/utils"
)

// Spawn spawns the given components and all their required components that are not declared in the component parameters. Return the associated entityId on success.
//
// An error is returned when any of the given components are of the same type.
//
// Calling Spawn without any components to generate an entityId is allowed.
func Spawn(world *world, components ...IComponent) (entityId, error) {
	componentTypes := toComponentTypes(components)

	// check for duplicates
	duplicate, duplicateIndexA, duplicateIndexB := utils.GetFirstDuplicate(componentTypes)
	if duplicate != nil {
		return 0, fmt.Errorf("can not spawn duplicate component: %s at positions %d and %d", *duplicate, duplicateIndexA, duplicateIndexB)
	}

	// get required components
	requiredComponents := getAllRequiredComponents(&componentTypes, components)
	components = append(components, requiredComponents...)

	// spawn components
	world.entityIdCounter++
	entityId := world.entityIdCounter
	world.entities[entityId] = &entry{
		components: components,
	}

	return entityId, nil
}
