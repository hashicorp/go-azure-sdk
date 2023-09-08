package applicationsynchronizationtemplateschemadirectory

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ApplicationSynchronizationTemplateId{}

func TestNewApplicationSynchronizationTemplateID(t *testing.T) {
	id := NewApplicationSynchronizationTemplateID("applicationIdValue", "synchronizationTemplateIdValue")

	if id.ApplicationId != "applicationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ApplicationId'", id.ApplicationId, "applicationIdValue")
	}

	if id.SynchronizationTemplateId != "synchronizationTemplateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SynchronizationTemplateId'", id.SynchronizationTemplateId, "synchronizationTemplateIdValue")
	}
}

func TestFormatApplicationSynchronizationTemplateID(t *testing.T) {
	actual := NewApplicationSynchronizationTemplateID("applicationIdValue", "synchronizationTemplateIdValue").ID()
	expected := "/applications/applicationIdValue/synchronization/templates/synchronizationTemplateIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseApplicationSynchronizationTemplateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationSynchronizationTemplateId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue/synchronization",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue/synchronization/templates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationIdValue/synchronization/templates/synchronizationTemplateIdValue",
			Expected: &ApplicationSynchronizationTemplateId{
				ApplicationId:             "applicationIdValue",
				SynchronizationTemplateId: "synchronizationTemplateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationIdValue/synchronization/templates/synchronizationTemplateIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationSynchronizationTemplateID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ApplicationId != v.Expected.ApplicationId {
			t.Fatalf("Expected %q but got %q for ApplicationId", v.Expected.ApplicationId, actual.ApplicationId)
		}

		if actual.SynchronizationTemplateId != v.Expected.SynchronizationTemplateId {
			t.Fatalf("Expected %q but got %q for SynchronizationTemplateId", v.Expected.SynchronizationTemplateId, actual.SynchronizationTemplateId)
		}

	}
}

func TestParseApplicationSynchronizationTemplateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationSynchronizationTemplateId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue/synchronization",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/sYnChRoNiZaTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue/synchronization/templates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/sYnChRoNiZaTiOn/tEmPlAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationIdValue/synchronization/templates/synchronizationTemplateIdValue",
			Expected: &ApplicationSynchronizationTemplateId{
				ApplicationId:             "applicationIdValue",
				SynchronizationTemplateId: "synchronizationTemplateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationIdValue/synchronization/templates/synchronizationTemplateIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/sYnChRoNiZaTiOn/tEmPlAtEs/sYnChRoNiZaTiOnTeMpLaTeIdVaLuE",
			Expected: &ApplicationSynchronizationTemplateId{
				ApplicationId:             "aPpLiCaTiOnIdVaLuE",
				SynchronizationTemplateId: "sYnChRoNiZaTiOnTeMpLaTeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/sYnChRoNiZaTiOn/tEmPlAtEs/sYnChRoNiZaTiOnTeMpLaTeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationSynchronizationTemplateIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ApplicationId != v.Expected.ApplicationId {
			t.Fatalf("Expected %q but got %q for ApplicationId", v.Expected.ApplicationId, actual.ApplicationId)
		}

		if actual.SynchronizationTemplateId != v.Expected.SynchronizationTemplateId {
			t.Fatalf("Expected %q but got %q for SynchronizationTemplateId", v.Expected.SynchronizationTemplateId, actual.SynchronizationTemplateId)
		}

	}
}

func TestSegmentsForApplicationSynchronizationTemplateId(t *testing.T) {
	segments := ApplicationSynchronizationTemplateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ApplicationSynchronizationTemplateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
