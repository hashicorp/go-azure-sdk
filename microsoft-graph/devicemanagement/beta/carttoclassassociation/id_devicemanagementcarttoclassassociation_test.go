package carttoclassassociation

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCartToClassAssociationId{}

func TestNewDeviceManagementCartToClassAssociationID(t *testing.T) {
	id := NewDeviceManagementCartToClassAssociationID("cartToClassAssociationIdValue")

	if id.CartToClassAssociationId != "cartToClassAssociationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CartToClassAssociationId'", id.CartToClassAssociationId, "cartToClassAssociationIdValue")
	}
}

func TestFormatDeviceManagementCartToClassAssociationID(t *testing.T) {
	actual := NewDeviceManagementCartToClassAssociationID("cartToClassAssociationIdValue").ID()
	expected := "/deviceManagement/cartToClassAssociations/cartToClassAssociationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementCartToClassAssociationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementCartToClassAssociationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/cartToClassAssociations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/cartToClassAssociations/cartToClassAssociationIdValue",
			Expected: &DeviceManagementCartToClassAssociationId{
				CartToClassAssociationId: "cartToClassAssociationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/cartToClassAssociations/cartToClassAssociationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementCartToClassAssociationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CartToClassAssociationId != v.Expected.CartToClassAssociationId {
			t.Fatalf("Expected %q but got %q for CartToClassAssociationId", v.Expected.CartToClassAssociationId, actual.CartToClassAssociationId)
		}

	}
}

func TestParseDeviceManagementCartToClassAssociationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementCartToClassAssociationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/cartToClassAssociations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cArTtOcLaSsAsSoCiAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/cartToClassAssociations/cartToClassAssociationIdValue",
			Expected: &DeviceManagementCartToClassAssociationId{
				CartToClassAssociationId: "cartToClassAssociationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/cartToClassAssociations/cartToClassAssociationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cArTtOcLaSsAsSoCiAtIoNs/cArTtOcLaSsAsSoCiAtIoNiDvAlUe",
			Expected: &DeviceManagementCartToClassAssociationId{
				CartToClassAssociationId: "cArTtOcLaSsAsSoCiAtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cArTtOcLaSsAsSoCiAtIoNs/cArTtOcLaSsAsSoCiAtIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementCartToClassAssociationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CartToClassAssociationId != v.Expected.CartToClassAssociationId {
			t.Fatalf("Expected %q but got %q for CartToClassAssociationId", v.Expected.CartToClassAssociationId, actual.CartToClassAssociationId)
		}

	}
}

func TestSegmentsForDeviceManagementCartToClassAssociationId(t *testing.T) {
	segments := DeviceManagementCartToClassAssociationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementCartToClassAssociationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
