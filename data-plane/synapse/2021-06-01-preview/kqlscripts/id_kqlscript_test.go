package kqlscripts

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &KqlScriptId{}

func TestNewKqlScriptID(t *testing.T) {
	id := NewKqlScriptID("https://endpoint-url.example.com", "kqlScriptName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.KqlScriptName != "kqlScriptName" {
		t.Fatalf("Expected %q but got %q for Segment 'KqlScriptName'", id.KqlScriptName, "kqlScriptName")
	}
}

func TestFormatKqlScriptID(t *testing.T) {
	actual := NewKqlScriptID("https://endpoint-url.example.com", "kqlScriptName").ID()
	expected := "https://endpoint-url.example.com/kqlScripts/kqlScriptName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseKqlScriptID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *KqlScriptId
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
			Input: "https://endpoint-url.example.com/kqlScripts",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/kqlScripts/kqlScriptName",
			Expected: &KqlScriptId{
				BaseURI:       "https://endpoint-url.example.com",
				KqlScriptName: "kqlScriptName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/kqlScripts/kqlScriptName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseKqlScriptID(v.Input)
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

		if actual.KqlScriptName != v.Expected.KqlScriptName {
			t.Fatalf("Expected %q but got %q for KqlScriptName", v.Expected.KqlScriptName, actual.KqlScriptName)
		}

	}
}

func TestParseKqlScriptIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *KqlScriptId
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
			Input: "https://endpoint-url.example.com/kqlScripts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/kQlScRiPtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/kqlScripts/kqlScriptName",
			Expected: &KqlScriptId{
				BaseURI:       "https://endpoint-url.example.com",
				KqlScriptName: "kqlScriptName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/kqlScripts/kqlScriptName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/kQlScRiPtS/kQlScRiPtNaMe",
			Expected: &KqlScriptId{
				BaseURI:       "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				KqlScriptName: "kQlScRiPtNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/kQlScRiPtS/kQlScRiPtNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseKqlScriptIDInsensitively(v.Input)
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

		if actual.KqlScriptName != v.Expected.KqlScriptName {
			t.Fatalf("Expected %q but got %q for KqlScriptName", v.Expected.KqlScriptName, actual.KqlScriptName)
		}

	}
}

func TestSegmentsForKqlScriptId(t *testing.T) {
	segments := KqlScriptId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("KqlScriptId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
