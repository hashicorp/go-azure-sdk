package extensions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &VersionId{}

func TestNewVersionID(t *testing.T) {
	id := NewVersionID("locationName", "publisherName", "extensionTypeName", "versionName")

	if id.LocationName != "locationName" {
		t.Fatalf("Expected %q but got %q for Segment 'LocationName'", id.LocationName, "locationName")
	}

	if id.PublisherName != "publisherName" {
		t.Fatalf("Expected %q but got %q for Segment 'PublisherName'", id.PublisherName, "publisherName")
	}

	if id.ExtensionTypeName != "extensionTypeName" {
		t.Fatalf("Expected %q but got %q for Segment 'ExtensionTypeName'", id.ExtensionTypeName, "extensionTypeName")
	}

	if id.VersionName != "versionName" {
		t.Fatalf("Expected %q but got %q for Segment 'VersionName'", id.VersionName, "versionName")
	}
}

func TestFormatVersionID(t *testing.T) {
	actual := NewVersionID("locationName", "publisherName", "extensionTypeName", "versionName").ID()
	expected := "/providers/Microsoft.HybridCompute/locations/locationName/publishers/publisherName/extensionTypes/extensionTypeName/versions/versionName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VersionId
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
			Input: "/providers/Microsoft.HybridCompute",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.HybridCompute/locations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.HybridCompute/locations/locationName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.HybridCompute/locations/locationName/publishers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.HybridCompute/locations/locationName/publishers/publisherName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.HybridCompute/locations/locationName/publishers/publisherName/extensionTypes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.HybridCompute/locations/locationName/publishers/publisherName/extensionTypes/extensionTypeName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.HybridCompute/locations/locationName/publishers/publisherName/extensionTypes/extensionTypeName/versions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.HybridCompute/locations/locationName/publishers/publisherName/extensionTypes/extensionTypeName/versions/versionName",
			Expected: &VersionId{
				LocationName:      "locationName",
				PublisherName:     "publisherName",
				ExtensionTypeName: "extensionTypeName",
				VersionName:       "versionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.HybridCompute/locations/locationName/publishers/publisherName/extensionTypes/extensionTypeName/versions/versionName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVersionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.LocationName != v.Expected.LocationName {
			t.Fatalf("Expected %q but got %q for LocationName", v.Expected.LocationName, actual.LocationName)
		}

		if actual.PublisherName != v.Expected.PublisherName {
			t.Fatalf("Expected %q but got %q for PublisherName", v.Expected.PublisherName, actual.PublisherName)
		}

		if actual.ExtensionTypeName != v.Expected.ExtensionTypeName {
			t.Fatalf("Expected %q but got %q for ExtensionTypeName", v.Expected.ExtensionTypeName, actual.ExtensionTypeName)
		}

		if actual.VersionName != v.Expected.VersionName {
			t.Fatalf("Expected %q but got %q for VersionName", v.Expected.VersionName, actual.VersionName)
		}

	}
}

func TestParseVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VersionId
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
			Input: "/providers/Microsoft.HybridCompute",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.HybridCompute/locations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/lOcAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.HybridCompute/locations/locationName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/lOcAtIoNs/lOcAtIoNnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.HybridCompute/locations/locationName/publishers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/pUbLiShErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.HybridCompute/locations/locationName/publishers/publisherName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/pUbLiShErS/pUbLiShErNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.HybridCompute/locations/locationName/publishers/publisherName/extensionTypes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/pUbLiShErS/pUbLiShErNaMe/eXtEnSiOnTyPeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.HybridCompute/locations/locationName/publishers/publisherName/extensionTypes/extensionTypeName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/pUbLiShErS/pUbLiShErNaMe/eXtEnSiOnTyPeS/eXtEnSiOnTyPeNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.HybridCompute/locations/locationName/publishers/publisherName/extensionTypes/extensionTypeName/versions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/pUbLiShErS/pUbLiShErNaMe/eXtEnSiOnTyPeS/eXtEnSiOnTyPeNaMe/vErSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.HybridCompute/locations/locationName/publishers/publisherName/extensionTypes/extensionTypeName/versions/versionName",
			Expected: &VersionId{
				LocationName:      "locationName",
				PublisherName:     "publisherName",
				ExtensionTypeName: "extensionTypeName",
				VersionName:       "versionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.HybridCompute/locations/locationName/publishers/publisherName/extensionTypes/extensionTypeName/versions/versionName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/pUbLiShErS/pUbLiShErNaMe/eXtEnSiOnTyPeS/eXtEnSiOnTyPeNaMe/vErSiOnS/vErSiOnNaMe",
			Expected: &VersionId{
				LocationName:      "lOcAtIoNnAmE",
				PublisherName:     "pUbLiShErNaMe",
				ExtensionTypeName: "eXtEnSiOnTyPeNaMe",
				VersionName:       "vErSiOnNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/pUbLiShErS/pUbLiShErNaMe/eXtEnSiOnTyPeS/eXtEnSiOnTyPeNaMe/vErSiOnS/vErSiOnNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVersionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.LocationName != v.Expected.LocationName {
			t.Fatalf("Expected %q but got %q for LocationName", v.Expected.LocationName, actual.LocationName)
		}

		if actual.PublisherName != v.Expected.PublisherName {
			t.Fatalf("Expected %q but got %q for PublisherName", v.Expected.PublisherName, actual.PublisherName)
		}

		if actual.ExtensionTypeName != v.Expected.ExtensionTypeName {
			t.Fatalf("Expected %q but got %q for ExtensionTypeName", v.Expected.ExtensionTypeName, actual.ExtensionTypeName)
		}

		if actual.VersionName != v.Expected.VersionName {
			t.Fatalf("Expected %q but got %q for VersionName", v.Expected.VersionName, actual.VersionName)
		}

	}
}

func TestSegmentsForVersionId(t *testing.T) {
	segments := VersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("VersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
