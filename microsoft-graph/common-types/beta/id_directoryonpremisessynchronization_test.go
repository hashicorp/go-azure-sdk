package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryOnPremisesSynchronizationId{}

func TestNewDirectoryOnPremisesSynchronizationID(t *testing.T) {
	id := NewDirectoryOnPremisesSynchronizationID("onPremisesDirectorySynchronizationIdValue")

	if id.OnPremisesDirectorySynchronizationId != "onPremisesDirectorySynchronizationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OnPremisesDirectorySynchronizationId'", id.OnPremisesDirectorySynchronizationId, "onPremisesDirectorySynchronizationIdValue")
	}
}

func TestFormatDirectoryOnPremisesSynchronizationID(t *testing.T) {
	actual := NewDirectoryOnPremisesSynchronizationID("onPremisesDirectorySynchronizationIdValue").ID()
	expected := "/directory/onPremisesSynchronization/onPremisesDirectorySynchronizationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryOnPremisesSynchronizationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryOnPremisesSynchronizationId
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
			Input: "/directory/onPremisesSynchronization",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/onPremisesSynchronization/onPremisesDirectorySynchronizationIdValue",
			Expected: &DirectoryOnPremisesSynchronizationId{
				OnPremisesDirectorySynchronizationId: "onPremisesDirectorySynchronizationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/onPremisesSynchronization/onPremisesDirectorySynchronizationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryOnPremisesSynchronizationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OnPremisesDirectorySynchronizationId != v.Expected.OnPremisesDirectorySynchronizationId {
			t.Fatalf("Expected %q but got %q for OnPremisesDirectorySynchronizationId", v.Expected.OnPremisesDirectorySynchronizationId, actual.OnPremisesDirectorySynchronizationId)
		}

	}
}

func TestParseDirectoryOnPremisesSynchronizationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryOnPremisesSynchronizationId
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
			Input: "/directory/onPremisesSynchronization",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/oNpReMiSeSsYnChRoNiZaTiOn",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/onPremisesSynchronization/onPremisesDirectorySynchronizationIdValue",
			Expected: &DirectoryOnPremisesSynchronizationId{
				OnPremisesDirectorySynchronizationId: "onPremisesDirectorySynchronizationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/onPremisesSynchronization/onPremisesDirectorySynchronizationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/oNpReMiSeSsYnChRoNiZaTiOn/oNpReMiSeSdIrEcToRySyNcHrOnIzAtIoNiDvAlUe",
			Expected: &DirectoryOnPremisesSynchronizationId{
				OnPremisesDirectorySynchronizationId: "oNpReMiSeSdIrEcToRySyNcHrOnIzAtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/oNpReMiSeSsYnChRoNiZaTiOn/oNpReMiSeSdIrEcToRySyNcHrOnIzAtIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryOnPremisesSynchronizationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OnPremisesDirectorySynchronizationId != v.Expected.OnPremisesDirectorySynchronizationId {
			t.Fatalf("Expected %q but got %q for OnPremisesDirectorySynchronizationId", v.Expected.OnPremisesDirectorySynchronizationId, actual.OnPremisesDirectorySynchronizationId)
		}

	}
}

func TestSegmentsForDirectoryOnPremisesSynchronizationId(t *testing.T) {
	segments := DirectoryOnPremisesSynchronizationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryOnPremisesSynchronizationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
