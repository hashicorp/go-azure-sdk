package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryCustomSecurityAttributeDefinitionId{}

func TestNewDirectoryCustomSecurityAttributeDefinitionID(t *testing.T) {
	id := NewDirectoryCustomSecurityAttributeDefinitionID("customSecurityAttributeDefinitionId")

	if id.CustomSecurityAttributeDefinitionId != "customSecurityAttributeDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'CustomSecurityAttributeDefinitionId'", id.CustomSecurityAttributeDefinitionId, "customSecurityAttributeDefinitionId")
	}
}

func TestFormatDirectoryCustomSecurityAttributeDefinitionID(t *testing.T) {
	actual := NewDirectoryCustomSecurityAttributeDefinitionID("customSecurityAttributeDefinitionId").ID()
	expected := "/directory/customSecurityAttributeDefinitions/customSecurityAttributeDefinitionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryCustomSecurityAttributeDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryCustomSecurityAttributeDefinitionId
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
			Input: "/directory/customSecurityAttributeDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/customSecurityAttributeDefinitions/customSecurityAttributeDefinitionId",
			Expected: &DirectoryCustomSecurityAttributeDefinitionId{
				CustomSecurityAttributeDefinitionId: "customSecurityAttributeDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/customSecurityAttributeDefinitions/customSecurityAttributeDefinitionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryCustomSecurityAttributeDefinitionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CustomSecurityAttributeDefinitionId != v.Expected.CustomSecurityAttributeDefinitionId {
			t.Fatalf("Expected %q but got %q for CustomSecurityAttributeDefinitionId", v.Expected.CustomSecurityAttributeDefinitionId, actual.CustomSecurityAttributeDefinitionId)
		}

	}
}

func TestParseDirectoryCustomSecurityAttributeDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryCustomSecurityAttributeDefinitionId
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
			Input: "/directory/customSecurityAttributeDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/cUsToMsEcUrItYaTtRiBuTeDeFiNiTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/customSecurityAttributeDefinitions/customSecurityAttributeDefinitionId",
			Expected: &DirectoryCustomSecurityAttributeDefinitionId{
				CustomSecurityAttributeDefinitionId: "customSecurityAttributeDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/customSecurityAttributeDefinitions/customSecurityAttributeDefinitionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/cUsToMsEcUrItYaTtRiBuTeDeFiNiTiOnS/cUsToMsEcUrItYaTtRiBuTeDeFiNiTiOnId",
			Expected: &DirectoryCustomSecurityAttributeDefinitionId{
				CustomSecurityAttributeDefinitionId: "cUsToMsEcUrItYaTtRiBuTeDeFiNiTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/cUsToMsEcUrItYaTtRiBuTeDeFiNiTiOnS/cUsToMsEcUrItYaTtRiBuTeDeFiNiTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryCustomSecurityAttributeDefinitionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CustomSecurityAttributeDefinitionId != v.Expected.CustomSecurityAttributeDefinitionId {
			t.Fatalf("Expected %q but got %q for CustomSecurityAttributeDefinitionId", v.Expected.CustomSecurityAttributeDefinitionId, actual.CustomSecurityAttributeDefinitionId)
		}

	}
}

func TestSegmentsForDirectoryCustomSecurityAttributeDefinitionId(t *testing.T) {
	segments := DirectoryCustomSecurityAttributeDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryCustomSecurityAttributeDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
