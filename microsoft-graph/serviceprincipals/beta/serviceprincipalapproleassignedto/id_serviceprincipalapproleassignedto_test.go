package serviceprincipalapproleassignedto

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalAppRoleAssignedToId{}

func TestNewServicePrincipalAppRoleAssignedToID(t *testing.T) {
	id := NewServicePrincipalAppRoleAssignedToID("servicePrincipalIdValue", "appRoleAssignmentIdValue")

	if id.ServicePrincipalId != "servicePrincipalIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalIdValue")
	}

	if id.AppRoleAssignmentId != "appRoleAssignmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AppRoleAssignmentId'", id.AppRoleAssignmentId, "appRoleAssignmentIdValue")
	}
}

func TestFormatServicePrincipalAppRoleAssignedToID(t *testing.T) {
	actual := NewServicePrincipalAppRoleAssignedToID("servicePrincipalIdValue", "appRoleAssignmentIdValue").ID()
	expected := "/servicePrincipals/servicePrincipalIdValue/appRoleAssignedTo/appRoleAssignmentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServicePrincipalAppRoleAssignedToID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalAppRoleAssignedToId
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
			Input: "/servicePrincipals/servicePrincipalIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalIdValue/appRoleAssignedTo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalIdValue/appRoleAssignedTo/appRoleAssignmentIdValue",
			Expected: &ServicePrincipalAppRoleAssignedToId{
				ServicePrincipalId:  "servicePrincipalIdValue",
				AppRoleAssignmentId: "appRoleAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalIdValue/appRoleAssignedTo/appRoleAssignmentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalAppRoleAssignedToID(v.Input)
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

func TestParseServicePrincipalAppRoleAssignedToIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalAppRoleAssignedToId
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
			Input: "/servicePrincipals/servicePrincipalIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalIdValue/appRoleAssignedTo",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/aPpRoLeAsSiGnEdTo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalIdValue/appRoleAssignedTo/appRoleAssignmentIdValue",
			Expected: &ServicePrincipalAppRoleAssignedToId{
				ServicePrincipalId:  "servicePrincipalIdValue",
				AppRoleAssignmentId: "appRoleAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalIdValue/appRoleAssignedTo/appRoleAssignmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/aPpRoLeAsSiGnEdTo/aPpRoLeAsSiGnMeNtIdVaLuE",
			Expected: &ServicePrincipalAppRoleAssignedToId{
				ServicePrincipalId:  "sErViCePrInCiPaLiDvAlUe",
				AppRoleAssignmentId: "aPpRoLeAsSiGnMeNtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/aPpRoLeAsSiGnEdTo/aPpRoLeAsSiGnMeNtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalAppRoleAssignedToIDInsensitively(v.Input)
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

func TestSegmentsForServicePrincipalAppRoleAssignedToId(t *testing.T) {
	segments := ServicePrincipalAppRoleAssignedToId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServicePrincipalAppRoleAssignedToId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
