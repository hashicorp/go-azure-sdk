package privateendpoints

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ResourceId{}

// ResourceId is a struct representing the Resource ID for a Resource
type ResourceId struct {
	SubscriptionId           string
	ResourceGroupName        string
	DigitalTwinsInstanceName string
	ResourceId               string
}

// NewResourceID returns a new ResourceId struct
func NewResourceID(subscriptionId string, resourceGroupName string, digitalTwinsInstanceName string, resourceId string) ResourceId {
	return ResourceId{
		SubscriptionId:           subscriptionId,
		ResourceGroupName:        resourceGroupName,
		DigitalTwinsInstanceName: digitalTwinsInstanceName,
		ResourceId:               resourceId,
	}
}

// ParseResourceID parses 'input' into a ResourceId
func ParseResourceID(input string) (*ResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(ResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ResourceId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseResourceIDInsensitively parses 'input' case-insensitively into a ResourceId
// note: this method should only be used for API response data and not user input
func ParseResourceIDInsensitively(input string) (*ResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(ResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ResourceId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ResourceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.DigitalTwinsInstanceName, ok = input.Parsed["digitalTwinsInstanceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "digitalTwinsInstanceName", input)
	}

	if id.ResourceId, ok = input.Parsed["resourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceId", input)
	}

	return nil
}

// ValidateResourceID checks that 'input' can be parsed as a Resource ID
func ValidateResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Resource ID
func (id ResourceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DigitalTwins/digitalTwinsInstances/%s/privateLinkResources/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.DigitalTwinsInstanceName, strings.TrimPrefix(id.ResourceId, "/"))
}

// Segments returns a slice of Resource ID Segments which comprise this Resource ID
func (id ResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDigitalTwins", "Microsoft.DigitalTwins", "Microsoft.DigitalTwins"),
		resourceids.StaticSegment("staticDigitalTwinsInstances", "digitalTwinsInstances", "digitalTwinsInstances"),
		resourceids.UserSpecifiedSegment("digitalTwinsInstanceName", "digitalTwinsInstanceValue"),
		resourceids.StaticSegment("staticPrivateLinkResources", "privateLinkResources", "privateLinkResources"),
		resourceids.ScopeSegment("resourceId", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
	}
}

// String returns a human-readable description of this Resource ID
func (id ResourceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Digital Twins Instance Name: %q", id.DigitalTwinsInstanceName),
		fmt.Sprintf("Resource: %q", id.ResourceId),
	}
	return fmt.Sprintf("Resource (%s)", strings.Join(components, "\n"))
}
