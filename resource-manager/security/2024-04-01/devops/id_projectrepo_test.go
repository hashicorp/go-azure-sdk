package devops

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ProjectRepoId{}

func TestNewProjectRepoID(t *testing.T) {
	id := NewProjectRepoID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorValue", "azureDevOpsOrgValue", "projectValue", "repoValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.SecurityConnectorName != "securityConnectorValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SecurityConnectorName'", id.SecurityConnectorName, "securityConnectorValue")
	}

	if id.AzureDevOpsOrgName != "azureDevOpsOrgValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AzureDevOpsOrgName'", id.AzureDevOpsOrgName, "azureDevOpsOrgValue")
	}

	if id.ProjectName != "projectValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ProjectName'", id.ProjectName, "projectValue")
	}

	if id.RepoName != "repoValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RepoName'", id.RepoName, "repoValue")
	}
}

func TestFormatProjectRepoID(t *testing.T) {
	actual := NewProjectRepoID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorValue", "azureDevOpsOrgValue", "projectValue", "repoValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorValue/devops/default/azureDevOpsOrgs/azureDevOpsOrgValue/projects/projectValue/repos/repoValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseProjectRepoID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ProjectRepoId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorValue/devops",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorValue/devops/default",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorValue/devops/default/azureDevOpsOrgs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorValue/devops/default/azureDevOpsOrgs/azureDevOpsOrgValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorValue/devops/default/azureDevOpsOrgs/azureDevOpsOrgValue/projects",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorValue/devops/default/azureDevOpsOrgs/azureDevOpsOrgValue/projects/projectValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorValue/devops/default/azureDevOpsOrgs/azureDevOpsOrgValue/projects/projectValue/repos",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorValue/devops/default/azureDevOpsOrgs/azureDevOpsOrgValue/projects/projectValue/repos/repoValue",
			Expected: &ProjectRepoId{
				SubscriptionId:        "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:     "example-resource-group",
				SecurityConnectorName: "securityConnectorValue",
				AzureDevOpsOrgName:    "azureDevOpsOrgValue",
				ProjectName:           "projectValue",
				RepoName:              "repoValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorValue/devops/default/azureDevOpsOrgs/azureDevOpsOrgValue/projects/projectValue/repos/repoValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProjectRepoID(v.Input)
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

		if actual.AzureDevOpsOrgName != v.Expected.AzureDevOpsOrgName {
			t.Fatalf("Expected %q but got %q for AzureDevOpsOrgName", v.Expected.AzureDevOpsOrgName, actual.AzureDevOpsOrgName)
		}

		if actual.ProjectName != v.Expected.ProjectName {
			t.Fatalf("Expected %q but got %q for ProjectName", v.Expected.ProjectName, actual.ProjectName)
		}

		if actual.RepoName != v.Expected.RepoName {
			t.Fatalf("Expected %q but got %q for RepoName", v.Expected.RepoName, actual.RepoName)
		}

	}
}

func TestParseProjectRepoIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ProjectRepoId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/sEcUrItYcOnNeCtOrS/sEcUrItYcOnNeCtOrVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorValue/devops",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/sEcUrItYcOnNeCtOrS/sEcUrItYcOnNeCtOrVaLuE/dEvOpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorValue/devops/default",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/sEcUrItYcOnNeCtOrS/sEcUrItYcOnNeCtOrVaLuE/dEvOpS/dEfAuLt",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorValue/devops/default/azureDevOpsOrgs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/sEcUrItYcOnNeCtOrS/sEcUrItYcOnNeCtOrVaLuE/dEvOpS/dEfAuLt/aZuReDeVoPsOrGs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorValue/devops/default/azureDevOpsOrgs/azureDevOpsOrgValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/sEcUrItYcOnNeCtOrS/sEcUrItYcOnNeCtOrVaLuE/dEvOpS/dEfAuLt/aZuReDeVoPsOrGs/aZuReDeVoPsOrGvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorValue/devops/default/azureDevOpsOrgs/azureDevOpsOrgValue/projects",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/sEcUrItYcOnNeCtOrS/sEcUrItYcOnNeCtOrVaLuE/dEvOpS/dEfAuLt/aZuReDeVoPsOrGs/aZuReDeVoPsOrGvAlUe/pRoJeCtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorValue/devops/default/azureDevOpsOrgs/azureDevOpsOrgValue/projects/projectValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/sEcUrItYcOnNeCtOrS/sEcUrItYcOnNeCtOrVaLuE/dEvOpS/dEfAuLt/aZuReDeVoPsOrGs/aZuReDeVoPsOrGvAlUe/pRoJeCtS/pRoJeCtVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorValue/devops/default/azureDevOpsOrgs/azureDevOpsOrgValue/projects/projectValue/repos",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/sEcUrItYcOnNeCtOrS/sEcUrItYcOnNeCtOrVaLuE/dEvOpS/dEfAuLt/aZuReDeVoPsOrGs/aZuReDeVoPsOrGvAlUe/pRoJeCtS/pRoJeCtVaLuE/rEpOs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorValue/devops/default/azureDevOpsOrgs/azureDevOpsOrgValue/projects/projectValue/repos/repoValue",
			Expected: &ProjectRepoId{
				SubscriptionId:        "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:     "example-resource-group",
				SecurityConnectorName: "securityConnectorValue",
				AzureDevOpsOrgName:    "azureDevOpsOrgValue",
				ProjectName:           "projectValue",
				RepoName:              "repoValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Security/securityConnectors/securityConnectorValue/devops/default/azureDevOpsOrgs/azureDevOpsOrgValue/projects/projectValue/repos/repoValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/sEcUrItYcOnNeCtOrS/sEcUrItYcOnNeCtOrVaLuE/dEvOpS/dEfAuLt/aZuReDeVoPsOrGs/aZuReDeVoPsOrGvAlUe/pRoJeCtS/pRoJeCtVaLuE/rEpOs/rEpOvAlUe",
			Expected: &ProjectRepoId{
				SubscriptionId:        "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:     "eXaMpLe-rEsOuRcE-GrOuP",
				SecurityConnectorName: "sEcUrItYcOnNeCtOrVaLuE",
				AzureDevOpsOrgName:    "aZuReDeVoPsOrGvAlUe",
				ProjectName:           "pRoJeCtVaLuE",
				RepoName:              "rEpOvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sEcUrItY/sEcUrItYcOnNeCtOrS/sEcUrItYcOnNeCtOrVaLuE/dEvOpS/dEfAuLt/aZuReDeVoPsOrGs/aZuReDeVoPsOrGvAlUe/pRoJeCtS/pRoJeCtVaLuE/rEpOs/rEpOvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProjectRepoIDInsensitively(v.Input)
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

		if actual.AzureDevOpsOrgName != v.Expected.AzureDevOpsOrgName {
			t.Fatalf("Expected %q but got %q for AzureDevOpsOrgName", v.Expected.AzureDevOpsOrgName, actual.AzureDevOpsOrgName)
		}

		if actual.ProjectName != v.Expected.ProjectName {
			t.Fatalf("Expected %q but got %q for ProjectName", v.Expected.ProjectName, actual.ProjectName)
		}

		if actual.RepoName != v.Expected.RepoName {
			t.Fatalf("Expected %q but got %q for RepoName", v.Expected.RepoName, actual.RepoName)
		}

	}
}

func TestSegmentsForProjectRepoId(t *testing.T) {
	segments := ProjectRepoId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ProjectRepoId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
