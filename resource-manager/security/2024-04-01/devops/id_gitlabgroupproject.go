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
	recaser.RegisterResourceId(&GitLabGroupProjectId{})
}

var _ resourceids.ResourceId = &GitLabGroupProjectId{}

// GitLabGroupProjectId is a struct representing the Resource ID for a Git Lab Group Project
type GitLabGroupProjectId struct {
	SubscriptionId        string
	ResourceGroupName     string
	SecurityConnectorName string
	GitLabGroupName       string
	ProjectName           string
}

// NewGitLabGroupProjectID returns a new GitLabGroupProjectId struct
func NewGitLabGroupProjectID(subscriptionId string, resourceGroupName string, securityConnectorName string, gitLabGroupName string, projectName string) GitLabGroupProjectId {
	return GitLabGroupProjectId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		SecurityConnectorName: securityConnectorName,
		GitLabGroupName:       gitLabGroupName,
		ProjectName:           projectName,
	}
}

// ParseGitLabGroupProjectID parses 'input' into a GitLabGroupProjectId
func ParseGitLabGroupProjectID(input string) (*GitLabGroupProjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GitLabGroupProjectId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GitLabGroupProjectId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGitLabGroupProjectIDInsensitively parses 'input' case-insensitively into a GitLabGroupProjectId
// note: this method should only be used for API response data and not user input
func ParseGitLabGroupProjectIDInsensitively(input string) (*GitLabGroupProjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GitLabGroupProjectId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GitLabGroupProjectId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GitLabGroupProjectId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.GitLabGroupName, ok = input.Parsed["gitLabGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "gitLabGroupName", input)
	}

	if id.ProjectName, ok = input.Parsed["projectName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "projectName", input)
	}

	return nil
}

// ValidateGitLabGroupProjectID checks that 'input' can be parsed as a Git Lab Group Project ID
func ValidateGitLabGroupProjectID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGitLabGroupProjectID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Git Lab Group Project ID
func (id GitLabGroupProjectId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Security/securityConnectors/%s/devops/default/gitLabGroups/%s/projects/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SecurityConnectorName, id.GitLabGroupName, id.ProjectName)
}

// Segments returns a slice of Resource ID Segments which comprise this Git Lab Group Project ID
func (id GitLabGroupProjectId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticGitLabGroups", "gitLabGroups", "gitLabGroups"),
		resourceids.UserSpecifiedSegment("gitLabGroupName", "gitLabGroupValue"),
		resourceids.StaticSegment("staticProjects", "projects", "projects"),
		resourceids.UserSpecifiedSegment("projectName", "projectValue"),
	}
}

// String returns a human-readable description of this Git Lab Group Project ID
func (id GitLabGroupProjectId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Security Connector Name: %q", id.SecurityConnectorName),
		fmt.Sprintf("Git Lab Group Name: %q", id.GitLabGroupName),
		fmt.Sprintf("Project Name: %q", id.ProjectName),
	}
	return fmt.Sprintf("Git Lab Group Project (%s)", strings.Join(components, "\n"))
}
