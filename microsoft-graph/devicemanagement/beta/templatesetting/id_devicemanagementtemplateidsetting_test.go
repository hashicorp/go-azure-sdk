package templatesetting

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTemplateIdSettingId{}

func TestNewDeviceManagementTemplateIdSettingID(t *testing.T) {
	id := NewDeviceManagementTemplateIdSettingID("deviceManagementTemplateIdValue", "deviceManagementSettingInstanceIdValue")

	if id.DeviceManagementTemplateId != "deviceManagementTemplateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementTemplateId'", id.DeviceManagementTemplateId, "deviceManagementTemplateIdValue")
	}

	if id.DeviceManagementSettingInstanceId != "deviceManagementSettingInstanceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementSettingInstanceId'", id.DeviceManagementSettingInstanceId, "deviceManagementSettingInstanceIdValue")
	}
}

func TestFormatDeviceManagementTemplateIdSettingID(t *testing.T) {
	actual := NewDeviceManagementTemplateIdSettingID("deviceManagementTemplateIdValue", "deviceManagementSettingInstanceIdValue").ID()
	expected := "/deviceManagement/templates/deviceManagementTemplateIdValue/settings/deviceManagementSettingInstanceIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementTemplateIdSettingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTemplateIdSettingId
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
			Input: "/deviceManagement/templates",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue/settings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue/settings/deviceManagementSettingInstanceIdValue",
			Expected: &DeviceManagementTemplateIdSettingId{
				DeviceManagementTemplateId:        "deviceManagementTemplateIdValue",
				DeviceManagementSettingInstanceId: "deviceManagementSettingInstanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue/settings/deviceManagementSettingInstanceIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTemplateIdSettingID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementTemplateId != v.Expected.DeviceManagementTemplateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementTemplateId", v.Expected.DeviceManagementTemplateId, actual.DeviceManagementTemplateId)
		}

		if actual.DeviceManagementSettingInstanceId != v.Expected.DeviceManagementSettingInstanceId {
			t.Fatalf("Expected %q but got %q for DeviceManagementSettingInstanceId", v.Expected.DeviceManagementSettingInstanceId, actual.DeviceManagementSettingInstanceId)
		}

	}
}

func TestParseDeviceManagementTemplateIdSettingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTemplateIdSettingId
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
			Input: "/deviceManagement/templates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue/settings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiDvAlUe/sEtTiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue/settings/deviceManagementSettingInstanceIdValue",
			Expected: &DeviceManagementTemplateIdSettingId{
				DeviceManagementTemplateId:        "deviceManagementTemplateIdValue",
				DeviceManagementSettingInstanceId: "deviceManagementSettingInstanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/templates/deviceManagementTemplateIdValue/settings/deviceManagementSettingInstanceIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiDvAlUe/sEtTiNgS/dEvIcEmAnAgEmEnTsEtTiNgInStAnCeIdVaLuE",
			Expected: &DeviceManagementTemplateIdSettingId{
				DeviceManagementTemplateId:        "dEvIcEmAnAgEmEnTtEmPlAtEiDvAlUe",
				DeviceManagementSettingInstanceId: "dEvIcEmAnAgEmEnTsEtTiNgInStAnCeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEs/dEvIcEmAnAgEmEnTtEmPlAtEiDvAlUe/sEtTiNgS/dEvIcEmAnAgEmEnTsEtTiNgInStAnCeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTemplateIdSettingIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementTemplateId != v.Expected.DeviceManagementTemplateId {
			t.Fatalf("Expected %q but got %q for DeviceManagementTemplateId", v.Expected.DeviceManagementTemplateId, actual.DeviceManagementTemplateId)
		}

		if actual.DeviceManagementSettingInstanceId != v.Expected.DeviceManagementSettingInstanceId {
			t.Fatalf("Expected %q but got %q for DeviceManagementSettingInstanceId", v.Expected.DeviceManagementSettingInstanceId, actual.DeviceManagementSettingInstanceId)
		}

	}
}

func TestSegmentsForDeviceManagementTemplateIdSettingId(t *testing.T) {
	segments := DeviceManagementTemplateIdSettingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementTemplateIdSettingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
