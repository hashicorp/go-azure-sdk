package skillsets

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SkillsetId{}

func TestNewSkillsetID(t *testing.T) {
	id := NewSkillsetID("https://endpoint-url.example.com", "skillsetName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.SkillsetName != "skillsetName" {
		t.Fatalf("Expected %q but got %q for Segment 'SkillsetName'", id.SkillsetName, "skillsetName")
	}
}

func TestFormatSkillsetID(t *testing.T) {
	actual := NewSkillsetID("https://endpoint-url.example.com", "skillsetName").ID()
	expected := "https://endpoint-url.example.com/skillsets/skillsetName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSkillsetID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SkillsetId
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
			Input: "https://endpoint-url.example.com/skillsets",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/skillsets/skillsetName",
			Expected: &SkillsetId{
				BaseURI:      "https://endpoint-url.example.com",
				SkillsetName: "skillsetName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/skillsets/skillsetName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSkillsetID(v.Input)
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

		if actual.SkillsetName != v.Expected.SkillsetName {
			t.Fatalf("Expected %q but got %q for SkillsetName", v.Expected.SkillsetName, actual.SkillsetName)
		}

	}
}

func TestParseSkillsetIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SkillsetId
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
			Input: "https://endpoint-url.example.com/skillsets",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sKiLlSeTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/skillsets/skillsetName",
			Expected: &SkillsetId{
				BaseURI:      "https://endpoint-url.example.com",
				SkillsetName: "skillsetName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/skillsets/skillsetName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sKiLlSeTs/sKiLlSeTnAmE",
			Expected: &SkillsetId{
				BaseURI:      "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				SkillsetName: "sKiLlSeTnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sKiLlSeTs/sKiLlSeTnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSkillsetIDInsensitively(v.Input)
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

		if actual.SkillsetName != v.Expected.SkillsetName {
			t.Fatalf("Expected %q but got %q for SkillsetName", v.Expected.SkillsetName, actual.SkillsetName)
		}

	}
}

func TestSegmentsForSkillsetId(t *testing.T) {
	segments := SkillsetId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SkillsetId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
