package agentversions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &OsTypeId{}

func TestNewOsTypeID(t *testing.T) {
	id := NewOsTypeID("osTypeValue")

	if id.OsTypeName != "osTypeValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OsTypeName'", id.OsTypeName, "osTypeValue")
	}
}

func TestFormatOsTypeID(t *testing.T) {
	actual := NewOsTypeID("osTypeValue").ID()
	expected := "/providers/Microsoft.HybridCompute/osType/osTypeValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseOsTypeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *OsTypeId
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
			Input: "/providers/Microsoft.HybridCompute/osType",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.HybridCompute/osType/osTypeValue",
			Expected: &OsTypeId{
				OsTypeName: "osTypeValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.HybridCompute/osType/osTypeValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseOsTypeID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OsTypeName != v.Expected.OsTypeName {
			t.Fatalf("Expected %q but got %q for OsTypeName", v.Expected.OsTypeName, actual.OsTypeName)
		}

	}
}

func TestParseOsTypeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *OsTypeId
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
			Input: "/providers/Microsoft.HybridCompute/osType",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/oStYpE",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.HybridCompute/osType/osTypeValue",
			Expected: &OsTypeId{
				OsTypeName: "osTypeValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.HybridCompute/osType/osTypeValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/oStYpE/oStYpEvAlUe",
			Expected: &OsTypeId{
				OsTypeName: "oStYpEvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/oStYpE/oStYpEvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseOsTypeIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OsTypeName != v.Expected.OsTypeName {
			t.Fatalf("Expected %q but got %q for OsTypeName", v.Expected.OsTypeName, actual.OsTypeName)
		}

	}
}

func TestSegmentsForOsTypeId(t *testing.T) {
	segments := OsTypeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("OsTypeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
