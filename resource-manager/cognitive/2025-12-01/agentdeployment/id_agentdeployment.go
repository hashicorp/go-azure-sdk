package agentdeployment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&AgentDeploymentId{})
}

var _ resourceids.ResourceId = &AgentDeploymentId{}

// AgentDeploymentId is a struct representing the Resource ID for a Agent Deployment
type AgentDeploymentId struct {
	SubscriptionId      string
	ResourceGroupName   string
	AccountName         string
	ProjectName         string
	ApplicationName     string
	AgentDeploymentName string
}

// NewAgentDeploymentID returns a new AgentDeploymentId struct
func NewAgentDeploymentID(subscriptionId string, resourceGroupName string, accountName string, projectName string, applicationName string, agentDeploymentName string) AgentDeploymentId {
	return AgentDeploymentId{
		SubscriptionId:      subscriptionId,
		ResourceGroupName:   resourceGroupName,
		AccountName:         accountName,
		ProjectName:         projectName,
		ApplicationName:     applicationName,
		AgentDeploymentName: agentDeploymentName,
	}
}

// ParseAgentDeploymentID parses 'input' into a AgentDeploymentId
func ParseAgentDeploymentID(input string) (*AgentDeploymentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AgentDeploymentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AgentDeploymentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAgentDeploymentIDInsensitively parses 'input' case-insensitively into a AgentDeploymentId
// note: this method should only be used for API response data and not user input
func ParseAgentDeploymentIDInsensitively(input string) (*AgentDeploymentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AgentDeploymentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AgentDeploymentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AgentDeploymentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.AccountName, ok = input.Parsed["accountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accountName", input)
	}

	if id.ProjectName, ok = input.Parsed["projectName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "projectName", input)
	}

	if id.ApplicationName, ok = input.Parsed["applicationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "applicationName", input)
	}

	if id.AgentDeploymentName, ok = input.Parsed["agentDeploymentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "agentDeploymentName", input)
	}

	return nil
}

// ValidateAgentDeploymentID checks that 'input' can be parsed as a Agent Deployment ID
func ValidateAgentDeploymentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAgentDeploymentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Agent Deployment ID
func (id AgentDeploymentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.CognitiveServices/accounts/%s/projects/%s/applications/%s/agentDeployments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, id.ProjectName, id.ApplicationName, id.AgentDeploymentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Agent Deployment ID
func (id AgentDeploymentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCognitiveServices", "Microsoft.CognitiveServices", "Microsoft.CognitiveServices"),
		resourceids.StaticSegment("staticAccounts", "accounts", "accounts"),
		resourceids.UserSpecifiedSegment("accountName", "accountName"),
		resourceids.StaticSegment("staticProjects", "projects", "projects"),
		resourceids.UserSpecifiedSegment("projectName", "projectName"),
		resourceids.StaticSegment("staticApplications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationName", "applicationName"),
		resourceids.StaticSegment("staticAgentDeployments", "agentDeployments", "agentDeployments"),
		resourceids.UserSpecifiedSegment("agentDeploymentName", "agentDeploymentName"),
	}
}

// String returns a human-readable description of this Agent Deployment ID
func (id AgentDeploymentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Project Name: %q", id.ProjectName),
		fmt.Sprintf("Application Name: %q", id.ApplicationName),
		fmt.Sprintf("Agent Deployment Name: %q", id.AgentDeploymentName),
	}
	return fmt.Sprintf("Agent Deployment (%s)", strings.Join(components, "\n"))
}
