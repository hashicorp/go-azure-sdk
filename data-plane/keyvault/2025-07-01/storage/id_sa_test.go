package storage

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SaId{}

func TestNewSaID(t *testing.T) {
	id := NewSaID("storageName", "saName")

	if id.StorageName != "storageName" {
		t.Fatalf("Expected %q but got %q for Segment 'StorageName'", id.StorageName, "storageName")
	}

	if id.SaName != "saName" {
		t.Fatalf("Expected %q but got %q for Segment 'SaName'", id.SaName, "saName")
	}
}

func TestFormatSaID(t *testing.T) {
	actual := NewSaID("storageName", "saName").ID()
	expected := "/storage/storageName/sas/saName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSaID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SaId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/storage",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/storage/storageName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/storage/storageName/sas",
			Error: true,
		},
		{
			// Valid URI
			Input: "/storage/storageName/sas/saName",
			Expected: &SaId{
				StorageName: "storageName",
				SaName:      "saName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/storage/storageName/sas/saName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSaID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.StorageName != v.Expected.StorageName {
			t.Fatalf("Expected %q but got %q for StorageName", v.Expected.StorageName, actual.StorageName)
		}

		if actual.SaName != v.Expected.SaName {
			t.Fatalf("Expected %q but got %q for SaName", v.Expected.SaName, actual.SaName)
		}

	}
}

func TestParseSaIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SaId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/storage",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sToRaGe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/storage/storageName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sToRaGe/sToRaGeNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/storage/storageName/sas",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sToRaGe/sToRaGeNaMe/sAs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/storage/storageName/sas/saName",
			Expected: &SaId{
				StorageName: "storageName",
				SaName:      "saName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/storage/storageName/sas/saName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sToRaGe/sToRaGeNaMe/sAs/sAnAmE",
			Expected: &SaId{
				StorageName: "sToRaGeNaMe",
				SaName:      "sAnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sToRaGe/sToRaGeNaMe/sAs/sAnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSaIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.StorageName != v.Expected.StorageName {
			t.Fatalf("Expected %q but got %q for StorageName", v.Expected.StorageName, actual.StorageName)
		}

		if actual.SaName != v.Expected.SaName {
			t.Fatalf("Expected %q but got %q for SaName", v.Expected.SaName, actual.SaName)
		}

	}
}

func TestSegmentsForSaId(t *testing.T) {
	segments := SaId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SaId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
