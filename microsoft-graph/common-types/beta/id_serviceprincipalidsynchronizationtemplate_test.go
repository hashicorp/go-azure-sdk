package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdSynchronizationTemplateId{}

func TestNewServicePrincipalIdSynchronizationTemplateID(t *testing.T) {
	id := NewServicePrincipalIdSynchronizationTemplateID("servicePrincipalId", "synchronizationTemplateId")

	if id.ServicePrincipalId != "servicePrincipalId" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalId")
	}

	if id.SynchronizationTemplateId != "synchronizationTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'SynchronizationTemplateId'", id.SynchronizationTemplateId, "synchronizationTemplateId")
	}
}

func TestFormatServicePrincipalIdSynchronizationTemplateID(t *testing.T) {
	actual := NewServicePrincipalIdSynchronizationTemplateID("servicePrincipalId", "synchronizationTemplateId").ID()
	expected := "/servicePrincipals/servicePrincipalId/synchronization/templates/synchronizationTemplateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServicePrincipalIdSynchronizationTemplateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdSynchronizationTemplateId
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
			Input: "/servicePrincipals/servicePrincipalId/synchronization/templates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/synchronization/templates/synchronizationTemplateId",
			Expected: &ServicePrincipalIdSynchronizationTemplateId{
				ServicePrincipalId:        "servicePrincipalId",
				SynchronizationTemplateId: "synchronizationTemplateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/synchronization/templates/synchronizationTemplateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdSynchronizationTemplateID(v.Input)
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

		if actual.SynchronizationTemplateId != v.Expected.SynchronizationTemplateId {
			t.Fatalf("Expected %q but got %q for SynchronizationTemplateId", v.Expected.SynchronizationTemplateId, actual.SynchronizationTemplateId)
		}

	}
}

func TestParseServicePrincipalIdSynchronizationTemplateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdSynchronizationTemplateId
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
			Input: "/servicePrincipals/servicePrincipalId/synchronization/templates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/sYnChRoNiZaTiOn/tEmPlAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/synchronization/templates/synchronizationTemplateId",
			Expected: &ServicePrincipalIdSynchronizationTemplateId{
				ServicePrincipalId:        "servicePrincipalId",
				SynchronizationTemplateId: "synchronizationTemplateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/synchronization/templates/synchronizationTemplateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/sYnChRoNiZaTiOn/tEmPlAtEs/sYnChRoNiZaTiOnTeMpLaTeId",
			Expected: &ServicePrincipalIdSynchronizationTemplateId{
				ServicePrincipalId:        "sErViCePrInCiPaLiD",
				SynchronizationTemplateId: "sYnChRoNiZaTiOnTeMpLaTeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/sYnChRoNiZaTiOn/tEmPlAtEs/sYnChRoNiZaTiOnTeMpLaTeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdSynchronizationTemplateIDInsensitively(v.Input)
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

		if actual.SynchronizationTemplateId != v.Expected.SynchronizationTemplateId {
			t.Fatalf("Expected %q but got %q for SynchronizationTemplateId", v.Expected.SynchronizationTemplateId, actual.SynchronizationTemplateId)
		}

	}
}

func TestSegmentsForServicePrincipalIdSynchronizationTemplateId(t *testing.T) {
	segments := ServicePrincipalIdSynchronizationTemplateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServicePrincipalIdSynchronizationTemplateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
