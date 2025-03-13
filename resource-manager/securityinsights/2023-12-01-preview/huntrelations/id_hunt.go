package huntrelations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&HuntId{})
}

var _ resourceids.ResourceId = &HuntId{}

// HuntId is a struct representing the Resource ID for a Hunt
type HuntId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	HuntId            string
}

// NewHuntID returns a new HuntId struct
func NewHuntID(subscriptionId string, resourceGroupName string, workspaceName string, huntId string) HuntId {
	return HuntId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		HuntId:            huntId,
	}
}

// ParseHuntID parses 'input' into a HuntId
func ParseHuntID(input string) (*HuntId, error) {
	parser := resourceids.NewParserFromResourceIdType(&HuntId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := HuntId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseHuntIDInsensitively parses 'input' case-insensitively into a HuntId
// note: this method should only be used for API response data and not user input
func ParseHuntIDInsensitively(input string) (*HuntId, error) {
	parser := resourceids.NewParserFromResourceIdType(&HuntId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := HuntId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *HuntId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.WorkspaceName, ok = input.Parsed["workspaceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workspaceName", input)
	}

	if id.HuntId, ok = input.Parsed["huntId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "huntId", input)
	}

	return nil
}

// ValidateHuntID checks that 'input' can be parsed as a Hunt ID
func ValidateHuntID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseHuntID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Hunt ID
func (id HuntId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/hunts/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.HuntId)
}

// Segments returns a slice of Resource ID Segments which comprise this Hunt ID
func (id HuntId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftOperationalInsights", "Microsoft.OperationalInsights", "Microsoft.OperationalInsights"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceName"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurityInsights", "Microsoft.SecurityInsights", "Microsoft.SecurityInsights"),
		resourceids.StaticSegment("staticHunts", "hunts", "hunts"),
		resourceids.UserSpecifiedSegment("huntId", "huntId"),
	}
}

// String returns a human-readable description of this Hunt ID
func (id HuntId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Hunt: %q", id.HuntId),
	}
	return fmt.Sprintf("Hunt (%s)", strings.Join(components, "\n"))
}
