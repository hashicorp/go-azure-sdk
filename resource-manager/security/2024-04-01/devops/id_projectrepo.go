package devops

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ProjectRepoId{})
}

var _ resourceids.ResourceId = &ProjectRepoId{}

// ProjectRepoId is a struct representing the Resource ID for a Project Repo
type ProjectRepoId struct {
	SubscriptionId        string
	ResourceGroupName     string
	SecurityConnectorName string
	AzureDevOpsOrgName    string
	ProjectName           string
	RepoName              string
}

// NewProjectRepoID returns a new ProjectRepoId struct
func NewProjectRepoID(subscriptionId string, resourceGroupName string, securityConnectorName string, azureDevOpsOrgName string, projectName string, repoName string) ProjectRepoId {
	return ProjectRepoId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		SecurityConnectorName: securityConnectorName,
		AzureDevOpsOrgName:    azureDevOpsOrgName,
		ProjectName:           projectName,
		RepoName:              repoName,
	}
}

// ParseProjectRepoID parses 'input' into a ProjectRepoId
func ParseProjectRepoID(input string) (*ProjectRepoId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ProjectRepoId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ProjectRepoId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseProjectRepoIDInsensitively parses 'input' case-insensitively into a ProjectRepoId
// note: this method should only be used for API response data and not user input
func ParseProjectRepoIDInsensitively(input string) (*ProjectRepoId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ProjectRepoId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ProjectRepoId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ProjectRepoId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.SecurityConnectorName, ok = input.Parsed["securityConnectorName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "securityConnectorName", input)
	}

	if id.AzureDevOpsOrgName, ok = input.Parsed["azureDevOpsOrgName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "azureDevOpsOrgName", input)
	}

	if id.ProjectName, ok = input.Parsed["projectName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "projectName", input)
	}

	if id.RepoName, ok = input.Parsed["repoName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "repoName", input)
	}

	return nil
}

// ValidateProjectRepoID checks that 'input' can be parsed as a Project Repo ID
func ValidateProjectRepoID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProjectRepoID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Project Repo ID
func (id ProjectRepoId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Security/securityConnectors/%s/devops/default/azureDevOpsOrgs/%s/projects/%s/repos/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SecurityConnectorName, id.AzureDevOpsOrgName, id.ProjectName, id.RepoName)
}

// Segments returns a slice of Resource ID Segments which comprise this Project Repo ID
func (id ProjectRepoId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticSecurityConnectors", "securityConnectors", "securityConnectors"),
		resourceids.UserSpecifiedSegment("securityConnectorName", "securityConnectorValue"),
		resourceids.StaticSegment("staticDevops", "devops", "devops"),
		resourceids.StaticSegment("staticDefault", "default", "default"),
		resourceids.StaticSegment("staticAzureDevOpsOrgs", "azureDevOpsOrgs", "azureDevOpsOrgs"),
		resourceids.UserSpecifiedSegment("azureDevOpsOrgName", "azureDevOpsOrgValue"),
		resourceids.StaticSegment("staticProjects", "projects", "projects"),
		resourceids.UserSpecifiedSegment("projectName", "projectValue"),
		resourceids.StaticSegment("staticRepos", "repos", "repos"),
		resourceids.UserSpecifiedSegment("repoName", "repoValue"),
	}
}

// String returns a human-readable description of this Project Repo ID
func (id ProjectRepoId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Security Connector Name: %q", id.SecurityConnectorName),
		fmt.Sprintf("Azure Dev Ops Org Name: %q", id.AzureDevOpsOrgName),
		fmt.Sprintf("Project Name: %q", id.ProjectName),
		fmt.Sprintf("Repo Name: %q", id.RepoName),
	}
	return fmt.Sprintf("Project Repo (%s)", strings.Join(components, "\n"))
}
