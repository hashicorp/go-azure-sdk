package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdLicenseDetailId{}

func TestNewServicePrincipalIdLicenseDetailID(t *testing.T) {
	id := NewServicePrincipalIdLicenseDetailID("servicePrincipalId", "licenseDetailsId")

	if id.ServicePrincipalId != "servicePrincipalId" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalId")
	}

	if id.LicenseDetailsId != "licenseDetailsId" {
		t.Fatalf("Expected %q but got %q for Segment 'LicenseDetailsId'", id.LicenseDetailsId, "licenseDetailsId")
	}
}

func TestFormatServicePrincipalIdLicenseDetailID(t *testing.T) {
	actual := NewServicePrincipalIdLicenseDetailID("servicePrincipalId", "licenseDetailsId").ID()
	expected := "/servicePrincipals/servicePrincipalId/licenseDetails/licenseDetailsId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServicePrincipalIdLicenseDetailID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdLicenseDetailId
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
			Input: "/servicePrincipals/servicePrincipalId/licenseDetails",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/licenseDetails/licenseDetailsId",
			Expected: &ServicePrincipalIdLicenseDetailId{
				ServicePrincipalId: "servicePrincipalId",
				LicenseDetailsId:   "licenseDetailsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/licenseDetails/licenseDetailsId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdLicenseDetailID(v.Input)
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

		if actual.LicenseDetailsId != v.Expected.LicenseDetailsId {
			t.Fatalf("Expected %q but got %q for LicenseDetailsId", v.Expected.LicenseDetailsId, actual.LicenseDetailsId)
		}

	}
}

func TestParseServicePrincipalIdLicenseDetailIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdLicenseDetailId
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
			Input: "/servicePrincipals/servicePrincipalId/licenseDetails",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/lIcEnSeDeTaIlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/licenseDetails/licenseDetailsId",
			Expected: &ServicePrincipalIdLicenseDetailId{
				ServicePrincipalId: "servicePrincipalId",
				LicenseDetailsId:   "licenseDetailsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/licenseDetails/licenseDetailsId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/lIcEnSeDeTaIlS/lIcEnSeDeTaIlSiD",
			Expected: &ServicePrincipalIdLicenseDetailId{
				ServicePrincipalId: "sErViCePrInCiPaLiD",
				LicenseDetailsId:   "lIcEnSeDeTaIlSiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/lIcEnSeDeTaIlS/lIcEnSeDeTaIlSiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdLicenseDetailIDInsensitively(v.Input)
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

		if actual.LicenseDetailsId != v.Expected.LicenseDetailsId {
			t.Fatalf("Expected %q but got %q for LicenseDetailsId", v.Expected.LicenseDetailsId, actual.LicenseDetailsId)
		}

	}
}

func TestSegmentsForServicePrincipalIdLicenseDetailId(t *testing.T) {
	segments := ServicePrincipalIdLicenseDetailId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServicePrincipalIdLicenseDetailId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
