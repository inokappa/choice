package main

import (
	"bytes"
	_ "fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestSimpleChoice(t *testing.T) {
	actual := Choice("foo,foo,foo", ",")
	expected := "foo"
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestSpPipeChoice(t *testing.T) {
	actual := Choice("foo|foo|foo", "|")
	expected := "foo"
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestSpPeriodChoice(t *testing.T) {
	actual := Choice("foo.foo.foo", ".")
	expected := "foo"
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestStdoutSimpleChoice(t *testing.T) {
	cmd := exec.Command("sh", "tests/simple_choice.sh")
	stdout := new(bytes.Buffer)
	cmd.Stdout = stdout
	output := "foo"

	_ = cmd.Run()

	if !strings.Contains(stdout.String(), output) {
		t.Fatal("Failed Test")
	}
}

func TestStdoutSpPipeChoice(t *testing.T) {
	cmd := exec.Command("sh", "tests/sp_pipe_choice.sh")
	stdout := new(bytes.Buffer)
	cmd.Stdout = stdout
	output := "foo"

	_ = cmd.Run()

	if !strings.Contains(stdout.String(), output) {
		t.Fatal("Failed Test")
	}
}

func TestStdoutMultilineChoice(t *testing.T) {
	cmd := exec.Command("sh", "tests/multiline_choice.sh")
	stdout := new(bytes.Buffer)
	cmd.Stdout = stdout
	output := "foo"

	_ = cmd.Run()

	if !strings.Contains(stdout.String(), output) {
		t.Fatal("Failed Test")
	}
}
