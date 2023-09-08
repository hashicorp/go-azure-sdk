package identityb2xuserflowuserflowidentityprovider

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = IdentityB2xUserFlowId{}

func TestNewIdentityB2xUserFlowID(t *testing.T) {
	id := NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

	if id.B2xIdentityUserFlowId != "b2xIdentityUserFlowIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'B2xIdentityUserFlowId'", id.B2xIdentityUserFlowId, "b2xIdentityUserFlowIdValue")
	}
}

func TestFormatIdentityB2xUserFlowID(t *testing.T) {
	actual := NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue").ID()
	expected := "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityB2xUserFlowID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityB2xUserFlowId
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
			// Valid URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue",
			Expected: &IdentityB2xUserFlowId{
				B2xIdentityUserFlowId: "b2xIdentityUserFlowIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityB2xUserFlowID(v.Input)
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

	}
}

func TestParseIdentityB2xUserFlowIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityB2xUserFlowId
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
			// Valid URI
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue",
			Expected: &IdentityB2xUserFlowId{
				B2xIdentityUserFlowId: "b2xIdentityUserFlowIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identity/b2xUserFlows/b2xIdentityUserFlowIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwIdVaLuE",
			Expected: &IdentityB2xUserFlowId{
				B2xIdentityUserFlowId: "b2xIdEnTiTyUsErFlOwIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItY/b2xUsErFlOwS/b2xIdEnTiTyUsErFlOwIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityB2xUserFlowIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForIdentityB2xUserFlowId(t *testing.T) {
	segments := IdentityB2xUserFlowId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityB2xUserFlowId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
