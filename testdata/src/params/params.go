package params

import "context"

type NamedParam interface {
	SingleParam(context.Context) error
}
