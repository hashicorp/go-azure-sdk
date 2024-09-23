package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdAppRoleAssignedToId{}

func TestNewServicePrincipalIdAppRoleAssignedToID(t *testing.T) {
	id := NewServicePrincipalIdAppRoleAssignedToID("servicePrincipalId", "appRoleAssignmentId")

	if id.ServicePrincipalId != "servicePrincipalId" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalId")
	}

	if id.AppRoleAssignmentId != "appRoleAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'AppRoleAssignmentId'", id.AppRoleAssignmentId, "appRoleAssignmentId")
	}
}

func TestFormatServicePrincipalIdAppRoleAssignedToID(t *testing.T) {
	actual := NewServicePrincipalIdAppRoleAssignedToID("servicePrincipalId", "appRoleAssignmentId").ID()
	expected := "/servicePrincipals/servicePrincipalId/appRoleAssignedTo/appRoleAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServicePrincipalIdAppRoleAssignedToID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdAppRoleAssignedToId
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
			Input: "/servicePrincipals/servicePrincipalId/appRoleAssignedTo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/appRoleAssignedTo/appRoleAssignmentId",
			Expected: &ServicePrincipalIdAppRoleAssignedToId{
				ServicePrincipalId:  "servicePrincipalId",
				AppRoleAssignmentId: "appRoleAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/appRoleAssignedTo/appRoleAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdAppRoleAssignedToID(v.Input)
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

func TestParseServicePrincipalIdAppRoleAssignedToIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdAppRoleAssignedToId
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
			Input: "/servicePrincipals/servicePrincipalId/appRoleAssignedTo",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/aPpRoLeAsSiGnEdTo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/appRoleAssignedTo/appRoleAssignmentId",
			Expected: &ServicePrincipalIdAppRoleAssignedToId{
				ServicePrincipalId:  "servicePrincipalId",
				AppRoleAssignmentId: "appRoleAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/appRoleAssignedTo/appRoleAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/aPpRoLeAsSiGnEdTo/aPpRoLeAsSiGnMeNtId",
			Expected: &ServicePrincipalIdAppRoleAssignedToId{
				ServicePrincipalId:  "sErViCePrInCiPaLiD",
				AppRoleAssignmentId: "aPpRoLeAsSiGnMeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/aPpRoLeAsSiGnEdTo/aPpRoLeAsSiGnMeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdAppRoleAssignedToIDInsensitively(v.Input)
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

func TestSegmentsForServicePrincipalIdAppRoleAssignedToId(t *testing.T) {
	segments := ServicePrincipalIdAppRoleAssignedToId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServicePrincipalIdAppRoleAssignedToId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
