package notificationmessagetemplatelocalizednotificationmessage

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementNotificationMessageTemplateId{}

func TestNewDeviceManagementNotificationMessageTemplateID(t *testing.T) {
	id := NewDeviceManagementNotificationMessageTemplateID("notificationMessageTemplateIdValue")

	if id.NotificationMessageTemplateId != "notificationMessageTemplateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'NotificationMessageTemplateId'", id.NotificationMessageTemplateId, "notificationMessageTemplateIdValue")
	}
}

func TestFormatDeviceManagementNotificationMessageTemplateID(t *testing.T) {
	actual := NewDeviceManagementNotificationMessageTemplateID("notificationMessageTemplateIdValue").ID()
	expected := "/deviceManagement/notificationMessageTemplates/notificationMessageTemplateIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementNotificationMessageTemplateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementNotificationMessageTemplateId
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
			Input: "/deviceManagement/notificationMessageTemplates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/notificationMessageTemplates/notificationMessageTemplateIdValue",
			Expected: &DeviceManagementNotificationMessageTemplateId{
				NotificationMessageTemplateId: "notificationMessageTemplateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/notificationMessageTemplates/notificationMessageTemplateIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementNotificationMessageTemplateID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.NotificationMessageTemplateId != v.Expected.NotificationMessageTemplateId {
			t.Fatalf("Expected %q but got %q for NotificationMessageTemplateId", v.Expected.NotificationMessageTemplateId, actual.NotificationMessageTemplateId)
		}

	}
}

func TestParseDeviceManagementNotificationMessageTemplateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementNotificationMessageTemplateId
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
			Input: "/deviceManagement/notificationMessageTemplates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/nOtIfIcAtIoNmEsSaGeTeMpLaTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/notificationMessageTemplates/notificationMessageTemplateIdValue",
			Expected: &DeviceManagementNotificationMessageTemplateId{
				NotificationMessageTemplateId: "notificationMessageTemplateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/notificationMessageTemplates/notificationMessageTemplateIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/nOtIfIcAtIoNmEsSaGeTeMpLaTeS/nOtIfIcAtIoNmEsSaGeTeMpLaTeIdVaLuE",
			Expected: &DeviceManagementNotificationMessageTemplateId{
				NotificationMessageTemplateId: "nOtIfIcAtIoNmEsSaGeTeMpLaTeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/nOtIfIcAtIoNmEsSaGeTeMpLaTeS/nOtIfIcAtIoNmEsSaGeTeMpLaTeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementNotificationMessageTemplateIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.NotificationMessageTemplateId != v.Expected.NotificationMessageTemplateId {
			t.Fatalf("Expected %q but got %q for NotificationMessageTemplateId", v.Expected.NotificationMessageTemplateId, actual.NotificationMessageTemplateId)
		}

	}
}

func TestSegmentsForDeviceManagementNotificationMessageTemplateId(t *testing.T) {
	segments := DeviceManagementNotificationMessageTemplateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementNotificationMessageTemplateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
