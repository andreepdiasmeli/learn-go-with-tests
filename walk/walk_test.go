package walk

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Profile Profile
}

type Profile struct {
	Age int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct{
		Name string
		Input interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct{
				Name string
			}{"André"},
			[]string{"André"},
		},
		{
			"struct with two string fields",
			struct{
				Name string
				City string
			}{"André", "Joinville"},
			[]string{"André", "Joinville"},
		},
		{
			"struct with non string field",
			struct{
				Name string
				Age int
			}{"André", 24},
			[]string{"André"},

		},
		{
			"nested fields",
			Person{
				"André",
				Profile{24, "Joinville"},
			},
			[]string{"André", "Joinville"},
		},
		{
			"pointer to things",
			&Person{
				"André",
				Profile{24, "Joinville"},
			},
			[]string{"André", "Joinville"},
		},
		{
			"slices",
			[]Profile{
				{24, "André"},
				{14, "Eduardo"},
			},
			[]string{"André", "Eduardo"},
		},
		{
			"arrays",
			[2]Profile{
				{24, "Pacheco"},
				{14, "Dias"},
			},
			[]string{"Pacheco", "Dias"},
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
	
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"a": "André",
			"d": "Dias",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "André")
		assertContains(t, got, "Dias")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func(){
			aChannel <- Profile{24, "Joinville"}
			aChannel <- Profile{14, "Florianópolis"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Joinville", "Florianópolis"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{24, "Joinville"}, Profile{14, "Florianópolis"}
		}

		var got []string
		want := []string{"Joinville", "Florianópolis"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}


	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expepcted %+v to contain %q but it did not", haystack, needle)
	}
}