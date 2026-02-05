package deletedstorage

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeletedstorageSaId{}

func TestNewDeletedstorageSaID(t *testing.T) {
	id := NewDeletedstorageSaID("https://endpoint-url.example.com", "deletedstorageName", "saName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.DeletedstorageName != "deletedstorageName" {
		t.Fatalf("Expected %q but got %q for Segment 'DeletedstorageName'", id.DeletedstorageName, "deletedstorageName")
	}

	if id.SaName != "saName" {
		t.Fatalf("Expected %q but got %q for Segment 'SaName'", id.SaName, "saName")
	}
}

func TestFormatDeletedstorageSaID(t *testing.T) {
	actual := NewDeletedstorageSaID("https://endpoint-url.example.com", "deletedstorageName", "saName").ID()
	expected := "https://endpoint-url.example.com/deletedstorage/deletedstorageName/sas/saName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeletedstorageSaID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeletedstorageSaId
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
			Input: "https://endpoint-url.example.com/deletedstorage",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/deletedstorage/deletedstorageName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/deletedstorage/deletedstorageName/sas",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/deletedstorage/deletedstorageName/sas/saName",
			Expected: &DeletedstorageSaId{
				BaseURI:            "https://endpoint-url.example.com",
				DeletedstorageName: "deletedstorageName",
				SaName:             "saName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/deletedstorage/deletedstorageName/sas/saName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeletedstorageSaID(v.Input)
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

		if actual.DeletedstorageName != v.Expected.DeletedstorageName {
			t.Fatalf("Expected %q but got %q for DeletedstorageName", v.Expected.DeletedstorageName, actual.DeletedstorageName)
		}

		if actual.SaName != v.Expected.SaName {
			t.Fatalf("Expected %q but got %q for SaName", v.Expected.SaName, actual.SaName)
		}

	}
}

func TestParseDeletedstorageSaIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeletedstorageSaId
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
			Input: "https://endpoint-url.example.com/deletedstorage",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dElEtEdStOrAgE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/deletedstorage/deletedstorageName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dElEtEdStOrAgE/dElEtEdStOrAgEnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/deletedstorage/deletedstorageName/sas",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dElEtEdStOrAgE/dElEtEdStOrAgEnAmE/sAs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/deletedstorage/deletedstorageName/sas/saName",
			Expected: &DeletedstorageSaId{
				BaseURI:            "https://endpoint-url.example.com",
				DeletedstorageName: "deletedstorageName",
				SaName:             "saName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/deletedstorage/deletedstorageName/sas/saName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dElEtEdStOrAgE/dElEtEdStOrAgEnAmE/sAs/sAnAmE",
			Expected: &DeletedstorageSaId{
				BaseURI:            "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				DeletedstorageName: "dElEtEdStOrAgEnAmE",
				SaName:             "sAnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dElEtEdStOrAgE/dElEtEdStOrAgEnAmE/sAs/sAnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeletedstorageSaIDInsensitively(v.Input)
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

		if actual.DeletedstorageName != v.Expected.DeletedstorageName {
			t.Fatalf("Expected %q but got %q for DeletedstorageName", v.Expected.DeletedstorageName, actual.DeletedstorageName)
		}

		if actual.SaName != v.Expected.SaName {
			t.Fatalf("Expected %q but got %q for SaName", v.Expected.SaName, actual.SaName)
		}

	}
}

func TestSegmentsForDeletedstorageSaId(t *testing.T) {
	segments := DeletedstorageSaId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeletedstorageSaId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
