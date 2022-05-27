package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Fero"},
			[]string{"Fero"},
		},
		{
			"struct with two string field",
			struct {
				Name string
				City string
			}{"Fero", "Zurich"},
			[]string{"Fero", "Zurich"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Fero", 36},
			[]string{"Fero"},
		},
		{
			"nested fields",
			Person{
				"Fero",
				Profile{36, "Zurich"},
			},
			[]string{"Fero", "Zurich"},
		},
		{
			"pointer to things",
			&Person{
				"Chris",
				Profile{22, "London"},
			},
			[]string{"Chris", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
