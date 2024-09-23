package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementExchangeTransitiveRoleAssignmentId{}

func TestNewRoleManagementExchangeTransitiveRoleAssignmentID(t *testing.T) {
	id := NewRoleManagementExchangeTransitiveRoleAssignmentID("unifiedRoleAssignmentId")

	if id.UnifiedRoleAssignmentId != "unifiedRoleAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'UnifiedRoleAssignmentId'", id.UnifiedRoleAssignmentId, "unifiedRoleAssignmentId")
	}
}

func TestFormatRoleManagementExchangeTransitiveRoleAssignmentID(t *testing.T) {
	actual := NewRoleManagementExchangeTransitiveRoleAssignmentID("unifiedRoleAssignmentId").ID()
	expected := "/roleManagement/exchange/transitiveRoleAssignments/unifiedRoleAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleManagementExchangeTransitiveRoleAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementExchangeTransitiveRoleAssignmentId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/exchange",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/exchange/transitiveRoleAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/exchange/transitiveRoleAssignments/unifiedRoleAssignmentId",
			Expected: &RoleManagementExchangeTransitiveRoleAssignmentId{
				UnifiedRoleAssignmentId: "unifiedRoleAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/exchange/transitiveRoleAssignments/unifiedRoleAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementExchangeTransitiveRoleAssignmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleAssignmentId != v.Expected.UnifiedRoleAssignmentId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleAssignmentId", v.Expected.UnifiedRoleAssignmentId, actual.UnifiedRoleAssignmentId)
		}

	}
}

func TestParseRoleManagementExchangeTransitiveRoleAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleManagementExchangeTransitiveRoleAssignmentId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/exchange",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eXcHaNgE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/roleManagement/exchange/transitiveRoleAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eXcHaNgE/tRaNsItIvErOlEaSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/roleManagement/exchange/transitiveRoleAssignments/unifiedRoleAssignmentId",
			Expected: &RoleManagementExchangeTransitiveRoleAssignmentId{
				UnifiedRoleAssignmentId: "unifiedRoleAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/roleManagement/exchange/transitiveRoleAssignments/unifiedRoleAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eXcHaNgE/tRaNsItIvErOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtId",
			Expected: &RoleManagementExchangeTransitiveRoleAssignmentId{
				UnifiedRoleAssignmentId: "uNiFiEdRoLeAsSiGnMeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rOlEmAnAgEmEnT/eXcHaNgE/tRaNsItIvErOlEaSsIgNmEnTs/uNiFiEdRoLeAsSiGnMeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleManagementExchangeTransitiveRoleAssignmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UnifiedRoleAssignmentId != v.Expected.UnifiedRoleAssignmentId {
			t.Fatalf("Expected %q but got %q for UnifiedRoleAssignmentId", v.Expected.UnifiedRoleAssignmentId, actual.UnifiedRoleAssignmentId)
		}

	}
}

func TestSegmentsForRoleManagementExchangeTransitiveRoleAssignmentId(t *testing.T) {
	segments := RoleManagementExchangeTransitiveRoleAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleManagementExchangeTransitiveRoleAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
