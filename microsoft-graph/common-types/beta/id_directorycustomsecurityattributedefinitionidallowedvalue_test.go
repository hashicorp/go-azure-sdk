package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId{}

func TestNewDirectoryCustomSecurityAttributeDefinitionIdAllowedValueID(t *testing.T) {
	id := NewDirectoryCustomSecurityAttributeDefinitionIdAllowedValueID("customSecurityAttributeDefinitionIdValue", "allowedValueIdValue")

	if id.CustomSecurityAttributeDefinitionId != "customSecurityAttributeDefinitionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CustomSecurityAttributeDefinitionId'", id.CustomSecurityAttributeDefinitionId, "customSecurityAttributeDefinitionIdValue")
	}

	if id.AllowedValueId != "allowedValueIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AllowedValueId'", id.AllowedValueId, "allowedValueIdValue")
	}
}

func TestFormatDirectoryCustomSecurityAttributeDefinitionIdAllowedValueID(t *testing.T) {
	actual := NewDirectoryCustomSecurityAttributeDefinitionIdAllowedValueID("customSecurityAttributeDefinitionIdValue", "allowedValueIdValue").ID()
	expected := "/directory/customSecurityAttributeDefinitions/customSecurityAttributeDefinitionIdValue/allowedValues/allowedValueIdValue"
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
			Input: "/directory/customSecurityAttributeDefinitions/customSecurityAttributeDefinitionIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/customSecurityAttributeDefinitions/customSecurityAttributeDefinitionIdValue/allowedValues",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/customSecurityAttributeDefinitions/customSecurityAttributeDefinitionIdValue/allowedValues/allowedValueIdValue",
			Expected: &DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId{
				CustomSecurityAttributeDefinitionId: "customSecurityAttributeDefinitionIdValue",
				AllowedValueId:                      "allowedValueIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/customSecurityAttributeDefinitions/customSecurityAttributeDefinitionIdValue/allowedValues/allowedValueIdValue/extra",
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
			Input: "/directory/customSecurityAttributeDefinitions/customSecurityAttributeDefinitionIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/cUsToMsEcUrItYaTtRiBuTeDeFiNiTiOnS/cUsToMsEcUrItYaTtRiBuTeDeFiNiTiOnIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/customSecurityAttributeDefinitions/customSecurityAttributeDefinitionIdValue/allowedValues",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/cUsToMsEcUrItYaTtRiBuTeDeFiNiTiOnS/cUsToMsEcUrItYaTtRiBuTeDeFiNiTiOnIdVaLuE/aLlOwEdVaLuEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/customSecurityAttributeDefinitions/customSecurityAttributeDefinitionIdValue/allowedValues/allowedValueIdValue",
			Expected: &DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId{
				CustomSecurityAttributeDefinitionId: "customSecurityAttributeDefinitionIdValue",
				AllowedValueId:                      "allowedValueIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/customSecurityAttributeDefinitions/customSecurityAttributeDefinitionIdValue/allowedValues/allowedValueIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/cUsToMsEcUrItYaTtRiBuTeDeFiNiTiOnS/cUsToMsEcUrItYaTtRiBuTeDeFiNiTiOnIdVaLuE/aLlOwEdVaLuEs/aLlOwEdVaLuEiDvAlUe",
			Expected: &DirectoryCustomSecurityAttributeDefinitionIdAllowedValueId{
				CustomSecurityAttributeDefinitionId: "cUsToMsEcUrItYaTtRiBuTeDeFiNiTiOnIdVaLuE",
				AllowedValueId:                      "aLlOwEdVaLuEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/cUsToMsEcUrItYaTtRiBuTeDeFiNiTiOnS/cUsToMsEcUrItYaTtRiBuTeDeFiNiTiOnIdVaLuE/aLlOwEdVaLuEs/aLlOwEdVaLuEiDvAlUe/extra",
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
