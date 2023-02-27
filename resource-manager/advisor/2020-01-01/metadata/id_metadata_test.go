// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package metadata

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = MetadataId{}

func TestNewMetadataID(t *testing.T) {
	id := NewMetadataID("metadataValue")

	if id.MetadataName != "metadataValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MetadataName'", id.MetadataName, "metadataValue")
	}
}

func TestFormatMetadataID(t *testing.T) {
	actual := NewMetadataID("metadataValue").ID()
	expected := "/providers/Microsoft.Advisor/metadata/metadataValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMetadataID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MetadataId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Advisor",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Advisor/metadata",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Advisor/metadata/metadataValue",
			Expected: &MetadataId{
				MetadataName: "metadataValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Advisor/metadata/metadataValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMetadataID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MetadataName != v.Expected.MetadataName {
			t.Fatalf("Expected %q but got %q for MetadataName", v.Expected.MetadataName, actual.MetadataName)
		}

	}
}

func TestParseMetadataIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MetadataId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Advisor",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aDvIsOr",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Advisor/metadata",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aDvIsOr/mEtAdAtA",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Advisor/metadata/metadataValue",
			Expected: &MetadataId{
				MetadataName: "metadataValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Advisor/metadata/metadataValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aDvIsOr/mEtAdAtA/mEtAdAtAvAlUe",
			Expected: &MetadataId{
				MetadataName: "mEtAdAtAvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aDvIsOr/mEtAdAtA/mEtAdAtAvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMetadataIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MetadataName != v.Expected.MetadataName {
			t.Fatalf("Expected %q but got %q for MetadataName", v.Expected.MetadataName, actual.MetadataName)
		}

	}
}

func TestSegmentsForMetadataId(t *testing.T) {
	segments := MetadataId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MetadataId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
