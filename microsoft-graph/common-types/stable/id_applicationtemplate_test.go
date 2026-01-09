package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApplicationTemplateId{}

func TestNewApplicationTemplateID(t *testing.T) {
	id := NewApplicationTemplateID("applicationTemplateId")

	if id.ApplicationTemplateId != "applicationTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'ApplicationTemplateId'", id.ApplicationTemplateId, "applicationTemplateId")
	}
}

func TestFormatApplicationTemplateID(t *testing.T) {
	actual := NewApplicationTemplateID("applicationTemplateId").ID()
	expected := "/applicationTemplates/applicationTemplateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseApplicationTemplateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationTemplateId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applicationTemplates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applicationTemplates/applicationTemplateId",
			Expected: &ApplicationTemplateId{
				ApplicationTemplateId: "applicationTemplateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applicationTemplates/applicationTemplateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationTemplateID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ApplicationTemplateId != v.Expected.ApplicationTemplateId {
			t.Fatalf("Expected %q but got %q for ApplicationTemplateId", v.Expected.ApplicationTemplateId, actual.ApplicationTemplateId)
		}

	}
}

func TestParseApplicationTemplateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationTemplateId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applicationTemplates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnTeMpLaTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applicationTemplates/applicationTemplateId",
			Expected: &ApplicationTemplateId{
				ApplicationTemplateId: "applicationTemplateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applicationTemplates/applicationTemplateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnTeMpLaTeS/aPpLiCaTiOnTeMpLaTeId",
			Expected: &ApplicationTemplateId{
				ApplicationTemplateId: "aPpLiCaTiOnTeMpLaTeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnTeMpLaTeS/aPpLiCaTiOnTeMpLaTeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationTemplateIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ApplicationTemplateId != v.Expected.ApplicationTemplateId {
			t.Fatalf("Expected %q but got %q for ApplicationTemplateId", v.Expected.ApplicationTemplateId, actual.ApplicationTemplateId)
		}

	}
}

func TestSegmentsForApplicationTemplateId(t *testing.T) {
	segments := ApplicationTemplateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ApplicationTemplateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
