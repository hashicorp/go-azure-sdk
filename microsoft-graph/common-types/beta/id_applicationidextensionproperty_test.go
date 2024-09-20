package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApplicationIdExtensionPropertyId{}

func TestNewApplicationIdExtensionPropertyID(t *testing.T) {
	id := NewApplicationIdExtensionPropertyID("applicationId", "extensionPropertyId")

	if id.ApplicationId != "applicationId" {
		t.Fatalf("Expected %q but got %q for Segment 'ApplicationId'", id.ApplicationId, "applicationId")
	}

	if id.ExtensionPropertyId != "extensionPropertyId" {
		t.Fatalf("Expected %q but got %q for Segment 'ExtensionPropertyId'", id.ExtensionPropertyId, "extensionPropertyId")
	}
}

func TestFormatApplicationIdExtensionPropertyID(t *testing.T) {
	actual := NewApplicationIdExtensionPropertyID("applicationId", "extensionPropertyId").ID()
	expected := "/applications/applicationId/extensionProperties/extensionPropertyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseApplicationIdExtensionPropertyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationIdExtensionPropertyId
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
			Input: "/applications/applicationId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationId/extensionProperties",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationId/extensionProperties/extensionPropertyId",
			Expected: &ApplicationIdExtensionPropertyId{
				ApplicationId:       "applicationId",
				ExtensionPropertyId: "extensionPropertyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationId/extensionProperties/extensionPropertyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationIdExtensionPropertyID(v.Input)
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

		if actual.ExtensionPropertyId != v.Expected.ExtensionPropertyId {
			t.Fatalf("Expected %q but got %q for ExtensionPropertyId", v.Expected.ExtensionPropertyId, actual.ExtensionPropertyId)
		}

	}
}

func TestParseApplicationIdExtensionPropertyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationIdExtensionPropertyId
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
			Input: "/applications/applicationId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationId/extensionProperties",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/eXtEnSiOnPrOpErTiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationId/extensionProperties/extensionPropertyId",
			Expected: &ApplicationIdExtensionPropertyId{
				ApplicationId:       "applicationId",
				ExtensionPropertyId: "extensionPropertyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationId/extensionProperties/extensionPropertyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/eXtEnSiOnPrOpErTiEs/eXtEnSiOnPrOpErTyId",
			Expected: &ApplicationIdExtensionPropertyId{
				ApplicationId:       "aPpLiCaTiOnId",
				ExtensionPropertyId: "eXtEnSiOnPrOpErTyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/eXtEnSiOnPrOpErTiEs/eXtEnSiOnPrOpErTyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationIdExtensionPropertyIDInsensitively(v.Input)
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

		if actual.ExtensionPropertyId != v.Expected.ExtensionPropertyId {
			t.Fatalf("Expected %q but got %q for ExtensionPropertyId", v.Expected.ExtensionPropertyId, actual.ExtensionPropertyId)
		}

	}
}

func TestSegmentsForApplicationIdExtensionPropertyId(t *testing.T) {
	segments := ApplicationIdExtensionPropertyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ApplicationIdExtensionPropertyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
