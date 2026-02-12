package keyvalues

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &KvId{}

func TestNewKvID(t *testing.T) {
	id := NewKvID("https://endpoint-url.example.com", "kvName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.KvName != "kvName" {
		t.Fatalf("Expected %q but got %q for Segment 'KvName'", id.KvName, "kvName")
	}
}

func TestFormatKvID(t *testing.T) {
	actual := NewKvID("https://endpoint-url.example.com", "kvName").ID()
	expected := "https://endpoint-url.example.com/kv/kvName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseKvID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *KvId
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
			Input: "https://endpoint-url.example.com/kv",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/kv/kvName",
			Expected: &KvId{
				BaseURI: "https://endpoint-url.example.com",
				KvName:  "kvName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/kv/kvName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseKvID(v.Input)
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

		if actual.KvName != v.Expected.KvName {
			t.Fatalf("Expected %q but got %q for KvName", v.Expected.KvName, actual.KvName)
		}

	}
}

func TestParseKvIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *KvId
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
			Input: "https://endpoint-url.example.com/kv",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/kV",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/kv/kvName",
			Expected: &KvId{
				BaseURI: "https://endpoint-url.example.com",
				KvName:  "kvName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/kv/kvName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/kV/kVnAmE",
			Expected: &KvId{
				BaseURI: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				KvName:  "kVnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/kV/kVnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseKvIDInsensitively(v.Input)
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

		if actual.KvName != v.Expected.KvName {
			t.Fatalf("Expected %q but got %q for KvName", v.Expected.KvName, actual.KvName)
		}

	}
}

func TestSegmentsForKvId(t *testing.T) {
	segments := KvId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("KvId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
