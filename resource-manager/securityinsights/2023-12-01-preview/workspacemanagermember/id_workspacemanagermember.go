package workspacemanagermember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&WorkspaceManagerMemberId{})
}

var _ resourceids.ResourceId = &WorkspaceManagerMemberId{}

// WorkspaceManagerMemberId is a struct representing the Resource ID for a Workspace Manager Member
type WorkspaceManagerMemberId struct {
	SubscriptionId             string
	ResourceGroupName          string
	WorkspaceName              string
	WorkspaceManagerMemberName string
}

// NewWorkspaceManagerMemberID returns a new WorkspaceManagerMemberId struct
func NewWorkspaceManagerMemberID(subscriptionId string, resourceGroupName string, workspaceName string, workspaceManagerMemberName string) WorkspaceManagerMemberId {
	return WorkspaceManagerMemberId{
		SubscriptionId:             subscriptionId,
		ResourceGroupName:          resourceGroupName,
		WorkspaceName:              workspaceName,
		WorkspaceManagerMemberName: workspaceManagerMemberName,
	}
}

// ParseWorkspaceManagerMemberID parses 'input' into a WorkspaceManagerMemberId
func ParseWorkspaceManagerMemberID(input string) (*WorkspaceManagerMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&WorkspaceManagerMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := WorkspaceManagerMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseWorkspaceManagerMemberIDInsensitively parses 'input' case-insensitively into a WorkspaceManagerMemberId
// note: this method should only be used for API response data and not user input
func ParseWorkspaceManagerMemberIDInsensitively(input string) (*WorkspaceManagerMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&WorkspaceManagerMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := WorkspaceManagerMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *WorkspaceManagerMemberId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.WorkspaceManagerMemberName, ok = input.Parsed["workspaceManagerMemberName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workspaceManagerMemberName", input)
	}

	return nil
}

// ValidateWorkspaceManagerMemberID checks that 'input' can be parsed as a Workspace Manager Member ID
func ValidateWorkspaceManagerMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseWorkspaceManagerMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Workspace Manager Member ID
func (id WorkspaceManagerMemberId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/workspaceManagerMembers/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.WorkspaceManagerMemberName)
}

// Segments returns a slice of Resource ID Segments which comprise this Workspace Manager Member ID
func (id WorkspaceManagerMemberId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticWorkspaceManagerMembers", "workspaceManagerMembers", "workspaceManagerMembers"),
		resourceids.UserSpecifiedSegment("workspaceManagerMemberName", "workspaceManagerMemberName"),
	}
}

// String returns a human-readable description of this Workspace Manager Member ID
func (id WorkspaceManagerMemberId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Workspace Manager Member Name: %q", id.WorkspaceManagerMemberName),
	}
	return fmt.Sprintf("Workspace Manager Member (%s)", strings.Join(components, "\n"))
}
