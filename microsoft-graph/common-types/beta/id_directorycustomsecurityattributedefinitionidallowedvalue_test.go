package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId{}

func TestNewDirectoryCustomSecurityAttributeDefinitionIdAllowedValueID(t *testing.T) {
	id := NewDirectoryCustomSecurityAttributeDefinitionIdAllowedValueID("customSecurityAttributeDefinitionId", "allowedValueId")

	if id.CustomSecurityAttributeDefinitionId != "customSecurityAttributeDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'CustomSecurityAttributeDefinitionId'", id.CustomSecurityAttributeDefinitionId, "customSecurityAttributeDefinitionId")
	}

	if id.AllowedValueId != "allowedValueId" {
		t.Fatalf("Expected %q but got %q for Segment 'AllowedValueId'", id.AllowedValueId, "allowedValueId")
	}
}

func TestFormatDirectoryCustomSecurityAttributeDefinitionIdAllowedValueID(t *testing.T) {
	actual := NewDirectoryCustomSecurityAttributeDefinitionIdAllowedValueID("customSecurityAttributeDefinitionId", "allowedValueId").ID()
	expected := "/directory/customSecurityAttributeDefinitions/customSecurityAttributeDefinitionId/allowedValues/allowedValueId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryCustomSecurityAttributeDefinitionIdAllowedValueID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId
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
			// Incomplete URI
			Input: "/directory/customSecurityAttributeDefinitions/customSecurityAttributeDefinitionId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/customSecurityAttributeDefinitions/customSecurityAttributeDefinitionId/allowedValues",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/customSecurityAttributeDefinitions/customSecurityAttributeDefinitionId/allowedValues/allowedValueId",
			Expected: &DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId{
				CustomSecurityAttributeDefinitionId: "customSecurityAttributeDefinitionId",
				AllowedValueId:                      "allowedValueId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/customSecurityAttributeDefinitions/customSecurityAttributeDefinitionId/allowedValues/allowedValueId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryCustomSecurityAttributeDefinitionIdAllowedValueID(v.Input)
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

		if actual.AllowedValueId != v.Expected.AllowedValueId {
			t.Fatalf("Expected %q but got %q for AllowedValueId", v.Expected.AllowedValueId, actual.AllowedValueId)
		}

	}
}

func TestParseDirectoryCustomSecurityAttributeDefinitionIdAllowedValueIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId
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
			// Incomplete URI
			Input: "/directory/customSecurityAttributeDefinitions/customSecurityAttributeDefinitionId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/cUsToMsEcUrItYaTtRiBuTeDeFiNiTiOnS/cUsToMsEcUrItYaTtRiBuTeDeFiNiTiOnId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/customSecurityAttributeDefinitions/customSecurityAttributeDefinitionId/allowedValues",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/cUsToMsEcUrItYaTtRiBuTeDeFiNiTiOnS/cUsToMsEcUrItYaTtRiBuTeDeFiNiTiOnId/aLlOwEdVaLuEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/customSecurityAttributeDefinitions/customSecurityAttributeDefinitionId/allowedValues/allowedValueId",
			Expected: &DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId{
				CustomSecurityAttributeDefinitionId: "customSecurityAttributeDefinitionId",
				AllowedValueId:                      "allowedValueId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/customSecurityAttributeDefinitions/customSecurityAttributeDefinitionId/allowedValues/allowedValueId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/cUsToMsEcUrItYaTtRiBuTeDeFiNiTiOnS/cUsToMsEcUrItYaTtRiBuTeDeFiNiTiOnId/aLlOwEdVaLuEs/aLlOwEdVaLuEiD",
			Expected: &DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId{
				CustomSecurityAttributeDefinitionId: "cUsToMsEcUrItYaTtRiBuTeDeFiNiTiOnId",
				AllowedValueId:                      "aLlOwEdVaLuEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/cUsToMsEcUrItYaTtRiBuTeDeFiNiTiOnS/cUsToMsEcUrItYaTtRiBuTeDeFiNiTiOnId/aLlOwEdVaLuEs/aLlOwEdVaLuEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryCustomSecurityAttributeDefinitionIdAllowedValueIDInsensitively(v.Input)
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

		if actual.AllowedValueId != v.Expected.AllowedValueId {
			t.Fatalf("Expected %q but got %q for AllowedValueId", v.Expected.AllowedValueId, actual.AllowedValueId)
		}

	}
}

func TestSegmentsForDirectoryCustomSecurityAttributeDefinitionIdAllowedValueId(t *testing.T) {
	segments := DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
