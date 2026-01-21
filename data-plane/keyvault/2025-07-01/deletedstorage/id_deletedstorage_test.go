package deletedstorage

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeletedstorageId{}

func TestNewDeletedstorageID(t *testing.T) {
	id := NewDeletedstorageID("deletedstorageName")

	if id.DeletedstorageName != "deletedstorageName" {
		t.Fatalf("Expected %q but got %q for Segment 'DeletedstorageName'", id.DeletedstorageName, "deletedstorageName")
	}
}

func TestFormatDeletedstorageID(t *testing.T) {
	actual := NewDeletedstorageID("deletedstorageName").ID()
	expected := "/deletedstorage/deletedstorageName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeletedstorageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeletedstorageId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deletedstorage",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deletedstorage/deletedstorageName",
			Expected: &DeletedstorageId{
				DeletedstorageName: "deletedstorageName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deletedstorage/deletedstorageName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeletedstorageID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeletedstorageName != v.Expected.DeletedstorageName {
			t.Fatalf("Expected %q but got %q for DeletedstorageName", v.Expected.DeletedstorageName, actual.DeletedstorageName)
		}

	}
}

func TestParseDeletedstorageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeletedstorageId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deletedstorage",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dElEtEdStOrAgE",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deletedstorage/deletedstorageName",
			Expected: &DeletedstorageId{
				DeletedstorageName: "deletedstorageName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deletedstorage/deletedstorageName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dElEtEdStOrAgE/dElEtEdStOrAgEnAmE",
			Expected: &DeletedstorageId{
				DeletedstorageName: "dElEtEdStOrAgEnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dElEtEdStOrAgE/dElEtEdStOrAgEnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeletedstorageIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeletedstorageName != v.Expected.DeletedstorageName {
			t.Fatalf("Expected %q but got %q for DeletedstorageName", v.Expected.DeletedstorageName, actual.DeletedstorageName)
		}

	}
}

func TestSegmentsForDeletedstorageId(t *testing.T) {
	segments := DeletedstorageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeletedstorageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
