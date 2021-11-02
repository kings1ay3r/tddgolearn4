package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	expected := "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})
	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	}

	if got[0] != expected {
		t.Errorf("Incorrect Value Recieved")
	}

}

func TestWalk2(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"Nested fields",
			struct {
				Name    string
				Profile struct {
					Age     int
					City    string
					Address struct {
						zip string
						pin int
					}
				}
			}{
				"Kevin",
				struct {
					Age     int
					City    string
					Address struct {
						zip string
						pin int
					}
				}{
					33,
					"Wayanad",
					struct {
						zip string
						pin int
					}{"xyz", 123},
				},
			},
			[]string{"Kevin", "Wayanad", "xyz"},
		},
		{
			"Pointers to things",
			&struct {
				Name string
				City string
			}{"Chris", "London"},
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

/* //
	{
		"Nested fields",
		struct {
			Name    string
			Profile struct {
				Name string
				Age  int
				City string
			}
		}{
			Name: "Chris",
			Profile: {
				Name: "Chris",
				Age:  33,
				City: "London",
			},
		[]string{"Chris", "London"},
	},
}
*/
