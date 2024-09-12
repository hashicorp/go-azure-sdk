package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeInferenceClassificationOverrideId{}

func TestNewMeInferenceClassificationOverrideID(t *testing.T) {
	id := NewMeInferenceClassificationOverrideID("inferenceClassificationOverrideIdValue")

	if id.InferenceClassificationOverrideId != "inferenceClassificationOverrideIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'InferenceClassificationOverrideId'", id.InferenceClassificationOverrideId, "inferenceClassificationOverrideIdValue")
	}
}

func TestFormatMeInferenceClassificationOverrideID(t *testing.T) {
	actual := NewMeInferenceClassificationOverrideID("inferenceClassificationOverrideIdValue").ID()
	expected := "/me/inferenceClassification/overrides/inferenceClassificationOverrideIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeInferenceClassificationOverrideID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeInferenceClassificationOverrideId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/inferenceClassification",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/inferenceClassification/overrides",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/inferenceClassification/overrides/inferenceClassificationOverrideIdValue",
			Expected: &MeInferenceClassificationOverrideId{
				InferenceClassificationOverrideId: "inferenceClassificationOverrideIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/inferenceClassification/overrides/inferenceClassificationOverrideIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeInferenceClassificationOverrideID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.InferenceClassificationOverrideId != v.Expected.InferenceClassificationOverrideId {
			t.Fatalf("Expected %q but got %q for InferenceClassificationOverrideId", v.Expected.InferenceClassificationOverrideId, actual.InferenceClassificationOverrideId)
		}

	}
}

func TestParseMeInferenceClassificationOverrideIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeInferenceClassificationOverrideId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/inferenceClassification",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfErEnCeClAsSiFiCaTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/inferenceClassification/overrides",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfErEnCeClAsSiFiCaTiOn/oVeRrIdEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/inferenceClassification/overrides/inferenceClassificationOverrideIdValue",
			Expected: &MeInferenceClassificationOverrideId{
				InferenceClassificationOverrideId: "inferenceClassificationOverrideIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/inferenceClassification/overrides/inferenceClassificationOverrideIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfErEnCeClAsSiFiCaTiOn/oVeRrIdEs/iNfErEnCeClAsSiFiCaTiOnOvErRiDeIdVaLuE",
			Expected: &MeInferenceClassificationOverrideId{
				InferenceClassificationOverrideId: "iNfErEnCeClAsSiFiCaTiOnOvErRiDeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfErEnCeClAsSiFiCaTiOn/oVeRrIdEs/iNfErEnCeClAsSiFiCaTiOnOvErRiDeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeInferenceClassificationOverrideIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.InferenceClassificationOverrideId != v.Expected.InferenceClassificationOverrideId {
			t.Fatalf("Expected %q but got %q for InferenceClassificationOverrideId", v.Expected.InferenceClassificationOverrideId, actual.InferenceClassificationOverrideId)
		}

	}
}

func TestSegmentsForMeInferenceClassificationOverrideId(t *testing.T) {
	segments := MeInferenceClassificationOverrideId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeInferenceClassificationOverrideId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
