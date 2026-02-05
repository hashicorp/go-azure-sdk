package storage

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &StorageId{}

func TestNewStorageID(t *testing.T) {
	id := NewStorageID("https://endpoint-url.example.com", "storageName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.StorageName != "storageName" {
		t.Fatalf("Expected %q but got %q for Segment 'StorageName'", id.StorageName, "storageName")
	}
}

func TestFormatStorageID(t *testing.T) {
	actual := NewStorageID("https://endpoint-url.example.com", "storageName").ID()
	expected := "https://endpoint-url.example.com/storage/storageName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseStorageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *StorageId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/storage",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/storage/storageName",
			Expected: &StorageId{
				BaseURI:     "https://endpoint-url.example.com",
				StorageName: "storageName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/storage/storageName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseStorageID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BaseURI != v.Expected.BaseURI {
			t.Fatalf("Expected %q but got %q for BaseURI", v.Expected.BaseURI, actual.BaseURI)
		}

		if actual.StorageName != v.Expected.StorageName {
			t.Fatalf("Expected %q but got %q for StorageName", v.Expected.StorageName, actual.StorageName)
		}

	}
}

func TestParseStorageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *StorageId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/storage",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sToRaGe",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/storage/storageName",
			Expected: &StorageId{
				BaseURI:     "https://endpoint-url.example.com",
				StorageName: "storageName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/storage/storageName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sToRaGe/sToRaGeNaMe",
			Expected: &StorageId{
				BaseURI:     "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				StorageName: "sToRaGeNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sToRaGe/sToRaGeNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseStorageIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BaseURI != v.Expected.BaseURI {
			t.Fatalf("Expected %q but got %q for BaseURI", v.Expected.BaseURI, actual.BaseURI)
		}

		if actual.StorageName != v.Expected.StorageName {
			t.Fatalf("Expected %q but got %q for StorageName", v.Expected.StorageName, actual.StorageName)
		}

	}
}

func TestSegmentsForStorageId(t *testing.T) {
	segments := StorageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("StorageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
