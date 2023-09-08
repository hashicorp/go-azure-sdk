package serviceprincipalsynchronizationjob

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalSynchronizationJobId{}

func TestNewServicePrincipalSynchronizationJobID(t *testing.T) {
	id := NewServicePrincipalSynchronizationJobID("servicePrincipalIdValue", "synchronizationJobIdValue")

	if id.ServicePrincipalId != "servicePrincipalIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalIdValue")
	}

	if id.SynchronizationJobId != "synchronizationJobIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SynchronizationJobId'", id.SynchronizationJobId, "synchronizationJobIdValue")
	}
}

func TestFormatServicePrincipalSynchronizationJobID(t *testing.T) {
	actual := NewServicePrincipalSynchronizationJobID("servicePrincipalIdValue", "synchronizationJobIdValue").ID()
	expected := "/servicePrincipals/servicePrincipalIdValue/synchronization/jobs/synchronizationJobIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServicePrincipalSynchronizationJobID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalSynchronizationJobId
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
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization/jobs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization/jobs/synchronizationJobIdValue",
			Expected: &ServicePrincipalSynchronizationJobId{
				ServicePrincipalId:   "servicePrincipalIdValue",
				SynchronizationJobId: "synchronizationJobIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization/jobs/synchronizationJobIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalSynchronizationJobID(v.Input)
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

func TestParseServicePrincipalSynchronizationJobIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalSynchronizationJobId
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
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/sYnChRoNiZaTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization/jobs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/sYnChRoNiZaTiOn/jObS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization/jobs/synchronizationJobIdValue",
			Expected: &ServicePrincipalSynchronizationJobId{
				ServicePrincipalId:   "servicePrincipalIdValue",
				SynchronizationJobId: "synchronizationJobIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization/jobs/synchronizationJobIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/sYnChRoNiZaTiOn/jObS/sYnChRoNiZaTiOnJoBiDvAlUe",
			Expected: &ServicePrincipalSynchronizationJobId{
				ServicePrincipalId:   "sErViCePrInCiPaLiDvAlUe",
				SynchronizationJobId: "sYnChRoNiZaTiOnJoBiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/sYnChRoNiZaTiOn/jObS/sYnChRoNiZaTiOnJoBiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalSynchronizationJobIDInsensitively(v.Input)
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

func TestSegmentsForServicePrincipalSynchronizationJobId(t *testing.T) {
	segments := ServicePrincipalSynchronizationJobId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServicePrincipalSynchronizationJobId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
