package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdSynchronizationJobId{}

func TestNewServicePrincipalIdSynchronizationJobID(t *testing.T) {
	id := NewServicePrincipalIdSynchronizationJobID("servicePrincipalId", "synchronizationJobId")

	if id.ServicePrincipalId != "servicePrincipalId" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalId")
	}

	if id.SynchronizationJobId != "synchronizationJobId" {
		t.Fatalf("Expected %q but got %q for Segment 'SynchronizationJobId'", id.SynchronizationJobId, "synchronizationJobId")
	}
}

func TestFormatServicePrincipalIdSynchronizationJobID(t *testing.T) {
	actual := NewServicePrincipalIdSynchronizationJobID("servicePrincipalId", "synchronizationJobId").ID()
	expected := "/servicePrincipals/servicePrincipalId/synchronization/jobs/synchronizationJobId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServicePrincipalIdSynchronizationJobID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdSynchronizationJobId
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
			Input: "/servicePrincipals/servicePrincipalId/synchronization",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/synchronization/jobs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/synchronization/jobs/synchronizationJobId",
			Expected: &ServicePrincipalIdSynchronizationJobId{
				ServicePrincipalId:   "servicePrincipalId",
				SynchronizationJobId: "synchronizationJobId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/synchronization/jobs/synchronizationJobId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdSynchronizationJobID(v.Input)
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

		if actual.SynchronizationJobId != v.Expected.SynchronizationJobId {
			t.Fatalf("Expected %q but got %q for SynchronizationJobId", v.Expected.SynchronizationJobId, actual.SynchronizationJobId)
		}

	}
}

func TestParseServicePrincipalIdSynchronizationJobIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdSynchronizationJobId
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
			Input: "/servicePrincipals/servicePrincipalId/synchronization",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/sYnChRoNiZaTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/synchronization/jobs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/sYnChRoNiZaTiOn/jObS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/synchronization/jobs/synchronizationJobId",
			Expected: &ServicePrincipalIdSynchronizationJobId{
				ServicePrincipalId:   "servicePrincipalId",
				SynchronizationJobId: "synchronizationJobId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/synchronization/jobs/synchronizationJobId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/sYnChRoNiZaTiOn/jObS/sYnChRoNiZaTiOnJoBiD",
			Expected: &ServicePrincipalIdSynchronizationJobId{
				ServicePrincipalId:   "sErViCePrInCiPaLiD",
				SynchronizationJobId: "sYnChRoNiZaTiOnJoBiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/sYnChRoNiZaTiOn/jObS/sYnChRoNiZaTiOnJoBiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdSynchronizationJobIDInsensitively(v.Input)
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

		if actual.SynchronizationJobId != v.Expected.SynchronizationJobId {
			t.Fatalf("Expected %q but got %q for SynchronizationJobId", v.Expected.SynchronizationJobId, actual.SynchronizationJobId)
		}

	}
}

func TestSegmentsForServicePrincipalIdSynchronizationJobId(t *testing.T) {
	segments := ServicePrincipalIdSynchronizationJobId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServicePrincipalIdSynchronizationJobId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
