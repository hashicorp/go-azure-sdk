package serviceprincipaldelegatedpermissionclassification

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalDelegatedPermissionClassificationId{}

func TestNewServicePrincipalDelegatedPermissionClassificationID(t *testing.T) {
	id := NewServicePrincipalDelegatedPermissionClassificationID("servicePrincipalIdValue", "delegatedPermissionClassificationIdValue")

	if id.ServicePrincipalId != "servicePrincipalIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalIdValue")
	}

	if id.DelegatedPermissionClassificationId != "delegatedPermissionClassificationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DelegatedPermissionClassificationId'", id.DelegatedPermissionClassificationId, "delegatedPermissionClassificationIdValue")
	}
}

func TestFormatServicePrincipalDelegatedPermissionClassificationID(t *testing.T) {
	actual := NewServicePrincipalDelegatedPermissionClassificationID("servicePrincipalIdValue", "delegatedPermissionClassificationIdValue").ID()
	expected := "/servicePrincipals/servicePrincipalIdValue/delegatedPermissionClassifications/delegatedPermissionClassificationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServicePrincipalDelegatedPermissionClassificationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalDelegatedPermissionClassificationId
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
			Input: "/servicePrincipals/servicePrincipalIdValue/delegatedPermissionClassifications",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalIdValue/delegatedPermissionClassifications/delegatedPermissionClassificationIdValue",
			Expected: &ServicePrincipalDelegatedPermissionClassificationId{
				ServicePrincipalId:                  "servicePrincipalIdValue",
				DelegatedPermissionClassificationId: "delegatedPermissionClassificationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalIdValue/delegatedPermissionClassifications/delegatedPermissionClassificationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalDelegatedPermissionClassificationID(v.Input)
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

		if actual.DelegatedPermissionClassificationId != v.Expected.DelegatedPermissionClassificationId {
			t.Fatalf("Expected %q but got %q for DelegatedPermissionClassificationId", v.Expected.DelegatedPermissionClassificationId, actual.DelegatedPermissionClassificationId)
		}

	}
}

func TestParseServicePrincipalDelegatedPermissionClassificationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalDelegatedPermissionClassificationId
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
			Input: "/servicePrincipals/servicePrincipalIdValue/delegatedPermissionClassifications",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/dElEgAtEdPeRmIsSiOnClAsSiFiCaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalIdValue/delegatedPermissionClassifications/delegatedPermissionClassificationIdValue",
			Expected: &ServicePrincipalDelegatedPermissionClassificationId{
				ServicePrincipalId:                  "servicePrincipalIdValue",
				DelegatedPermissionClassificationId: "delegatedPermissionClassificationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalIdValue/delegatedPermissionClassifications/delegatedPermissionClassificationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/dElEgAtEdPeRmIsSiOnClAsSiFiCaTiOnS/dElEgAtEdPeRmIsSiOnClAsSiFiCaTiOnIdVaLuE",
			Expected: &ServicePrincipalDelegatedPermissionClassificationId{
				ServicePrincipalId:                  "sErViCePrInCiPaLiDvAlUe",
				DelegatedPermissionClassificationId: "dElEgAtEdPeRmIsSiOnClAsSiFiCaTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/dElEgAtEdPeRmIsSiOnClAsSiFiCaTiOnS/dElEgAtEdPeRmIsSiOnClAsSiFiCaTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalDelegatedPermissionClassificationIDInsensitively(v.Input)
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

		if actual.DelegatedPermissionClassificationId != v.Expected.DelegatedPermissionClassificationId {
			t.Fatalf("Expected %q but got %q for DelegatedPermissionClassificationId", v.Expected.DelegatedPermissionClassificationId, actual.DelegatedPermissionClassificationId)
		}

	}
}

func TestSegmentsForServicePrincipalDelegatedPermissionClassificationId(t *testing.T) {
	segments := ServicePrincipalDelegatedPermissionClassificationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServicePrincipalDelegatedPermissionClassificationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
