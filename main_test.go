package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestLoadCodes(t *testing.T) {
	codes, err := loadCodes()
	if err != nil {
		t.Fatalf("loadCodes() failed: %v", err)
	}

	if len(codes) == 0 {
		t.Fatal("loadCodes() returned empty map")
	}

	// Test some common status codes
	testCases := []struct {
		code    string
		wantMsg string
	}{
		{"200", "OK"},
		{"404", "Not Found"},
		{"500", "Internal Server Error"},
	}

	for _, tc := range testCases {
		t.Run("code_"+tc.code, func(t *testing.T) {
			sc, found := codes[tc.code]
			if !found {
				t.Errorf("Status code %s not found in codes map", tc.code)
				return
			}
			if sc.Message != tc.wantMsg {
				t.Errorf("Status code %s: got message %q, want %q", tc.code, sc.Message, tc.wantMsg)
			}
		})
	}
}

func TestLoadCodesStructure(t *testing.T) {
	codes, err := loadCodes()
	if err != nil {
		t.Fatalf("loadCodes() failed: %v", err)
	}

	// Verify each status code has required fields
	for code, sc := range codes {
		if sc.Message == "" {
			t.Errorf("Status code %s has empty Message field", code)
		}
		if sc.Description == "" {
			t.Errorf("Status code %s has empty Description field", code)
		}
		if sc.Code == nil {
			t.Errorf("Status code %s has nil Code field", code)
		}
	}
}

func TestDisplayStatusCode(t *testing.T) {
	codes, err := loadCodes()
	if err != nil {
		t.Fatalf("loadCodes() failed: %v", err)
	}

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Test displaying a status code
	if sc, found := codes["200"]; found {
		displayStatusCode(sc)
	} else {
		t.Fatal("Status code 200 not found")
	}

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Read captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Verify output contains expected content
	if !strings.Contains(output, "200") {
		t.Error("Output should contain status code 200")
	}
	if !strings.Contains(output, "OK") {
		t.Error("Output should contain status message 'OK'")
	}
	if !strings.Contains(output, "Source:") {
		t.Error("Output should contain source attribution")
	}
}

func TestVersionConstant(t *testing.T) {
	if version == "" {
		t.Error("version constant should not be empty")
	}
}

func TestShowVersion(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	showVersion()

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Read captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Verify output contains version
	if !strings.Contains(output, "httpcodes version") {
		t.Error("showVersion() output should contain 'httpcodes version'")
	}
	if !strings.Contains(output, version) {
		t.Errorf("showVersion() output should contain version %q", version)
	}
}

func TestShowUsage(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	showUsage()

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Read captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Verify output contains expected content
	expectedStrings := []string{
		"Usage:",
		"httpcodes",
		"Examples:",
		"--version",
		"--help",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(output, expected) {
			t.Errorf("showUsage() output should contain %q", expected)
		}
	}
}
