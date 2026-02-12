package iotcentrals

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &OrganizationId{}

func TestNewOrganizationID(t *testing.T) {
	id := NewOrganizationID("https://endpoint-url.example.com", "organizationId")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.OrganizationId != "organizationId" {
		t.Fatalf("Expected %q but got %q for Segment 'OrganizationId'", id.OrganizationId, "organizationId")
	}
}

func TestFormatOrganizationID(t *testing.T) {
	actual := NewOrganizationID("https://endpoint-url.example.com", "organizationId").ID()
	expected := "https://endpoint-url.example.com/organizations/organizationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseOrganizationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *OrganizationId
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
			Input: "https://endpoint-url.example.com/organizations",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/organizations/organizationId",
			Expected: &OrganizationId{
				BaseURI:        "https://endpoint-url.example.com",
				OrganizationId: "organizationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/organizations/organizationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseOrganizationID(v.Input)
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

		if actual.OrganizationId != v.Expected.OrganizationId {
			t.Fatalf("Expected %q but got %q for OrganizationId", v.Expected.OrganizationId, actual.OrganizationId)
		}

	}
}

func TestParseOrganizationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *OrganizationId
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
			Input: "https://endpoint-url.example.com/organizations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/oRgAnIzAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/organizations/organizationId",
			Expected: &OrganizationId{
				BaseURI:        "https://endpoint-url.example.com",
				OrganizationId: "organizationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/organizations/organizationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/oRgAnIzAtIoNs/oRgAnIzAtIoNiD",
			Expected: &OrganizationId{
				BaseURI:        "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				OrganizationId: "oRgAnIzAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/oRgAnIzAtIoNs/oRgAnIzAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseOrganizationIDInsensitively(v.Input)
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

		if actual.OrganizationId != v.Expected.OrganizationId {
			t.Fatalf("Expected %q but got %q for OrganizationId", v.Expected.OrganizationId, actual.OrganizationId)
		}

	}
}

func TestSegmentsForOrganizationId(t *testing.T) {
	segments := OrganizationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("OrganizationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
