package serviceprincipallicensedetail

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalLicenseDetailId{}

func TestNewServicePrincipalLicenseDetailID(t *testing.T) {
	id := NewServicePrincipalLicenseDetailID("servicePrincipalIdValue", "licenseDetailsIdValue")

	if id.ServicePrincipalId != "servicePrincipalIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalIdValue")
	}

	if id.LicenseDetailsId != "licenseDetailsIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'LicenseDetailsId'", id.LicenseDetailsId, "licenseDetailsIdValue")
	}
}

func TestFormatServicePrincipalLicenseDetailID(t *testing.T) {
	actual := NewServicePrincipalLicenseDetailID("servicePrincipalIdValue", "licenseDetailsIdValue").ID()
	expected := "/servicePrincipals/servicePrincipalIdValue/licenseDetails/licenseDetailsIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServicePrincipalLicenseDetailID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalLicenseDetailId
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
			Input: "/servicePrincipals/servicePrincipalIdValue/licenseDetails",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalIdValue/licenseDetails/licenseDetailsIdValue",
			Expected: &ServicePrincipalLicenseDetailId{
				ServicePrincipalId: "servicePrincipalIdValue",
				LicenseDetailsId:   "licenseDetailsIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalIdValue/licenseDetails/licenseDetailsIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalLicenseDetailID(v.Input)
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

func TestParseServicePrincipalLicenseDetailIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalLicenseDetailId
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
			Input: "/servicePrincipals/servicePrincipalIdValue/licenseDetails",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/lIcEnSeDeTaIlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalIdValue/licenseDetails/licenseDetailsIdValue",
			Expected: &ServicePrincipalLicenseDetailId{
				ServicePrincipalId: "servicePrincipalIdValue",
				LicenseDetailsId:   "licenseDetailsIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalIdValue/licenseDetails/licenseDetailsIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/lIcEnSeDeTaIlS/lIcEnSeDeTaIlSiDvAlUe",
			Expected: &ServicePrincipalLicenseDetailId{
				ServicePrincipalId: "sErViCePrInCiPaLiDvAlUe",
				LicenseDetailsId:   "lIcEnSeDeTaIlSiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/lIcEnSeDeTaIlS/lIcEnSeDeTaIlSiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalLicenseDetailIDInsensitively(v.Input)
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

func TestSegmentsForServicePrincipalLicenseDetailId(t *testing.T) {
	segments := ServicePrincipalLicenseDetailId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServicePrincipalLicenseDetailId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
