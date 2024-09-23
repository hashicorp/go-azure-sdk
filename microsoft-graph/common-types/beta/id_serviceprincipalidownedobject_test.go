package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdOwnedObjectId{}

func TestNewServicePrincipalIdOwnedObjectID(t *testing.T) {
	id := NewServicePrincipalIdOwnedObjectID("servicePrincipalId", "directoryObjectId")

	if id.ServicePrincipalId != "servicePrincipalId" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalId")
	}

	if id.DirectoryObjectId != "directoryObjectId" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectId")
	}
}

func TestFormatServicePrincipalIdOwnedObjectID(t *testing.T) {
	actual := NewServicePrincipalIdOwnedObjectID("servicePrincipalId", "directoryObjectId").ID()
	expected := "/servicePrincipals/servicePrincipalId/ownedObjects/directoryObjectId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServicePrincipalIdOwnedObjectID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdOwnedObjectId
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
			Input: "/servicePrincipals/servicePrincipalId/ownedObjects",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/ownedObjects/directoryObjectId",
			Expected: &ServicePrincipalIdOwnedObjectId{
				ServicePrincipalId: "servicePrincipalId",
				DirectoryObjectId:  "directoryObjectId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/ownedObjects/directoryObjectId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdOwnedObjectID(v.Input)
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

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestParseServicePrincipalIdOwnedObjectIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdOwnedObjectId
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
			Input: "/servicePrincipals/servicePrincipalId/ownedObjects",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/oWnEdObJeCtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/ownedObjects/directoryObjectId",
			Expected: &ServicePrincipalIdOwnedObjectId{
				ServicePrincipalId: "servicePrincipalId",
				DirectoryObjectId:  "directoryObjectId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/ownedObjects/directoryObjectId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/oWnEdObJeCtS/dIrEcToRyObJeCtId",
			Expected: &ServicePrincipalIdOwnedObjectId{
				ServicePrincipalId: "sErViCePrInCiPaLiD",
				DirectoryObjectId:  "dIrEcToRyObJeCtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/oWnEdObJeCtS/dIrEcToRyObJeCtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdOwnedObjectIDInsensitively(v.Input)
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

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestSegmentsForServicePrincipalIdOwnedObjectId(t *testing.T) {
	segments := ServicePrincipalIdOwnedObjectId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServicePrincipalIdOwnedObjectId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
