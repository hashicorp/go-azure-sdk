package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityConditionalAccessTemplateId{}

func TestNewIdentityConditionalAccessTemplateID(t *testing.T) {
	id := NewIdentityConditionalAccessTemplateID("conditionalAccessTemplateId")

	if id.ConditionalAccessTemplateId != "conditionalAccessTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'ConditionalAccessTemplateId'", id.ConditionalAccessTemplateId, "conditionalAccessTemplateId")
	}
}

func TestFormatIdentityConditionalAccessTemplateID(t *testing.T) {
	actual := NewIdentityConditionalAccessTemplateID("conditionalAccessTemplateId").ID()
	expected := "/identity/conditionalAccess/templates/conditionalAccessTemplateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityConditionalAccessTemplateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityConditionalAccessTemplateId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/conditionalAccess",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/conditionalAccess/templates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/conditionalAccess/templates/conditionalAccessTemplateId",
			Expected: &IdentityConditionalAccessTemplateId{
				ConditionalAccessTemplateId: "conditionalAccessTemplateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/conditionalAccess/templates/conditionalAccessTemplateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityConditionalAccessTemplateID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ConditionalAccessTemplateId != v.Expected.ConditionalAccessTemplateId {
			t.Fatalf("Expected %q but got %q for ConditionalAccessTemplateId", v.Expected.ConditionalAccessTemplateId, actual.ConditionalAccessTemplateId)
		}

	}
}

func TestParseIdentityConditionalAccessTemplateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityConditionalAccessTemplateId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/conditionalAccess",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/conditionalAccess/templates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/tEmPlAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/conditionalAccess/templates/conditionalAccessTemplateId",
			Expected: &IdentityConditionalAccessTemplateId{
				ConditionalAccessTemplateId: "conditionalAccessTemplateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/conditionalAccess/templates/conditionalAccessTemplateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/tEmPlAtEs/cOnDiTiOnAlAcCeSsTeMpLaTeId",
			Expected: &IdentityConditionalAccessTemplateId{
				ConditionalAccessTemplateId: "cOnDiTiOnAlAcCeSsTeMpLaTeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/tEmPlAtEs/cOnDiTiOnAlAcCeSsTeMpLaTeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityConditionalAccessTemplateIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ConditionalAccessTemplateId != v.Expected.ConditionalAccessTemplateId {
			t.Fatalf("Expected %q but got %q for ConditionalAccessTemplateId", v.Expected.ConditionalAccessTemplateId, actual.ConditionalAccessTemplateId)
		}

	}
}

func TestSegmentsForIdentityConditionalAccessTemplateId(t *testing.T) {
	segments := IdentityConditionalAccessTemplateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityConditionalAccessTemplateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
