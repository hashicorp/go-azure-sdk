package computerps

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &VersionId{}

func TestNewVersionID(t *testing.T) {
	id := NewVersionID("12345678-1234-9876-4563-123456789012", "locationName", "publisherName", "typeName", "versionName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.LocationName != "locationName" {
		t.Fatalf("Expected %q but got %q for Segment 'LocationName'", id.LocationName, "locationName")
	}

	if id.PublisherName != "publisherName" {
		t.Fatalf("Expected %q but got %q for Segment 'PublisherName'", id.PublisherName, "publisherName")
	}

	if id.TypeName != "typeName" {
		t.Fatalf("Expected %q but got %q for Segment 'TypeName'", id.TypeName, "typeName")
	}

	if id.VersionName != "versionName" {
		t.Fatalf("Expected %q but got %q for Segment 'VersionName'", id.VersionName, "versionName")
	}
}

func TestFormatVersionID(t *testing.T) {
	actual := NewVersionID("12345678-1234-9876-4563-123456789012", "locationName", "publisherName", "typeName", "versionName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/publishers/publisherName/artifactTypes/vmExtension/types/typeName/versions/versionName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VersionId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/publishers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/publishers/publisherName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/publishers/publisherName/artifactTypes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/publishers/publisherName/artifactTypes/vmExtension",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/publishers/publisherName/artifactTypes/vmExtension/types",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/publishers/publisherName/artifactTypes/vmExtension/types/typeName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/publishers/publisherName/artifactTypes/vmExtension/types/typeName/versions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/publishers/publisherName/artifactTypes/vmExtension/types/typeName/versions/versionName",
			Expected: &VersionId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				LocationName:   "locationName",
				PublisherName:  "publisherName",
				TypeName:       "typeName",
				VersionName:    "versionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/publishers/publisherName/artifactTypes/vmExtension/types/typeName/versions/versionName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVersionID(v.Input)
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

		if actual.PublisherName != v.Expected.PublisherName {
			t.Fatalf("Expected %q but got %q for PublisherName", v.Expected.PublisherName, actual.PublisherName)
		}

		if actual.TypeName != v.Expected.TypeName {
			t.Fatalf("Expected %q but got %q for TypeName", v.Expected.TypeName, actual.TypeName)
		}

		if actual.VersionName != v.Expected.VersionName {
			t.Fatalf("Expected %q but got %q for VersionName", v.Expected.VersionName, actual.VersionName)
		}

	}
}

func TestParseVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VersionId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/publishers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/pUbLiShErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/publishers/publisherName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/pUbLiShErS/pUbLiShErNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/publishers/publisherName/artifactTypes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/pUbLiShErS/pUbLiShErNaMe/aRtIfAcTtYpEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/publishers/publisherName/artifactTypes/vmExtension",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/pUbLiShErS/pUbLiShErNaMe/aRtIfAcTtYpEs/vMeXtEnSiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/publishers/publisherName/artifactTypes/vmExtension/types",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/pUbLiShErS/pUbLiShErNaMe/aRtIfAcTtYpEs/vMeXtEnSiOn/tYpEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/publishers/publisherName/artifactTypes/vmExtension/types/typeName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/pUbLiShErS/pUbLiShErNaMe/aRtIfAcTtYpEs/vMeXtEnSiOn/tYpEs/tYpEnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/publishers/publisherName/artifactTypes/vmExtension/types/typeName/versions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/pUbLiShErS/pUbLiShErNaMe/aRtIfAcTtYpEs/vMeXtEnSiOn/tYpEs/tYpEnAmE/vErSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/publishers/publisherName/artifactTypes/vmExtension/types/typeName/versions/versionName",
			Expected: &VersionId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				LocationName:   "locationName",
				PublisherName:  "publisherName",
				TypeName:       "typeName",
				VersionName:    "versionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationName/publishers/publisherName/artifactTypes/vmExtension/types/typeName/versions/versionName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/pUbLiShErS/pUbLiShErNaMe/aRtIfAcTtYpEs/vMeXtEnSiOn/tYpEs/tYpEnAmE/vErSiOnS/vErSiOnNaMe",
			Expected: &VersionId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				LocationName:   "lOcAtIoNnAmE",
				PublisherName:  "pUbLiShErNaMe",
				TypeName:       "tYpEnAmE",
				VersionName:    "vErSiOnNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNnAmE/pUbLiShErS/pUbLiShErNaMe/aRtIfAcTtYpEs/vMeXtEnSiOn/tYpEs/tYpEnAmE/vErSiOnS/vErSiOnNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVersionIDInsensitively(v.Input)
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

		if actual.PublisherName != v.Expected.PublisherName {
			t.Fatalf("Expected %q but got %q for PublisherName", v.Expected.PublisherName, actual.PublisherName)
		}

		if actual.TypeName != v.Expected.TypeName {
			t.Fatalf("Expected %q but got %q for TypeName", v.Expected.TypeName, actual.TypeName)
		}

		if actual.VersionName != v.Expected.VersionName {
			t.Fatalf("Expected %q but got %q for VersionName", v.Expected.VersionName, actual.VersionName)
		}

	}
}

func TestSegmentsForVersionId(t *testing.T) {
	segments := VersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("VersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
