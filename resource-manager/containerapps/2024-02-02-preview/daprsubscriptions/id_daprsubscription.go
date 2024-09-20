package daprsubscriptions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&DaprSubscriptionId{})
}

var _ resourceids.ResourceId = &DaprSubscriptionId{}

// DaprSubscriptionId is a struct representing the Resource ID for a Dapr Subscription
type DaprSubscriptionId struct {
	SubscriptionId         string
	ResourceGroupName      string
	ManagedEnvironmentName string
	DaprSubscriptionName   string
}

// NewDaprSubscriptionID returns a new DaprSubscriptionId struct
func NewDaprSubscriptionID(subscriptionId string, resourceGroupName string, managedEnvironmentName string, daprSubscriptionName string) DaprSubscriptionId {
	return DaprSubscriptionId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		ManagedEnvironmentName: managedEnvironmentName,
		DaprSubscriptionName:   daprSubscriptionName,
	}
}

// ParseDaprSubscriptionID parses 'input' into a DaprSubscriptionId
func ParseDaprSubscriptionID(input string) (*DaprSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DaprSubscriptionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DaprSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDaprSubscriptionIDInsensitively parses 'input' case-insensitively into a DaprSubscriptionId
// note: this method should only be used for API response data and not user input
func ParseDaprSubscriptionIDInsensitively(input string) (*DaprSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DaprSubscriptionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DaprSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DaprSubscriptionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ManagedEnvironmentName, ok = input.Parsed["managedEnvironmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedEnvironmentName", input)
	}

	if id.DaprSubscriptionName, ok = input.Parsed["daprSubscriptionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "daprSubscriptionName", input)
	}

	return nil
}

// ValidateDaprSubscriptionID checks that 'input' can be parsed as a Dapr Subscription ID
func ValidateDaprSubscriptionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDaprSubscriptionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Dapr Subscription ID
func (id DaprSubscriptionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.App/managedEnvironments/%s/daprSubscriptions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ManagedEnvironmentName, id.DaprSubscriptionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Dapr Subscription ID
func (id DaprSubscriptionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApp", "Microsoft.App", "Microsoft.App"),
		resourceids.StaticSegment("staticManagedEnvironments", "managedEnvironments", "managedEnvironments"),
		resourceids.UserSpecifiedSegment("managedEnvironmentName", "environmentName"),
		resourceids.StaticSegment("staticDaprSubscriptions", "daprSubscriptions", "daprSubscriptions"),
		resourceids.UserSpecifiedSegment("daprSubscriptionName", "name"),
	}
}

// String returns a human-readable description of this Dapr Subscription ID
func (id DaprSubscriptionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Managed Environment Name: %q", id.ManagedEnvironmentName),
		fmt.Sprintf("Dapr Subscription Name: %q", id.DaprSubscriptionName),
	}
	return fmt.Sprintf("Dapr Subscription (%s)", strings.Join(components, "\n"))
}
