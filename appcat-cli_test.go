package main

import (
	"reflect"
	"strings"
	"testing"

	"github.com/vshn/appcat-cli/internal/util"
)

func TestFormatInput(t *testing.T) {
	type cleanInputTest struct {
		arg      []string
		expected []string
		hasError bool
	}

	var cleanInputTests = []cleanInputTest{
		{[]string{"--foo=bar=baz", "--foo", "bar"}, []string{"--foo", "bar=baz", "--foo", "bar"}, false},
		{[]string{"--foo=bar", "--baz=2", "--fooby", "bary"}, []string{"--foo", "bar", "--baz", "2", "--fooby", "bary"}, false},
		{[]string{"--fooby", "bary", "--foo=bar", "--baz=2"}, []string{"--fooby", "bary", "--foo", "bar", "--baz", "2"}, false},
		{[]string{"--foo-", "--bar-", "--foobar-"}, []string{"--foo-", "--bar-", "--foobar-"}, false},
	}

	for _, test := range cleanInputTests {
		t.Run("input: "+strings.Join(test.arg, " "), func(t *testing.T) {
			output := util.FormatInputArguments(test.arg)
			if !reflect.DeepEqual(output, test.expected) {
				t.Errorf("Output %q not equal to expected %q", output, test.expected)
			}
		})
	}
}

func TestCheckForMissingValues(t *testing.T) {
	type cleanInputTest struct {
		arg      []string
		hasError bool
	}

	var cleanInputTests = []cleanInputTest{
		// missing value
		{[]string{"--foo", "--bar"}, true},
		{[]string{"--foo", "--bar-"}, true},
		{[]string{"--foo-", "--baz"}, true},
		// missing key
		{[]string{"foo-", "--baz", "bar"}, true},
		{[]string{"--foo", "foobar", "bar", "--foo"}, true},
		{[]string{"--foo,", "bar", "--barry-", "fooby"}, true},
		{[]string{"--foo", "foobar", "--foo-", "bar"}, true},
		// valid input
		{[]string{"--foo-"}, false},
		{[]string{"--foo", "bar"}, false},
		{[]string{"--foo-", "--baz", "bar", "--foobar-"}, false},
		{[]string{"--foo", "foobar", "--bar", "foo"}, false},
		{[]string{"--foo", "bar", "--baz-", "--barry-", "--fooby", "bary"}, false},
	}

	for _, test := range cleanInputTests {
		t.Run("input: "+strings.Join(test.arg, " "), func(t *testing.T) {
			output := util.CheckForMissingValues(test.arg)
			hasGeneratedError := output != nil
			if hasGeneratedError == false && test.hasError == true {
				t.Errorf("%v\nError should have been thrown for input %s", output, test.arg)
			} else if hasGeneratedError && test.hasError == false {
				t.Errorf("%v\nError should not have been thrown for input %s", output, test.arg)
			}
		})
	}
}

func TestSetFields(t *testing.T) {
	type testStruct struct {
		StringField  string
		Int8Field    int8
		Int16Field   int16
		Int32Field   int32
		Int64Field   int64
		UInt8Field   uint8
		UInt16Field  uint16
		UInt32Field  uint32
		UInt64Field  uint64
		Float32Field float32
		Float64Field float64
		BoolField    bool
		Unsupported  interface{}
	}

	tests := []struct {
		fieldName string
		value     util.Input
		expected  interface{}
		hasErr    bool
	}{
		{"StringField", util.Input{ParameterHierarchy: []string{}, Value: "foo", Unset: false}, "foo", false},
		{"Int8Field", util.Input{ParameterHierarchy: []string{}, Value: "042", Unset: false}, int64(42), false},
		{"Int16Field", util.Input{ParameterHierarchy: []string{}, Value: "042", Unset: false}, int64(42), false},
		{"Int32Field", util.Input{ParameterHierarchy: []string{}, Value: "042", Unset: false}, int64(42), false},
		{"Int64Field", util.Input{ParameterHierarchy: []string{}, Value: "042", Unset: false}, int64(42), false},
		{"UInt8Field", util.Input{ParameterHierarchy: []string{}, Value: "0123", Unset: false}, uint64(123), false},
		{"UInt16Field", util.Input{ParameterHierarchy: []string{}, Value: "0123", Unset: false}, uint64(123), false},
		{"UInt32Field", util.Input{ParameterHierarchy: []string{}, Value: "0123", Unset: false}, uint64(123), false},
		{"UInt64Field", util.Input{ParameterHierarchy: []string{}, Value: "0123", Unset: false}, uint64(123), false},
		{"Float32Field", util.Input{ParameterHierarchy: []string{}, Value: "03.14", Unset: false}, float64(3.14), false},
		{"Float64Field", util.Input{ParameterHierarchy: []string{}, Value: "03.14", Unset: false}, float64(3.14), false},
		{"BoolField", util.Input{ParameterHierarchy: []string{}, Value: "true", Unset: false}, true, false},
		{"Unsupported", util.Input{ParameterHierarchy: []string{}, Value: "bar", Unset: false}, nil, true},

		{"Int64Field", util.Input{ParameterHierarchy: []string{}, Value: "abc", Unset: false}, false, true},
		{"UInt64Field", util.Input{ParameterHierarchy: []string{}, Value: "abc", Unset: false}, false, true},
		{"Float64Field", util.Input{ParameterHierarchy: []string{}, Value: "ab.c", Unset: false}, false, true},
		{"BoolField", util.Input{ParameterHierarchy: []string{}, Value: "abc", Unset: false}, false, true},
	}

	for _, test := range tests {
		t.Run(test.fieldName, func(t *testing.T) {
			// Create a new instance of the test struct
			testFields := testStruct{}

			// Get the field we want to set
			field := reflect.ValueOf(&testFields).Elem().FieldByName(test.fieldName)

			// Call the setFields function with the field and value
			err := util.SetFields(field, test.value)

			hasErrReturn := err != nil
			// Check that error return is correct
			if !hasErrReturn {
				if test.hasErr {
					t.Errorf("Error should have been returned")
				}
				// Check that the field was set as expected
				actual := field.Interface()
				expected := reflect.ValueOf(test.expected).Convert(field.Type()).Interface()
				if !reflect.DeepEqual(actual, expected) {
					t.Errorf("setFields(%s, %q) = %v; want %v", test.fieldName, test.value.Value, actual, expected)
				}
			} else {
				if !test.hasErr {
					t.Errorf("setFields returned an unexpected error: %v", err)
				}
			}

		})
	}
}
