package applications

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApplicationIdId{}

func TestNewApplicationIdID(t *testing.T) {
	id := NewApplicationIdID("applicationIdValue")

	if id.ApplicationId != "applicationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ApplicationId'", id.ApplicationId, "applicationIdValue")
	}
}

func TestFormatApplicationIdID(t *testing.T) {
	actual := NewApplicationIdID("applicationIdValue").ID()
	expected := "/applicationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseApplicationIdID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationIdId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applicationIdValue",
			Expected: &ApplicationIdId{
				ApplicationId: "applicationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applicationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationIdID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ApplicationId != v.Expected.ApplicationId {
			t.Fatalf("Expected %q but got %q for ApplicationId", v.Expected.ApplicationId, actual.ApplicationId)
		}

	}
}

func TestParseApplicationIdIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationIdId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applicationIdValue",
			Expected: &ApplicationIdId{
				ApplicationId: "applicationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applicationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnIdVaLuE",
			Expected: &ApplicationIdId{
				ApplicationId: "aPpLiCaTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationIdIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ApplicationId != v.Expected.ApplicationId {
			t.Fatalf("Expected %q but got %q for ApplicationId", v.Expected.ApplicationId, actual.ApplicationId)
		}

	}
}

func TestSegmentsForApplicationIdId(t *testing.T) {
	segments := ApplicationIdId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ApplicationIdId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
