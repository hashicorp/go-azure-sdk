package tags

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &TagValueId{}

func TestNewTagValueID(t *testing.T) {
	id := NewTagValueID("12345678-1234-9876-4563-123456789012", "tagValue", "tagValueValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.TagName != "tagValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TagName'", id.TagName, "tagValue")
	}

	if id.TagValueName != "tagValueValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TagValueName'", id.TagValueName, "tagValueValue")
	}
}

func TestFormatTagValueID(t *testing.T) {
	actual := NewTagValueID("12345678-1234-9876-4563-123456789012", "tagValue", "tagValueValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/tagNames/tagValue/tagValues/tagValueValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseTagValueID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *TagValueId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/tagNames",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/tagNames/tagValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/tagNames/tagValue/tagValues",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/tagNames/tagValue/tagValues/tagValueValue",
			Expected: &TagValueId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				TagName:        "tagValue",
				TagValueName:   "tagValueValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/tagNames/tagValue/tagValues/tagValueValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseTagValueID(v.Input)
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

		if actual.TagName != v.Expected.TagName {
			t.Fatalf("Expected %q but got %q for TagName", v.Expected.TagName, actual.TagName)
		}

		if actual.TagValueName != v.Expected.TagValueName {
			t.Fatalf("Expected %q but got %q for TagValueName", v.Expected.TagValueName, actual.TagValueName)
		}

	}
}

func TestParseTagValueIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *TagValueId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/tagNames",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/tAgNaMeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/tagNames/tagValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/tAgNaMeS/tAgVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/tagNames/tagValue/tagValues",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/tAgNaMeS/tAgVaLuE/tAgVaLuEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/tagNames/tagValue/tagValues/tagValueValue",
			Expected: &TagValueId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				TagName:        "tagValue",
				TagValueName:   "tagValueValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/tagNames/tagValue/tagValues/tagValueValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/tAgNaMeS/tAgVaLuE/tAgVaLuEs/tAgVaLuEvAlUe",
			Expected: &TagValueId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				TagName:        "tAgVaLuE",
				TagValueName:   "tAgVaLuEvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/tAgNaMeS/tAgVaLuE/tAgVaLuEs/tAgVaLuEvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseTagValueIDInsensitively(v.Input)
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

		if actual.TagName != v.Expected.TagName {
			t.Fatalf("Expected %q but got %q for TagName", v.Expected.TagName, actual.TagName)
		}

		if actual.TagValueName != v.Expected.TagValueName {
			t.Fatalf("Expected %q but got %q for TagValueName", v.Expected.TagValueName, actual.TagValueName)
		}

	}
}

func TestSegmentsForTagValueId(t *testing.T) {
	segments := TagValueId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("TagValueId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
