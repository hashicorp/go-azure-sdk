package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdInferenceClassificationOverrideId{}

func TestNewUserIdInferenceClassificationOverrideID(t *testing.T) {
	id := NewUserIdInferenceClassificationOverrideID("userId", "inferenceClassificationOverrideId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.InferenceClassificationOverrideId != "inferenceClassificationOverrideId" {
		t.Fatalf("Expected %q but got %q for Segment 'InferenceClassificationOverrideId'", id.InferenceClassificationOverrideId, "inferenceClassificationOverrideId")
	}
}

func TestFormatUserIdInferenceClassificationOverrideID(t *testing.T) {
	actual := NewUserIdInferenceClassificationOverrideID("userId", "inferenceClassificationOverrideId").ID()
	expected := "/users/userId/inferenceClassification/overrides/inferenceClassificationOverrideId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdInferenceClassificationOverrideID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdInferenceClassificationOverrideId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/inferenceClassification",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/inferenceClassification/overrides",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/inferenceClassification/overrides/inferenceClassificationOverrideId",
			Expected: &UserIdInferenceClassificationOverrideId{
				UserId:                            "userId",
				InferenceClassificationOverrideId: "inferenceClassificationOverrideId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/inferenceClassification/overrides/inferenceClassificationOverrideId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdInferenceClassificationOverrideID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.InferenceClassificationOverrideId != v.Expected.InferenceClassificationOverrideId {
			t.Fatalf("Expected %q but got %q for InferenceClassificationOverrideId", v.Expected.InferenceClassificationOverrideId, actual.InferenceClassificationOverrideId)
		}

	}
}

func TestParseUserIdInferenceClassificationOverrideIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdInferenceClassificationOverrideId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/inferenceClassification",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/iNfErEnCeClAsSiFiCaTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/inferenceClassification/overrides",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/iNfErEnCeClAsSiFiCaTiOn/oVeRrIdEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/inferenceClassification/overrides/inferenceClassificationOverrideId",
			Expected: &UserIdInferenceClassificationOverrideId{
				UserId:                            "userId",
				InferenceClassificationOverrideId: "inferenceClassificationOverrideId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/inferenceClassification/overrides/inferenceClassificationOverrideId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/iNfErEnCeClAsSiFiCaTiOn/oVeRrIdEs/iNfErEnCeClAsSiFiCaTiOnOvErRiDeId",
			Expected: &UserIdInferenceClassificationOverrideId{
				UserId:                            "uSeRiD",
				InferenceClassificationOverrideId: "iNfErEnCeClAsSiFiCaTiOnOvErRiDeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/iNfErEnCeClAsSiFiCaTiOn/oVeRrIdEs/iNfErEnCeClAsSiFiCaTiOnOvErRiDeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdInferenceClassificationOverrideIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.InferenceClassificationOverrideId != v.Expected.InferenceClassificationOverrideId {
			t.Fatalf("Expected %q but got %q for InferenceClassificationOverrideId", v.Expected.InferenceClassificationOverrideId, actual.InferenceClassificationOverrideId)
		}

	}
}

func TestSegmentsForUserIdInferenceClassificationOverrideId(t *testing.T) {
	segments := UserIdInferenceClassificationOverrideId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdInferenceClassificationOverrideId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
