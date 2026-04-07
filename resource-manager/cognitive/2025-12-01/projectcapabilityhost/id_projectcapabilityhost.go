package projectcapabilityhost

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ProjectCapabilityHostId{})
}

var _ resourceids.ResourceId = &ProjectCapabilityHostId{}

// ProjectCapabilityHostId is a struct representing the Resource ID for a Project Capability Host
type ProjectCapabilityHostId struct {
	SubscriptionId     string
	ResourceGroupName  string
	AccountName        string
	ProjectName        string
	CapabilityHostName string
}

// NewProjectCapabilityHostID returns a new ProjectCapabilityHostId struct
func NewProjectCapabilityHostID(subscriptionId string, resourceGroupName string, accountName string, projectName string, capabilityHostName string) ProjectCapabilityHostId {
	return ProjectCapabilityHostId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		AccountName:        accountName,
		ProjectName:        projectName,
		CapabilityHostName: capabilityHostName,
	}
}

// ParseProjectCapabilityHostID parses 'input' into a ProjectCapabilityHostId
func ParseProjectCapabilityHostID(input string) (*ProjectCapabilityHostId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ProjectCapabilityHostId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ProjectCapabilityHostId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseProjectCapabilityHostIDInsensitively parses 'input' case-insensitively into a ProjectCapabilityHostId
// note: this method should only be used for API response data and not user input
func ParseProjectCapabilityHostIDInsensitively(input string) (*ProjectCapabilityHostId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ProjectCapabilityHostId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ProjectCapabilityHostId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ProjectCapabilityHostId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.CapabilityHostName, ok = input.Parsed["capabilityHostName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "capabilityHostName", input)
	}

	return nil
}

// ValidateProjectCapabilityHostID checks that 'input' can be parsed as a Project Capability Host ID
func ValidateProjectCapabilityHostID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProjectCapabilityHostID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Project Capability Host ID
func (id ProjectCapabilityHostId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.CognitiveServices/accounts/%s/projects/%s/capabilityHosts/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, id.ProjectName, id.CapabilityHostName)
}

// Segments returns a slice of Resource ID Segments which comprise this Project Capability Host ID
func (id ProjectCapabilityHostId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticCapabilityHosts", "capabilityHosts", "capabilityHosts"),
		resourceids.UserSpecifiedSegment("capabilityHostName", "capabilityHostName"),
	}
}

// String returns a human-readable description of this Project Capability Host ID
func (id ProjectCapabilityHostId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Project Name: %q", id.ProjectName),
		fmt.Sprintf("Capability Host Name: %q", id.CapabilityHostName),
	}
	return fmt.Sprintf("Project Capability Host (%s)", strings.Join(components, "\n"))
}
