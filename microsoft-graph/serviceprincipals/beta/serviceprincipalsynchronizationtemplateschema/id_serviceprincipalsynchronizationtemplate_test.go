package serviceprincipalsynchronizationtemplateschema

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalSynchronizationTemplateId{}

func TestNewServicePrincipalSynchronizationTemplateID(t *testing.T) {
	id := NewServicePrincipalSynchronizationTemplateID("servicePrincipalIdValue", "synchronizationTemplateIdValue")

	if id.ServicePrincipalId != "servicePrincipalIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalIdValue")
	}

	if id.SynchronizationTemplateId != "synchronizationTemplateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SynchronizationTemplateId'", id.SynchronizationTemplateId, "synchronizationTemplateIdValue")
	}
}

func TestFormatServicePrincipalSynchronizationTemplateID(t *testing.T) {
	actual := NewServicePrincipalSynchronizationTemplateID("servicePrincipalIdValue", "synchronizationTemplateIdValue").ID()
	expected := "/servicePrincipals/servicePrincipalIdValue/synchronization/templates/synchronizationTemplateIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServicePrincipalSynchronizationTemplateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalSynchronizationTemplateId
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
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization/templates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization/templates/synchronizationTemplateIdValue",
			Expected: &ServicePrincipalSynchronizationTemplateId{
				ServicePrincipalId:        "servicePrincipalIdValue",
				SynchronizationTemplateId: "synchronizationTemplateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization/templates/synchronizationTemplateIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalSynchronizationTemplateID(v.Input)
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

func TestParseServicePrincipalSynchronizationTemplateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalSynchronizationTemplateId
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
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization/templates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/sYnChRoNiZaTiOn/tEmPlAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization/templates/synchronizationTemplateIdValue",
			Expected: &ServicePrincipalSynchronizationTemplateId{
				ServicePrincipalId:        "servicePrincipalIdValue",
				SynchronizationTemplateId: "synchronizationTemplateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization/templates/synchronizationTemplateIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/sYnChRoNiZaTiOn/tEmPlAtEs/sYnChRoNiZaTiOnTeMpLaTeIdVaLuE",
			Expected: &ServicePrincipalSynchronizationTemplateId{
				ServicePrincipalId:        "sErViCePrInCiPaLiDvAlUe",
				SynchronizationTemplateId: "sYnChRoNiZaTiOnTeMpLaTeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/sYnChRoNiZaTiOn/tEmPlAtEs/sYnChRoNiZaTiOnTeMpLaTeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalSynchronizationTemplateIDInsensitively(v.Input)
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

func TestSegmentsForServicePrincipalSynchronizationTemplateId(t *testing.T) {
	segments := ServicePrincipalSynchronizationTemplateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServicePrincipalSynchronizationTemplateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
