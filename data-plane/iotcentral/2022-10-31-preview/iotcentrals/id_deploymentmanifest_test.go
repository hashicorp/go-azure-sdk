package iotcentrals

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeploymentManifestId{}

func TestNewDeploymentManifestID(t *testing.T) {
	id := NewDeploymentManifestID("https://endpoint-url.example.com", "deploymentManifestId")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.DeploymentManifestId != "deploymentManifestId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeploymentManifestId'", id.DeploymentManifestId, "deploymentManifestId")
	}
}

func TestFormatDeploymentManifestID(t *testing.T) {
	actual := NewDeploymentManifestID("https://endpoint-url.example.com", "deploymentManifestId").ID()
	expected := "https://endpoint-url.example.com/deploymentManifests/deploymentManifestId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeploymentManifestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeploymentManifestId
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
			Input: "https://endpoint-url.example.com/deploymentManifests",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/deploymentManifests/deploymentManifestId",
			Expected: &DeploymentManifestId{
				BaseURI:              "https://endpoint-url.example.com",
				DeploymentManifestId: "deploymentManifestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/deploymentManifests/deploymentManifestId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeploymentManifestID(v.Input)
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

		if actual.DeploymentManifestId != v.Expected.DeploymentManifestId {
			t.Fatalf("Expected %q but got %q for DeploymentManifestId", v.Expected.DeploymentManifestId, actual.DeploymentManifestId)
		}

	}
}

func TestParseDeploymentManifestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeploymentManifestId
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
			Input: "https://endpoint-url.example.com/deploymentManifests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dEpLoYmEnTmAnIfEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/deploymentManifests/deploymentManifestId",
			Expected: &DeploymentManifestId{
				BaseURI:              "https://endpoint-url.example.com",
				DeploymentManifestId: "deploymentManifestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/deploymentManifests/deploymentManifestId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dEpLoYmEnTmAnIfEsTs/dEpLoYmEnTmAnIfEsTiD",
			Expected: &DeploymentManifestId{
				BaseURI:              "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				DeploymentManifestId: "dEpLoYmEnTmAnIfEsTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dEpLoYmEnTmAnIfEsTs/dEpLoYmEnTmAnIfEsTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeploymentManifestIDInsensitively(v.Input)
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

		if actual.DeploymentManifestId != v.Expected.DeploymentManifestId {
			t.Fatalf("Expected %q but got %q for DeploymentManifestId", v.Expected.DeploymentManifestId, actual.DeploymentManifestId)
		}

	}
}

func TestSegmentsForDeploymentManifestId(t *testing.T) {
	segments := DeploymentManifestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeploymentManifestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
