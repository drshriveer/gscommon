package gerrors_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/drshriveer/gtools/gerrors/gen"
)

func TestSimpleEnumGeneration(t *testing.T) {
	generator := gen.Generate{
		InFile:  "./custom_error.go",
		OutFile: "./custom_error.gerror.go",
		Types:   []string{"GRPCError"},
	}

	require.NoError(t, generator.Parse())
}

//
// func TestExtendedError_Equality(t *testing.T) {
// 	err1 := internal.L1()
// 	err2 := internal.L1()
// 	assert.True(t, errors.Is(internal.ErrExtendedExample, internal.ErrExtendedExample))
//
// 	_, ok := err1.(interface{ Is(error) bool })
// 	require.True(t, ok)
//
// 	assert.True(t, err1.(gerrors.Error).Is(err2))
// 	assert.True(t, err2.(gerrors.Error).Is(err1))
// 	assert.True(t, errors.Is(err1, err2))
// 	assert.True(t, errors.Is(err2, err1))
// 	assert.True(t, errors.Is(err1, internal.ErrExtendedExample))
// 	assert.True(t, errors.Is(internal.ErrExtendedExample, err1))
//
// 	errToConvert := errors.New("random error")
// 	convertedErr := internal.ErrExtendedExample.Convert(errToConvert)
// 	assert.True(t, errors.Is(convertedErr, internal.ErrExtendedExample))
// 	assert.True(t, errors.Is(convertedErr, errToConvert))
// 	assert.True(t, errors.Is(convertedErr, errToConvert))
//
// 	switch errors.Unwrap(err2) {
// 	case internal.ErrExtendedExample:
// 	default:
// 		assert.Fail(t, "was supposed to reach case above")
// 	}
//
// }
//
// func TestExtendedError_CorrectlyLogged(t *testing.T) {
// 	err, ok := internal.L1().(gerrors.Error)
// 	require.Truef(t, ok, "error must implement the gerror.Error interface")
// 	assert.Contains(t, err.Error(), "GRPCStatus: InvalidArgument, ")
// 	assert.Contains(t, err.Error(), "SomeMessage: Print this message, ")
// 	assert.NotContains(t, err.Error(), "DoNotPrint")
// 	assert.NotContains(t, err.Error(), "this is for internal issue only")
// 	assert.Equal(t, "ErrExtendedExample", err.ErrName())
// 	assert.Equal(t, "extended error example", err.ErrMessage())
// 	assert.Equal(t, "gerrors_test:L3", err.ErrSource())
// 	assert.Equal(t, "", err.DTag())
// }
