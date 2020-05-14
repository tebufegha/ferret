package testing

import (
	"context"
	"github.com/MontFerret/ferret/pkg/runtime/core"
)

// Gte asserts that an actual value is lesser than or equal to an expected one.
// @param (Mixed) - Actual value.
// @param (Mixed) - Expected value.
// @param (String) - Message to display on error.
func Lte(_ context.Context, args ...core.Value) (core.Value, error) {
	return compare(args, LessOrEqualOp)
}