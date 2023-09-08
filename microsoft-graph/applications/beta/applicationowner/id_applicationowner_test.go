package applicationowner

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ApplicationOwnerId{}

func TestNewApplicationOwnerID(t *testing.T) {
	id := NewApplicationOwnerID("applicationIdValue", "directoryObjectIdValue")

	if id.ApplicationId != "applicationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ApplicationId'", id.ApplicationId, "applicationIdValue")
	}

	if id.DirectoryObjectId != "directoryObjectIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectIdValue")
	}
}

func TestFormatApplicationOwnerID(t *testing.T) {
	actual := NewApplicationOwnerID("applicationIdValue", "directoryObjectIdValue").ID()
	expected := "/applications/applicationIdValue/owners/directoryObjectIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseApplicationOwnerID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationOwnerId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue/owners",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationIdValue/owners/directoryObjectIdValue",
			Expected: &ApplicationOwnerId{
				ApplicationId:     "applicationIdValue",
				DirectoryObjectId: "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationIdValue/owners/directoryObjectIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationOwnerID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ApplicationId != v.Expected.ApplicationId {
			t.Fatalf("Expected %q but got %q for ApplicationId", v.Expected.ApplicationId, actual.ApplicationId)
		}

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestParseApplicationOwnerIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationOwnerId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue/owners",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/oWnErS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationIdValue/owners/directoryObjectIdValue",
			Expected: &ApplicationOwnerId{
				ApplicationId:     "applicationIdValue",
				DirectoryObjectId: "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationIdValue/owners/directoryObjectIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/oWnErS/dIrEcToRyObJeCtIdVaLuE",
			Expected: &ApplicationOwnerId{
				ApplicationId:     "aPpLiCaTiOnIdVaLuE",
				DirectoryObjectId: "dIrEcToRyObJeCtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/oWnErS/dIrEcToRyObJeCtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationOwnerIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ApplicationId != v.Expected.ApplicationId {
			t.Fatalf("Expected %q but got %q for ApplicationId", v.Expected.ApplicationId, actual.ApplicationId)
		}

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestSegmentsForApplicationOwnerId(t *testing.T) {
	segments := ApplicationOwnerId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ApplicationOwnerId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
