package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeManagedAppLogCollectionRequestId{}

func TestNewMeManagedAppLogCollectionRequestID(t *testing.T) {
	id := NewMeManagedAppLogCollectionRequestID("managedAppLogCollectionRequestId")

	if id.ManagedAppLogCollectionRequestId != "managedAppLogCollectionRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedAppLogCollectionRequestId'", id.ManagedAppLogCollectionRequestId, "managedAppLogCollectionRequestId")
	}
}

func TestFormatMeManagedAppLogCollectionRequestID(t *testing.T) {
	actual := NewMeManagedAppLogCollectionRequestID("managedAppLogCollectionRequestId").ID()
	expected := "/me/managedAppLogCollectionRequests/managedAppLogCollectionRequestId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeManagedAppLogCollectionRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeManagedAppLogCollectionRequestId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedAppLogCollectionRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/managedAppLogCollectionRequests/managedAppLogCollectionRequestId",
			Expected: &MeManagedAppLogCollectionRequestId{
				ManagedAppLogCollectionRequestId: "managedAppLogCollectionRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/managedAppLogCollectionRequests/managedAppLogCollectionRequestId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeManagedAppLogCollectionRequestID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagedAppLogCollectionRequestId != v.Expected.ManagedAppLogCollectionRequestId {
			t.Fatalf("Expected %q but got %q for ManagedAppLogCollectionRequestId", v.Expected.ManagedAppLogCollectionRequestId, actual.ManagedAppLogCollectionRequestId)
		}

	}
}

func TestParseMeManagedAppLogCollectionRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeManagedAppLogCollectionRequestId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedAppLogCollectionRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdApPlOgCoLlEcTiOnReQuEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/managedAppLogCollectionRequests/managedAppLogCollectionRequestId",
			Expected: &MeManagedAppLogCollectionRequestId{
				ManagedAppLogCollectionRequestId: "managedAppLogCollectionRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/managedAppLogCollectionRequests/managedAppLogCollectionRequestId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdApPlOgCoLlEcTiOnReQuEsTs/mAnAgEdApPlOgCoLlEcTiOnReQuEsTiD",
			Expected: &MeManagedAppLogCollectionRequestId{
				ManagedAppLogCollectionRequestId: "mAnAgEdApPlOgCoLlEcTiOnReQuEsTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdApPlOgCoLlEcTiOnReQuEsTs/mAnAgEdApPlOgCoLlEcTiOnReQuEsTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeManagedAppLogCollectionRequestIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagedAppLogCollectionRequestId != v.Expected.ManagedAppLogCollectionRequestId {
			t.Fatalf("Expected %q but got %q for ManagedAppLogCollectionRequestId", v.Expected.ManagedAppLogCollectionRequestId, actual.ManagedAppLogCollectionRequestId)
		}

	}
}

func TestSegmentsForMeManagedAppLogCollectionRequestId(t *testing.T) {
	segments := MeManagedAppLogCollectionRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeManagedAppLogCollectionRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
