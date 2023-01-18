package roleeligibilityschedulerequests

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ScopedRoleEligibilityScheduleRequestId{}

func TestNewScopedRoleEligibilityScheduleRequestID(t *testing.T) {
	id := NewScopedRoleEligibilityScheduleRequestID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "roleEligibilityScheduleRequestValue")

	if id.Scope != "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'Scope'", id.Scope, "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")
	}

	if id.RoleEligibilityScheduleRequestName != "roleEligibilityScheduleRequestValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RoleEligibilityScheduleRequestName'", id.RoleEligibilityScheduleRequestName, "roleEligibilityScheduleRequestValue")
	}
}

func TestFormatScopedRoleEligibilityScheduleRequestID(t *testing.T) {
	actual := NewScopedRoleEligibilityScheduleRequestID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "roleEligibilityScheduleRequestValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleEligibilityScheduleRequests/roleEligibilityScheduleRequestValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseScopedRoleEligibilityScheduleRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedRoleEligibilityScheduleRequestId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleEligibilityScheduleRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleEligibilityScheduleRequests/roleEligibilityScheduleRequestValue",
			Expected: &ScopedRoleEligibilityScheduleRequestId{
				Scope:                              "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				RoleEligibilityScheduleRequestName: "roleEligibilityScheduleRequestValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleEligibilityScheduleRequests/roleEligibilityScheduleRequestValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedRoleEligibilityScheduleRequestID(v.Input)
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

		if actual.RoleEligibilityScheduleRequestName != v.Expected.RoleEligibilityScheduleRequestName {
			t.Fatalf("Expected %q but got %q for RoleEligibilityScheduleRequestName", v.Expected.RoleEligibilityScheduleRequestName, actual.RoleEligibilityScheduleRequestName)
		}

	}
}

func TestParseScopedRoleEligibilityScheduleRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedRoleEligibilityScheduleRequestId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleEligibilityScheduleRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/rOlEeLiGiBiLiTyScHeDuLeReQuEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleEligibilityScheduleRequests/roleEligibilityScheduleRequestValue",
			Expected: &ScopedRoleEligibilityScheduleRequestId{
				Scope:                              "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				RoleEligibilityScheduleRequestName: "roleEligibilityScheduleRequestValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/roleEligibilityScheduleRequests/roleEligibilityScheduleRequestValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/rOlEeLiGiBiLiTyScHeDuLeReQuEsTs/rOlEeLiGiBiLiTyScHeDuLeReQuEsTvAlUe",
			Expected: &ScopedRoleEligibilityScheduleRequestId{
				Scope:                              "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
				RoleEligibilityScheduleRequestName: "rOlEeLiGiBiLiTyScHeDuLeReQuEsTvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/rOlEeLiGiBiLiTyScHeDuLeReQuEsTs/rOlEeLiGiBiLiTyScHeDuLeReQuEsTvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedRoleEligibilityScheduleRequestIDInsensitively(v.Input)
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

		if actual.RoleEligibilityScheduleRequestName != v.Expected.RoleEligibilityScheduleRequestName {
			t.Fatalf("Expected %q but got %q for RoleEligibilityScheduleRequestName", v.Expected.RoleEligibilityScheduleRequestName, actual.RoleEligibilityScheduleRequestName)
		}

	}
}

func TestSegmentsForScopedRoleEligibilityScheduleRequestId(t *testing.T) {
	segments := ScopedRoleEligibilityScheduleRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ScopedRoleEligibilityScheduleRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
