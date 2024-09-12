package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectorySharedEmailDomainId{}

func TestNewDirectorySharedEmailDomainID(t *testing.T) {
	id := NewDirectorySharedEmailDomainID("sharedEmailDomainIdValue")

	if id.SharedEmailDomainId != "sharedEmailDomainIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SharedEmailDomainId'", id.SharedEmailDomainId, "sharedEmailDomainIdValue")
	}
}

func TestFormatDirectorySharedEmailDomainID(t *testing.T) {
	actual := NewDirectorySharedEmailDomainID("sharedEmailDomainIdValue").ID()
	expected := "/directory/sharedEmailDomains/sharedEmailDomainIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectorySharedEmailDomainID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectorySharedEmailDomainId
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
			Input: "/directory/sharedEmailDomains",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/sharedEmailDomains/sharedEmailDomainIdValue",
			Expected: &DirectorySharedEmailDomainId{
				SharedEmailDomainId: "sharedEmailDomainIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/sharedEmailDomains/sharedEmailDomainIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectorySharedEmailDomainID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SharedEmailDomainId != v.Expected.SharedEmailDomainId {
			t.Fatalf("Expected %q but got %q for SharedEmailDomainId", v.Expected.SharedEmailDomainId, actual.SharedEmailDomainId)
		}

	}
}

func TestParseDirectorySharedEmailDomainIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectorySharedEmailDomainId
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
			Input: "/directory/sharedEmailDomains",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/sHaReDeMaIlDoMaInS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/sharedEmailDomains/sharedEmailDomainIdValue",
			Expected: &DirectorySharedEmailDomainId{
				SharedEmailDomainId: "sharedEmailDomainIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/sharedEmailDomains/sharedEmailDomainIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/sHaReDeMaIlDoMaInS/sHaReDeMaIlDoMaInIdVaLuE",
			Expected: &DirectorySharedEmailDomainId{
				SharedEmailDomainId: "sHaReDeMaIlDoMaInIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/sHaReDeMaIlDoMaInS/sHaReDeMaIlDoMaInIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectorySharedEmailDomainIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SharedEmailDomainId != v.Expected.SharedEmailDomainId {
			t.Fatalf("Expected %q but got %q for SharedEmailDomainId", v.Expected.SharedEmailDomainId, actual.SharedEmailDomainId)
		}

	}
}

func TestSegmentsForDirectorySharedEmailDomainId(t *testing.T) {
	segments := DirectorySharedEmailDomainId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectorySharedEmailDomainId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
