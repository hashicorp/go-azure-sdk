package virtualmachineimages

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &VMImageOfferId{}

func TestNewVMImageOfferID(t *testing.T) {
	id := NewVMImageOfferID("12345678-1234-9876-4563-123456789012", "location", "edgeZone", "publisherName", "offer")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.LocationName != "location" {
		t.Fatalf("Expected %q but got %q for Segment 'LocationName'", id.LocationName, "location")
	}

	if id.EdgeZoneName != "edgeZone" {
		t.Fatalf("Expected %q but got %q for Segment 'EdgeZoneName'", id.EdgeZoneName, "edgeZone")
	}

	if id.PublisherName != "publisherName" {
		t.Fatalf("Expected %q but got %q for Segment 'PublisherName'", id.PublisherName, "publisherName")
	}

	if id.OfferName != "offer" {
		t.Fatalf("Expected %q but got %q for Segment 'OfferName'", id.OfferName, "offer")
	}
}

func TestFormatVMImageOfferID(t *testing.T) {
	actual := NewVMImageOfferID("12345678-1234-9876-4563-123456789012", "location", "edgeZone", "publisherName", "offer").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/location/edgeZones/edgeZone/publishers/publisherName/artifactTypes/vmImage/offers/offer"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseVMImageOfferID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VMImageOfferId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/location",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/location/edgeZones",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/location/edgeZones/edgeZone",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/location/edgeZones/edgeZone/publishers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/location/edgeZones/edgeZone/publishers/publisherName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/location/edgeZones/edgeZone/publishers/publisherName/artifactTypes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/location/edgeZones/edgeZone/publishers/publisherName/artifactTypes/vmImage",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/location/edgeZones/edgeZone/publishers/publisherName/artifactTypes/vmImage/offers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/location/edgeZones/edgeZone/publishers/publisherName/artifactTypes/vmImage/offers/offer",
			Expected: &VMImageOfferId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				LocationName:   "location",
				EdgeZoneName:   "edgeZone",
				PublisherName:  "publisherName",
				OfferName:      "offer",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/location/edgeZones/edgeZone/publishers/publisherName/artifactTypes/vmImage/offers/offer/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVMImageOfferID(v.Input)
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

	}
}

func TestParseVMImageOfferIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VMImageOfferId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/location",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoN",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/location/edgeZones",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoN/eDgEzOnEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/location/edgeZones/edgeZone",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoN/eDgEzOnEs/eDgEzOnE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/location/edgeZones/edgeZone/publishers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoN/eDgEzOnEs/eDgEzOnE/pUbLiShErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/location/edgeZones/edgeZone/publishers/publisherName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoN/eDgEzOnEs/eDgEzOnE/pUbLiShErS/pUbLiShErNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/location/edgeZones/edgeZone/publishers/publisherName/artifactTypes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoN/eDgEzOnEs/eDgEzOnE/pUbLiShErS/pUbLiShErNaMe/aRtIfAcTtYpEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/location/edgeZones/edgeZone/publishers/publisherName/artifactTypes/vmImage",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoN/eDgEzOnEs/eDgEzOnE/pUbLiShErS/pUbLiShErNaMe/aRtIfAcTtYpEs/vMiMaGe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/location/edgeZones/edgeZone/publishers/publisherName/artifactTypes/vmImage/offers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoN/eDgEzOnEs/eDgEzOnE/pUbLiShErS/pUbLiShErNaMe/aRtIfAcTtYpEs/vMiMaGe/oFfErS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/location/edgeZones/edgeZone/publishers/publisherName/artifactTypes/vmImage/offers/offer",
			Expected: &VMImageOfferId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				LocationName:   "location",
				EdgeZoneName:   "edgeZone",
				PublisherName:  "publisherName",
				OfferName:      "offer",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/location/edgeZones/edgeZone/publishers/publisherName/artifactTypes/vmImage/offers/offer/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoN/eDgEzOnEs/eDgEzOnE/pUbLiShErS/pUbLiShErNaMe/aRtIfAcTtYpEs/vMiMaGe/oFfErS/oFfEr",
			Expected: &VMImageOfferId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				LocationName:   "lOcAtIoN",
				EdgeZoneName:   "eDgEzOnE",
				PublisherName:  "pUbLiShErNaMe",
				OfferName:      "oFfEr",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoN/eDgEzOnEs/eDgEzOnE/pUbLiShErS/pUbLiShErNaMe/aRtIfAcTtYpEs/vMiMaGe/oFfErS/oFfEr/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVMImageOfferIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForVMImageOfferId(t *testing.T) {
	segments := VMImageOfferId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("VMImageOfferId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
