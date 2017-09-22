package main

import "testing"
import "os"

func TestMain(m *testing.M) {

	m.Run()

	// Teardown section

	os.Exit(0)
}

func TestAPIInfoEnpoint(t *testing.T) {

}
