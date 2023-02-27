package usersession

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserSessionId{}

func TestNewUserSessionID(t *testing.T) {
	id := NewUserSessionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostPoolValue", "sessionHostValue", "userSessionIdValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.HostPoolName != "hostPoolValue" {
		t.Fatalf("Expected %q but got %q for Segment 'HostPoolName'", id.HostPoolName, "hostPoolValue")
	}

	if id.SessionHostName != "sessionHostValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SessionHostName'", id.SessionHostName, "sessionHostValue")
	}

	if id.UserSessionId != "userSessionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserSessionId'", id.UserSessionId, "userSessionIdValue")
	}
}

func TestFormatUserSessionID(t *testing.T) {
	actual := NewUserSessionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostPoolValue", "sessionHostValue", "userSessionIdValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DesktopVirtualization/hostPools/hostPoolValue/sessionHosts/sessionHostValue/userSessions/userSessionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserSessionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserSessionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DesktopVirtualization",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DesktopVirtualization/hostPools",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DesktopVirtualization/hostPools/hostPoolValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DesktopVirtualization/hostPools/hostPoolValue/sessionHosts",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DesktopVirtualization/hostPools/hostPoolValue/sessionHosts/sessionHostValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DesktopVirtualization/hostPools/hostPoolValue/sessionHosts/sessionHostValue/userSessions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DesktopVirtualization/hostPools/hostPoolValue/sessionHosts/sessionHostValue/userSessions/userSessionIdValue",
			Expected: &UserSessionId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "example-resource-group",
				HostPoolName:      "hostPoolValue",
				SessionHostName:   "sessionHostValue",
				UserSessionId:     "userSessionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DesktopVirtualization/hostPools/hostPoolValue/sessionHosts/sessionHostValue/userSessions/userSessionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserSessionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.HostPoolName != v.Expected.HostPoolName {
			t.Fatalf("Expected %q but got %q for HostPoolName", v.Expected.HostPoolName, actual.HostPoolName)
		}

		if actual.SessionHostName != v.Expected.SessionHostName {
			t.Fatalf("Expected %q but got %q for SessionHostName", v.Expected.SessionHostName, actual.SessionHostName)
		}

		if actual.UserSessionId != v.Expected.UserSessionId {
			t.Fatalf("Expected %q but got %q for UserSessionId", v.Expected.UserSessionId, actual.UserSessionId)
		}

	}
}

func TestParseUserSessionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserSessionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DesktopVirtualization",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dEsKtOpViRtUaLiZaTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DesktopVirtualization/hostPools",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dEsKtOpViRtUaLiZaTiOn/hOsTpOoLs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DesktopVirtualization/hostPools/hostPoolValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dEsKtOpViRtUaLiZaTiOn/hOsTpOoLs/hOsTpOoLvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DesktopVirtualization/hostPools/hostPoolValue/sessionHosts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dEsKtOpViRtUaLiZaTiOn/hOsTpOoLs/hOsTpOoLvAlUe/sEsSiOnHoStS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DesktopVirtualization/hostPools/hostPoolValue/sessionHosts/sessionHostValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dEsKtOpViRtUaLiZaTiOn/hOsTpOoLs/hOsTpOoLvAlUe/sEsSiOnHoStS/sEsSiOnHoStVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DesktopVirtualization/hostPools/hostPoolValue/sessionHosts/sessionHostValue/userSessions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dEsKtOpViRtUaLiZaTiOn/hOsTpOoLs/hOsTpOoLvAlUe/sEsSiOnHoStS/sEsSiOnHoStVaLuE/uSeRsEsSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DesktopVirtualization/hostPools/hostPoolValue/sessionHosts/sessionHostValue/userSessions/userSessionIdValue",
			Expected: &UserSessionId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "example-resource-group",
				HostPoolName:      "hostPoolValue",
				SessionHostName:   "sessionHostValue",
				UserSessionId:     "userSessionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.DesktopVirtualization/hostPools/hostPoolValue/sessionHosts/sessionHostValue/userSessions/userSessionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dEsKtOpViRtUaLiZaTiOn/hOsTpOoLs/hOsTpOoLvAlUe/sEsSiOnHoStS/sEsSiOnHoStVaLuE/uSeRsEsSiOnS/uSeRsEsSiOnIdVaLuE",
			Expected: &UserSessionId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "eXaMpLe-rEsOuRcE-GrOuP",
				HostPoolName:      "hOsTpOoLvAlUe",
				SessionHostName:   "sEsSiOnHoStVaLuE",
				UserSessionId:     "uSeRsEsSiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dEsKtOpViRtUaLiZaTiOn/hOsTpOoLs/hOsTpOoLvAlUe/sEsSiOnHoStS/sEsSiOnHoStVaLuE/uSeRsEsSiOnS/uSeRsEsSiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserSessionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.HostPoolName != v.Expected.HostPoolName {
			t.Fatalf("Expected %q but got %q for HostPoolName", v.Expected.HostPoolName, actual.HostPoolName)
		}

		if actual.SessionHostName != v.Expected.SessionHostName {
			t.Fatalf("Expected %q but got %q for SessionHostName", v.Expected.SessionHostName, actual.SessionHostName)
		}

		if actual.UserSessionId != v.Expected.UserSessionId {
			t.Fatalf("Expected %q but got %q for UserSessionId", v.Expected.UserSessionId, actual.UserSessionId)
		}

	}
}

func TestSegmentsForUserSessionId(t *testing.T) {
	segments := UserSessionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserSessionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
