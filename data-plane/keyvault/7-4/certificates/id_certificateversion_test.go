package certificates

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &CertificateversionId{}

func TestNewCertificateversionID(t *testing.T) {
	id := NewCertificateversionID("https://endpoint-url.example.com", "certificateName", "certificateversion")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.CertificateName != "certificateName" {
		t.Fatalf("Expected %q but got %q for Segment 'CertificateName'", id.CertificateName, "certificateName")
	}

	if id.Certificateversion != "certificateversion" {
		t.Fatalf("Expected %q but got %q for Segment 'Certificateversion'", id.Certificateversion, "certificateversion")
	}
}

func TestFormatCertificateversionID(t *testing.T) {
	actual := NewCertificateversionID("https://endpoint-url.example.com", "certificateName", "certificateversion").ID()
	expected := "https://endpoint-url.example.com/certificates/certificateName/certificateversion"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseCertificateversionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CertificateversionId
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
			Input: "https://endpoint-url.example.com/certificates",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/certificates/certificateName",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/certificates/certificateName/certificateversion",
			Expected: &CertificateversionId{
				BaseURI:            "https://endpoint-url.example.com",
				CertificateName:    "certificateName",
				Certificateversion: "certificateversion",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/certificates/certificateName/certificateversion/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCertificateversionID(v.Input)
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

		if actual.CertificateName != v.Expected.CertificateName {
			t.Fatalf("Expected %q but got %q for CertificateName", v.Expected.CertificateName, actual.CertificateName)
		}

		if actual.Certificateversion != v.Expected.Certificateversion {
			t.Fatalf("Expected %q but got %q for Certificateversion", v.Expected.Certificateversion, actual.Certificateversion)
		}

	}
}

func TestParseCertificateversionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CertificateversionId
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
			Input: "https://endpoint-url.example.com/certificates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/cErTiFiCaTeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/certificates/certificateName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/cErTiFiCaTeS/cErTiFiCaTeNaMe",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/certificates/certificateName/certificateversion",
			Expected: &CertificateversionId{
				BaseURI:            "https://endpoint-url.example.com",
				CertificateName:    "certificateName",
				Certificateversion: "certificateversion",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/certificates/certificateName/certificateversion/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/cErTiFiCaTeS/cErTiFiCaTeNaMe/cErTiFiCaTeVeRsIoN",
			Expected: &CertificateversionId{
				BaseURI:            "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				CertificateName:    "cErTiFiCaTeNaMe",
				Certificateversion: "cErTiFiCaTeVeRsIoN",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/cErTiFiCaTeS/cErTiFiCaTeNaMe/cErTiFiCaTeVeRsIoN/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCertificateversionIDInsensitively(v.Input)
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

		if actual.CertificateName != v.Expected.CertificateName {
			t.Fatalf("Expected %q but got %q for CertificateName", v.Expected.CertificateName, actual.CertificateName)
		}

		if actual.Certificateversion != v.Expected.Certificateversion {
			t.Fatalf("Expected %q but got %q for Certificateversion", v.Expected.Certificateversion, actual.Certificateversion)
		}

	}
}

func TestSegmentsForCertificateversionId(t *testing.T) {
	segments := CertificateversionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("CertificateversionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
