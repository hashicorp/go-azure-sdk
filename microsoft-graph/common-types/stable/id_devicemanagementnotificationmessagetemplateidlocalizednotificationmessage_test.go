package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId{}

func TestNewDeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageID(t *testing.T) {
	id := NewDeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageID("notificationMessageTemplateIdValue", "localizedNotificationMessageIdValue")

	if id.NotificationMessageTemplateId != "notificationMessageTemplateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'NotificationMessageTemplateId'", id.NotificationMessageTemplateId, "notificationMessageTemplateIdValue")
	}

	if id.LocalizedNotificationMessageId != "localizedNotificationMessageIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'LocalizedNotificationMessageId'", id.LocalizedNotificationMessageId, "localizedNotificationMessageIdValue")
	}
}

func TestFormatDeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageID(t *testing.T) {
	actual := NewDeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageID("notificationMessageTemplateIdValue", "localizedNotificationMessageIdValue").ID()
	expected := "/deviceManagement/notificationMessageTemplates/notificationMessageTemplateIdValue/localizedNotificationMessages/localizedNotificationMessageIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId
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
			// Incomplete URI
			Input: "/deviceManagement/notificationMessageTemplates/notificationMessageTemplateIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/notificationMessageTemplates/notificationMessageTemplateIdValue/localizedNotificationMessages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/notificationMessageTemplates/notificationMessageTemplateIdValue/localizedNotificationMessages/localizedNotificationMessageIdValue",
			Expected: &DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId{
				NotificationMessageTemplateId:  "notificationMessageTemplateIdValue",
				LocalizedNotificationMessageId: "localizedNotificationMessageIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/notificationMessageTemplates/notificationMessageTemplateIdValue/localizedNotificationMessages/localizedNotificationMessageIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageID(v.Input)
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

		if actual.LocalizedNotificationMessageId != v.Expected.LocalizedNotificationMessageId {
			t.Fatalf("Expected %q but got %q for LocalizedNotificationMessageId", v.Expected.LocalizedNotificationMessageId, actual.LocalizedNotificationMessageId)
		}

	}
}

func TestParseDeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId
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
			// Incomplete URI
			Input: "/deviceManagement/notificationMessageTemplates/notificationMessageTemplateIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/nOtIfIcAtIoNmEsSaGeTeMpLaTeS/nOtIfIcAtIoNmEsSaGeTeMpLaTeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/notificationMessageTemplates/notificationMessageTemplateIdValue/localizedNotificationMessages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/nOtIfIcAtIoNmEsSaGeTeMpLaTeS/nOtIfIcAtIoNmEsSaGeTeMpLaTeIdVaLuE/lOcAlIzEdNoTiFiCaTiOnMeSsAgEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/notificationMessageTemplates/notificationMessageTemplateIdValue/localizedNotificationMessages/localizedNotificationMessageIdValue",
			Expected: &DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId{
				NotificationMessageTemplateId:  "notificationMessageTemplateIdValue",
				LocalizedNotificationMessageId: "localizedNotificationMessageIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/notificationMessageTemplates/notificationMessageTemplateIdValue/localizedNotificationMessages/localizedNotificationMessageIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/nOtIfIcAtIoNmEsSaGeTeMpLaTeS/nOtIfIcAtIoNmEsSaGeTeMpLaTeIdVaLuE/lOcAlIzEdNoTiFiCaTiOnMeSsAgEs/lOcAlIzEdNoTiFiCaTiOnMeSsAgEiDvAlUe",
			Expected: &DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId{
				NotificationMessageTemplateId:  "nOtIfIcAtIoNmEsSaGeTeMpLaTeIdVaLuE",
				LocalizedNotificationMessageId: "lOcAlIzEdNoTiFiCaTiOnMeSsAgEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/nOtIfIcAtIoNmEsSaGeTeMpLaTeS/nOtIfIcAtIoNmEsSaGeTeMpLaTeIdVaLuE/lOcAlIzEdNoTiFiCaTiOnMeSsAgEs/lOcAlIzEdNoTiFiCaTiOnMeSsAgEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageIDInsensitively(v.Input)
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

		if actual.LocalizedNotificationMessageId != v.Expected.LocalizedNotificationMessageId {
			t.Fatalf("Expected %q but got %q for LocalizedNotificationMessageId", v.Expected.LocalizedNotificationMessageId, actual.LocalizedNotificationMessageId)
		}

	}
}

func TestSegmentsForDeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId(t *testing.T) {
	segments := DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
