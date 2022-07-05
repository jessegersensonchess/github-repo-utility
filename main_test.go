package main

import "testing"

// to do:
// func TestListGithubReleases(t *testing.T) {
//}
//func TestListGithubPulls(t *testing.T) {
//}

func TestValidateGithubApiEndpoint(t *testing.T) {
	input := "releases"
	expectedOutput := true

	output := ValidateGithubApiEndpoint(input)

	if expectedOutput != output {
		t.Errorf("Failed ! got %v want %v", output, expectedOutput)
	} else {
		t.Logf("Success !")
	}

	input = "pulls"
	expectedOutput = true

	output = ValidateGithubApiEndpoint(input)

	if expectedOutput != output {
		t.Errorf("Failed ! got %v want %v", output, expectedOutput)
	} else {
		t.Logf("Success !")
	}

	input = "invalid"
	expectedOutput = false

	output = ValidateGithubApiEndpoint(input)

	if expectedOutput != output {
		t.Errorf("Failed ! got %v want %v", output, expectedOutput)
	} else {
		t.Logf("Success !")
	}
}
