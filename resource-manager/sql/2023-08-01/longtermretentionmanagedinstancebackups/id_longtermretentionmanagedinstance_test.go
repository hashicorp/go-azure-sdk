package longtermretentionmanagedinstancebackups

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &LongTermRetentionManagedInstanceId{}

func TestNewLongTermRetentionManagedInstanceID(t *testing.T) {
	id := NewLongTermRetentionManagedInstanceID("12345678-1234-9876-4563-123456789012", "locationName", "longTermRetentionManagedInstanceName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.LocationName != "locationName" {
		t.Fatalf("Expected %q but got %q for Segment 'LocationName'", id.LocationName, "locationName")
	}

	if id.LongTermRetentionManagedInstanceName != "longTermRetentionManagedInstanceName" {
		t.Fatalf("Expected %q but got %q for Segment 'LongTermRetentionManagedInstanceName'", id.LongTermRetentionManagedInstanceName, "longTermRetentionManagedInstanceName")
	}
}

func TestFormatLongTermRetentionManagedInstanceID(t *testing.T) {
	actual := NewLongTermRetentionManagedInstanceID("12345678-1234-9876-4563-123456789012", "locationName", "longTermRetentionManagedInstanceName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationName/longTermRetentionManagedInstances/longTermRetentionManagedInstanceName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseLongTermRetentionManagedInstanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LongTermRetentionManagedInstanceId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationName/longTermRetentionManagedInstances",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationName/longTermRetentionManagedInstances/longTermRetentionManagedInstanceName",
			Expected: &LongTermRetentionManagedInstanceId{
				SubscriptionId:                       "12345678-1234-9876-4563-123456789012",
				LocationName:                         "locationName",
				LongTermRetentionManagedInstanceName: "longTermRetentionManagedInstanceName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationName/longTermRetentionManagedInstances/longTermRetentionManagedInstanceName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLongTermRetentionManagedInstanceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.LocationName != v.Expected.LocationName {
			t.Fatalf("Expected %q but got %q for LocationName", v.Expected.LocationName, actual.LocationName)
		}

		if actual.LongTermRetentionManagedInstanceName != v.Expected.LongTermRetentionManagedInstanceName {
			t.Fatalf("Expected %q but got %q for LongTermRetentionManagedInstanceName", v.Expected.LongTermRetentionManagedInstanceName, actual.LongTermRetentionManagedInstanceName)
		}

	}
}

func TestParseLongTermRetentionManagedInstanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LongTermRetentionManagedInstanceId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationName/longTermRetentionManagedInstances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNnAmE/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationName/longTermRetentionManagedInstances/longTermRetentionManagedInstanceName",
			Expected: &LongTermRetentionManagedInstanceId{
				SubscriptionId:                       "12345678-1234-9876-4563-123456789012",
				LocationName:                         "locationName",
				LongTermRetentionManagedInstanceName: "longTermRetentionManagedInstanceName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Sql/locations/locationName/longTermRetentionManagedInstances/longTermRetentionManagedInstanceName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNnAmE/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEs/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEnAmE",
			Expected: &LongTermRetentionManagedInstanceId{
				SubscriptionId:                       "12345678-1234-9876-4563-123456789012",
				LocationName:                         "lOcAtIoNnAmE",
				LongTermRetentionManagedInstanceName: "lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sQl/lOcAtIoNs/lOcAtIoNnAmE/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEs/lOnGtErMrEtEnTiOnMaNaGeDiNsTaNcEnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLongTermRetentionManagedInstanceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.LocationName != v.Expected.LocationName {
			t.Fatalf("Expected %q but got %q for LocationName", v.Expected.LocationName, actual.LocationName)
		}

		if actual.LongTermRetentionManagedInstanceName != v.Expected.LongTermRetentionManagedInstanceName {
			t.Fatalf("Expected %q but got %q for LongTermRetentionManagedInstanceName", v.Expected.LongTermRetentionManagedInstanceName, actual.LongTermRetentionManagedInstanceName)
		}

	}
}

func TestSegmentsForLongTermRetentionManagedInstanceId(t *testing.T) {
	segments := LongTermRetentionManagedInstanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("LongTermRetentionManagedInstanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
