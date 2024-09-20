package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityB2xUserFlowIdUserAttributeAssignmentId{}

func TestNewIdentityB2xUserFlowIdUserAttributeAssignmentID(t *testing.T) {
	id := NewIdentityB2xUserFlowIdUserAttributeAssignmentID("b2xIdentityUserFlowId", "identityUserFlowAttributeAssignmentId")

	if id.B2xIdentityUserFlowId != "b2xIdentityUserFlowId" {
		t.Fatalf("Expected %q but got %q for Segment 'B2xIdentityUserFlowId'", id.B2xIdentityUserFlowId, "b2xIdentityUserFlowId")
	}

	if id.IdentityUserFlowAttributeAssignmentId != "identityUserFlowAttributeAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'IdentityUserFlowAttributeAssignmentId'", id.IdentityUserFlowAttributeAssignmentId, "identityUserFlowAttributeAssignmentId")
	}
}

func TestFormatIdentityB2xUserFlowIdUserAttributeAssignmentID(t *testing.T) {
	actual := NewIdentityB2xUserFlowIdUserAttributeAssignmentID("b2xIdentityUserFlowId", "identityUserFlowAttributeAssignmentId").ID()
	expected := "/identity/b2xUserFlows/b2xIdentityUserFlowId/userAttributeAssignments/identityUserFlowAttributeAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityB2xUserFlowIdUserAttributeAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityB2xUserFlowIdUserAttributeAssignmentId
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
			Input: "/identity/b2xUserFlows",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/userAttributeAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/userAttributeAssignments/identityUserFlowAttributeAssignmentId",
			Expected: &IdentityB2xUserFlowIdUserAttributeAssignmentId{
				B2xIdentityUserFlowId:                 "b2xIdentityUserFlowId",
				IdentityUserFlowAttributeAssignmentId: "identityUserFlowAttributeAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/userAttributeAssignments/identityUserFlowAttributeAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityB2xUserFlowIdUserAttributeAssignmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.B2xIdentityUserFlowId != v.Expected.B2xIdentityUserFlowId {
			t.Fatalf("Expected %q but got %q for B2xIdentityUserFlowId", v.Expected.B2xIdentityUserFlowId, actual.B2xIdentityUserFlowId)
		}

		if actual.IdentityUserFlowAttributeAssignmentId != v.Expected.IdentityUserFlowAttributeAssignmentId {
			t.Fatalf("Expected %q but got %q for IdentityUserFlowAttributeAssignmentId", v.Expected.IdentityUserFlowAttributeAssignmentId, actual.IdentityUserFlowAttributeAssignmentId)
		}

	}
}

func TestParseIdentityB2xUserFlowIdUserAttributeAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityB2xUserFlowIdUserAttributeAssignmentId
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
			Input: "/identity/b2xUserFlows",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/userAttributeAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwId/uSeRaTtRiBuTeAsSiGnMeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/userAttributeAssignments/identityUserFlowAttributeAssignmentId",
			Expected: &IdentityB2xUserFlowIdUserAttributeAssignmentId{
				B2xIdentityUserFlowId:                 "b2xIdentityUserFlowId",
				IdentityUserFlowAttributeAssignmentId: "identityUserFlowAttributeAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowId/userAttributeAssignments/identityUserFlowAttributeAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwId/uSeRaTtRiBuTeAsSiGnMeNtS/iDeNtItYuSeRfLoWaTtRiBuTeAsSiGnMeNtId",
			Expected: &IdentityB2xUserFlowIdUserAttributeAssignmentId{
				B2xIdentityUserFlowId:                 "b2xIdEnTiTyUsErFlOwId",
				IdentityUserFlowAttributeAssignmentId: "iDeNtItYuSeRfLoWaTtRiBuTeAsSiGnMeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwId/uSeRaTtRiBuTeAsSiGnMeNtS/iDeNtItYuSeRfLoWaTtRiBuTeAsSiGnMeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityB2xUserFlowIdUserAttributeAssignmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.B2xIdentityUserFlowId != v.Expected.B2xIdentityUserFlowId {
			t.Fatalf("Expected %q but got %q for B2xIdentityUserFlowId", v.Expected.B2xIdentityUserFlowId, actual.B2xIdentityUserFlowId)
		}

		if actual.IdentityUserFlowAttributeAssignmentId != v.Expected.IdentityUserFlowAttributeAssignmentId {
			t.Fatalf("Expected %q but got %q for IdentityUserFlowAttributeAssignmentId", v.Expected.IdentityUserFlowAttributeAssignmentId, actual.IdentityUserFlowAttributeAssignmentId)
		}

	}
}

func TestSegmentsForIdentityB2xUserFlowIdUserAttributeAssignmentId(t *testing.T) {
	segments := IdentityB2xUserFlowIdUserAttributeAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityB2xUserFlowIdUserAttributeAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
