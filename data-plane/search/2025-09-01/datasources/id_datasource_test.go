package datasources

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DatasourceId{}

func TestNewDatasourceID(t *testing.T) {
	id := NewDatasourceID("https://endpoint-url.example.com", "datasourceName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.DatasourceName != "datasourceName" {
		t.Fatalf("Expected %q but got %q for Segment 'DatasourceName'", id.DatasourceName, "datasourceName")
	}
}

func TestFormatDatasourceID(t *testing.T) {
	actual := NewDatasourceID("https://endpoint-url.example.com", "datasourceName").ID()
	expected := "https://endpoint-url.example.com/datasources/datasourceName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDatasourceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DatasourceId
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
			Input: "https://endpoint-url.example.com/datasources",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/datasources/datasourceName",
			Expected: &DatasourceId{
				BaseURI:        "https://endpoint-url.example.com",
				DatasourceName: "datasourceName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/datasources/datasourceName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDatasourceID(v.Input)
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

		if actual.DatasourceName != v.Expected.DatasourceName {
			t.Fatalf("Expected %q but got %q for DatasourceName", v.Expected.DatasourceName, actual.DatasourceName)
		}

	}
}

func TestParseDatasourceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DatasourceId
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
			Input: "https://endpoint-url.example.com/datasources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dAtAsOuRcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/datasources/datasourceName",
			Expected: &DatasourceId{
				BaseURI:        "https://endpoint-url.example.com",
				DatasourceName: "datasourceName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/datasources/datasourceName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dAtAsOuRcEs/dAtAsOuRcEnAmE",
			Expected: &DatasourceId{
				BaseURI:        "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				DatasourceName: "dAtAsOuRcEnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dAtAsOuRcEs/dAtAsOuRcEnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDatasourceIDInsensitively(v.Input)
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

		if actual.DatasourceName != v.Expected.DatasourceName {
			t.Fatalf("Expected %q but got %q for DatasourceName", v.Expected.DatasourceName, actual.DatasourceName)
		}

	}
}

func TestSegmentsForDatasourceId(t *testing.T) {
	segments := DatasourceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DatasourceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
