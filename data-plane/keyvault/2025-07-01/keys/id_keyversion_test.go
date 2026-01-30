package keys

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &KeyversionId{}

func TestNewKeyversionID(t *testing.T) {
	id := NewKeyversionID("https://endpoint-url.example.com", "keyName", "keyversion")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.KeyName != "keyName" {
		t.Fatalf("Expected %q but got %q for Segment 'KeyName'", id.KeyName, "keyName")
	}

	if id.Keyversion != "keyversion" {
		t.Fatalf("Expected %q but got %q for Segment 'Keyversion'", id.Keyversion, "keyversion")
	}
}

func TestFormatKeyversionID(t *testing.T) {
	actual := NewKeyversionID("https://endpoint-url.example.com", "keyName", "keyversion").ID()
	expected := "https://endpoint-url.example.com/keys/keyName/keyversion"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseKeyversionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *KeyversionId
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
			Input: "https://endpoint-url.example.com/keys",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/keys/keyName",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/keys/keyName/keyversion",
			Expected: &KeyversionId{
				BaseURI:    "https://endpoint-url.example.com",
				KeyName:    "keyName",
				Keyversion: "keyversion",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/keys/keyName/keyversion/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseKeyversionID(v.Input)
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

		if actual.KeyName != v.Expected.KeyName {
			t.Fatalf("Expected %q but got %q for KeyName", v.Expected.KeyName, actual.KeyName)
		}

		if actual.Keyversion != v.Expected.Keyversion {
			t.Fatalf("Expected %q but got %q for Keyversion", v.Expected.Keyversion, actual.Keyversion)
		}

	}
}

func TestParseKeyversionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *KeyversionId
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
			Input: "https://endpoint-url.example.com/keys",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/kEyS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/keys/keyName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/kEyS/kEyNaMe",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/keys/keyName/keyversion",
			Expected: &KeyversionId{
				BaseURI:    "https://endpoint-url.example.com",
				KeyName:    "keyName",
				Keyversion: "keyversion",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/keys/keyName/keyversion/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/kEyS/kEyNaMe/kEyVeRsIoN",
			Expected: &KeyversionId{
				BaseURI:    "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				KeyName:    "kEyNaMe",
				Keyversion: "kEyVeRsIoN",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/kEyS/kEyNaMe/kEyVeRsIoN/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseKeyversionIDInsensitively(v.Input)
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

		if actual.KeyName != v.Expected.KeyName {
			t.Fatalf("Expected %q but got %q for KeyName", v.Expected.KeyName, actual.KeyName)
		}

		if actual.Keyversion != v.Expected.Keyversion {
			t.Fatalf("Expected %q but got %q for Keyversion", v.Expected.Keyversion, actual.Keyversion)
		}

	}
}

func TestSegmentsForKeyversionId(t *testing.T) {
	segments := KeyversionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("KeyversionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
