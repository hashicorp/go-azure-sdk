package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApplicationIdAppManagementPolicyId{}

func TestNewApplicationIdAppManagementPolicyID(t *testing.T) {
	id := NewApplicationIdAppManagementPolicyID("applicationId", "appManagementPolicyId")

	if id.ApplicationId != "applicationId" {
		t.Fatalf("Expected %q but got %q for Segment 'ApplicationId'", id.ApplicationId, "applicationId")
	}

	if id.AppManagementPolicyId != "appManagementPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'AppManagementPolicyId'", id.AppManagementPolicyId, "appManagementPolicyId")
	}
}

func TestFormatApplicationIdAppManagementPolicyID(t *testing.T) {
	actual := NewApplicationIdAppManagementPolicyID("applicationId", "appManagementPolicyId").ID()
	expected := "/applications/applicationId/appManagementPolicies/appManagementPolicyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseApplicationIdAppManagementPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationIdAppManagementPolicyId
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
			Input: "/applications/applicationId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationId/appManagementPolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationId/appManagementPolicies/appManagementPolicyId",
			Expected: &ApplicationIdAppManagementPolicyId{
				ApplicationId:         "applicationId",
				AppManagementPolicyId: "appManagementPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationId/appManagementPolicies/appManagementPolicyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationIdAppManagementPolicyID(v.Input)
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

func TestParseApplicationIdAppManagementPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationIdAppManagementPolicyId
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
			Input: "/applications/applicationId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationId/appManagementPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/aPpMaNaGeMeNtPoLiCiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationId/appManagementPolicies/appManagementPolicyId",
			Expected: &ApplicationIdAppManagementPolicyId{
				ApplicationId:         "applicationId",
				AppManagementPolicyId: "appManagementPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationId/appManagementPolicies/appManagementPolicyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/aPpMaNaGeMeNtPoLiCiEs/aPpMaNaGeMeNtPoLiCyId",
			Expected: &ApplicationIdAppManagementPolicyId{
				ApplicationId:         "aPpLiCaTiOnId",
				AppManagementPolicyId: "aPpMaNaGeMeNtPoLiCyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/aPpMaNaGeMeNtPoLiCiEs/aPpMaNaGeMeNtPoLiCyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationIdAppManagementPolicyIDInsensitively(v.Input)
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

func TestSegmentsForApplicationIdAppManagementPolicyId(t *testing.T) {
	segments := ApplicationIdAppManagementPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ApplicationIdAppManagementPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
