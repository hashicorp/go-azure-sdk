package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdOAuth2PermissionGrantId{}

func TestNewServicePrincipalIdOAuth2PermissionGrantID(t *testing.T) {
	id := NewServicePrincipalIdOAuth2PermissionGrantID("servicePrincipalId", "oAuth2PermissionGrantId")

	if id.ServicePrincipalId != "servicePrincipalId" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalId")
	}

	if id.OAuth2PermissionGrantId != "oAuth2PermissionGrantId" {
		t.Fatalf("Expected %q but got %q for Segment 'OAuth2PermissionGrantId'", id.OAuth2PermissionGrantId, "oAuth2PermissionGrantId")
	}
}

func TestFormatServicePrincipalIdOAuth2PermissionGrantID(t *testing.T) {
	actual := NewServicePrincipalIdOAuth2PermissionGrantID("servicePrincipalId", "oAuth2PermissionGrantId").ID()
	expected := "/servicePrincipals/servicePrincipalId/oauth2PermissionGrants/oAuth2PermissionGrantId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServicePrincipalIdOAuth2PermissionGrantID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdOAuth2PermissionGrantId
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
			Input: "/servicePrincipals/servicePrincipalId/oauth2PermissionGrants",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/oauth2PermissionGrants/oAuth2PermissionGrantId",
			Expected: &ServicePrincipalIdOAuth2PermissionGrantId{
				ServicePrincipalId:      "servicePrincipalId",
				OAuth2PermissionGrantId: "oAuth2PermissionGrantId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/oauth2PermissionGrants/oAuth2PermissionGrantId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdOAuth2PermissionGrantID(v.Input)
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

func TestParseServicePrincipalIdOAuth2PermissionGrantIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdOAuth2PermissionGrantId
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
			Input: "/servicePrincipals/servicePrincipalId/oauth2PermissionGrants",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/oAuTh2pErMiSsIoNgRaNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/oauth2PermissionGrants/oAuth2PermissionGrantId",
			Expected: &ServicePrincipalIdOAuth2PermissionGrantId{
				ServicePrincipalId:      "servicePrincipalId",
				OAuth2PermissionGrantId: "oAuth2PermissionGrantId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/oauth2PermissionGrants/oAuth2PermissionGrantId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/oAuTh2pErMiSsIoNgRaNtS/oAuTh2pErMiSsIoNgRaNtId",
			Expected: &ServicePrincipalIdOAuth2PermissionGrantId{
				ServicePrincipalId:      "sErViCePrInCiPaLiD",
				OAuth2PermissionGrantId: "oAuTh2pErMiSsIoNgRaNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/oAuTh2pErMiSsIoNgRaNtS/oAuTh2pErMiSsIoNgRaNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdOAuth2PermissionGrantIDInsensitively(v.Input)
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

func TestSegmentsForServicePrincipalIdOAuth2PermissionGrantId(t *testing.T) {
	segments := ServicePrincipalIdOAuth2PermissionGrantId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServicePrincipalIdOAuth2PermissionGrantId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
