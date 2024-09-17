package containerappsbuilds

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ContainerAppBuildId{})
}

var _ resourceids.ResourceId = &ContainerAppBuildId{}

// ContainerAppBuildId is a struct representing the Resource ID for a Container App Build
type ContainerAppBuildId struct {
	SubscriptionId    string
	ResourceGroupName string
	ContainerAppName  string
	BuildName         string
}

// NewContainerAppBuildID returns a new ContainerAppBuildId struct
func NewContainerAppBuildID(subscriptionId string, resourceGroupName string, containerAppName string, buildName string) ContainerAppBuildId {
	return ContainerAppBuildId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ContainerAppName:  containerAppName,
		BuildName:         buildName,
	}
}

// ParseContainerAppBuildID parses 'input' into a ContainerAppBuildId
func ParseContainerAppBuildID(input string) (*ContainerAppBuildId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ContainerAppBuildId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ContainerAppBuildId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseContainerAppBuildIDInsensitively parses 'input' case-insensitively into a ContainerAppBuildId
// note: this method should only be used for API response data and not user input
func ParseContainerAppBuildIDInsensitively(input string) (*ContainerAppBuildId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ContainerAppBuildId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ContainerAppBuildId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ContainerAppBuildId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.BuildName, ok = input.Parsed["buildName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "buildName", input)
	}

	return nil
}

// ValidateContainerAppBuildID checks that 'input' can be parsed as a Container App Build ID
func ValidateContainerAppBuildID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseContainerAppBuildID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Container App Build ID
func (id ContainerAppBuildId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.App/containerApps/%s/builds/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ContainerAppName, id.BuildName)
}

// Segments returns a slice of Resource ID Segments which comprise this Container App Build ID
func (id ContainerAppBuildId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApp", "Microsoft.App", "Microsoft.App"),
		resourceids.StaticSegment("staticContainerApps", "containerApps", "containerApps"),
		resourceids.UserSpecifiedSegment("containerAppName", "containerAppValue"),
		resourceids.StaticSegment("staticBuilds", "builds", "builds"),
		resourceids.UserSpecifiedSegment("buildName", "buildValue"),
	}
}

// String returns a human-readable description of this Container App Build ID
func (id ContainerAppBuildId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Container App Name: %q", id.ContainerAppName),
		fmt.Sprintf("Build Name: %q", id.BuildName),
	}
	return fmt.Sprintf("Container App Build (%s)", strings.Join(components, "\n"))
}
