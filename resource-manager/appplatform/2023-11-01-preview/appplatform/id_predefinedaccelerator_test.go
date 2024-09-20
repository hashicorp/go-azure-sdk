package appplatform

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PredefinedAcceleratorId{}

func TestNewPredefinedAcceleratorID(t *testing.T) {
	id := NewPredefinedAcceleratorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceName", "applicationAcceleratorName", "predefinedAcceleratorName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.SpringName != "serviceName" {
		t.Fatalf("Expected %q but got %q for Segment 'SpringName'", id.SpringName, "serviceName")
	}

	if id.ApplicationAcceleratorName != "applicationAcceleratorName" {
		t.Fatalf("Expected %q but got %q for Segment 'ApplicationAcceleratorName'", id.ApplicationAcceleratorName, "applicationAcceleratorName")
	}

	if id.PredefinedAcceleratorName != "predefinedAcceleratorName" {
		t.Fatalf("Expected %q but got %q for Segment 'PredefinedAcceleratorName'", id.PredefinedAcceleratorName, "predefinedAcceleratorName")
	}
}

func TestFormatPredefinedAcceleratorID(t *testing.T) {
	actual := NewPredefinedAcceleratorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceName", "applicationAcceleratorName", "predefinedAcceleratorName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AppPlatform/spring/serviceName/applicationAccelerators/applicationAcceleratorName/predefinedAccelerators/predefinedAcceleratorName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePredefinedAcceleratorID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PredefinedAcceleratorId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AppPlatform",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AppPlatform/spring",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AppPlatform/spring/serviceName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AppPlatform/spring/serviceName/applicationAccelerators",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AppPlatform/spring/serviceName/applicationAccelerators/applicationAcceleratorName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AppPlatform/spring/serviceName/applicationAccelerators/applicationAcceleratorName/predefinedAccelerators",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AppPlatform/spring/serviceName/applicationAccelerators/applicationAcceleratorName/predefinedAccelerators/predefinedAcceleratorName",
			Expected: &PredefinedAcceleratorId{
				SubscriptionId:             "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:          "example-resource-group",
				SpringName:                 "serviceName",
				ApplicationAcceleratorName: "applicationAcceleratorName",
				PredefinedAcceleratorName:  "predefinedAcceleratorName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AppPlatform/spring/serviceName/applicationAccelerators/applicationAcceleratorName/predefinedAccelerators/predefinedAcceleratorName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePredefinedAcceleratorID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.SpringName != v.Expected.SpringName {
			t.Fatalf("Expected %q but got %q for SpringName", v.Expected.SpringName, actual.SpringName)
		}

		if actual.ApplicationAcceleratorName != v.Expected.ApplicationAcceleratorName {
			t.Fatalf("Expected %q but got %q for ApplicationAcceleratorName", v.Expected.ApplicationAcceleratorName, actual.ApplicationAcceleratorName)
		}

		if actual.PredefinedAcceleratorName != v.Expected.PredefinedAcceleratorName {
			t.Fatalf("Expected %q but got %q for PredefinedAcceleratorName", v.Expected.PredefinedAcceleratorName, actual.PredefinedAcceleratorName)
		}

	}
}

func TestParsePredefinedAcceleratorIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PredefinedAcceleratorId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AppPlatform",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aPpPlAtFoRm",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AppPlatform/spring",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aPpPlAtFoRm/sPrInG",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AppPlatform/spring/serviceName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aPpPlAtFoRm/sPrInG/sErViCeNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AppPlatform/spring/serviceName/applicationAccelerators",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aPpPlAtFoRm/sPrInG/sErViCeNaMe/aPpLiCaTiOnAcCeLeRaToRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AppPlatform/spring/serviceName/applicationAccelerators/applicationAcceleratorName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aPpPlAtFoRm/sPrInG/sErViCeNaMe/aPpLiCaTiOnAcCeLeRaToRs/aPpLiCaTiOnAcCeLeRaToRnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AppPlatform/spring/serviceName/applicationAccelerators/applicationAcceleratorName/predefinedAccelerators",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aPpPlAtFoRm/sPrInG/sErViCeNaMe/aPpLiCaTiOnAcCeLeRaToRs/aPpLiCaTiOnAcCeLeRaToRnAmE/pReDeFiNeDaCcElErAtOrS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AppPlatform/spring/serviceName/applicationAccelerators/applicationAcceleratorName/predefinedAccelerators/predefinedAcceleratorName",
			Expected: &PredefinedAcceleratorId{
				SubscriptionId:             "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:          "example-resource-group",
				SpringName:                 "serviceName",
				ApplicationAcceleratorName: "applicationAcceleratorName",
				PredefinedAcceleratorName:  "predefinedAcceleratorName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AppPlatform/spring/serviceName/applicationAccelerators/applicationAcceleratorName/predefinedAccelerators/predefinedAcceleratorName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aPpPlAtFoRm/sPrInG/sErViCeNaMe/aPpLiCaTiOnAcCeLeRaToRs/aPpLiCaTiOnAcCeLeRaToRnAmE/pReDeFiNeDaCcElErAtOrS/pReDeFiNeDaCcElErAtOrNaMe",
			Expected: &PredefinedAcceleratorId{
				SubscriptionId:             "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:          "eXaMpLe-rEsOuRcE-GrOuP",
				SpringName:                 "sErViCeNaMe",
				ApplicationAcceleratorName: "aPpLiCaTiOnAcCeLeRaToRnAmE",
				PredefinedAcceleratorName:  "pReDeFiNeDaCcElErAtOrNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aPpPlAtFoRm/sPrInG/sErViCeNaMe/aPpLiCaTiOnAcCeLeRaToRs/aPpLiCaTiOnAcCeLeRaToRnAmE/pReDeFiNeDaCcElErAtOrS/pReDeFiNeDaCcElErAtOrNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePredefinedAcceleratorIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.SpringName != v.Expected.SpringName {
			t.Fatalf("Expected %q but got %q for SpringName", v.Expected.SpringName, actual.SpringName)
		}

		if actual.ApplicationAcceleratorName != v.Expected.ApplicationAcceleratorName {
			t.Fatalf("Expected %q but got %q for ApplicationAcceleratorName", v.Expected.ApplicationAcceleratorName, actual.ApplicationAcceleratorName)
		}

		if actual.PredefinedAcceleratorName != v.Expected.PredefinedAcceleratorName {
			t.Fatalf("Expected %q but got %q for PredefinedAcceleratorName", v.Expected.PredefinedAcceleratorName, actual.PredefinedAcceleratorName)
		}

	}
}

func TestSegmentsForPredefinedAcceleratorId(t *testing.T) {
	segments := PredefinedAcceleratorId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PredefinedAcceleratorId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
