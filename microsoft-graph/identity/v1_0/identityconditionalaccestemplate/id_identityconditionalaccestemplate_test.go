package identityconditionalaccestemplate

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = IdentityConditionalAccesTemplateId{}

func TestNewIdentityConditionalAccesTemplateID(t *testing.T) {
	id := NewIdentityConditionalAccesTemplateID("conditionalAccessTemplateIdValue")

	if id.ConditionalAccessTemplateId != "conditionalAccessTemplateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ConditionalAccessTemplateId'", id.ConditionalAccessTemplateId, "conditionalAccessTemplateIdValue")
	}
}

func TestFormatIdentityConditionalAccesTemplateID(t *testing.T) {
	actual := NewIdentityConditionalAccesTemplateID("conditionalAccessTemplateIdValue").ID()
	expected := "/identity/conditionalAccess/templates/conditionalAccessTemplateIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityConditionalAccesTemplateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityConditionalAccesTemplateId
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
			Input: "/identity/conditionalAccess/templates/conditionalAccessTemplateIdValue",
			Expected: &IdentityConditionalAccesTemplateId{
				ConditionalAccessTemplateId: "conditionalAccessTemplateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/conditionalAccess/templates/conditionalAccessTemplateIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityConditionalAccesTemplateID(v.Input)
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

func TestParseIdentityConditionalAccesTemplateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityConditionalAccesTemplateId
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
			Input: "/identity/conditionalAccess/templates/conditionalAccessTemplateIdValue",
			Expected: &IdentityConditionalAccesTemplateId{
				ConditionalAccessTemplateId: "conditionalAccessTemplateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/conditionalAccess/templates/conditionalAccessTemplateIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/tEmPlAtEs/cOnDiTiOnAlAcCeSsTeMpLaTeIdVaLuE",
			Expected: &IdentityConditionalAccesTemplateId{
				ConditionalAccessTemplateId: "cOnDiTiOnAlAcCeSsTeMpLaTeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/cOnDiTiOnAlAcCeSs/tEmPlAtEs/cOnDiTiOnAlAcCeSsTeMpLaTeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityConditionalAccesTemplateIDInsensitively(v.Input)
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

func TestSegmentsForIdentityConditionalAccesTemplateId(t *testing.T) {
	segments := IdentityConditionalAccesTemplateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityConditionalAccesTemplateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
