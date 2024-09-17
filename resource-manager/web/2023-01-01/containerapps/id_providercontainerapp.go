package containerapps

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ProviderContainerAppId{})
}

var _ resourceids.ResourceId = &ProviderContainerAppId{}

// ProviderContainerAppId is a struct representing the Resource ID for a Provider Container App
type ProviderContainerAppId struct {
	SubscriptionId    string
	ResourceGroupName string
	ContainerAppName  string
}

// NewProviderContainerAppID returns a new ProviderContainerAppId struct
func NewProviderContainerAppID(subscriptionId string, resourceGroupName string, containerAppName string) ProviderContainerAppId {
	return ProviderContainerAppId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ContainerAppName:  containerAppName,
	}
}

// ParseProviderContainerAppID parses 'input' into a ProviderContainerAppId
func ParseProviderContainerAppID(input string) (*ProviderContainerAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ProviderContainerAppId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ProviderContainerAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseProviderContainerAppIDInsensitively parses 'input' case-insensitively into a ProviderContainerAppId
// note: this method should only be used for API response data and not user input
func ParseProviderContainerAppIDInsensitively(input string) (*ProviderContainerAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ProviderContainerAppId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ProviderContainerAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ProviderContainerAppId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateProviderContainerAppID checks that 'input' can be parsed as a Provider Container App ID
func ValidateProviderContainerAppID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviderContainerAppID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Provider Container App ID
func (id ProviderContainerAppId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/containerApps/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ContainerAppName)
}

// Segments returns a slice of Resource ID Segments which comprise this Provider Container App ID
func (id ProviderContainerAppId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticContainerApps", "containerApps", "containerApps"),
		resourceids.UserSpecifiedSegment("containerAppName", "containerAppValue"),
	}
}

// String returns a human-readable description of this Provider Container App ID
func (id ProviderContainerAppId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Container App Name: %q", id.ContainerAppName),
	}
	return fmt.Sprintf("Provider Container App (%s)", strings.Join(components, "\n"))
}
