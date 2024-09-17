package workspacemanagedsqlserverdedicatedsqlminimaltlssettings

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&DedicatedSQLMinimalTLSSettingId{})
}

var _ resourceids.ResourceId = &DedicatedSQLMinimalTLSSettingId{}

// DedicatedSQLMinimalTLSSettingId is a struct representing the Resource ID for a Dedicated S Q L Minimal T L S Setting
type DedicatedSQLMinimalTLSSettingId struct {
	SubscriptionId                    string
	ResourceGroupName                 string
	WorkspaceName                     string
	DedicatedSQLMinimalTLSSettingName string
}

// NewDedicatedSQLMinimalTLSSettingID returns a new DedicatedSQLMinimalTLSSettingId struct
func NewDedicatedSQLMinimalTLSSettingID(subscriptionId string, resourceGroupName string, workspaceName string, dedicatedSQLMinimalTLSSettingName string) DedicatedSQLMinimalTLSSettingId {
	return DedicatedSQLMinimalTLSSettingId{
		SubscriptionId:                    subscriptionId,
		ResourceGroupName:                 resourceGroupName,
		WorkspaceName:                     workspaceName,
		DedicatedSQLMinimalTLSSettingName: dedicatedSQLMinimalTLSSettingName,
	}
}

// ParseDedicatedSQLMinimalTLSSettingID parses 'input' into a DedicatedSQLMinimalTLSSettingId
func ParseDedicatedSQLMinimalTLSSettingID(input string) (*DedicatedSQLMinimalTLSSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DedicatedSQLMinimalTLSSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DedicatedSQLMinimalTLSSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDedicatedSQLMinimalTLSSettingIDInsensitively parses 'input' case-insensitively into a DedicatedSQLMinimalTLSSettingId
// note: this method should only be used for API response data and not user input
func ParseDedicatedSQLMinimalTLSSettingIDInsensitively(input string) (*DedicatedSQLMinimalTLSSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DedicatedSQLMinimalTLSSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DedicatedSQLMinimalTLSSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DedicatedSQLMinimalTLSSettingId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.DedicatedSQLMinimalTLSSettingName, ok = input.Parsed["dedicatedSQLMinimalTLSSettingName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dedicatedSQLMinimalTLSSettingName", input)
	}

	return nil
}

// ValidateDedicatedSQLMinimalTLSSettingID checks that 'input' can be parsed as a Dedicated S Q L Minimal T L S Setting ID
func ValidateDedicatedSQLMinimalTLSSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDedicatedSQLMinimalTLSSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Dedicated S Q L Minimal T L S Setting ID
func (id DedicatedSQLMinimalTLSSettingId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Synapse/workspaces/%s/dedicatedSQLMinimalTLSSettings/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.DedicatedSQLMinimalTLSSettingName)
}

// Segments returns a slice of Resource ID Segments which comprise this Dedicated S Q L Minimal T L S Setting ID
func (id DedicatedSQLMinimalTLSSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSynapse", "Microsoft.Synapse", "Microsoft.Synapse"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticDedicatedSQLMinimalTLSSettings", "dedicatedSQLMinimalTLSSettings", "dedicatedSQLMinimalTLSSettings"),
		resourceids.UserSpecifiedSegment("dedicatedSQLMinimalTLSSettingName", "dedicatedSQLMinimalTLSSettingValue"),
	}
}

// String returns a human-readable description of this Dedicated S Q L Minimal T L S Setting ID
func (id DedicatedSQLMinimalTLSSettingId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Dedicated S Q L Minimal T L S Setting Name: %q", id.DedicatedSQLMinimalTLSSettingName),
	}
	return fmt.Sprintf("Dedicated S Q L Minimal T L S Setting (%s)", strings.Join(components, "\n"))
}
