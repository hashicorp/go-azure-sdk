package meapproleassignedresource

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeAppRoleAssignedResourceId{}

func TestNewMeAppRoleAssignedResourceID(t *testing.T) {
	id := NewMeAppRoleAssignedResourceID("servicePrincipalIdValue")

	if id.ServicePrincipalId != "servicePrincipalIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalIdValue")
	}
}

func TestFormatMeAppRoleAssignedResourceID(t *testing.T) {
	actual := NewMeAppRoleAssignedResourceID("servicePrincipalIdValue").ID()
	expected := "/me/appRoleAssignedResources/servicePrincipalIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeAppRoleAssignedResourceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAppRoleAssignedResourceId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/appRoleAssignedResources",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/appRoleAssignedResources/servicePrincipalIdValue",
			Expected: &MeAppRoleAssignedResourceId{
				ServicePrincipalId: "servicePrincipalIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/appRoleAssignedResources/servicePrincipalIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAppRoleAssignedResourceID(v.Input)
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

	}
}

func TestParseMeAppRoleAssignedResourceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAppRoleAssignedResourceId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/appRoleAssignedResources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aPpRoLeAsSiGnEdReSoUrCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/appRoleAssignedResources/servicePrincipalIdValue",
			Expected: &MeAppRoleAssignedResourceId{
				ServicePrincipalId: "servicePrincipalIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/appRoleAssignedResources/servicePrincipalIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aPpRoLeAsSiGnEdReSoUrCeS/sErViCePrInCiPaLiDvAlUe",
			Expected: &MeAppRoleAssignedResourceId{
				ServicePrincipalId: "sErViCePrInCiPaLiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/aPpRoLeAsSiGnEdReSoUrCeS/sErViCePrInCiPaLiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAppRoleAssignedResourceIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForMeAppRoleAssignedResourceId(t *testing.T) {
	segments := MeAppRoleAssignedResourceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeAppRoleAssignedResourceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
