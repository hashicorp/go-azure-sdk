package devops

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GitLabGroupProjectId{}

func TestNewGitLabGroupProjectID(t *testing.T) {
	id := NewGitLabGroupProjectID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName", "gitLabGroupName", "projectName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.SecurityConnectorName != "securityConnectorName" {
		t.Fatalf("Expected %q but got %q for Segment 'SecurityConnectorName'", id.SecurityConnectorName, "securityConnectorName")
	}

	if id.GitLabGroupName != "gitLabGroupName" {
		t.Fatalf("Expected %q but got %q for Segment 'GitLabGroupName'", id.GitLabGroupName, "gitLabGroupName")
	}

	if id.ProjectName != "projectName" {
		t.Fatalf("Expected %q but got %q for Segment 'ProjectName'", id.ProjectName, "projectName")
	}
}

func TestFormatGitLabGroupProjectID(t *testing.T) {
	actual := NewGitLabGroupProjectID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName", "gitLabGroupName", "projectName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorName/devops/default/gitLabGroups/gitLabGroupName/projects/projectName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGitLabGroupProjectID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GitLabGroupProjectId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorName/devops",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorName/devops/default",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorName/devops/default/gitLabGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorName/devops/default/gitLabGroups/gitLabGroupName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorName/devops/default/gitLabGroups/gitLabGroupName/projects",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorName/devops/default/gitLabGroups/gitLabGroupName/projects/projectName",
			Expected: &GitLabGroupProjectId{
				SubscriptionId:        "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:     "example-resource-group",
				SecurityConnectorName: "securityConnectorName",
				GitLabGroupName:       "gitLabGroupName",
				ProjectName:           "projectName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorName/devops/default/gitLabGroups/gitLabGroupName/projects/projectName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGitLabGroupProjectID(v.Input)
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

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.SecurityConnectorName != v.Expected.SecurityConnectorName {
			t.Fatalf("Expected %q but got %q for SecurityConnectorName", v.Expected.SecurityConnectorName, actual.SecurityConnectorName)
		}

		if actual.GitLabGroupName != v.Expected.GitLabGroupName {
			t.Fatalf("Expected %q but got %q for GitLabGroupName", v.Expected.GitLabGroupName, actual.GitLabGroupName)
		}

		if actual.ProjectName != v.Expected.ProjectName {
			t.Fatalf("Expected %q but got %q for ProjectName", v.Expected.ProjectName, actual.ProjectName)
		}

	}
}

func TestParseGitLabGroupProjectIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GitLabGroupProjectId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/sEcUrItYcOnNeCtOrS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/sEcUrItYcOnNeCtOrS/sEcUrItYcOnNeCtOrNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorName/devops",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/sEcUrItYcOnNeCtOrS/sEcUrItYcOnNeCtOrNaMe/dEvOpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorName/devops/default",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/sEcUrItYcOnNeCtOrS/sEcUrItYcOnNeCtOrNaMe/dEvOpS/dEfAuLt",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorName/devops/default/gitLabGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/sEcUrItYcOnNeCtOrS/sEcUrItYcOnNeCtOrNaMe/dEvOpS/dEfAuLt/gItLaBgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorName/devops/default/gitLabGroups/gitLabGroupName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/sEcUrItYcOnNeCtOrS/sEcUrItYcOnNeCtOrNaMe/dEvOpS/dEfAuLt/gItLaBgRoUpS/gItLaBgRoUpNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorName/devops/default/gitLabGroups/gitLabGroupName/projects",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/sEcUrItYcOnNeCtOrS/sEcUrItYcOnNeCtOrNaMe/dEvOpS/dEfAuLt/gItLaBgRoUpS/gItLaBgRoUpNaMe/pRoJeCtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorName/devops/default/gitLabGroups/gitLabGroupName/projects/projectName",
			Expected: &GitLabGroupProjectId{
				SubscriptionId:        "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:     "example-resource-group",
				SecurityConnectorName: "securityConnectorName",
				GitLabGroupName:       "gitLabGroupName",
				ProjectName:           "projectName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorName/devops/default/gitLabGroups/gitLabGroupName/projects/projectName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/sEcUrItYcOnNeCtOrS/sEcUrItYcOnNeCtOrNaMe/dEvOpS/dEfAuLt/gItLaBgRoUpS/gItLaBgRoUpNaMe/pRoJeCtS/pRoJeCtNaMe",
			Expected: &GitLabGroupProjectId{
				SubscriptionId:        "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:     "eXaMpLe-rEsOuRcE-GrOuP",
				SecurityConnectorName: "sEcUrItYcOnNeCtOrNaMe",
				GitLabGroupName:       "gItLaBgRoUpNaMe",
				ProjectName:           "pRoJeCtNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/sEcUrItYcOnNeCtOrS/sEcUrItYcOnNeCtOrNaMe/dEvOpS/dEfAuLt/gItLaBgRoUpS/gItLaBgRoUpNaMe/pRoJeCtS/pRoJeCtNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGitLabGroupProjectIDInsensitively(v.Input)
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

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.SecurityConnectorName != v.Expected.SecurityConnectorName {
			t.Fatalf("Expected %q but got %q for SecurityConnectorName", v.Expected.SecurityConnectorName, actual.SecurityConnectorName)
		}

		if actual.GitLabGroupName != v.Expected.GitLabGroupName {
			t.Fatalf("Expected %q but got %q for GitLabGroupName", v.Expected.GitLabGroupName, actual.GitLabGroupName)
		}

		if actual.ProjectName != v.Expected.ProjectName {
			t.Fatalf("Expected %q but got %q for ProjectName", v.Expected.ProjectName, actual.ProjectName)
		}

	}
}

func TestSegmentsForGitLabGroupProjectId(t *testing.T) {
	segments := GitLabGroupProjectId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GitLabGroupProjectId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
