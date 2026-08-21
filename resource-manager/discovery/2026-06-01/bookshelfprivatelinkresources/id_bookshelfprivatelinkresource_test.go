package bookshelfprivatelinkresources

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &BookshelfPrivateLinkResourceId{}

func TestNewBookshelfPrivateLinkResourceID(t *testing.T) {
	id := NewBookshelfPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bookshelfName", "privateLinkResourceName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.BookshelfName != "bookshelfName" {
		t.Fatalf("Expected %q but got %q for Segment 'BookshelfName'", id.BookshelfName, "bookshelfName")
	}

	if id.PrivateLinkResourceName != "privateLinkResourceName" {
		t.Fatalf("Expected %q but got %q for Segment 'PrivateLinkResourceName'", id.PrivateLinkResourceName, "privateLinkResourceName")
	}
}

func TestFormatBookshelfPrivateLinkResourceID(t *testing.T) {
	actual := NewBookshelfPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bookshelfName", "privateLinkResourceName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Discovery/bookshelves/bookshelfName/privateLinkResources/privateLinkResourceName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseBookshelfPrivateLinkResourceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BookshelfPrivateLinkResourceId
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
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Discovery",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Discovery/bookshelves",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Discovery/bookshelves/bookshelfName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Discovery/bookshelves/bookshelfName/privateLinkResources",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Discovery/bookshelves/bookshelfName/privateLinkResources/privateLinkResourceName",
			Expected: &BookshelfPrivateLinkResourceId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:       "example-resource-group",
				BookshelfName:           "bookshelfName",
				PrivateLinkResourceName: "privateLinkResourceName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Discovery/bookshelves/bookshelfName/privateLinkResources/privateLinkResourceName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBookshelfPrivateLinkResourceID(v.Input)
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

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.BookshelfName != v.Expected.BookshelfName {
			t.Fatalf("Expected %q but got %q for BookshelfName", v.Expected.BookshelfName, actual.BookshelfName)
		}

		if actual.PrivateLinkResourceName != v.Expected.PrivateLinkResourceName {
			t.Fatalf("Expected %q but got %q for PrivateLinkResourceName", v.Expected.PrivateLinkResourceName, actual.PrivateLinkResourceName)
		}

	}
}

func TestParseBookshelfPrivateLinkResourceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BookshelfPrivateLinkResourceId
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
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Discovery",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dIsCoVeRy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Discovery/bookshelves",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dIsCoVeRy/bOoKsHeLvEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Discovery/bookshelves/bookshelfName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dIsCoVeRy/bOoKsHeLvEs/bOoKsHeLfNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Discovery/bookshelves/bookshelfName/privateLinkResources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dIsCoVeRy/bOoKsHeLvEs/bOoKsHeLfNaMe/pRiVaTeLiNkReSoUrCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Discovery/bookshelves/bookshelfName/privateLinkResources/privateLinkResourceName",
			Expected: &BookshelfPrivateLinkResourceId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:       "example-resource-group",
				BookshelfName:           "bookshelfName",
				PrivateLinkResourceName: "privateLinkResourceName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Discovery/bookshelves/bookshelfName/privateLinkResources/privateLinkResourceName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dIsCoVeRy/bOoKsHeLvEs/bOoKsHeLfNaMe/pRiVaTeLiNkReSoUrCeS/pRiVaTeLiNkReSoUrCeNaMe",
			Expected: &BookshelfPrivateLinkResourceId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:       "eXaMpLe-rEsOuRcE-GrOuP",
				BookshelfName:           "bOoKsHeLfNaMe",
				PrivateLinkResourceName: "pRiVaTeLiNkReSoUrCeNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.dIsCoVeRy/bOoKsHeLvEs/bOoKsHeLfNaMe/pRiVaTeLiNkReSoUrCeS/pRiVaTeLiNkReSoUrCeNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBookshelfPrivateLinkResourceIDInsensitively(v.Input)
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

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.BookshelfName != v.Expected.BookshelfName {
			t.Fatalf("Expected %q but got %q for BookshelfName", v.Expected.BookshelfName, actual.BookshelfName)
		}

		if actual.PrivateLinkResourceName != v.Expected.PrivateLinkResourceName {
			t.Fatalf("Expected %q but got %q for PrivateLinkResourceName", v.Expected.PrivateLinkResourceName, actual.PrivateLinkResourceName)
		}

	}
}

func TestSegmentsForBookshelfPrivateLinkResourceId(t *testing.T) {
	segments := BookshelfPrivateLinkResourceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("BookshelfPrivateLinkResourceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %d unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
