package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId{}

func TestNewServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupID(t *testing.T) {
	id := NewServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupID("servicePrincipalIdValue", "targetDeviceGroupIdValue")

	if id.ServicePrincipalId != "servicePrincipalIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalIdValue")
	}

	if id.TargetDeviceGroupId != "targetDeviceGroupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TargetDeviceGroupId'", id.TargetDeviceGroupId, "targetDeviceGroupIdValue")
	}
}

func TestFormatServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupID(t *testing.T) {
	actual := NewServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupID("servicePrincipalIdValue", "targetDeviceGroupIdValue").ID()
	expected := "/servicePrincipals/servicePrincipalIdValue/remoteDesktopSecurityConfiguration/targetDeviceGroups/targetDeviceGroupIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId
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
			Input: "/servicePrincipals/servicePrincipalIdValue/remoteDesktopSecurityConfiguration",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalIdValue/remoteDesktopSecurityConfiguration/targetDeviceGroups",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalIdValue/remoteDesktopSecurityConfiguration/targetDeviceGroups/targetDeviceGroupIdValue",
			Expected: &ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId{
				ServicePrincipalId:  "servicePrincipalIdValue",
				TargetDeviceGroupId: "targetDeviceGroupIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalIdValue/remoteDesktopSecurityConfiguration/targetDeviceGroups/targetDeviceGroupIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupID(v.Input)
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

		if actual.TargetDeviceGroupId != v.Expected.TargetDeviceGroupId {
			t.Fatalf("Expected %q but got %q for TargetDeviceGroupId", v.Expected.TargetDeviceGroupId, actual.TargetDeviceGroupId)
		}

	}
}

func TestParseServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId
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
			Input: "/servicePrincipals/servicePrincipalIdValue/remoteDesktopSecurityConfiguration",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/rEmOtEdEsKtOpSeCuRiTyCoNfIgUrAtIoN",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalIdValue/remoteDesktopSecurityConfiguration/targetDeviceGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/rEmOtEdEsKtOpSeCuRiTyCoNfIgUrAtIoN/tArGeTdEvIcEgRoUpS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalIdValue/remoteDesktopSecurityConfiguration/targetDeviceGroups/targetDeviceGroupIdValue",
			Expected: &ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId{
				ServicePrincipalId:  "servicePrincipalIdValue",
				TargetDeviceGroupId: "targetDeviceGroupIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalIdValue/remoteDesktopSecurityConfiguration/targetDeviceGroups/targetDeviceGroupIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/rEmOtEdEsKtOpSeCuRiTyCoNfIgUrAtIoN/tArGeTdEvIcEgRoUpS/tArGeTdEvIcEgRoUpIdVaLuE",
			Expected: &ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId{
				ServicePrincipalId:  "sErViCePrInCiPaLiDvAlUe",
				TargetDeviceGroupId: "tArGeTdEvIcEgRoUpIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/rEmOtEdEsKtOpSeCuRiTyCoNfIgUrAtIoN/tArGeTdEvIcEgRoUpS/tArGeTdEvIcEgRoUpIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupIDInsensitively(v.Input)
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

		if actual.TargetDeviceGroupId != v.Expected.TargetDeviceGroupId {
			t.Fatalf("Expected %q but got %q for TargetDeviceGroupId", v.Expected.TargetDeviceGroupId, actual.TargetDeviceGroupId)
		}

	}
}

func TestSegmentsForServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId(t *testing.T) {
	segments := ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServicePrincipalIdRemoteDesktopSecurityConfigurationTargetDeviceGroupId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
