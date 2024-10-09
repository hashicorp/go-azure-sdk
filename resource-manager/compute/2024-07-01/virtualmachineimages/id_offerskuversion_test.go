package virtualmachineimages

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &OfferSkuVersionId{}

func TestNewOfferSkuVersionID(t *testing.T) {
	id := NewOfferSkuVersionID("12345678-1234-9876-4563-123456789012", "locationName", "edgeZoneName", "publisherName", "offerName", "skuName", "versionName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.LocationName != "locationName" {
		t.Fatalf("Expected %q but got %q for Segment 'LocationName'", id.LocationName, "locationName")
	}

	if id.EdgeZoneName != "edgeZoneName" {
		t.Fatalf("Expected %q but got %q for Segment 'EdgeZoneName'", id.EdgeZoneName, "edgeZoneName")
	}

	if id.PublisherName != "publisherName" {
		t.Fatalf("Expected %q but got %q for Segment 'PublisherName'", id.PublisherName, "publisherName")
	}

	if id.OfferName != "offerName" {
		t.Fatalf("Expected %q but got %q for Segment 'OfferName'", id.OfferName, "offerName")
	}

	if id.SkuName != "skuName" {
		t.Fatalf("Expected %q but got %q for Segment 'SkuName'", id.SkuName, "skuName")
	}

	if id.VersionName != "versionName" {
		t.Fatalf("Expected %q but got %q for Segment 'VersionName'", id.VersionName, "versionName")
	}
}

func TestFormatOfferSkuVersionID(t *testing.T) {
	actual := NewOfferSkuVersionID("12345678-1234-9876-4563-123456789012", "locationName", "edgeZoneName", "publisherName", "offerName", "skuName", "versionName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/edgeZones/edgeZoneName/publishers/publisherName/artifactTypes/vmImage/offers/offerName/skus/skuName/versions/versionName"
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/edgeZones",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/edgeZones/edgeZoneName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/edgeZones/edgeZoneName/publishers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/edgeZones/edgeZoneName/publishers/publisherName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/edgeZones/edgeZoneName/publishers/publisherName/artifactTypes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/edgeZones/edgeZoneName/publishers/publisherName/artifactTypes/vmImage",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/edgeZones/edgeZoneName/publishers/publisherName/artifactTypes/vmImage/offers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/edgeZones/edgeZoneName/publishers/publisherName/artifactTypes/vmImage/offers/offerName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/edgeZones/edgeZoneName/publishers/publisherName/artifactTypes/vmImage/offers/offerName/skus",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/edgeZones/edgeZoneName/publishers/publisherName/artifactTypes/vmImage/offers/offerName/skus/skuName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/edgeZones/edgeZoneName/publishers/publisherName/artifactTypes/vmImage/offers/offerName/skus/skuName/versions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/edgeZones/edgeZoneName/publishers/publisherName/artifactTypes/vmImage/offers/offerName/skus/skuName/versions/versionName",
			Expected: &OfferSkuVersionId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				LocationName:   "locationName",
				EdgeZoneName:   "edgeZoneName",
				PublisherName:  "publisherName",
				OfferName:      "offerName",
				SkuName:        "skuName",
				VersionName:    "versionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/edgeZones/edgeZoneName/publishers/publisherName/artifactTypes/vmImage/offers/offerName/skus/skuName/versions/versionName/extra",
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/edgeZones",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/eDgEzOnEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/edgeZones/edgeZoneName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/eDgEzOnEs/eDgEzOnEnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/edgeZones/edgeZoneName/publishers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/eDgEzOnEs/eDgEzOnEnAmE/pUbLiShErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/edgeZones/edgeZoneName/publishers/publisherName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/eDgEzOnEs/eDgEzOnEnAmE/pUbLiShErS/pUbLiShErNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/edgeZones/edgeZoneName/publishers/publisherName/artifactTypes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/eDgEzOnEs/eDgEzOnEnAmE/pUbLiShErS/pUbLiShErNaMe/aRtIfAcTtYpEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/edgeZones/edgeZoneName/publishers/publisherName/artifactTypes/vmImage",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/eDgEzOnEs/eDgEzOnEnAmE/pUbLiShErS/pUbLiShErNaMe/aRtIfAcTtYpEs/vMiMaGe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/edgeZones/edgeZoneName/publishers/publisherName/artifactTypes/vmImage/offers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/eDgEzOnEs/eDgEzOnEnAmE/pUbLiShErS/pUbLiShErNaMe/aRtIfAcTtYpEs/vMiMaGe/oFfErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/edgeZones/edgeZoneName/publishers/publisherName/artifactTypes/vmImage/offers/offerName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/eDgEzOnEs/eDgEzOnEnAmE/pUbLiShErS/pUbLiShErNaMe/aRtIfAcTtYpEs/vMiMaGe/oFfErS/oFfErNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/edgeZones/edgeZoneName/publishers/publisherName/artifactTypes/vmImage/offers/offerName/skus",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/eDgEzOnEs/eDgEzOnEnAmE/pUbLiShErS/pUbLiShErNaMe/aRtIfAcTtYpEs/vMiMaGe/oFfErS/oFfErNaMe/sKuS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/edgeZones/edgeZoneName/publishers/publisherName/artifactTypes/vmImage/offers/offerName/skus/skuName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/eDgEzOnEs/eDgEzOnEnAmE/pUbLiShErS/pUbLiShErNaMe/aRtIfAcTtYpEs/vMiMaGe/oFfErS/oFfErNaMe/sKuS/sKuNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/edgeZones/edgeZoneName/publishers/publisherName/artifactTypes/vmImage/offers/offerName/skus/skuName/versions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/eDgEzOnEs/eDgEzOnEnAmE/pUbLiShErS/pUbLiShErNaMe/aRtIfAcTtYpEs/vMiMaGe/oFfErS/oFfErNaMe/sKuS/sKuNaMe/vErSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/edgeZones/edgeZoneName/publishers/publisherName/artifactTypes/vmImage/offers/offerName/skus/skuName/versions/versionName",
			Expected: &OfferSkuVersionId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				LocationName:   "locationName",
				EdgeZoneName:   "edgeZoneName",
				PublisherName:  "publisherName",
				OfferName:      "offerName",
				SkuName:        "skuName",
				VersionName:    "versionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/edgeZones/edgeZoneName/publishers/publisherName/artifactTypes/vmImage/offers/offerName/skus/skuName/versions/versionName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/eDgEzOnEs/eDgEzOnEnAmE/pUbLiShErS/pUbLiShErNaMe/aRtIfAcTtYpEs/vMiMaGe/oFfErS/oFfErNaMe/sKuS/sKuNaMe/vErSiOnS/vErSiOnNaMe",
			Expected: &OfferSkuVersionId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				LocationName:   "lOcAtIoNnAmE",
				EdgeZoneName:   "eDgEzOnEnAmE",
				PublisherName:  "pUbLiShErNaMe",
				OfferName:      "oFfErNaMe",
				SkuName:        "sKuNaMe",
				VersionName:    "vErSiOnNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/eDgEzOnEs/eDgEzOnEnAmE/pUbLiShErS/pUbLiShErNaMe/aRtIfAcTtYpEs/vMiMaGe/oFfErS/oFfErNaMe/sKuS/sKuNaMe/vErSiOnS/vErSiOnNaMe/extra",
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
