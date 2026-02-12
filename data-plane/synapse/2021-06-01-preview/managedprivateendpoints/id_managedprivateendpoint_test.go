package managedprivateendpoints

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ManagedPrivateEndpointId{}

func TestNewManagedPrivateEndpointID(t *testing.T) {
	id := NewManagedPrivateEndpointID("https://endpoint-url.example.com", "managedVirtualNetworkName", "managedPrivateEndpointName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.ManagedVirtualNetworkName != "managedVirtualNetworkName" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedVirtualNetworkName'", id.ManagedVirtualNetworkName, "managedVirtualNetworkName")
	}

	if id.ManagedPrivateEndpointName != "managedPrivateEndpointName" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedPrivateEndpointName'", id.ManagedPrivateEndpointName, "managedPrivateEndpointName")
	}
}

func TestFormatManagedPrivateEndpointID(t *testing.T) {
	actual := NewManagedPrivateEndpointID("https://endpoint-url.example.com", "managedVirtualNetworkName", "managedPrivateEndpointName").ID()
	expected := "https://endpoint-url.example.com/managedVirtualNetworks/managedVirtualNetworkName/managedPrivateEndpoints/managedPrivateEndpointName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseManagedPrivateEndpointID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ManagedPrivateEndpointId
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
			// Incomplete URI
			Input: "https://endpoint-url.example.com/managedVirtualNetworks/managedVirtualNetworkName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/managedVirtualNetworks/managedVirtualNetworkName/managedPrivateEndpoints",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/managedVirtualNetworks/managedVirtualNetworkName/managedPrivateEndpoints/managedPrivateEndpointName",
			Expected: &ManagedPrivateEndpointId{
				BaseURI:                    "https://endpoint-url.example.com",
				ManagedVirtualNetworkName:  "managedVirtualNetworkName",
				ManagedPrivateEndpointName: "managedPrivateEndpointName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/managedVirtualNetworks/managedVirtualNetworkName/managedPrivateEndpoints/managedPrivateEndpointName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseManagedPrivateEndpointID(v.Input)
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

		if actual.ManagedPrivateEndpointName != v.Expected.ManagedPrivateEndpointName {
			t.Fatalf("Expected %q but got %q for ManagedPrivateEndpointName", v.Expected.ManagedPrivateEndpointName, actual.ManagedPrivateEndpointName)
		}

	}
}

func TestParseManagedPrivateEndpointIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ManagedPrivateEndpointId
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
			// Incomplete URI
			Input: "https://endpoint-url.example.com/managedVirtualNetworks/managedVirtualNetworkName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/mAnAgEdViRtUaLnEtWoRkS/mAnAgEdViRtUaLnEtWoRkNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/managedVirtualNetworks/managedVirtualNetworkName/managedPrivateEndpoints",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/mAnAgEdViRtUaLnEtWoRkS/mAnAgEdViRtUaLnEtWoRkNaMe/mAnAgEdPrIvAtEeNdPoInTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/managedVirtualNetworks/managedVirtualNetworkName/managedPrivateEndpoints/managedPrivateEndpointName",
			Expected: &ManagedPrivateEndpointId{
				BaseURI:                    "https://endpoint-url.example.com",
				ManagedVirtualNetworkName:  "managedVirtualNetworkName",
				ManagedPrivateEndpointName: "managedPrivateEndpointName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/managedVirtualNetworks/managedVirtualNetworkName/managedPrivateEndpoints/managedPrivateEndpointName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/mAnAgEdViRtUaLnEtWoRkS/mAnAgEdViRtUaLnEtWoRkNaMe/mAnAgEdPrIvAtEeNdPoInTs/mAnAgEdPrIvAtEeNdPoInTnAmE",
			Expected: &ManagedPrivateEndpointId{
				BaseURI:                    "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				ManagedVirtualNetworkName:  "mAnAgEdViRtUaLnEtWoRkNaMe",
				ManagedPrivateEndpointName: "mAnAgEdPrIvAtEeNdPoInTnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/mAnAgEdViRtUaLnEtWoRkS/mAnAgEdViRtUaLnEtWoRkNaMe/mAnAgEdPrIvAtEeNdPoInTs/mAnAgEdPrIvAtEeNdPoInTnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseManagedPrivateEndpointIDInsensitively(v.Input)
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

		if actual.ManagedPrivateEndpointName != v.Expected.ManagedPrivateEndpointName {
			t.Fatalf("Expected %q but got %q for ManagedPrivateEndpointName", v.Expected.ManagedPrivateEndpointName, actual.ManagedPrivateEndpointName)
		}

	}
}

func TestSegmentsForManagedPrivateEndpointId(t *testing.T) {
	segments := ManagedPrivateEndpointId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ManagedPrivateEndpointId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
