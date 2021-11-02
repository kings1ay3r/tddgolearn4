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
						Zip string
						Pin int
					}
				}
			}{
				"Kevin",
				struct {
					Age     int
					City    string
					Address struct {
						Zip string
						Pin int
					}
				}{
					33,
					"Wayanad",
					struct {
						Zip string
						Pin int
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
		{
			"Slices",
			[]struct {
				Age  int
				City string
			}{
				{33, "London"},
				{34, "Tokyo"},
			},
			[]string{"London", "Tokyo"},
		},
		{
			"Arrays",
			[2]string{"Japan", "Mauritius"},
			[]string{"Japan", "Mauritius"},
		},
		{
			"Maps",
			map[string]string{"Country": "Japan", "Capital": "Mauritius"},
			[]string{"Japan", "Mauritius"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if test.Name == "Maps" {
				assertMapsEqual(t, got, test.ExpectedCalls)
			} else {

				if !reflect.DeepEqual(got, test.ExpectedCalls) {
					t.Errorf("got %v, want %v", got, test.ExpectedCalls)
				}
			}
		})
	}
}
func assertMapsEqual(t *testing.T, got, want []string) {

	t.Helper()
	if len(got) != len(want) {
		t.Errorf("Different Length Slices")
		return
	}
	for _, x := range want {
		contains := false
		var y string
		for _, y = range got {
			if y == x {
				contains = true
			}
		}
		if !contains {
			t.Errorf("Expected %v to contain %v", y, want)
			break
		}
	}

}
