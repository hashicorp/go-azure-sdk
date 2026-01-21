package keys

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeletedkeyId{}

func TestNewDeletedkeyID(t *testing.T) {
	id := NewDeletedkeyID("deletedkeyName")

	if id.DeletedkeyName != "deletedkeyName" {
		t.Fatalf("Expected %q but got %q for Segment 'DeletedkeyName'", id.DeletedkeyName, "deletedkeyName")
	}
}

func TestFormatDeletedkeyID(t *testing.T) {
	actual := NewDeletedkeyID("deletedkeyName").ID()
	expected := "/deletedkeys/deletedkeyName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeletedkeyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeletedkeyId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deletedkeys",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deletedkeys/deletedkeyName",
			Expected: &DeletedkeyId{
				DeletedkeyName: "deletedkeyName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deletedkeys/deletedkeyName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeletedkeyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeletedkeyName != v.Expected.DeletedkeyName {
			t.Fatalf("Expected %q but got %q for DeletedkeyName", v.Expected.DeletedkeyName, actual.DeletedkeyName)
		}

	}
}

func TestParseDeletedkeyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeletedkeyId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deletedkeys",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dElEtEdKeYs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deletedkeys/deletedkeyName",
			Expected: &DeletedkeyId{
				DeletedkeyName: "deletedkeyName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deletedkeys/deletedkeyName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dElEtEdKeYs/dElEtEdKeYnAmE",
			Expected: &DeletedkeyId{
				DeletedkeyName: "dElEtEdKeYnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dElEtEdKeYs/dElEtEdKeYnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeletedkeyIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeletedkeyName != v.Expected.DeletedkeyName {
			t.Fatalf("Expected %q but got %q for DeletedkeyName", v.Expected.DeletedkeyName, actual.DeletedkeyName)
		}

	}
}

func TestSegmentsForDeletedkeyId(t *testing.T) {
	segments := DeletedkeyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeletedkeyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
