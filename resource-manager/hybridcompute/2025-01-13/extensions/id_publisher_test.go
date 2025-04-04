package extensions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PublisherId{}

func TestNewPublisherID(t *testing.T) {
	id := NewPublisherID("locationName", "publisherName")

	if id.LocationName != "locationName" {
		t.Fatalf("Expected %q but got %q for Segment 'LocationName'", id.LocationName, "locationName")
	}

	if id.PublisherName != "publisherName" {
		t.Fatalf("Expected %q but got %q for Segment 'PublisherName'", id.PublisherName, "publisherName")
	}
}

func TestFormatPublisherID(t *testing.T) {
	actual := NewPublisherID("locationName", "publisherName").ID()
	expected := "/providers/Microsoft.HybridCompute/locations/locationName/publishers/publisherName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePublisherID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PublisherId
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
			// Valid URI
			Input: "/providers/Microsoft.HybridCompute/locations/locationName/publishers/publisherName",
			Expected: &PublisherId{
				LocationName:  "locationName",
				PublisherName: "publisherName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.HybridCompute/locations/locationName/publishers/publisherName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePublisherID(v.Input)
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

	}
}

func TestParsePublisherIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PublisherId
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
			// Valid URI
			Input: "/providers/Microsoft.HybridCompute/locations/locationName/publishers/publisherName",
			Expected: &PublisherId{
				LocationName:  "locationName",
				PublisherName: "publisherName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.HybridCompute/locations/locationName/publishers/publisherName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/pUbLiShErS/pUbLiShErNaMe",
			Expected: &PublisherId{
				LocationName:  "lOcAtIoNnAmE",
				PublisherName: "pUbLiShErNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/pUbLiShErS/pUbLiShErNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePublisherIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForPublisherId(t *testing.T) {
	segments := PublisherId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PublisherId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
