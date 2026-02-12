package managedprivateendpoints

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ManagedVirtualNetworkId{}

func TestNewManagedVirtualNetworkID(t *testing.T) {
	id := NewManagedVirtualNetworkID("https://endpoint-url.example.com", "managedVirtualNetworkName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.ManagedVirtualNetworkName != "managedVirtualNetworkName" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedVirtualNetworkName'", id.ManagedVirtualNetworkName, "managedVirtualNetworkName")
	}
}

func TestFormatManagedVirtualNetworkID(t *testing.T) {
	actual := NewManagedVirtualNetworkID("https://endpoint-url.example.com", "managedVirtualNetworkName").ID()
	expected := "https://endpoint-url.example.com/managedVirtualNetworks/managedVirtualNetworkName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseManagedVirtualNetworkID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ManagedVirtualNetworkId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/managedVirtualNetworks",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/managedVirtualNetworks/managedVirtualNetworkName",
			Expected: &ManagedVirtualNetworkId{
				BaseURI:                   "https://endpoint-url.example.com",
				ManagedVirtualNetworkName: "managedVirtualNetworkName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/managedVirtualNetworks/managedVirtualNetworkName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseManagedVirtualNetworkID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BaseURI != v.Expected.BaseURI {
			t.Fatalf("Expected %q but got %q for BaseURI", v.Expected.BaseURI, actual.BaseURI)
		}

		if actual.ManagedVirtualNetworkName != v.Expected.ManagedVirtualNetworkName {
			t.Fatalf("Expected %q but got %q for ManagedVirtualNetworkName", v.Expected.ManagedVirtualNetworkName, actual.ManagedVirtualNetworkName)
		}

	}
}

func TestParseManagedVirtualNetworkIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ManagedVirtualNetworkId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/managedVirtualNetworks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/mAnAgEdViRtUaLnEtWoRkS",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/managedVirtualNetworks/managedVirtualNetworkName",
			Expected: &ManagedVirtualNetworkId{
				BaseURI:                   "https://endpoint-url.example.com",
				ManagedVirtualNetworkName: "managedVirtualNetworkName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/managedVirtualNetworks/managedVirtualNetworkName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/mAnAgEdViRtUaLnEtWoRkS/mAnAgEdViRtUaLnEtWoRkNaMe",
			Expected: &ManagedVirtualNetworkId{
				BaseURI:                   "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				ManagedVirtualNetworkName: "mAnAgEdViRtUaLnEtWoRkNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/mAnAgEdViRtUaLnEtWoRkS/mAnAgEdViRtUaLnEtWoRkNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseManagedVirtualNetworkIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BaseURI != v.Expected.BaseURI {
			t.Fatalf("Expected %q but got %q for BaseURI", v.Expected.BaseURI, actual.BaseURI)
		}

		if actual.ManagedVirtualNetworkName != v.Expected.ManagedVirtualNetworkName {
			t.Fatalf("Expected %q but got %q for ManagedVirtualNetworkName", v.Expected.ManagedVirtualNetworkName, actual.ManagedVirtualNetworkName)
		}

	}
}

func TestSegmentsForManagedVirtualNetworkId(t *testing.T) {
	segments := ManagedVirtualNetworkId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ManagedVirtualNetworkId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
