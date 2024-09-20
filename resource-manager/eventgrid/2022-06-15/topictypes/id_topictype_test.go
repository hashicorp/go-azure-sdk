package topictypes

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &TopicTypeId{}

func TestNewTopicTypeID(t *testing.T) {
	id := NewTopicTypeID("topicTypeName")

	if id.TopicTypeName != "topicTypeName" {
		t.Fatalf("Expected %q but got %q for Segment 'TopicTypeName'", id.TopicTypeName, "topicTypeName")
	}
}

func TestFormatTopicTypeID(t *testing.T) {
	actual := NewTopicTypeID("topicTypeName").ID()
	expected := "/providers/Microsoft.EventGrid/topicTypes/topicTypeName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseTopicTypeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *TopicTypeId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.EventGrid",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.EventGrid/topicTypes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.EventGrid/topicTypes/topicTypeName",
			Expected: &TopicTypeId{
				TopicTypeName: "topicTypeName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.EventGrid/topicTypes/topicTypeName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseTopicTypeID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.TopicTypeName != v.Expected.TopicTypeName {
			t.Fatalf("Expected %q but got %q for TopicTypeName", v.Expected.TopicTypeName, actual.TopicTypeName)
		}

	}
}

func TestParseTopicTypeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *TopicTypeId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.EventGrid",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.eVeNtGrId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.EventGrid/topicTypes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.eVeNtGrId/tOpIcTyPeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.EventGrid/topicTypes/topicTypeName",
			Expected: &TopicTypeId{
				TopicTypeName: "topicTypeName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.EventGrid/topicTypes/topicTypeName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.eVeNtGrId/tOpIcTyPeS/tOpIcTyPeNaMe",
			Expected: &TopicTypeId{
				TopicTypeName: "tOpIcTyPeNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.eVeNtGrId/tOpIcTyPeS/tOpIcTyPeNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseTopicTypeIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.TopicTypeName != v.Expected.TopicTypeName {
			t.Fatalf("Expected %q but got %q for TopicTypeName", v.Expected.TopicTypeName, actual.TopicTypeName)
		}

	}
}

func TestSegmentsForTopicTypeId(t *testing.T) {
	segments := TopicTypeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("TopicTypeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
