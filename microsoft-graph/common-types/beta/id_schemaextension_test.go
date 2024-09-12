package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SchemaExtensionId{}

func TestNewSchemaExtensionID(t *testing.T) {
	id := NewSchemaExtensionID("schemaExtensionIdValue")

	if id.SchemaExtensionId != "schemaExtensionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SchemaExtensionId'", id.SchemaExtensionId, "schemaExtensionIdValue")
	}
}

func TestFormatSchemaExtensionID(t *testing.T) {
	actual := NewSchemaExtensionID("schemaExtensionIdValue").ID()
	expected := "/schemaExtensions/schemaExtensionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSchemaExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SchemaExtensionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/schemaExtensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/schemaExtensions/schemaExtensionIdValue",
			Expected: &SchemaExtensionId{
				SchemaExtensionId: "schemaExtensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/schemaExtensions/schemaExtensionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSchemaExtensionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SchemaExtensionId != v.Expected.SchemaExtensionId {
			t.Fatalf("Expected %q but got %q for SchemaExtensionId", v.Expected.SchemaExtensionId, actual.SchemaExtensionId)
		}

	}
}

func TestParseSchemaExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SchemaExtensionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/schemaExtensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sChEmAeXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/schemaExtensions/schemaExtensionIdValue",
			Expected: &SchemaExtensionId{
				SchemaExtensionId: "schemaExtensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/schemaExtensions/schemaExtensionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sChEmAeXtEnSiOnS/sChEmAeXtEnSiOnIdVaLuE",
			Expected: &SchemaExtensionId{
				SchemaExtensionId: "sChEmAeXtEnSiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sChEmAeXtEnSiOnS/sChEmAeXtEnSiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSchemaExtensionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SchemaExtensionId != v.Expected.SchemaExtensionId {
			t.Fatalf("Expected %q but got %q for SchemaExtensionId", v.Expected.SchemaExtensionId, actual.SchemaExtensionId)
		}

	}
}

func TestSegmentsForSchemaExtensionId(t *testing.T) {
	segments := SchemaExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SchemaExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
