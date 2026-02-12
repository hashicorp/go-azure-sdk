package policy

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AttestationTypeId{}

func TestNewAttestationTypeID(t *testing.T) {
	id := NewAttestationTypeID("https://endpoint-url.example.com", "OpenEnclave")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.AttestationType != "OpenEnclave" {
		t.Fatalf("Expected %q but got %q for Segment 'AttestationType'", id.AttestationType, "OpenEnclave")
	}
}

func TestFormatAttestationTypeID(t *testing.T) {
	actual := NewAttestationTypeID("https://endpoint-url.example.com", "OpenEnclave").ID()
	expected := "https://endpoint-url.example.com/policies/OpenEnclave"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAttestationTypeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AttestationTypeId
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
			Input: "https://endpoint-url.example.com/policies",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/policies/OpenEnclave",
			Expected: &AttestationTypeId{
				BaseURI:         "https://endpoint-url.example.com",
				AttestationType: "OpenEnclave",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/policies/OpenEnclave/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAttestationTypeID(v.Input)
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

		if actual.AttestationType != v.Expected.AttestationType {
			t.Fatalf("Expected %q but got %q for AttestationType", v.Expected.AttestationType, actual.AttestationType)
		}

	}
}

func TestParseAttestationTypeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AttestationTypeId
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
			Input: "https://endpoint-url.example.com/policies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/pOlIcIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/policies/OpenEnclave",
			Expected: &AttestationTypeId{
				BaseURI:         "https://endpoint-url.example.com",
				AttestationType: "OpenEnclave",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/policies/OpenEnclave/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/pOlIcIeS/oPeNeNcLaVe",
			Expected: &AttestationTypeId{
				BaseURI:         "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				AttestationType: "OpenEnclave",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/pOlIcIeS/oPeNeNcLaVe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAttestationTypeIDInsensitively(v.Input)
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

		if actual.AttestationType != v.Expected.AttestationType {
			t.Fatalf("Expected %q but got %q for AttestationType", v.Expected.AttestationType, actual.AttestationType)
		}

	}
}

func TestSegmentsForAttestationTypeId(t *testing.T) {
	segments := AttestationTypeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AttestationTypeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
