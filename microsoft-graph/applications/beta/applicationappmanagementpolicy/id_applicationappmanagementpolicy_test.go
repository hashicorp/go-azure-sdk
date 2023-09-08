package applicationappmanagementpolicy

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ApplicationAppManagementPolicyId{}

func TestNewApplicationAppManagementPolicyID(t *testing.T) {
	id := NewApplicationAppManagementPolicyID("applicationIdValue", "appManagementPolicyIdValue")

	if id.ApplicationId != "applicationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ApplicationId'", id.ApplicationId, "applicationIdValue")
	}

	if id.AppManagementPolicyId != "appManagementPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AppManagementPolicyId'", id.AppManagementPolicyId, "appManagementPolicyIdValue")
	}
}

func TestFormatApplicationAppManagementPolicyID(t *testing.T) {
	actual := NewApplicationAppManagementPolicyID("applicationIdValue", "appManagementPolicyIdValue").ID()
	expected := "/applications/applicationIdValue/appManagementPolicies/appManagementPolicyIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseApplicationAppManagementPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationAppManagementPolicyId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue/appManagementPolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationIdValue/appManagementPolicies/appManagementPolicyIdValue",
			Expected: &ApplicationAppManagementPolicyId{
				ApplicationId:         "applicationIdValue",
				AppManagementPolicyId: "appManagementPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationIdValue/appManagementPolicies/appManagementPolicyIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationAppManagementPolicyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ApplicationId != v.Expected.ApplicationId {
			t.Fatalf("Expected %q but got %q for ApplicationId", v.Expected.ApplicationId, actual.ApplicationId)
		}

		if actual.AppManagementPolicyId != v.Expected.AppManagementPolicyId {
			t.Fatalf("Expected %q but got %q for AppManagementPolicyId", v.Expected.AppManagementPolicyId, actual.AppManagementPolicyId)
		}

	}
}

func TestParseApplicationAppManagementPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationAppManagementPolicyId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue/appManagementPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/aPpMaNaGeMeNtPoLiCiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationIdValue/appManagementPolicies/appManagementPolicyIdValue",
			Expected: &ApplicationAppManagementPolicyId{
				ApplicationId:         "applicationIdValue",
				AppManagementPolicyId: "appManagementPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationIdValue/appManagementPolicies/appManagementPolicyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/aPpMaNaGeMeNtPoLiCiEs/aPpMaNaGeMeNtPoLiCyIdVaLuE",
			Expected: &ApplicationAppManagementPolicyId{
				ApplicationId:         "aPpLiCaTiOnIdVaLuE",
				AppManagementPolicyId: "aPpMaNaGeMeNtPoLiCyIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/aPpMaNaGeMeNtPoLiCiEs/aPpMaNaGeMeNtPoLiCyIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationAppManagementPolicyIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ApplicationId != v.Expected.ApplicationId {
			t.Fatalf("Expected %q but got %q for ApplicationId", v.Expected.ApplicationId, actual.ApplicationId)
		}

		if actual.AppManagementPolicyId != v.Expected.AppManagementPolicyId {
			t.Fatalf("Expected %q but got %q for AppManagementPolicyId", v.Expected.AppManagementPolicyId, actual.AppManagementPolicyId)
		}

	}
}

func TestSegmentsForApplicationAppManagementPolicyId(t *testing.T) {
	segments := ApplicationAppManagementPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ApplicationAppManagementPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
