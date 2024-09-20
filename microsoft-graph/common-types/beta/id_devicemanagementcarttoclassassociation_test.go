package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCartToClassAssociationId{}

func TestNewDeviceManagementCartToClassAssociationID(t *testing.T) {
	id := NewDeviceManagementCartToClassAssociationID("cartToClassAssociationId")

	if id.CartToClassAssociationId != "cartToClassAssociationId" {
		t.Fatalf("Expected %q but got %q for Segment 'CartToClassAssociationId'", id.CartToClassAssociationId, "cartToClassAssociationId")
	}
}

func TestFormatDeviceManagementCartToClassAssociationID(t *testing.T) {
	actual := NewDeviceManagementCartToClassAssociationID("cartToClassAssociationId").ID()
	expected := "/deviceManagement/cartToClassAssociations/cartToClassAssociationId"
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
			Input: "/deviceManagement/cartToClassAssociations/cartToClassAssociationId",
			Expected: &DeviceManagementCartToClassAssociationId{
				CartToClassAssociationId: "cartToClassAssociationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/cartToClassAssociations/cartToClassAssociationId/extra",
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
			Input: "/deviceManagement/cartToClassAssociations/cartToClassAssociationId",
			Expected: &DeviceManagementCartToClassAssociationId{
				CartToClassAssociationId: "cartToClassAssociationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/cartToClassAssociations/cartToClassAssociationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cArTtOcLaSsAsSoCiAtIoNs/cArTtOcLaSsAsSoCiAtIoNiD",
			Expected: &DeviceManagementCartToClassAssociationId{
				CartToClassAssociationId: "cArTtOcLaSsAsSoCiAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cArTtOcLaSsAsSoCiAtIoNs/cArTtOcLaSsAsSoCiAtIoNiD/extra",
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
