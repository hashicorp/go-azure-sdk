package settings

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SettingId{}

func TestNewSettingID(t *testing.T) {
	id := NewSettingID("https://endpoint-url.example.com", "settingName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.SettingName != "settingName" {
		t.Fatalf("Expected %q but got %q for Segment 'SettingName'", id.SettingName, "settingName")
	}
}

func TestFormatSettingID(t *testing.T) {
	actual := NewSettingID("https://endpoint-url.example.com", "settingName").ID()
	expected := "https://endpoint-url.example.com/settings/settingName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSettingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SettingId
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
			Input: "https://endpoint-url.example.com/settings",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/settings/settingName",
			Expected: &SettingId{
				BaseURI:     "https://endpoint-url.example.com",
				SettingName: "settingName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/settings/settingName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSettingID(v.Input)
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

		if actual.SettingName != v.Expected.SettingName {
			t.Fatalf("Expected %q but got %q for SettingName", v.Expected.SettingName, actual.SettingName)
		}

	}
}

func TestParseSettingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SettingId
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
			Input: "https://endpoint-url.example.com/settings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sEtTiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/settings/settingName",
			Expected: &SettingId{
				BaseURI:     "https://endpoint-url.example.com",
				SettingName: "settingName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/settings/settingName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sEtTiNgS/sEtTiNgNaMe",
			Expected: &SettingId{
				BaseURI:     "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				SettingName: "sEtTiNgNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sEtTiNgS/sEtTiNgNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSettingIDInsensitively(v.Input)
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

		if actual.SettingName != v.Expected.SettingName {
			t.Fatalf("Expected %q but got %q for SettingName", v.Expected.SettingName, actual.SettingName)
		}

	}
}

func TestSegmentsForSettingId(t *testing.T) {
	segments := SettingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SettingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
