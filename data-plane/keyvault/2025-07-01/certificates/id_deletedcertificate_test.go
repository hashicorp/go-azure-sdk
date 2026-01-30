package certificates

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeletedcertificateId{}

func TestNewDeletedcertificateID(t *testing.T) {
	id := NewDeletedcertificateID("https://endpoint-url.example.com", "deletedcertificateName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.DeletedcertificateName != "deletedcertificateName" {
		t.Fatalf("Expected %q but got %q for Segment 'DeletedcertificateName'", id.DeletedcertificateName, "deletedcertificateName")
	}
}

func TestFormatDeletedcertificateID(t *testing.T) {
	actual := NewDeletedcertificateID("https://endpoint-url.example.com", "deletedcertificateName").ID()
	expected := "https://endpoint-url.example.com/deletedcertificates/deletedcertificateName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeletedcertificateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeletedcertificateId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/deletedcertificates",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/deletedcertificates/deletedcertificateName",
			Expected: &DeletedcertificateId{
				BaseURI:                "https://endpoint-url.example.com",
				DeletedcertificateName: "deletedcertificateName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/deletedcertificates/deletedcertificateName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeletedcertificateID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BaseURI != v.Expected.BaseURI {
			t.Fatalf("Expected %q but got %q for BaseURI", v.Expected.BaseURI, actual.BaseURI)
		}

		if actual.DeletedcertificateName != v.Expected.DeletedcertificateName {
			t.Fatalf("Expected %q but got %q for DeletedcertificateName", v.Expected.DeletedcertificateName, actual.DeletedcertificateName)
		}

	}
}

func TestParseDeletedcertificateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeletedcertificateId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/deletedcertificates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dElEtEdCeRtIfIcAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/deletedcertificates/deletedcertificateName",
			Expected: &DeletedcertificateId{
				BaseURI:                "https://endpoint-url.example.com",
				DeletedcertificateName: "deletedcertificateName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/deletedcertificates/deletedcertificateName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dElEtEdCeRtIfIcAtEs/dElEtEdCeRtIfIcAtEnAmE",
			Expected: &DeletedcertificateId{
				BaseURI:                "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				DeletedcertificateName: "dElEtEdCeRtIfIcAtEnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dElEtEdCeRtIfIcAtEs/dElEtEdCeRtIfIcAtEnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeletedcertificateIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BaseURI != v.Expected.BaseURI {
			t.Fatalf("Expected %q but got %q for BaseURI", v.Expected.BaseURI, actual.BaseURI)
		}

		if actual.DeletedcertificateName != v.Expected.DeletedcertificateName {
			t.Fatalf("Expected %q but got %q for DeletedcertificateName", v.Expected.DeletedcertificateName, actual.DeletedcertificateName)
		}

	}
}

func TestSegmentsForDeletedcertificateId(t *testing.T) {
	segments := DeletedcertificateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeletedcertificateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
