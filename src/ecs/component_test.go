package ecs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Compiler check to verify that `Component` satisfies `IComponent`
var _ IComponent = Component{}

/*****************************
 * Component without require *
 *****************************/
type componentWithoutRequire struct{ Component }

/***************************************
 * Components that requires each other *
 ***************************************/
type componentThatRequiresEachOtherA struct{ Component }
type componentThatRequiresEachOtherB struct{ Component }

func (a componentThatRequiresEachOtherA) requiredComponents() []IComponent {
	return []IComponent{componentThatRequiresEachOtherB{}}
}
func (a componentThatRequiresEachOtherB) requiredComponents() []IComponent {
	return []IComponent{componentThatRequiresEachOtherA{}}
}

/**********************************
 * Component that requires itself *
 **********************************/
type componentThatRequiresItself struct{ Component }

func (a componentThatRequiresItself) requiredComponents() []IComponent {
	return []IComponent{componentThatRequiresItself{}}
}

/****************************************************
 * Components with a tree structure *
 ****************************************************
 *
 *				    ---> 2A <---
 *			  	  /			   	 \
 *				 /				  \
 *		   ---> 1A ----> 2B -----> 3A
 *		 /			 	  \
 * 0A --			   	    -----> 3B
 *		 \
 *		  \
 *		    --> 1B ----> 2C -----> 3C
 *						/
 *					   /
 * 0B <---------------
 *
 ****************************************************/
type componentTree0A struct{ Component }
type componentTree0B struct{ Component }
type componentTree1A struct{ Component }
type componentTree1B struct{ Component }
type componentTree2A struct{ Component }
type componentTree2B struct{ Component }
type componentTree2C struct{ Component }
type componentTree3A struct{ Component }
type componentTree3B struct{ Component }
type componentTree3C struct{ Component }

func (a componentTree0A) requiredComponents() []IComponent {
	return []IComponent{componentTree1A{}, componentTree1B{}}
}
func (a componentTree1B) requiredComponents() []IComponent {
	return []IComponent{componentTree2C{}}
}
func (a componentTree2C) requiredComponents() []IComponent {
	return []IComponent{componentTree3C{}, componentTree0B{}}
}
func (a componentTree1A) requiredComponents() []IComponent {
	return []IComponent{componentTree2A{}, componentTree2B{}}
}
func (a componentTree2B) requiredComponents() []IComponent {
	return []IComponent{componentTree3A{}, componentTree3B{}}
}
func (a componentTree3A) requiredComponents() []IComponent {
	return []IComponent{componentTree2A{}}
}

func TestGetAllRequiredComponents(t *testing.T) {
	scenarios := []struct {
		description       string
		components        []IComponent
		nrExpectedResults int
	}{
		{
			description:       "handles empty list of components",
			components:        []IComponent{},
			nrExpectedResults: 0,
		},
		{
			description:       "handles component with recursive require",
			components:        []IComponent{componentThatRequiresEachOtherA{}},
			nrExpectedResults: 1,
		},
		{
			description:       "handles components that require each other",
			components:        []IComponent{componentThatRequiresEachOtherA{}, componentThatRequiresEachOtherB{}},
			nrExpectedResults: 0,
		},
		{
			description:       "handles component that requires itself",
			components:        []IComponent{componentThatRequiresItself{}},
			nrExpectedResults: 0,
		},
		{
			description:       "returns empty for component without required components",
			components:        []IComponent{componentWithoutRequire{}},
			nrExpectedResults: 0,
		},
		{
			description:       "handles complex tree of required components",
			components:        []IComponent{componentTree0A{}, componentTree0B{}},
			nrExpectedResults: 8,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			typesToExclude := toComponentTypes(scenario.components)
			result := getAllRequiredComponents(&typesToExclude, scenario.components)
			assert.Equal(t, scenario.nrExpectedResults, len(result))
		})
	}
}
