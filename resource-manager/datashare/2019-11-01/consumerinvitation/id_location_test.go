package consumerinvitation

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &LocationId{}

func TestNewLocationID(t *testing.T) {
	id := NewLocationID("location")

	if id.LocationName != "location" {
		t.Fatalf("Expected %q but got %q for Segment 'LocationName'", id.LocationName, "location")
	}
}

func TestFormatLocationID(t *testing.T) {
	actual := NewLocationID("location").ID()
	expected := "/providers/Microsoft.DataShare/locations/location"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseLocationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LocationId
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
			Input: "/providers/Microsoft.DataShare",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.DataShare/locations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.DataShare/locations/location",
			Expected: &LocationId{
				LocationName: "location",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.DataShare/locations/location/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLocationID(v.Input)
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

	}
}

func TestParseLocationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LocationId
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
			Input: "/providers/Microsoft.DataShare",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.dAtAsHaRe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.DataShare/locations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.dAtAsHaRe/lOcAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.DataShare/locations/location",
			Expected: &LocationId{
				LocationName: "location",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.DataShare/locations/location/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.dAtAsHaRe/lOcAtIoNs/lOcAtIoN",
			Expected: &LocationId{
				LocationName: "lOcAtIoN",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.dAtAsHaRe/lOcAtIoNs/lOcAtIoN/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLocationIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForLocationId(t *testing.T) {
	segments := LocationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("LocationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
