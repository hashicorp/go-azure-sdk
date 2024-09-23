package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdAppRoleAssignmentId{}

func TestNewServicePrincipalIdAppRoleAssignmentID(t *testing.T) {
	id := NewServicePrincipalIdAppRoleAssignmentID("servicePrincipalId", "appRoleAssignmentId")

	if id.ServicePrincipalId != "servicePrincipalId" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalId")
	}

	if id.AppRoleAssignmentId != "appRoleAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'AppRoleAssignmentId'", id.AppRoleAssignmentId, "appRoleAssignmentId")
	}
}

func TestFormatServicePrincipalIdAppRoleAssignmentID(t *testing.T) {
	actual := NewServicePrincipalIdAppRoleAssignmentID("servicePrincipalId", "appRoleAssignmentId").ID()
	expected := "/servicePrincipals/servicePrincipalId/appRoleAssignments/appRoleAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServicePrincipalIdAppRoleAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdAppRoleAssignmentId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/appRoleAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/appRoleAssignments/appRoleAssignmentId",
			Expected: &ServicePrincipalIdAppRoleAssignmentId{
				ServicePrincipalId:  "servicePrincipalId",
				AppRoleAssignmentId: "appRoleAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/appRoleAssignments/appRoleAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdAppRoleAssignmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ServicePrincipalId != v.Expected.ServicePrincipalId {
			t.Fatalf("Expected %q but got %q for ServicePrincipalId", v.Expected.ServicePrincipalId, actual.ServicePrincipalId)
		}

		if actual.AppRoleAssignmentId != v.Expected.AppRoleAssignmentId {
			t.Fatalf("Expected %q but got %q for AppRoleAssignmentId", v.Expected.AppRoleAssignmentId, actual.AppRoleAssignmentId)
		}

	}
}

func TestParseServicePrincipalIdAppRoleAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdAppRoleAssignmentId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/appRoleAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/aPpRoLeAsSiGnMeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/appRoleAssignments/appRoleAssignmentId",
			Expected: &ServicePrincipalIdAppRoleAssignmentId{
				ServicePrincipalId:  "servicePrincipalId",
				AppRoleAssignmentId: "appRoleAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/appRoleAssignments/appRoleAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/aPpRoLeAsSiGnMeNtS/aPpRoLeAsSiGnMeNtId",
			Expected: &ServicePrincipalIdAppRoleAssignmentId{
				ServicePrincipalId:  "sErViCePrInCiPaLiD",
				AppRoleAssignmentId: "aPpRoLeAsSiGnMeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/aPpRoLeAsSiGnMeNtS/aPpRoLeAsSiGnMeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdAppRoleAssignmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ServicePrincipalId != v.Expected.ServicePrincipalId {
			t.Fatalf("Expected %q but got %q for ServicePrincipalId", v.Expected.ServicePrincipalId, actual.ServicePrincipalId)
		}

		if actual.AppRoleAssignmentId != v.Expected.AppRoleAssignmentId {
			t.Fatalf("Expected %q but got %q for AppRoleAssignmentId", v.Expected.AppRoleAssignmentId, actual.AppRoleAssignmentId)
		}

	}
}

func TestSegmentsForServicePrincipalIdAppRoleAssignmentId(t *testing.T) {
	segments := ServicePrincipalIdAppRoleAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServicePrincipalIdAppRoleAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
