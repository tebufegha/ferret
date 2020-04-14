package testing

import (
	"context"
	"github.com/MontFerret/ferret/pkg/stdlib/strings"

	"github.com/MontFerret/ferret/pkg/runtime/core"
	"github.com/MontFerret/ferret/pkg/runtime/values"
)

// NotMatch asserts that value does not match the regular expression.
// @param (Mixed) - Actual value.
// @param (String) - Regular expression.
// @param (String) - Message to display on error.
func NotMatch(ctx context.Context, args ...core.Value) (core.Value, error) {
	err := core.ValidateArgs(args, 2, 3)

	if err != nil {
		return values.None, err
	}

	value := args[0]
	regexp := args[1]

	out, err := strings.RegexMatch(ctx, value, regexp)

	if err != nil {
		return values.None, err
	}

	if out.Compare(values.False) == 0 {
		return values.None, nil
	}

	if len(args) > 2 {
		return values.None, core.Error(ErrAssertion, args[2].String())
	}

	return values.None, core.Errorf(ErrAssertion, "expected %s to not match regular expression", value, regexp)
}
