package logicapps

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&LogicAppId{})
}

var _ resourceids.ResourceId = &LogicAppId{}

// LogicAppId is a struct representing the Resource ID for a Logic App
type LogicAppId struct {
	SubscriptionId    string
	ResourceGroupName string
	ContainerAppName  string
	LogicAppName      string
}

// NewLogicAppID returns a new LogicAppId struct
func NewLogicAppID(subscriptionId string, resourceGroupName string, containerAppName string, logicAppName string) LogicAppId {
	return LogicAppId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ContainerAppName:  containerAppName,
		LogicAppName:      logicAppName,
	}
}

// ParseLogicAppID parses 'input' into a LogicAppId
func ParseLogicAppID(input string) (*LogicAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LogicAppId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LogicAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseLogicAppIDInsensitively parses 'input' case-insensitively into a LogicAppId
// note: this method should only be used for API response data and not user input
func ParseLogicAppIDInsensitively(input string) (*LogicAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LogicAppId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LogicAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *LogicAppId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ContainerAppName, ok = input.Parsed["containerAppName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "containerAppName", input)
	}

	if id.LogicAppName, ok = input.Parsed["logicAppName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "logicAppName", input)
	}

	return nil
}

// ValidateLogicAppID checks that 'input' can be parsed as a Logic App ID
func ValidateLogicAppID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseLogicAppID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Logic App ID
func (id LogicAppId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.App/containerApps/%s/providers/Microsoft.App/logicApps/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ContainerAppName, id.LogicAppName)
}

// Segments returns a slice of Resource ID Segments which comprise this Logic App ID
func (id LogicAppId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApp", "Microsoft.App", "Microsoft.App"),
		resourceids.StaticSegment("staticContainerApps", "containerApps", "containerApps"),
		resourceids.UserSpecifiedSegment("containerAppName", "containerAppValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApp2", "Microsoft.App", "Microsoft.App"),
		resourceids.StaticSegment("staticLogicApps", "logicApps", "logicApps"),
		resourceids.UserSpecifiedSegment("logicAppName", "logicAppValue"),
	}
}

// String returns a human-readable description of this Logic App ID
func (id LogicAppId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Container App Name: %q", id.ContainerAppName),
		fmt.Sprintf("Logic App Name: %q", id.LogicAppName),
	}
	return fmt.Sprintf("Logic App (%s)", strings.Join(components, "\n"))
}
