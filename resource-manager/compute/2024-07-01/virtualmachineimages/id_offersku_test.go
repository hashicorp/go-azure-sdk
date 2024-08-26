package virtualmachineimages

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &OfferSkuId{}

func TestNewOfferSkuID(t *testing.T) {
	id := NewOfferSkuID("12345678-1234-9876-4563-123456789012", "locationValue", "edgeZoneValue", "publisherValue", "offerValue", "skuValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.LocationName != "locationValue" {
		t.Fatalf("Expected %q but got %q for Segment 'LocationName'", id.LocationName, "locationValue")
	}

	if id.EdgeZoneName != "edgeZoneValue" {
		t.Fatalf("Expected %q but got %q for Segment 'EdgeZoneName'", id.EdgeZoneName, "edgeZoneValue")
	}

	if id.PublisherName != "publisherValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PublisherName'", id.PublisherName, "publisherValue")
	}

	if id.OfferName != "offerValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OfferName'", id.OfferName, "offerValue")
	}

	if id.SkuName != "skuValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SkuName'", id.SkuName, "skuValue")
	}
}

func TestFormatOfferSkuID(t *testing.T) {
	actual := NewOfferSkuID("12345678-1234-9876-4563-123456789012", "locationValue", "edgeZoneValue", "publisherValue", "offerValue", "skuValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers/publisherValue/artifactTypes/vmImage/offers/offerValue/skus/skuValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseOfferSkuID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *OfferSkuId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers/publisherValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers/publisherValue/artifactTypes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers/publisherValue/artifactTypes/vmImage",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers/publisherValue/artifactTypes/vmImage/offers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers/publisherValue/artifactTypes/vmImage/offers/offerValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers/publisherValue/artifactTypes/vmImage/offers/offerValue/skus",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers/publisherValue/artifactTypes/vmImage/offers/offerValue/skus/skuValue",
			Expected: &OfferSkuId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				LocationName:   "locationValue",
				EdgeZoneName:   "edgeZoneValue",
				PublisherName:  "publisherValue",
				OfferName:      "offerValue",
				SkuName:        "skuValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers/publisherValue/artifactTypes/vmImage/offers/offerValue/skus/skuValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseOfferSkuID(v.Input)
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

		if actual.LocationName != v.Expected.LocationName {
			t.Fatalf("Expected %q but got %q for LocationName", v.Expected.LocationName, actual.LocationName)
		}

		if actual.EdgeZoneName != v.Expected.EdgeZoneName {
			t.Fatalf("Expected %q but got %q for EdgeZoneName", v.Expected.EdgeZoneName, actual.EdgeZoneName)
		}

		if actual.PublisherName != v.Expected.PublisherName {
			t.Fatalf("Expected %q but got %q for PublisherName", v.Expected.PublisherName, actual.PublisherName)
		}

		if actual.OfferName != v.Expected.OfferName {
			t.Fatalf("Expected %q but got %q for OfferName", v.Expected.OfferName, actual.OfferName)
		}

		if actual.SkuName != v.Expected.SkuName {
			t.Fatalf("Expected %q but got %q for SkuName", v.Expected.SkuName, actual.SkuName)
		}

	}
}

func TestParseOfferSkuIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *OfferSkuId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe/eDgEzOnEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe/eDgEzOnEs/eDgEzOnEvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe/eDgEzOnEs/eDgEzOnEvAlUe/pUbLiShErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers/publisherValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe/eDgEzOnEs/eDgEzOnEvAlUe/pUbLiShErS/pUbLiShErVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers/publisherValue/artifactTypes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe/eDgEzOnEs/eDgEzOnEvAlUe/pUbLiShErS/pUbLiShErVaLuE/aRtIfAcTtYpEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers/publisherValue/artifactTypes/vmImage",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe/eDgEzOnEs/eDgEzOnEvAlUe/pUbLiShErS/pUbLiShErVaLuE/aRtIfAcTtYpEs/vMiMaGe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers/publisherValue/artifactTypes/vmImage/offers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe/eDgEzOnEs/eDgEzOnEvAlUe/pUbLiShErS/pUbLiShErVaLuE/aRtIfAcTtYpEs/vMiMaGe/oFfErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers/publisherValue/artifactTypes/vmImage/offers/offerValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe/eDgEzOnEs/eDgEzOnEvAlUe/pUbLiShErS/pUbLiShErVaLuE/aRtIfAcTtYpEs/vMiMaGe/oFfErS/oFfErVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers/publisherValue/artifactTypes/vmImage/offers/offerValue/skus",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe/eDgEzOnEs/eDgEzOnEvAlUe/pUbLiShErS/pUbLiShErVaLuE/aRtIfAcTtYpEs/vMiMaGe/oFfErS/oFfErVaLuE/sKuS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers/publisherValue/artifactTypes/vmImage/offers/offerValue/skus/skuValue",
			Expected: &OfferSkuId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				LocationName:   "locationValue",
				EdgeZoneName:   "edgeZoneValue",
				PublisherName:  "publisherValue",
				OfferName:      "offerValue",
				SkuName:        "skuValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers/publisherValue/artifactTypes/vmImage/offers/offerValue/skus/skuValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe/eDgEzOnEs/eDgEzOnEvAlUe/pUbLiShErS/pUbLiShErVaLuE/aRtIfAcTtYpEs/vMiMaGe/oFfErS/oFfErVaLuE/sKuS/sKuVaLuE",
			Expected: &OfferSkuId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				LocationName:   "lOcAtIoNvAlUe",
				EdgeZoneName:   "eDgEzOnEvAlUe",
				PublisherName:  "pUbLiShErVaLuE",
				OfferName:      "oFfErVaLuE",
				SkuName:        "sKuVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe/eDgEzOnEs/eDgEzOnEvAlUe/pUbLiShErS/pUbLiShErVaLuE/aRtIfAcTtYpEs/vMiMaGe/oFfErS/oFfErVaLuE/sKuS/sKuVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseOfferSkuIDInsensitively(v.Input)
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

		if actual.LocationName != v.Expected.LocationName {
			t.Fatalf("Expected %q but got %q for LocationName", v.Expected.LocationName, actual.LocationName)
		}

		if actual.EdgeZoneName != v.Expected.EdgeZoneName {
			t.Fatalf("Expected %q but got %q for EdgeZoneName", v.Expected.EdgeZoneName, actual.EdgeZoneName)
		}

		if actual.PublisherName != v.Expected.PublisherName {
			t.Fatalf("Expected %q but got %q for PublisherName", v.Expected.PublisherName, actual.PublisherName)
		}

		if actual.OfferName != v.Expected.OfferName {
			t.Fatalf("Expected %q but got %q for OfferName", v.Expected.OfferName, actual.OfferName)
		}

		if actual.SkuName != v.Expected.SkuName {
			t.Fatalf("Expected %q but got %q for SkuName", v.Expected.SkuName, actual.SkuName)
		}

	}
}

func TestSegmentsForOfferSkuId(t *testing.T) {
	segments := OfferSkuId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("OfferSkuId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
