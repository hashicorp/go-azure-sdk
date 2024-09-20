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
	recaser.RegisterResourceId(&RepoId{})
}

var _ resourceids.ResourceId = &RepoId{}

// RepoId is a struct representing the Resource ID for a Repo
type RepoId struct {
	SubscriptionId        string
	ResourceGroupName     string
	SecurityConnectorName string
	GitHubOwnerName       string
	RepoName              string
}

// NewRepoID returns a new RepoId struct
func NewRepoID(subscriptionId string, resourceGroupName string, securityConnectorName string, gitHubOwnerName string, repoName string) RepoId {
	return RepoId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		SecurityConnectorName: securityConnectorName,
		GitHubOwnerName:       gitHubOwnerName,
		RepoName:              repoName,
	}
}

// ParseRepoID parses 'input' into a RepoId
func ParseRepoID(input string) (*RepoId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RepoId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RepoId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRepoIDInsensitively parses 'input' case-insensitively into a RepoId
// note: this method should only be used for API response data and not user input
func ParseRepoIDInsensitively(input string) (*RepoId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RepoId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RepoId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RepoId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.GitHubOwnerName, ok = input.Parsed["gitHubOwnerName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "gitHubOwnerName", input)
	}

	if id.RepoName, ok = input.Parsed["repoName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "repoName", input)
	}

	return nil
}

// ValidateRepoID checks that 'input' can be parsed as a Repo ID
func ValidateRepoID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRepoID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Repo ID
func (id RepoId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Security/securityConnectors/%s/devops/default/gitHubOwners/%s/repos/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SecurityConnectorName, id.GitHubOwnerName, id.RepoName)
}

// Segments returns a slice of Resource ID Segments which comprise this Repo ID
func (id RepoId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticSecurityConnectors", "securityConnectors", "securityConnectors"),
		resourceids.UserSpecifiedSegment("securityConnectorName", "securityConnectorName"),
		resourceids.StaticSegment("staticDevops", "devops", "devops"),
		resourceids.StaticSegment("staticDefault", "default", "default"),
		resourceids.StaticSegment("staticGitHubOwners", "gitHubOwners", "gitHubOwners"),
		resourceids.UserSpecifiedSegment("gitHubOwnerName", "ownerName"),
		resourceids.StaticSegment("staticRepos", "repos", "repos"),
		resourceids.UserSpecifiedSegment("repoName", "repoName"),
	}
}

// String returns a human-readable description of this Repo ID
func (id RepoId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Security Connector Name: %q", id.SecurityConnectorName),
		fmt.Sprintf("Git Hub Owner Name: %q", id.GitHubOwnerName),
		fmt.Sprintf("Repo Name: %q", id.RepoName),
	}
	return fmt.Sprintf("Repo (%s)", strings.Join(components, "\n"))
}
