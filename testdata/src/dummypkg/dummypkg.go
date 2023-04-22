package dummypkg

import "context"

type xtruct struct {
	a int
}

type Another interface {
	Get() string
}

type NamedParam interface {
	Void()

	NoArgs() string

	WithName(ctx context.Context, number int, toggle bool, xtruct *xtruct, getter Another) (bool, error)

	WithoutName(
		context.Context, // want "interface method WithoutName must have named param for type context.Context"
		int, // want "interface method WithoutName must have named param for type int"
		bool, // want "interface method WithoutName must have named param for type bool"
		xtruct, // want "interface method WithoutName must have named param for type xtruct"
		Another, // want "interface method WithoutName must have named param for type Another"
		struct{ b bool }, // want "interface method WithoutName must have all named params"
	)
}
