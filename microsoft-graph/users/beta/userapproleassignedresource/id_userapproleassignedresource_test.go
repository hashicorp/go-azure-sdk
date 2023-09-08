package userapproleassignedresource

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserAppRoleAssignedResourceId{}

func TestNewUserAppRoleAssignedResourceID(t *testing.T) {
	id := NewUserAppRoleAssignedResourceID("userIdValue", "servicePrincipalIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.ServicePrincipalId != "servicePrincipalIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalIdValue")
	}
}

func TestFormatUserAppRoleAssignedResourceID(t *testing.T) {
	actual := NewUserAppRoleAssignedResourceID("userIdValue", "servicePrincipalIdValue").ID()
	expected := "/users/userIdValue/appRoleAssignedResources/servicePrincipalIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserAppRoleAssignedResourceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserAppRoleAssignedResourceId
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
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/appRoleAssignedResources",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/appRoleAssignedResources/servicePrincipalIdValue",
			Expected: &UserAppRoleAssignedResourceId{
				UserId:             "userIdValue",
				ServicePrincipalId: "servicePrincipalIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/appRoleAssignedResources/servicePrincipalIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserAppRoleAssignedResourceID(v.Input)
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

func TestParseUserAppRoleAssignedResourceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserAppRoleAssignedResourceId
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
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/appRoleAssignedResources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aPpRoLeAsSiGnEdReSoUrCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/appRoleAssignedResources/servicePrincipalIdValue",
			Expected: &UserAppRoleAssignedResourceId{
				UserId:             "userIdValue",
				ServicePrincipalId: "servicePrincipalIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/appRoleAssignedResources/servicePrincipalIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aPpRoLeAsSiGnEdReSoUrCeS/sErViCePrInCiPaLiDvAlUe",
			Expected: &UserAppRoleAssignedResourceId{
				UserId:             "uSeRiDvAlUe",
				ServicePrincipalId: "sErViCePrInCiPaLiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aPpRoLeAsSiGnEdReSoUrCeS/sErViCePrInCiPaLiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserAppRoleAssignedResourceIDInsensitively(v.Input)
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

func TestSegmentsForUserAppRoleAssignedResourceId(t *testing.T) {
	segments := UserAppRoleAssignedResourceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserAppRoleAssignedResourceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
