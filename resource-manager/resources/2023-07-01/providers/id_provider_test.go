package providers

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ProviderId{}

func TestNewProviderID(t *testing.T) {
	id := NewProviderID("resourceProviderNamespace")

	if id.ProviderName != "resourceProviderNamespace" {
		t.Fatalf("Expected %q but got %q for Segment 'ProviderName'", id.ProviderName, "resourceProviderNamespace")
	}
}

func TestFormatProviderID(t *testing.T) {
	actual := NewProviderID("resourceProviderNamespace").ID()
	expected := "/providers/resourceProviderNamespace"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseProviderID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ProviderId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/resourceProviderNamespace",
			Expected: &ProviderId{
				ProviderName: "resourceProviderNamespace",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/resourceProviderNamespace/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviderID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ProviderName != v.Expected.ProviderName {
			t.Fatalf("Expected %q but got %q for ProviderName", v.Expected.ProviderName, actual.ProviderName)
		}

	}
}

func TestParseProviderIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ProviderId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/resourceProviderNamespace",
			Expected: &ProviderId{
				ProviderName: "resourceProviderNamespace",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/resourceProviderNamespace/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/rEsOuRcEpRoViDeRnAmEsPaCe",
			Expected: &ProviderId{
				ProviderName: "rEsOuRcEpRoViDeRnAmEsPaCe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/rEsOuRcEpRoViDeRnAmEsPaCe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviderIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ProviderName != v.Expected.ProviderName {
			t.Fatalf("Expected %q but got %q for ProviderName", v.Expected.ProviderName, actual.ProviderName)
		}

	}
}

func TestSegmentsForProviderId(t *testing.T) {
	segments := ProviderId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ProviderId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
