package registrationassignments

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ScopedRegistrationAssignmentId{}

func TestNewScopedRegistrationAssignmentID(t *testing.T) {
	id := NewScopedRegistrationAssignmentID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "registrationAssignmentIdValue")

	if id.Scope != "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'Scope'", id.Scope, "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")
	}

	if id.RegistrationAssignmentId != "registrationAssignmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RegistrationAssignmentId'", id.RegistrationAssignmentId, "registrationAssignmentIdValue")
	}
}

func TestFormatScopedRegistrationAssignmentID(t *testing.T) {
	actual := NewScopedRegistrationAssignmentID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "registrationAssignmentIdValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.ManagedServices/registrationAssignments/registrationAssignmentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseScopedRegistrationAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedRegistrationAssignmentId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.ManagedServices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.ManagedServices/registrationAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.ManagedServices/registrationAssignments/registrationAssignmentIdValue",
			Expected: &ScopedRegistrationAssignmentId{
				Scope:                    "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				RegistrationAssignmentId: "registrationAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.ManagedServices/registrationAssignments/registrationAssignmentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedRegistrationAssignmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.Scope != v.Expected.Scope {
			t.Fatalf("Expected %q but got %q for Scope", v.Expected.Scope, actual.Scope)
		}

		if actual.RegistrationAssignmentId != v.Expected.RegistrationAssignmentId {
			t.Fatalf("Expected %q but got %q for RegistrationAssignmentId", v.Expected.RegistrationAssignmentId, actual.RegistrationAssignmentId)
		}

	}
}

func TestParseScopedRegistrationAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedRegistrationAssignmentId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.ManagedServices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.mAnAgEdSeRvIcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.ManagedServices/registrationAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.mAnAgEdSeRvIcEs/rEgIsTrAtIoNaSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.ManagedServices/registrationAssignments/registrationAssignmentIdValue",
			Expected: &ScopedRegistrationAssignmentId{
				Scope:                    "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				RegistrationAssignmentId: "registrationAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.ManagedServices/registrationAssignments/registrationAssignmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.mAnAgEdSeRvIcEs/rEgIsTrAtIoNaSsIgNmEnTs/rEgIsTrAtIoNaSsIgNmEnTiDvAlUe",
			Expected: &ScopedRegistrationAssignmentId{
				Scope:                    "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
				RegistrationAssignmentId: "rEgIsTrAtIoNaSsIgNmEnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.mAnAgEdSeRvIcEs/rEgIsTrAtIoNaSsIgNmEnTs/rEgIsTrAtIoNaSsIgNmEnTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedRegistrationAssignmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.Scope != v.Expected.Scope {
			t.Fatalf("Expected %q but got %q for Scope", v.Expected.Scope, actual.Scope)
		}

		if actual.RegistrationAssignmentId != v.Expected.RegistrationAssignmentId {
			t.Fatalf("Expected %q but got %q for RegistrationAssignmentId", v.Expected.RegistrationAssignmentId, actual.RegistrationAssignmentId)
		}

	}
}

func TestSegmentsForScopedRegistrationAssignmentId(t *testing.T) {
	segments := ScopedRegistrationAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ScopedRegistrationAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
