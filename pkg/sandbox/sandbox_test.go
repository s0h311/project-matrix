package sandbox

import "testing"

func TestLs(t *testing.T) {
	firstList, err := Ls()

	if err != nil {
		t.Fatal(err)
	}

	newSandbox := CreateOptions{
		Name:       "new-shiny-sandbox",
		Agent:      "claude",
		Workspaces: []string{"."},
	}

	for _, v := range firstList {
		assertNotEqual(t, v.Name, newSandbox.Name)
	}

	err = Create(&newSandbox)

	if err != nil {
		t.Fatal(err)
	}

	secondList, err := Ls()

	if err != nil {
		t.Fatal(err)
	}

	found := false

	for _, v := range secondList {
		if v.Name == newSandbox.Name {
			found = true
		}
	}

	assertTrue(t, found)

	err = Rm(&[]string{newSandbox.Name})

	if err != nil {
		t.Fatal(err)
	}

	thirdList, err := Ls()

	if err != nil {
		t.Fatal(err)
	}

	for _, v := range thirdList {
		assertNotEqual(t, v.Name, newSandbox.Name)
	}
}

func assertEqual[T comparable](t *testing.T, actual T, expected T) {
	if actual != expected {
		t.Errorf("Actual was: %+v, Expected: %+v", actual, expected)
	}
}

func assertNotEqual[T comparable](t *testing.T, actual T, expected T) {
	if actual == expected {
		t.Errorf("Actual was: %+v, Expected: %+v", actual, expected)
	}
}

func assertTrue(t *testing.T, actual bool) {
	if actual != true {
		t.Errorf("Actual was: %+v, Expected: true", actual)
	}
}
