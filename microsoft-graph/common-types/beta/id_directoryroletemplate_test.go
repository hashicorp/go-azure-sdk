package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryRoleTemplateId{}

func TestNewDirectoryRoleTemplateID(t *testing.T) {
	id := NewDirectoryRoleTemplateID("directoryRoleTemplateId")

	if id.DirectoryRoleTemplateId != "directoryRoleTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryRoleTemplateId'", id.DirectoryRoleTemplateId, "directoryRoleTemplateId")
	}
}

func TestFormatDirectoryRoleTemplateID(t *testing.T) {
	actual := NewDirectoryRoleTemplateID("directoryRoleTemplateId").ID()
	expected := "/directoryRoleTemplates/directoryRoleTemplateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryRoleTemplateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryRoleTemplateId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directoryRoleTemplates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directoryRoleTemplates/directoryRoleTemplateId",
			Expected: &DirectoryRoleTemplateId{
				DirectoryRoleTemplateId: "directoryRoleTemplateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directoryRoleTemplates/directoryRoleTemplateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryRoleTemplateID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DirectoryRoleTemplateId != v.Expected.DirectoryRoleTemplateId {
			t.Fatalf("Expected %q but got %q for DirectoryRoleTemplateId", v.Expected.DirectoryRoleTemplateId, actual.DirectoryRoleTemplateId)
		}

	}
}

func TestParseDirectoryRoleTemplateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryRoleTemplateId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directoryRoleTemplates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRyRoLeTeMpLaTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directoryRoleTemplates/directoryRoleTemplateId",
			Expected: &DirectoryRoleTemplateId{
				DirectoryRoleTemplateId: "directoryRoleTemplateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directoryRoleTemplates/directoryRoleTemplateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRyRoLeTeMpLaTeS/dIrEcToRyRoLeTeMpLaTeId",
			Expected: &DirectoryRoleTemplateId{
				DirectoryRoleTemplateId: "dIrEcToRyRoLeTeMpLaTeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRyRoLeTeMpLaTeS/dIrEcToRyRoLeTeMpLaTeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryRoleTemplateIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DirectoryRoleTemplateId != v.Expected.DirectoryRoleTemplateId {
			t.Fatalf("Expected %q but got %q for DirectoryRoleTemplateId", v.Expected.DirectoryRoleTemplateId, actual.DirectoryRoleTemplateId)
		}

	}
}

func TestSegmentsForDirectoryRoleTemplateId(t *testing.T) {
	segments := DirectoryRoleTemplateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryRoleTemplateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
