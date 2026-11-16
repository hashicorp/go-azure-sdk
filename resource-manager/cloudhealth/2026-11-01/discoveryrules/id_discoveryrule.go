package discoveryrules

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&DiscoveryruleId{})
}

var _ resourceids.ResourceId = &DiscoveryruleId{}

// DiscoveryruleId is a struct representing the Resource ID for a Discoveryrule
type DiscoveryruleId struct {
	SubscriptionId    string
	ResourceGroupName string
	HealthmodelName   string
	DiscoveryruleName string
}

// NewDiscoveryruleID returns a new DiscoveryruleId struct
func NewDiscoveryruleID(subscriptionId string, resourceGroupName string, healthmodelName string, discoveryruleName string) DiscoveryruleId {
	return DiscoveryruleId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		HealthmodelName:   healthmodelName,
		DiscoveryruleName: discoveryruleName,
	}
}

// ParseDiscoveryruleID parses 'input' into a DiscoveryruleId
func ParseDiscoveryruleID(input string) (*DiscoveryruleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DiscoveryruleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DiscoveryruleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDiscoveryruleIDInsensitively parses 'input' case-insensitively into a DiscoveryruleId
// note: this method should only be used for API response data and not user input
func ParseDiscoveryruleIDInsensitively(input string) (*DiscoveryruleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DiscoveryruleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DiscoveryruleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DiscoveryruleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.HealthmodelName, ok = input.Parsed["healthmodelName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "healthmodelName", input)
	}

	if id.DiscoveryruleName, ok = input.Parsed["discoveryruleName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "discoveryruleName", input)
	}

	return nil
}

// ValidateDiscoveryruleID checks that 'input' can be parsed as a Discoveryrule ID
func ValidateDiscoveryruleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDiscoveryruleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Discoveryrule ID
func (id DiscoveryruleId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.CloudHealth/healthmodels/%s/discoveryrules/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.HealthmodelName, id.DiscoveryruleName)
}

// Segments returns a slice of Resource ID Segments which comprise this Discoveryrule ID
func (id DiscoveryruleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCloudHealth", "Microsoft.CloudHealth", "Microsoft.CloudHealth"),
		resourceids.StaticSegment("staticHealthmodels", "healthmodels", "healthmodels"),
		resourceids.UserSpecifiedSegment("healthmodelName", "healthmodelName"),
		resourceids.StaticSegment("staticDiscoveryrules", "discoveryrules", "discoveryrules"),
		resourceids.UserSpecifiedSegment("discoveryruleName", "discoveryruleName"),
	}
}

// String returns a human-readable description of this Discoveryrule ID
func (id DiscoveryruleId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Healthmodel Name: %q", id.HealthmodelName),
		fmt.Sprintf("Discoveryrule Name: %q", id.DiscoveryruleName),
	}
	return fmt.Sprintf("Discoveryrule (%s)", strings.Join(components, "\n"))
}
