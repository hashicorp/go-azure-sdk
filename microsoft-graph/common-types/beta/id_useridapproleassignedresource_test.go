package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAppRoleAssignedResourceId{}

func TestNewUserIdAppRoleAssignedResourceID(t *testing.T) {
	id := NewUserIdAppRoleAssignedResourceID("userId", "servicePrincipalId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.ServicePrincipalId != "servicePrincipalId" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalId")
	}
}

func TestFormatUserIdAppRoleAssignedResourceID(t *testing.T) {
	actual := NewUserIdAppRoleAssignedResourceID("userId", "servicePrincipalId").ID()
	expected := "/users/userId/appRoleAssignedResources/servicePrincipalId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdAppRoleAssignedResourceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAppRoleAssignedResourceId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/appRoleAssignedResources",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/appRoleAssignedResources/servicePrincipalId",
			Expected: &UserIdAppRoleAssignedResourceId{
				UserId:             "userId",
				ServicePrincipalId: "servicePrincipalId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/appRoleAssignedResources/servicePrincipalId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAppRoleAssignedResourceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.ServicePrincipalId != v.Expected.ServicePrincipalId {
			t.Fatalf("Expected %q but got %q for ServicePrincipalId", v.Expected.ServicePrincipalId, actual.ServicePrincipalId)
		}

	}
}

func TestParseUserIdAppRoleAssignedResourceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAppRoleAssignedResourceId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/appRoleAssignedResources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aPpRoLeAsSiGnEdReSoUrCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/appRoleAssignedResources/servicePrincipalId",
			Expected: &UserIdAppRoleAssignedResourceId{
				UserId:             "userId",
				ServicePrincipalId: "servicePrincipalId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/appRoleAssignedResources/servicePrincipalId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aPpRoLeAsSiGnEdReSoUrCeS/sErViCePrInCiPaLiD",
			Expected: &UserIdAppRoleAssignedResourceId{
				UserId:             "uSeRiD",
				ServicePrincipalId: "sErViCePrInCiPaLiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aPpRoLeAsSiGnEdReSoUrCeS/sErViCePrInCiPaLiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAppRoleAssignedResourceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.ServicePrincipalId != v.Expected.ServicePrincipalId {
			t.Fatalf("Expected %q but got %q for ServicePrincipalId", v.Expected.ServicePrincipalId, actual.ServicePrincipalId)
		}

	}
}

func TestSegmentsForUserIdAppRoleAssignedResourceId(t *testing.T) {
	segments := UserIdAppRoleAssignedResourceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdAppRoleAssignedResourceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
