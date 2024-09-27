package standardassignments

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ScopedStandardAssignmentId{}

func TestNewScopedStandardAssignmentID(t *testing.T) {
	id := NewScopedStandardAssignmentID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "standardAssignmentName")

	if id.ResourceId != "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceId'", id.ResourceId, "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")
	}

	if id.StandardAssignmentName != "standardAssignmentName" {
		t.Fatalf("Expected %q but got %q for Segment 'StandardAssignmentName'", id.StandardAssignmentName, "standardAssignmentName")
	}
}

func TestFormatScopedStandardAssignmentID(t *testing.T) {
	actual := NewScopedStandardAssignmentID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "standardAssignmentName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/standardAssignments/standardAssignmentName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseScopedStandardAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedStandardAssignmentId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/standardAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/standardAssignments/standardAssignmentName",
			Expected: &ScopedStandardAssignmentId{
				ResourceId:             "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				StandardAssignmentName: "standardAssignmentName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/standardAssignments/standardAssignmentName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedStandardAssignmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ResourceId != v.Expected.ResourceId {
			t.Fatalf("Expected %q but got %q for ResourceId", v.Expected.ResourceId, actual.ResourceId)
		}

		if actual.StandardAssignmentName != v.Expected.StandardAssignmentName {
			t.Fatalf("Expected %q but got %q for StandardAssignmentName", v.Expected.StandardAssignmentName, actual.StandardAssignmentName)
		}

	}
}

func TestParseScopedStandardAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedStandardAssignmentId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.sEcUrItY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/standardAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.sEcUrItY/sTaNdArDaSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/standardAssignments/standardAssignmentName",
			Expected: &ScopedStandardAssignmentId{
				ResourceId:             "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				StandardAssignmentName: "standardAssignmentName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/standardAssignments/standardAssignmentName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.sEcUrItY/sTaNdArDaSsIgNmEnTs/sTaNdArDaSsIgNmEnTnAmE",
			Expected: &ScopedStandardAssignmentId{
				ResourceId:             "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
				StandardAssignmentName: "sTaNdArDaSsIgNmEnTnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.sEcUrItY/sTaNdArDaSsIgNmEnTs/sTaNdArDaSsIgNmEnTnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedStandardAssignmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ResourceId != v.Expected.ResourceId {
			t.Fatalf("Expected %q but got %q for ResourceId", v.Expected.ResourceId, actual.ResourceId)
		}

		if actual.StandardAssignmentName != v.Expected.StandardAssignmentName {
			t.Fatalf("Expected %q but got %q for StandardAssignmentName", v.Expected.StandardAssignmentName, actual.StandardAssignmentName)
		}

	}
}

func TestSegmentsForScopedStandardAssignmentId(t *testing.T) {
	segments := ScopedStandardAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ScopedStandardAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
