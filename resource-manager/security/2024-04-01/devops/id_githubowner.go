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
	recaser.RegisterResourceId(&GitHubOwnerId{})
}

var _ resourceids.ResourceId = &GitHubOwnerId{}

// GitHubOwnerId is a struct representing the Resource ID for a Git Hub Owner
type GitHubOwnerId struct {
	SubscriptionId        string
	ResourceGroupName     string
	SecurityConnectorName string
	GitHubOwnerName       string
}

// NewGitHubOwnerID returns a new GitHubOwnerId struct
func NewGitHubOwnerID(subscriptionId string, resourceGroupName string, securityConnectorName string, gitHubOwnerName string) GitHubOwnerId {
	return GitHubOwnerId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		SecurityConnectorName: securityConnectorName,
		GitHubOwnerName:       gitHubOwnerName,
	}
}

// ParseGitHubOwnerID parses 'input' into a GitHubOwnerId
func ParseGitHubOwnerID(input string) (*GitHubOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GitHubOwnerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GitHubOwnerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGitHubOwnerIDInsensitively parses 'input' case-insensitively into a GitHubOwnerId
// note: this method should only be used for API response data and not user input
func ParseGitHubOwnerIDInsensitively(input string) (*GitHubOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GitHubOwnerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GitHubOwnerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GitHubOwnerId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateGitHubOwnerID checks that 'input' can be parsed as a Git Hub Owner ID
func ValidateGitHubOwnerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGitHubOwnerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Git Hub Owner ID
func (id GitHubOwnerId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Security/securityConnectors/%s/devops/default/gitHubOwners/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SecurityConnectorName, id.GitHubOwnerName)
}

// Segments returns a slice of Resource ID Segments which comprise this Git Hub Owner ID
func (id GitHubOwnerId) Segments() []resourceids.Segment {
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
	}
}

// String returns a human-readable description of this Git Hub Owner ID
func (id GitHubOwnerId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Security Connector Name: %q", id.SecurityConnectorName),
		fmt.Sprintf("Git Hub Owner Name: %q", id.GitHubOwnerName),
	}
	return fmt.Sprintf("Git Hub Owner (%s)", strings.Join(components, "\n"))
}
