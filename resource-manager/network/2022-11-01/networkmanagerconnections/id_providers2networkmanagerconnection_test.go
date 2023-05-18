package networkmanagerconnections

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = Providers2NetworkManagerConnectionId{}

func TestNewProviders2NetworkManagerConnectionID(t *testing.T) {
	id := NewProviders2NetworkManagerConnectionID("managementGroupIdValue", "networkManagerConnectionValue")

	if id.ManagementGroupId != "managementGroupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagementGroupId'", id.ManagementGroupId, "managementGroupIdValue")
	}

	if id.NetworkManagerConnectionName != "networkManagerConnectionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'NetworkManagerConnectionName'", id.NetworkManagerConnectionName, "networkManagerConnectionValue")
	}
}

func TestFormatProviders2NetworkManagerConnectionID(t *testing.T) {
	actual := NewProviders2NetworkManagerConnectionID("managementGroupIdValue", "networkManagerConnectionValue").ID()
	expected := "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Network/networkManagerConnections/networkManagerConnectionValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseProviders2NetworkManagerConnectionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *Providers2NetworkManagerConnectionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Network",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Network/networkManagerConnections",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Network/networkManagerConnections/networkManagerConnectionValue",
			Expected: &Providers2NetworkManagerConnectionId{
				ManagementGroupId:            "managementGroupIdValue",
				NetworkManagerConnectionName: "networkManagerConnectionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Network/networkManagerConnections/networkManagerConnectionValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviders2NetworkManagerConnectionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagementGroupId != v.Expected.ManagementGroupId {
			t.Fatalf("Expected %q but got %q for ManagementGroupId", v.Expected.ManagementGroupId, actual.ManagementGroupId)
		}

		if actual.NetworkManagerConnectionName != v.Expected.NetworkManagerConnectionName {
			t.Fatalf("Expected %q but got %q for NetworkManagerConnectionName", v.Expected.NetworkManagerConnectionName, actual.NetworkManagerConnectionName)
		}

	}
}

func TestParseProviders2NetworkManagerConnectionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *Providers2NetworkManagerConnectionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpIdVaLuE/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Network",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpIdVaLuE/pRoViDeRs/mIcRoSoFt.nEtWoRk",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Network/networkManagerConnections",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpIdVaLuE/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkMaNaGeRcOnNeCtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Network/networkManagerConnections/networkManagerConnectionValue",
			Expected: &Providers2NetworkManagerConnectionId{
				ManagementGroupId:            "managementGroupIdValue",
				NetworkManagerConnectionName: "networkManagerConnectionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Network/networkManagerConnections/networkManagerConnectionValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpIdVaLuE/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkMaNaGeRcOnNeCtIoNs/nEtWoRkMaNaGeRcOnNeCtIoNvAlUe",
			Expected: &Providers2NetworkManagerConnectionId{
				ManagementGroupId:            "mAnAgEmEnTgRoUpIdVaLuE",
				NetworkManagerConnectionName: "nEtWoRkMaNaGeRcOnNeCtIoNvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpIdVaLuE/pRoViDeRs/mIcRoSoFt.nEtWoRk/nEtWoRkMaNaGeRcOnNeCtIoNs/nEtWoRkMaNaGeRcOnNeCtIoNvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviders2NetworkManagerConnectionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagementGroupId != v.Expected.ManagementGroupId {
			t.Fatalf("Expected %q but got %q for ManagementGroupId", v.Expected.ManagementGroupId, actual.ManagementGroupId)
		}

		if actual.NetworkManagerConnectionName != v.Expected.NetworkManagerConnectionName {
			t.Fatalf("Expected %q but got %q for NetworkManagerConnectionName", v.Expected.NetworkManagerConnectionName, actual.NetworkManagerConnectionName)
		}

	}
}

func TestSegmentsForProviders2NetworkManagerConnectionId(t *testing.T) {
	segments := Providers2NetworkManagerConnectionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("Providers2NetworkManagerConnectionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
