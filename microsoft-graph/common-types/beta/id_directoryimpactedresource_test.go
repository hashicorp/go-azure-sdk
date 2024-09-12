package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryImpactedResourceId{}

func TestNewDirectoryImpactedResourceID(t *testing.T) {
	id := NewDirectoryImpactedResourceID("impactedResourceIdValue")

	if id.ImpactedResourceId != "impactedResourceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ImpactedResourceId'", id.ImpactedResourceId, "impactedResourceIdValue")
	}
}

func TestFormatDirectoryImpactedResourceID(t *testing.T) {
	actual := NewDirectoryImpactedResourceID("impactedResourceIdValue").ID()
	expected := "/directory/impactedResources/impactedResourceIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryImpactedResourceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryImpactedResourceId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/impactedResources",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/impactedResources/impactedResourceIdValue",
			Expected: &DirectoryImpactedResourceId{
				ImpactedResourceId: "impactedResourceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/impactedResources/impactedResourceIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryImpactedResourceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ImpactedResourceId != v.Expected.ImpactedResourceId {
			t.Fatalf("Expected %q but got %q for ImpactedResourceId", v.Expected.ImpactedResourceId, actual.ImpactedResourceId)
		}

	}
}

func TestParseDirectoryImpactedResourceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryImpactedResourceId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/impactedResources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/iMpAcTeDrEsOuRcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/impactedResources/impactedResourceIdValue",
			Expected: &DirectoryImpactedResourceId{
				ImpactedResourceId: "impactedResourceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/impactedResources/impactedResourceIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/iMpAcTeDrEsOuRcEs/iMpAcTeDrEsOuRcEiDvAlUe",
			Expected: &DirectoryImpactedResourceId{
				ImpactedResourceId: "iMpAcTeDrEsOuRcEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/iMpAcTeDrEsOuRcEs/iMpAcTeDrEsOuRcEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryImpactedResourceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ImpactedResourceId != v.Expected.ImpactedResourceId {
			t.Fatalf("Expected %q but got %q for ImpactedResourceId", v.Expected.ImpactedResourceId, actual.ImpactedResourceId)
		}

	}
}

func TestSegmentsForDirectoryImpactedResourceId(t *testing.T) {
	segments := DirectoryImpactedResourceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryImpactedResourceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
