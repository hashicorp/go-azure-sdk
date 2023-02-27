package virtualmachineimages

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = OfferSkuVersionId{}

func TestNewOfferSkuVersionID(t *testing.T) {
	id := NewOfferSkuVersionID("12345678-1234-9876-4563-123456789012", "locationValue", "edgeZoneValue", "publisherValue", "offerValue", "skuValue", "versionValue")

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

	if id.VersionName != "versionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'VersionName'", id.VersionName, "versionValue")
	}
}

func TestFormatOfferSkuVersionID(t *testing.T) {
	actual := NewOfferSkuVersionID("12345678-1234-9876-4563-123456789012", "locationValue", "edgeZoneValue", "publisherValue", "offerValue", "skuValue", "versionValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers/publisherValue/artifactTypes/vmImage/offers/offerValue/skus/skuValue/versions/versionValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseOfferSkuVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *OfferSkuVersionId
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
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers/publisherValue/artifactTypes/vmImage/offers/offerValue/skus/skuValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers/publisherValue/artifactTypes/vmImage/offers/offerValue/skus/skuValue/versions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers/publisherValue/artifactTypes/vmImage/offers/offerValue/skus/skuValue/versions/versionValue",
			Expected: &OfferSkuVersionId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				LocationName:   "locationValue",
				EdgeZoneName:   "edgeZoneValue",
				PublisherName:  "publisherValue",
				OfferName:      "offerValue",
				SkuName:        "skuValue",
				VersionName:    "versionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers/publisherValue/artifactTypes/vmImage/offers/offerValue/skus/skuValue/versions/versionValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseOfferSkuVersionID(v.Input)
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

		if actual.VersionName != v.Expected.VersionName {
			t.Fatalf("Expected %q but got %q for VersionName", v.Expected.VersionName, actual.VersionName)
		}

	}
}

func TestParseOfferSkuVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *OfferSkuVersionId
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
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers/publisherValue/artifactTypes/vmImage/offers/offerValue/skus/skuValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe/eDgEzOnEs/eDgEzOnEvAlUe/pUbLiShErS/pUbLiShErVaLuE/aRtIfAcTtYpEs/vMiMaGe/oFfErS/oFfErVaLuE/sKuS/sKuVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers/publisherValue/artifactTypes/vmImage/offers/offerValue/skus/skuValue/versions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe/eDgEzOnEs/eDgEzOnEvAlUe/pUbLiShErS/pUbLiShErVaLuE/aRtIfAcTtYpEs/vMiMaGe/oFfErS/oFfErVaLuE/sKuS/sKuVaLuE/vErSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers/publisherValue/artifactTypes/vmImage/offers/offerValue/skus/skuValue/versions/versionValue",
			Expected: &OfferSkuVersionId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				LocationName:   "locationValue",
				EdgeZoneName:   "edgeZoneValue",
				PublisherName:  "publisherValue",
				OfferName:      "offerValue",
				SkuName:        "skuValue",
				VersionName:    "versionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/edgeZones/edgeZoneValue/publishers/publisherValue/artifactTypes/vmImage/offers/offerValue/skus/skuValue/versions/versionValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe/eDgEzOnEs/eDgEzOnEvAlUe/pUbLiShErS/pUbLiShErVaLuE/aRtIfAcTtYpEs/vMiMaGe/oFfErS/oFfErVaLuE/sKuS/sKuVaLuE/vErSiOnS/vErSiOnVaLuE",
			Expected: &OfferSkuVersionId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				LocationName:   "lOcAtIoNvAlUe",
				EdgeZoneName:   "eDgEzOnEvAlUe",
				PublisherName:  "pUbLiShErVaLuE",
				OfferName:      "oFfErVaLuE",
				SkuName:        "sKuVaLuE",
				VersionName:    "vErSiOnVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe/eDgEzOnEs/eDgEzOnEvAlUe/pUbLiShErS/pUbLiShErVaLuE/aRtIfAcTtYpEs/vMiMaGe/oFfErS/oFfErVaLuE/sKuS/sKuVaLuE/vErSiOnS/vErSiOnVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseOfferSkuVersionIDInsensitively(v.Input)
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

		if actual.VersionName != v.Expected.VersionName {
			t.Fatalf("Expected %q but got %q for VersionName", v.Expected.VersionName, actual.VersionName)
		}

	}
}

func TestSegmentsForOfferSkuVersionId(t *testing.T) {
	segments := OfferSkuVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("OfferSkuVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
