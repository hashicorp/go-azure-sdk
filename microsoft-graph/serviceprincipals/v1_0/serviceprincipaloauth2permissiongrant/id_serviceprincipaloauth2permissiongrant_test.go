package serviceprincipaloauth2permissiongrant

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalOauth2PermissionGrantId{}

func TestNewServicePrincipalOauth2PermissionGrantID(t *testing.T) {
	id := NewServicePrincipalOauth2PermissionGrantID("servicePrincipalIdValue", "oAuth2PermissionGrantIdValue")

	if id.ServicePrincipalId != "servicePrincipalIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalIdValue")
	}

	if id.OAuth2PermissionGrantId != "oAuth2PermissionGrantIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OAuth2PermissionGrantId'", id.OAuth2PermissionGrantId, "oAuth2PermissionGrantIdValue")
	}
}

func TestFormatServicePrincipalOauth2PermissionGrantID(t *testing.T) {
	actual := NewServicePrincipalOauth2PermissionGrantID("servicePrincipalIdValue", "oAuth2PermissionGrantIdValue").ID()
	expected := "/servicePrincipals/servicePrincipalIdValue/oauth2PermissionGrants/oAuth2PermissionGrantIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServicePrincipalOauth2PermissionGrantID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalOauth2PermissionGrantId
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
			Input: "/servicePrincipals/servicePrincipalIdValue/oauth2PermissionGrants",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalIdValue/oauth2PermissionGrants/oAuth2PermissionGrantIdValue",
			Expected: &ServicePrincipalOauth2PermissionGrantId{
				ServicePrincipalId:      "servicePrincipalIdValue",
				OAuth2PermissionGrantId: "oAuth2PermissionGrantIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalIdValue/oauth2PermissionGrants/oAuth2PermissionGrantIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalOauth2PermissionGrantID(v.Input)
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

		if actual.OAuth2PermissionGrantId != v.Expected.OAuth2PermissionGrantId {
			t.Fatalf("Expected %q but got %q for OAuth2PermissionGrantId", v.Expected.OAuth2PermissionGrantId, actual.OAuth2PermissionGrantId)
		}

	}
}

func TestParseServicePrincipalOauth2PermissionGrantIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalOauth2PermissionGrantId
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
			Input: "/servicePrincipals/servicePrincipalIdValue/oauth2PermissionGrants",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/oAuTh2pErMiSsIoNgRaNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalIdValue/oauth2PermissionGrants/oAuth2PermissionGrantIdValue",
			Expected: &ServicePrincipalOauth2PermissionGrantId{
				ServicePrincipalId:      "servicePrincipalIdValue",
				OAuth2PermissionGrantId: "oAuth2PermissionGrantIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalIdValue/oauth2PermissionGrants/oAuth2PermissionGrantIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/oAuTh2pErMiSsIoNgRaNtS/oAuTh2pErMiSsIoNgRaNtIdVaLuE",
			Expected: &ServicePrincipalOauth2PermissionGrantId{
				ServicePrincipalId:      "sErViCePrInCiPaLiDvAlUe",
				OAuth2PermissionGrantId: "oAuTh2pErMiSsIoNgRaNtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/oAuTh2pErMiSsIoNgRaNtS/oAuTh2pErMiSsIoNgRaNtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalOauth2PermissionGrantIDInsensitively(v.Input)
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

		if actual.OAuth2PermissionGrantId != v.Expected.OAuth2PermissionGrantId {
			t.Fatalf("Expected %q but got %q for OAuth2PermissionGrantId", v.Expected.OAuth2PermissionGrantId, actual.OAuth2PermissionGrantId)
		}

	}
}

func TestSegmentsForServicePrincipalOauth2PermissionGrantId(t *testing.T) {
	segments := ServicePrincipalOauth2PermissionGrantId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServicePrincipalOauth2PermissionGrantId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
