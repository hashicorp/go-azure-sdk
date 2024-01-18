package roleeligibilityscheduleinstances

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ScopedRoleEligibilityScheduleInstanceId{}

func TestNewScopedRoleEligibilityScheduleInstanceID(t *testing.T) {
	id := NewScopedRoleEligibilityScheduleInstanceID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "roleEligibilityScheduleInstanceValue")

	if id.Scope != "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'Scope'", id.Scope, "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")
	}

	if id.RoleEligibilityScheduleInstanceName != "roleEligibilityScheduleInstanceValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RoleEligibilityScheduleInstanceName'", id.RoleEligibilityScheduleInstanceName, "roleEligibilityScheduleInstanceValue")
	}
}

func TestFormatScopedRoleEligibilityScheduleInstanceID(t *testing.T) {
	actual := NewScopedRoleEligibilityScheduleInstanceID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "roleEligibilityScheduleInstanceValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleEligibilityScheduleInstances/roleEligibilityScheduleInstanceValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseScopedRoleEligibilityScheduleInstanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedRoleEligibilityScheduleInstanceId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleEligibilityScheduleInstances",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleEligibilityScheduleInstances/roleEligibilityScheduleInstanceValue",
			Expected: &ScopedRoleEligibilityScheduleInstanceId{
				Scope:                               "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				RoleEligibilityScheduleInstanceName: "roleEligibilityScheduleInstanceValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleEligibilityScheduleInstances/roleEligibilityScheduleInstanceValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedRoleEligibilityScheduleInstanceID(v.Input)
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

		if actual.RoleEligibilityScheduleInstanceName != v.Expected.RoleEligibilityScheduleInstanceName {
			t.Fatalf("Expected %q but got %q for RoleEligibilityScheduleInstanceName", v.Expected.RoleEligibilityScheduleInstanceName, actual.RoleEligibilityScheduleInstanceName)
		}

	}
}

func TestParseScopedRoleEligibilityScheduleInstanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedRoleEligibilityScheduleInstanceId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleEligibilityScheduleInstances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/rOlEeLiGiBiLiTyScHeDuLeInStAnCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleEligibilityScheduleInstances/roleEligibilityScheduleInstanceValue",
			Expected: &ScopedRoleEligibilityScheduleInstanceId{
				Scope:                               "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				RoleEligibilityScheduleInstanceName: "roleEligibilityScheduleInstanceValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleEligibilityScheduleInstances/roleEligibilityScheduleInstanceValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/rOlEeLiGiBiLiTyScHeDuLeInStAnCeS/rOlEeLiGiBiLiTyScHeDuLeInStAnCeVaLuE",
			Expected: &ScopedRoleEligibilityScheduleInstanceId{
				Scope:                               "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
				RoleEligibilityScheduleInstanceName: "rOlEeLiGiBiLiTyScHeDuLeInStAnCeVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/rOlEeLiGiBiLiTyScHeDuLeInStAnCeS/rOlEeLiGiBiLiTyScHeDuLeInStAnCeVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedRoleEligibilityScheduleInstanceIDInsensitively(v.Input)
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

		if actual.RoleEligibilityScheduleInstanceName != v.Expected.RoleEligibilityScheduleInstanceName {
			t.Fatalf("Expected %q but got %q for RoleEligibilityScheduleInstanceName", v.Expected.RoleEligibilityScheduleInstanceName, actual.RoleEligibilityScheduleInstanceName)
		}

	}
}

func TestSegmentsForScopedRoleEligibilityScheduleInstanceId(t *testing.T) {
	segments := ScopedRoleEligibilityScheduleInstanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ScopedRoleEligibilityScheduleInstanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
