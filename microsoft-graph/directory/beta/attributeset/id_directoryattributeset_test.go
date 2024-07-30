package attributeset

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryAttributeSetId{}

func TestNewDirectoryAttributeSetID(t *testing.T) {
	id := NewDirectoryAttributeSetID("attributeSetIdValue")

	if id.AttributeSetId != "attributeSetIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AttributeSetId'", id.AttributeSetId, "attributeSetIdValue")
	}
}

func TestFormatDirectoryAttributeSetID(t *testing.T) {
	actual := NewDirectoryAttributeSetID("attributeSetIdValue").ID()
	expected := "/directory/attributeSets/attributeSetIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryAttributeSetID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryAttributeSetId
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
			Input: "/directory/attributeSets",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/attributeSets/attributeSetIdValue",
			Expected: &DirectoryAttributeSetId{
				AttributeSetId: "attributeSetIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/attributeSets/attributeSetIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryAttributeSetID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AttributeSetId != v.Expected.AttributeSetId {
			t.Fatalf("Expected %q but got %q for AttributeSetId", v.Expected.AttributeSetId, actual.AttributeSetId)
		}

	}
}

func TestParseDirectoryAttributeSetIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryAttributeSetId
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
			Input: "/directory/attributeSets",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/aTtRiBuTeSeTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/attributeSets/attributeSetIdValue",
			Expected: &DirectoryAttributeSetId{
				AttributeSetId: "attributeSetIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/attributeSets/attributeSetIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/aTtRiBuTeSeTs/aTtRiBuTeSeTiDvAlUe",
			Expected: &DirectoryAttributeSetId{
				AttributeSetId: "aTtRiBuTeSeTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/aTtRiBuTeSeTs/aTtRiBuTeSeTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryAttributeSetIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AttributeSetId != v.Expected.AttributeSetId {
			t.Fatalf("Expected %q but got %q for AttributeSetId", v.Expected.AttributeSetId, actual.AttributeSetId)
		}

	}
}

func TestSegmentsForDirectoryAttributeSetId(t *testing.T) {
	segments := DirectoryAttributeSetId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryAttributeSetId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
