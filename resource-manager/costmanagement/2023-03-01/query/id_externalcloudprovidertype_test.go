package query

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ExternalCloudProviderTypeId{}

func TestNewExternalCloudProviderTypeID(t *testing.T) {
	id := NewExternalCloudProviderTypeID("example", "externalCloudProviderIdValue")

	if id.ExternalCloudProviderType != "example" {
		t.Fatalf("Expected %q but got %q for Segment 'ExternalCloudProviderType'", id.ExternalCloudProviderType, "example")
	}

	if id.ExternalCloudProviderId != "externalCloudProviderIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ExternalCloudProviderId'", id.ExternalCloudProviderId, "externalCloudProviderIdValue")
	}
}

func TestFormatExternalCloudProviderTypeID(t *testing.T) {
	actual := NewExternalCloudProviderTypeID("example", "externalCloudProviderIdValue").ID()
	expected := "/providers/Microsoft.CostManagement/example/externalCloudProviderIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseExternalCloudProviderTypeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ExternalCloudProviderTypeId
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
			Input: "/providers/Microsoft.CostManagement",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.CostManagement/example",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.CostManagement/example/externalCloudProviderIdValue",
			Expected: &ExternalCloudProviderTypeId{
				ExternalCloudProviderType: "example",
				ExternalCloudProviderId:   "externalCloudProviderIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.CostManagement/example/externalCloudProviderIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseExternalCloudProviderTypeID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ExternalCloudProviderType != v.Expected.ExternalCloudProviderType {
			t.Fatalf("Expected %q but got %q for ExternalCloudProviderType", v.Expected.ExternalCloudProviderType, actual.ExternalCloudProviderType)
		}

		if actual.ExternalCloudProviderId != v.Expected.ExternalCloudProviderId {
			t.Fatalf("Expected %q but got %q for ExternalCloudProviderId", v.Expected.ExternalCloudProviderId, actual.ExternalCloudProviderId)
		}

	}
}

func TestParseExternalCloudProviderTypeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ExternalCloudProviderTypeId
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
			Input: "/providers/Microsoft.CostManagement",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.cOsTmAnAgEmEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.CostManagement/example",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.cOsTmAnAgEmEnT/eXaMpLe",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.CostManagement/example/externalCloudProviderIdValue",
			Expected: &ExternalCloudProviderTypeId{
				ExternalCloudProviderType: "example",
				ExternalCloudProviderId:   "externalCloudProviderIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.CostManagement/example/externalCloudProviderIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.cOsTmAnAgEmEnT/eXaMpLe/eXtErNaLcLoUdPrOvIdErIdVaLuE",
			Expected: &ExternalCloudProviderTypeId{
				ExternalCloudProviderType: "example",
				ExternalCloudProviderId:   "eXtErNaLcLoUdPrOvIdErIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.cOsTmAnAgEmEnT/eXaMpLe/eXtErNaLcLoUdPrOvIdErIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseExternalCloudProviderTypeIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ExternalCloudProviderType != v.Expected.ExternalCloudProviderType {
			t.Fatalf("Expected %q but got %q for ExternalCloudProviderType", v.Expected.ExternalCloudProviderType, actual.ExternalCloudProviderType)
		}

		if actual.ExternalCloudProviderId != v.Expected.ExternalCloudProviderId {
			t.Fatalf("Expected %q but got %q for ExternalCloudProviderId", v.Expected.ExternalCloudProviderId, actual.ExternalCloudProviderId)
		}

	}
}

func TestSegmentsForExternalCloudProviderTypeId(t *testing.T) {
	segments := ExternalCloudProviderTypeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ExternalCloudProviderTypeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
