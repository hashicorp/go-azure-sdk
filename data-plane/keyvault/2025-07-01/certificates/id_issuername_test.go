package certificates

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IssuernameId{}

func TestNewIssuernameID(t *testing.T) {
	id := NewIssuernameID("https://endpoint-url.example.com", "issuerName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.IssuerName != "issuerName" {
		t.Fatalf("Expected %q but got %q for Segment 'IssuerName'", id.IssuerName, "issuerName")
	}
}

func TestFormatIssuernameID(t *testing.T) {
	actual := NewIssuernameID("https://endpoint-url.example.com", "issuerName").ID()
	expected := "https://endpoint-url.example.com/certificates/issuers/issuerName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIssuernameID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IssuernameId
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
			Input: "https://endpoint-url.example.com/certificates",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/certificates/issuers",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/certificates/issuers/issuerName",
			Expected: &IssuernameId{
				BaseURI:    "https://endpoint-url.example.com",
				IssuerName: "issuerName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/certificates/issuers/issuerName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIssuernameID(v.Input)
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

		if actual.IssuerName != v.Expected.IssuerName {
			t.Fatalf("Expected %q but got %q for IssuerName", v.Expected.IssuerName, actual.IssuerName)
		}

	}
}

func TestParseIssuernameIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IssuernameId
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
			Input: "https://endpoint-url.example.com/certificates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/cErTiFiCaTeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/certificates/issuers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/cErTiFiCaTeS/iSsUeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/certificates/issuers/issuerName",
			Expected: &IssuernameId{
				BaseURI:    "https://endpoint-url.example.com",
				IssuerName: "issuerName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/certificates/issuers/issuerName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/cErTiFiCaTeS/iSsUeRs/iSsUeRnAmE",
			Expected: &IssuernameId{
				BaseURI:    "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				IssuerName: "iSsUeRnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/cErTiFiCaTeS/iSsUeRs/iSsUeRnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIssuernameIDInsensitively(v.Input)
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

		if actual.IssuerName != v.Expected.IssuerName {
			t.Fatalf("Expected %q but got %q for IssuerName", v.Expected.IssuerName, actual.IssuerName)
		}

	}
}

func TestSegmentsForIssuernameId(t *testing.T) {
	segments := IssuernameId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IssuernameId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
