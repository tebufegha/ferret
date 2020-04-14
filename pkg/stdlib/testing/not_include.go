package testing

import (
	"context"

	"github.com/MontFerret/ferret/pkg/runtime/core"
	"github.com/MontFerret/ferret/pkg/runtime/values"
	"github.com/MontFerret/ferret/pkg/stdlib/collections"
)

// NotInclude asserts that haystack does not include needle.
// @param (String|Array|Object|Iterable) - Haystack value.
// @param (Mixed) - Needle value.
// @param (String) - Message to display on error.
func NotInclude(ctx context.Context, args ...core.Value) (core.Value, error) {
	err := core.ValidateArgs(args, 2, 3)

	if err != nil {
		return values.None, err
	}

	haystack := args[0]
	needle := args[1]

	out, err := collections.Includes(ctx, haystack, needle)

	if err != nil {
		return values.None, err
	}

	if out.Compare(values.False) == 0 {
		return values.None, nil
	}

	if len(args) > 2 {
		return values.None, core.Error(ErrAssertion, args[2].String())
	}

	return values.None, core.Errorf(ErrAssertion, "expected %s to not include %s", haystack, needle)
}
