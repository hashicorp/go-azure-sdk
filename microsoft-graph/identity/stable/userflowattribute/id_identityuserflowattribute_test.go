package userflowattribute

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityUserFlowAttributeId{}

func TestNewIdentityUserFlowAttributeID(t *testing.T) {
	id := NewIdentityUserFlowAttributeID("identityUserFlowAttributeIdValue")

	if id.IdentityUserFlowAttributeId != "identityUserFlowAttributeIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'IdentityUserFlowAttributeId'", id.IdentityUserFlowAttributeId, "identityUserFlowAttributeIdValue")
	}
}

func TestFormatIdentityUserFlowAttributeID(t *testing.T) {
	actual := NewIdentityUserFlowAttributeID("identityUserFlowAttributeIdValue").ID()
	expected := "/identity/userFlowAttributes/identityUserFlowAttributeIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityUserFlowAttributeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityUserFlowAttributeId
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
			Input: "/identity/userFlowAttributes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/userFlowAttributes/identityUserFlowAttributeIdValue",
			Expected: &IdentityUserFlowAttributeId{
				IdentityUserFlowAttributeId: "identityUserFlowAttributeIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/userFlowAttributes/identityUserFlowAttributeIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityUserFlowAttributeID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.IdentityUserFlowAttributeId != v.Expected.IdentityUserFlowAttributeId {
			t.Fatalf("Expected %q but got %q for IdentityUserFlowAttributeId", v.Expected.IdentityUserFlowAttributeId, actual.IdentityUserFlowAttributeId)
		}

	}
}

func TestParseIdentityUserFlowAttributeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityUserFlowAttributeId
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
			Input: "/identity/userFlowAttributes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/uSeRfLoWaTtRiBuTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/userFlowAttributes/identityUserFlowAttributeIdValue",
			Expected: &IdentityUserFlowAttributeId{
				IdentityUserFlowAttributeId: "identityUserFlowAttributeIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/userFlowAttributes/identityUserFlowAttributeIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/uSeRfLoWaTtRiBuTeS/iDeNtItYuSeRfLoWaTtRiBuTeIdVaLuE",
			Expected: &IdentityUserFlowAttributeId{
				IdentityUserFlowAttributeId: "iDeNtItYuSeRfLoWaTtRiBuTeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/uSeRfLoWaTtRiBuTeS/iDeNtItYuSeRfLoWaTtRiBuTeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityUserFlowAttributeIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.IdentityUserFlowAttributeId != v.Expected.IdentityUserFlowAttributeId {
			t.Fatalf("Expected %q but got %q for IdentityUserFlowAttributeId", v.Expected.IdentityUserFlowAttributeId, actual.IdentityUserFlowAttributeId)
		}

	}
}

func TestSegmentsForIdentityUserFlowAttributeId(t *testing.T) {
	segments := IdentityUserFlowAttributeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityUserFlowAttributeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
