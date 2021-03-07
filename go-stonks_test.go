package main

import "testing"

func TestGetQuote(t *testing.T) {
	// Test for AAPL
	appleResult := getQuote("AAPL")
	if appleResult.Name != "Apple Inc" {
		t.Errorf("getQuote failed, expected %v for Name field, got %v", "Apple Inc", appleResult.Name)
	}
}
