package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppId{}

func TestNewServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppID(t *testing.T) {
	id := NewServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppID("servicePrincipalId", "approvedClientAppId")

	if id.ServicePrincipalId != "servicePrincipalId" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalId")
	}

	if id.ApprovedClientAppId != "approvedClientAppId" {
		t.Fatalf("Expected %q but got %q for Segment 'ApprovedClientAppId'", id.ApprovedClientAppId, "approvedClientAppId")
	}
}

func TestFormatServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppID(t *testing.T) {
	actual := NewServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppID("servicePrincipalId", "approvedClientAppId").ID()
	expected := "/servicePrincipals/servicePrincipalId/remoteDesktopSecurityConfiguration/approvedClientApps/approvedClientAppId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppId
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
			Input: "/servicePrincipals/servicePrincipalId/remoteDesktopSecurityConfiguration",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/remoteDesktopSecurityConfiguration/approvedClientApps",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/remoteDesktopSecurityConfiguration/approvedClientApps/approvedClientAppId",
			Expected: &ServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppId{
				ServicePrincipalId:  "servicePrincipalId",
				ApprovedClientAppId: "approvedClientAppId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/remoteDesktopSecurityConfiguration/approvedClientApps/approvedClientAppId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppID(v.Input)
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

		if actual.ApprovedClientAppId != v.Expected.ApprovedClientAppId {
			t.Fatalf("Expected %q but got %q for ApprovedClientAppId", v.Expected.ApprovedClientAppId, actual.ApprovedClientAppId)
		}

	}
}

func TestParseServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppId
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
			Input: "/servicePrincipals/servicePrincipalId/remoteDesktopSecurityConfiguration",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/rEmOtEdEsKtOpSeCuRiTyCoNfIgUrAtIoN",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/remoteDesktopSecurityConfiguration/approvedClientApps",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/rEmOtEdEsKtOpSeCuRiTyCoNfIgUrAtIoN/aPpRoVeDcLiEnTaPpS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/remoteDesktopSecurityConfiguration/approvedClientApps/approvedClientAppId",
			Expected: &ServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppId{
				ServicePrincipalId:  "servicePrincipalId",
				ApprovedClientAppId: "approvedClientAppId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/remoteDesktopSecurityConfiguration/approvedClientApps/approvedClientAppId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/rEmOtEdEsKtOpSeCuRiTyCoNfIgUrAtIoN/aPpRoVeDcLiEnTaPpS/aPpRoVeDcLiEnTaPpId",
			Expected: &ServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppId{
				ServicePrincipalId:  "sErViCePrInCiPaLiD",
				ApprovedClientAppId: "aPpRoVeDcLiEnTaPpId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/rEmOtEdEsKtOpSeCuRiTyCoNfIgUrAtIoN/aPpRoVeDcLiEnTaPpS/aPpRoVeDcLiEnTaPpId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppIDInsensitively(v.Input)
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

		if actual.ApprovedClientAppId != v.Expected.ApprovedClientAppId {
			t.Fatalf("Expected %q but got %q for ApprovedClientAppId", v.Expected.ApprovedClientAppId, actual.ApprovedClientAppId)
		}

	}
}

func TestSegmentsForServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppId(t *testing.T) {
	segments := ServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServicePrincipalIdRemoteDesktopSecurityConfigurationApprovedClientAppId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
