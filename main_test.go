package main

import (
	"log"
	"os"
	"os/exec"
	"path"
	"testing"
)

const (
	binaryName = "advent2019"
)

func TestMain(m *testing.M) {
	cmd := exec.Command("make")
	err := cmd.Run()
	if err != nil {
		log.Fatal("could not run make:", err)
	}
	os.Exit(m.Run())
}

func TestIntegration(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command(path.Join(dir, binaryName), "-day", "1")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}

	want := `Day 1: https://adventofcode.com/2019/day/1
Part 1: 3421505
`
	got := string(output)
	if want != string(output) {
		t.Errorf("unexpected output: want %v got %v", want, got)
	}
}
