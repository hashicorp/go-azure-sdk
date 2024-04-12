package workspacesettings

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&WorkspaceSettingId{})
}

var _ resourceids.ResourceId = &WorkspaceSettingId{}

// WorkspaceSettingId is a struct representing the Resource ID for a Workspace Setting
type WorkspaceSettingId struct {
	SubscriptionId       string
	WorkspaceSettingName string
}

// NewWorkspaceSettingID returns a new WorkspaceSettingId struct
func NewWorkspaceSettingID(subscriptionId string, workspaceSettingName string) WorkspaceSettingId {
	return WorkspaceSettingId{
		SubscriptionId:       subscriptionId,
		WorkspaceSettingName: workspaceSettingName,
	}
}

// ParseWorkspaceSettingID parses 'input' into a WorkspaceSettingId
func ParseWorkspaceSettingID(input string) (*WorkspaceSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&WorkspaceSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := WorkspaceSettingId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseWorkspaceSettingIDInsensitively parses 'input' case-insensitively into a WorkspaceSettingId
// note: this method should only be used for API response data and not user input
func ParseWorkspaceSettingIDInsensitively(input string) (*WorkspaceSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&WorkspaceSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := WorkspaceSettingId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *WorkspaceSettingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.WorkspaceSettingName, ok = input.Parsed["workspaceSettingName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workspaceSettingName", input)
	}

	return nil
}

// ValidateWorkspaceSettingID checks that 'input' can be parsed as a Workspace Setting ID
func ValidateWorkspaceSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseWorkspaceSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Workspace Setting ID
func (id WorkspaceSettingId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Security/workspaceSettings/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.WorkspaceSettingName)
}

// Segments returns a slice of Resource ID Segments which comprise this Workspace Setting ID
func (id WorkspaceSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticWorkspaceSettings", "workspaceSettings", "workspaceSettings"),
		resourceids.UserSpecifiedSegment("workspaceSettingName", "workspaceSettingValue"),
	}
}

// String returns a human-readable description of this Workspace Setting ID
func (id WorkspaceSettingId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Workspace Setting Name: %q", id.WorkspaceSettingName),
	}
	return fmt.Sprintf("Workspace Setting (%s)", strings.Join(components, "\n"))
}
