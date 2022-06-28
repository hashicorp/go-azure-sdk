package resources

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = SubscriptionResourceGroupId{}

func TestNewSubscriptionResourceGroupID(t *testing.T) {
	id := NewSubscriptionResourceGroupID("12345678-1234-9876-4563-123456789012", "sourceResourceGroupValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.SourceResourceGroupName != "sourceResourceGroupValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SourceResourceGroupName'", id.SourceResourceGroupName, "sourceResourceGroupValue")
	}
}

func TestFormatSubscriptionResourceGroupID(t *testing.T) {
	actual := NewSubscriptionResourceGroupID("12345678-1234-9876-4563-123456789012", "sourceResourceGroupValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/sourceResourceGroupValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSubscriptionResourceGroupID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SubscriptionResourceGroupId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/sourceResourceGroupValue",
			Expected: &SubscriptionResourceGroupId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				SourceResourceGroupName: "sourceResourceGroupValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/sourceResourceGroupValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSubscriptionResourceGroupID(v.Input)
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

		if actual.SourceResourceGroupName != v.Expected.SourceResourceGroupName {
			t.Fatalf("Expected %q but got %q for SourceResourceGroupName", v.Expected.SourceResourceGroupName, actual.SourceResourceGroupName)
		}

	}
}

func TestParseSubscriptionResourceGroupIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SubscriptionResourceGroupId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/sourceResourceGroupValue",
			Expected: &SubscriptionResourceGroupId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				SourceResourceGroupName: "sourceResourceGroupValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/sourceResourceGroupValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOuRcErEsOuRcEgRoUpVaLuE",
			Expected: &SubscriptionResourceGroupId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				SourceResourceGroupName: "sOuRcErEsOuRcEgRoUpVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOuRcErEsOuRcEgRoUpVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSubscriptionResourceGroupIDInsensitively(v.Input)
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

		if actual.SourceResourceGroupName != v.Expected.SourceResourceGroupName {
			t.Fatalf("Expected %q but got %q for SourceResourceGroupName", v.Expected.SourceResourceGroupName, actual.SourceResourceGroupName)
		}

	}
}

func TestSegmentsForSubscriptionResourceGroupId(t *testing.T) {
	segments := SubscriptionResourceGroupId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SubscriptionResourceGroupId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
