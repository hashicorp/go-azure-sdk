package integrationruntimes

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IntegrationRuntimeId{}

func TestNewIntegrationRuntimeID(t *testing.T) {
	id := NewIntegrationRuntimeID("https://endpoint-url.example.com", "integrationRuntimeName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.IntegrationRuntimeName != "integrationRuntimeName" {
		t.Fatalf("Expected %q but got %q for Segment 'IntegrationRuntimeName'", id.IntegrationRuntimeName, "integrationRuntimeName")
	}
}

func TestFormatIntegrationRuntimeID(t *testing.T) {
	actual := NewIntegrationRuntimeID("https://endpoint-url.example.com", "integrationRuntimeName").ID()
	expected := "https://endpoint-url.example.com/integrationRuntimes/integrationRuntimeName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIntegrationRuntimeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IntegrationRuntimeId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/integrationRuntimes",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/integrationRuntimes/integrationRuntimeName",
			Expected: &IntegrationRuntimeId{
				BaseURI:                "https://endpoint-url.example.com",
				IntegrationRuntimeName: "integrationRuntimeName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/integrationRuntimes/integrationRuntimeName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIntegrationRuntimeID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BaseURI != v.Expected.BaseURI {
			t.Fatalf("Expected %q but got %q for BaseURI", v.Expected.BaseURI, actual.BaseURI)
		}

		if actual.IntegrationRuntimeName != v.Expected.IntegrationRuntimeName {
			t.Fatalf("Expected %q but got %q for IntegrationRuntimeName", v.Expected.IntegrationRuntimeName, actual.IntegrationRuntimeName)
		}

	}
}

func TestParseIntegrationRuntimeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IntegrationRuntimeId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/integrationRuntimes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/iNtEgRaTiOnRuNtImEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/integrationRuntimes/integrationRuntimeName",
			Expected: &IntegrationRuntimeId{
				BaseURI:                "https://endpoint-url.example.com",
				IntegrationRuntimeName: "integrationRuntimeName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/integrationRuntimes/integrationRuntimeName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/iNtEgRaTiOnRuNtImEs/iNtEgRaTiOnRuNtImEnAmE",
			Expected: &IntegrationRuntimeId{
				BaseURI:                "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				IntegrationRuntimeName: "iNtEgRaTiOnRuNtImEnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/iNtEgRaTiOnRuNtImEs/iNtEgRaTiOnRuNtImEnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIntegrationRuntimeIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BaseURI != v.Expected.BaseURI {
			t.Fatalf("Expected %q but got %q for BaseURI", v.Expected.BaseURI, actual.BaseURI)
		}

		if actual.IntegrationRuntimeName != v.Expected.IntegrationRuntimeName {
			t.Fatalf("Expected %q but got %q for IntegrationRuntimeName", v.Expected.IntegrationRuntimeName, actual.IntegrationRuntimeName)
		}

	}
}

func TestSegmentsForIntegrationRuntimeId(t *testing.T) {
	segments := IntegrationRuntimeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IntegrationRuntimeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
