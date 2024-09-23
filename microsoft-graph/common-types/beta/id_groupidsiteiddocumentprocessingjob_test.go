package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdDocumentProcessingJobId{}

func TestNewGroupIdSiteIdDocumentProcessingJobID(t *testing.T) {
	id := NewGroupIdSiteIdDocumentProcessingJobID("groupId", "siteId", "documentProcessingJobId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.SiteId != "siteId" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteId")
	}

	if id.DocumentProcessingJobId != "documentProcessingJobId" {
		t.Fatalf("Expected %q but got %q for Segment 'DocumentProcessingJobId'", id.DocumentProcessingJobId, "documentProcessingJobId")
	}
}

func TestFormatGroupIdSiteIdDocumentProcessingJobID(t *testing.T) {
	actual := NewGroupIdSiteIdDocumentProcessingJobID("groupId", "siteId", "documentProcessingJobId").ID()
	expected := "/groups/groupId/sites/siteId/documentProcessingJobs/documentProcessingJobId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdSiteIdDocumentProcessingJobID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdDocumentProcessingJobId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/documentProcessingJobs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/sites/siteId/documentProcessingJobs/documentProcessingJobId",
			Expected: &GroupIdSiteIdDocumentProcessingJobId{
				GroupId:                 "groupId",
				SiteId:                  "siteId",
				DocumentProcessingJobId: "documentProcessingJobId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/sites/siteId/documentProcessingJobs/documentProcessingJobId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdDocumentProcessingJobID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.SiteId != v.Expected.SiteId {
			t.Fatalf("Expected %q but got %q for SiteId", v.Expected.SiteId, actual.SiteId)
		}

		if actual.DocumentProcessingJobId != v.Expected.DocumentProcessingJobId {
			t.Fatalf("Expected %q but got %q for DocumentProcessingJobId", v.Expected.DocumentProcessingJobId, actual.DocumentProcessingJobId)
		}

	}
}

func TestParseGroupIdSiteIdDocumentProcessingJobIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdDocumentProcessingJobId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/documentProcessingJobs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/dOcUmEnTpRoCeSsInGjObS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/sites/siteId/documentProcessingJobs/documentProcessingJobId",
			Expected: &GroupIdSiteIdDocumentProcessingJobId{
				GroupId:                 "groupId",
				SiteId:                  "siteId",
				DocumentProcessingJobId: "documentProcessingJobId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/sites/siteId/documentProcessingJobs/documentProcessingJobId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/dOcUmEnTpRoCeSsInGjObS/dOcUmEnTpRoCeSsInGjObId",
			Expected: &GroupIdSiteIdDocumentProcessingJobId{
				GroupId:                 "gRoUpId",
				SiteId:                  "sItEiD",
				DocumentProcessingJobId: "dOcUmEnTpRoCeSsInGjObId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/dOcUmEnTpRoCeSsInGjObS/dOcUmEnTpRoCeSsInGjObId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdDocumentProcessingJobIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.SiteId != v.Expected.SiteId {
			t.Fatalf("Expected %q but got %q for SiteId", v.Expected.SiteId, actual.SiteId)
		}

		if actual.DocumentProcessingJobId != v.Expected.DocumentProcessingJobId {
			t.Fatalf("Expected %q but got %q for DocumentProcessingJobId", v.Expected.DocumentProcessingJobId, actual.DocumentProcessingJobId)
		}

	}
}

func TestSegmentsForGroupIdSiteIdDocumentProcessingJobId(t *testing.T) {
	segments := GroupIdSiteIdDocumentProcessingJobId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdSiteIdDocumentProcessingJobId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
