package settings

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&SettingId{})
}

var _ resourceids.ResourceId = &SettingId{}

// SettingId is a struct representing the Resource ID for a Setting
type SettingId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	SettingName       string
}

// NewSettingID returns a new SettingId struct
func NewSettingID(subscriptionId string, resourceGroupName string, workspaceName string, settingName string) SettingId {
	return SettingId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		SettingName:       settingName,
	}
}

// ParseSettingID parses 'input' into a SettingId
func ParseSettingID(input string) (*SettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSettingIDInsensitively parses 'input' case-insensitively into a SettingId
// note: this method should only be used for API response data and not user input
func ParseSettingIDInsensitively(input string) (*SettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SettingId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.SettingName, ok = input.Parsed["settingName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "settingName", input)
	}

	return nil
}

// ValidateSettingID checks that 'input' can be parsed as a Setting ID
func ValidateSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Setting ID
func (id SettingId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/settings/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.SettingName)
}

// Segments returns a slice of Resource ID Segments which comprise this Setting ID
func (id SettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftOperationalInsights", "Microsoft.OperationalInsights", "Microsoft.OperationalInsights"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurityInsights", "Microsoft.SecurityInsights", "Microsoft.SecurityInsights"),
		resourceids.StaticSegment("staticSettings", "settings", "settings"),
		resourceids.UserSpecifiedSegment("settingName", "settingValue"),
	}
}

// String returns a human-readable description of this Setting ID
func (id SettingId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Setting Name: %q", id.SettingName),
	}
	return fmt.Sprintf("Setting (%s)", strings.Join(components, "\n"))
}
