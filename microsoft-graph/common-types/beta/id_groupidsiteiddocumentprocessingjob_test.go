package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdDocumentProcessingJobId{}

func TestNewGroupIdSiteIdDocumentProcessingJobID(t *testing.T) {
	id := NewGroupIdSiteIdDocumentProcessingJobID("groupIdValue", "siteIdValue", "documentProcessingJobIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.SiteId != "siteIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteIdValue")
	}

	if id.DocumentProcessingJobId != "documentProcessingJobIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DocumentProcessingJobId'", id.DocumentProcessingJobId, "documentProcessingJobIdValue")
	}
}

func TestFormatGroupIdSiteIdDocumentProcessingJobID(t *testing.T) {
	actual := NewGroupIdSiteIdDocumentProcessingJobID("groupIdValue", "siteIdValue", "documentProcessingJobIdValue").ID()
	expected := "/groups/groupIdValue/sites/siteIdValue/documentProcessingJobs/documentProcessingJobIdValue"
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
			Input: "/groups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/documentProcessingJobs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/documentProcessingJobs/documentProcessingJobIdValue",
			Expected: &GroupIdSiteIdDocumentProcessingJobId{
				GroupId:                 "groupIdValue",
				SiteId:                  "siteIdValue",
				DocumentProcessingJobId: "documentProcessingJobIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/documentProcessingJobs/documentProcessingJobIdValue/extra",
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
			Input: "/groups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/documentProcessingJobs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/dOcUmEnTpRoCeSsInGjObS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/documentProcessingJobs/documentProcessingJobIdValue",
			Expected: &GroupIdSiteIdDocumentProcessingJobId{
				GroupId:                 "groupIdValue",
				SiteId:                  "siteIdValue",
				DocumentProcessingJobId: "documentProcessingJobIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/documentProcessingJobs/documentProcessingJobIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/dOcUmEnTpRoCeSsInGjObS/dOcUmEnTpRoCeSsInGjObIdVaLuE",
			Expected: &GroupIdSiteIdDocumentProcessingJobId{
				GroupId:                 "gRoUpIdVaLuE",
				SiteId:                  "sItEiDvAlUe",
				DocumentProcessingJobId: "dOcUmEnTpRoCeSsInGjObIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/dOcUmEnTpRoCeSsInGjObS/dOcUmEnTpRoCeSsInGjObIdVaLuE/extra",
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
