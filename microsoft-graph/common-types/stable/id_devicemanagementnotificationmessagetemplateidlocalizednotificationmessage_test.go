package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId{}

func TestNewDeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageID(t *testing.T) {
	id := NewDeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageID("notificationMessageTemplateId", "localizedNotificationMessageId")

	if id.NotificationMessageTemplateId != "notificationMessageTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'NotificationMessageTemplateId'", id.NotificationMessageTemplateId, "notificationMessageTemplateId")
	}

	if id.LocalizedNotificationMessageId != "localizedNotificationMessageId" {
		t.Fatalf("Expected %q but got %q for Segment 'LocalizedNotificationMessageId'", id.LocalizedNotificationMessageId, "localizedNotificationMessageId")
	}
}

func TestFormatDeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageID(t *testing.T) {
	actual := NewDeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageID("notificationMessageTemplateId", "localizedNotificationMessageId").ID()
	expected := "/deviceManagement/notificationMessageTemplates/notificationMessageTemplateId/localizedNotificationMessages/localizedNotificationMessageId"
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
			Input: "/deviceManagement/notificationMessageTemplates/notificationMessageTemplateId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/notificationMessageTemplates/notificationMessageTemplateId/localizedNotificationMessages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/notificationMessageTemplates/notificationMessageTemplateId/localizedNotificationMessages/localizedNotificationMessageId",
			Expected: &DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId{
				NotificationMessageTemplateId:  "notificationMessageTemplateId",
				LocalizedNotificationMessageId: "localizedNotificationMessageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/notificationMessageTemplates/notificationMessageTemplateId/localizedNotificationMessages/localizedNotificationMessageId/extra",
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
			Input: "/deviceManagement/notificationMessageTemplates/notificationMessageTemplateId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/nOtIfIcAtIoNmEsSaGeTeMpLaTeS/nOtIfIcAtIoNmEsSaGeTeMpLaTeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/notificationMessageTemplates/notificationMessageTemplateId/localizedNotificationMessages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/nOtIfIcAtIoNmEsSaGeTeMpLaTeS/nOtIfIcAtIoNmEsSaGeTeMpLaTeId/lOcAlIzEdNoTiFiCaTiOnMeSsAgEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/notificationMessageTemplates/notificationMessageTemplateId/localizedNotificationMessages/localizedNotificationMessageId",
			Expected: &DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId{
				NotificationMessageTemplateId:  "notificationMessageTemplateId",
				LocalizedNotificationMessageId: "localizedNotificationMessageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/notificationMessageTemplates/notificationMessageTemplateId/localizedNotificationMessages/localizedNotificationMessageId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/nOtIfIcAtIoNmEsSaGeTeMpLaTeS/nOtIfIcAtIoNmEsSaGeTeMpLaTeId/lOcAlIzEdNoTiFiCaTiOnMeSsAgEs/lOcAlIzEdNoTiFiCaTiOnMeSsAgEiD",
			Expected: &DeviceManagementNotificationMessageTemplateIdLocalizedNotificationMessageId{
				NotificationMessageTemplateId:  "nOtIfIcAtIoNmEsSaGeTeMpLaTeId",
				LocalizedNotificationMessageId: "lOcAlIzEdNoTiFiCaTiOnMeSsAgEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/nOtIfIcAtIoNmEsSaGeTeMpLaTeS/nOtIfIcAtIoNmEsSaGeTeMpLaTeId/lOcAlIzEdNoTiFiCaTiOnMeSsAgEs/lOcAlIzEdNoTiFiCaTiOnMeSsAgEiD/extra",
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
