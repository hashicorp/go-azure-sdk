package subscription

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectorySubscriptionId{}

func TestNewDirectorySubscriptionID(t *testing.T) {
	id := NewDirectorySubscriptionID("companySubscriptionIdValue")

	if id.CompanySubscriptionId != "companySubscriptionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CompanySubscriptionId'", id.CompanySubscriptionId, "companySubscriptionIdValue")
	}
}

func TestFormatDirectorySubscriptionID(t *testing.T) {
	actual := NewDirectorySubscriptionID("companySubscriptionIdValue").ID()
	expected := "/directory/subscriptions/companySubscriptionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectorySubscriptionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectorySubscriptionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/subscriptions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/subscriptions/companySubscriptionIdValue",
			Expected: &DirectorySubscriptionId{
				CompanySubscriptionId: "companySubscriptionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/subscriptions/companySubscriptionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectorySubscriptionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CompanySubscriptionId != v.Expected.CompanySubscriptionId {
			t.Fatalf("Expected %q but got %q for CompanySubscriptionId", v.Expected.CompanySubscriptionId, actual.CompanySubscriptionId)
		}

	}
}

func TestParseDirectorySubscriptionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectorySubscriptionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/subscriptions/companySubscriptionIdValue",
			Expected: &DirectorySubscriptionId{
				CompanySubscriptionId: "companySubscriptionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/subscriptions/companySubscriptionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/sUbScRiPtIoNs/cOmPaNySuBsCrIpTiOnIdVaLuE",
			Expected: &DirectorySubscriptionId{
				CompanySubscriptionId: "cOmPaNySuBsCrIpTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/sUbScRiPtIoNs/cOmPaNySuBsCrIpTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectorySubscriptionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CompanySubscriptionId != v.Expected.CompanySubscriptionId {
			t.Fatalf("Expected %q but got %q for CompanySubscriptionId", v.Expected.CompanySubscriptionId, actual.CompanySubscriptionId)
		}

	}
}

func TestSegmentsForDirectorySubscriptionId(t *testing.T) {
	segments := DirectorySubscriptionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectorySubscriptionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
