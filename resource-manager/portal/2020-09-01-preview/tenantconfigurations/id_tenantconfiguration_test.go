package tenantconfigurations

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &TenantConfigurationId{}

func TestNewTenantConfigurationID(t *testing.T) {
	id := NewTenantConfigurationID("tenantConfigurationName")

	if id.TenantConfigurationName != "tenantConfigurationName" {
		t.Fatalf("Expected %q but got %q for Segment 'TenantConfigurationName'", id.TenantConfigurationName, "tenantConfigurationName")
	}
}

func TestFormatTenantConfigurationID(t *testing.T) {
	actual := NewTenantConfigurationID("tenantConfigurationName").ID()
	expected := "/providers/Microsoft.Portal/tenantConfigurations/tenantConfigurationName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseTenantConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *TenantConfigurationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Portal",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Portal/tenantConfigurations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Portal/tenantConfigurations/tenantConfigurationName",
			Expected: &TenantConfigurationId{
				TenantConfigurationName: "tenantConfigurationName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Portal/tenantConfigurations/tenantConfigurationName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseTenantConfigurationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.TenantConfigurationName != v.Expected.TenantConfigurationName {
			t.Fatalf("Expected %q but got %q for TenantConfigurationName", v.Expected.TenantConfigurationName, actual.TenantConfigurationName)
		}

	}
}

func TestParseTenantConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *TenantConfigurationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Portal",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.pOrTaL",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Portal/tenantConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.pOrTaL/tEnAnTcOnFiGuRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Portal/tenantConfigurations/tenantConfigurationName",
			Expected: &TenantConfigurationId{
				TenantConfigurationName: "tenantConfigurationName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Portal/tenantConfigurations/tenantConfigurationName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.pOrTaL/tEnAnTcOnFiGuRaTiOnS/tEnAnTcOnFiGuRaTiOnNaMe",
			Expected: &TenantConfigurationId{
				TenantConfigurationName: "tEnAnTcOnFiGuRaTiOnNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.pOrTaL/tEnAnTcOnFiGuRaTiOnS/tEnAnTcOnFiGuRaTiOnNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseTenantConfigurationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.TenantConfigurationName != v.Expected.TenantConfigurationName {
			t.Fatalf("Expected %q but got %q for TenantConfigurationName", v.Expected.TenantConfigurationName, actual.TenantConfigurationName)
		}

	}
}

func TestSegmentsForTenantConfigurationId(t *testing.T) {
	segments := TenantConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("TenantConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
