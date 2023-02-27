package customapis

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = CustomApiId{}

// CustomApiId is a struct representing the Resource ID for a Custom Api
type CustomApiId struct {
	SubscriptionId    string
	ResourceGroupName string
	CustomApiName     string
}

// NewCustomApiID returns a new CustomApiId struct
func NewCustomApiID(subscriptionId string, resourceGroupName string, customApiName string) CustomApiId {
	return CustomApiId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		CustomApiName:     customApiName,
	}
}

// ParseCustomApiID parses 'input' into a CustomApiId
func ParseCustomApiID(input string) (*CustomApiId, error) {
	parser := resourceids.NewParserFromResourceIdType(CustomApiId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CustomApiId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.CustomApiName, ok = parsed.Parsed["customApiName"]; !ok {
		return nil, fmt.Errorf("the segment 'customApiName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseCustomApiIDInsensitively parses 'input' case-insensitively into a CustomApiId
// note: this method should only be used for API response data and not user input
func ParseCustomApiIDInsensitively(input string) (*CustomApiId, error) {
	parser := resourceids.NewParserFromResourceIdType(CustomApiId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CustomApiId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.CustomApiName, ok = parsed.Parsed["customApiName"]; !ok {
		return nil, fmt.Errorf("the segment 'customApiName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateCustomApiID checks that 'input' can be parsed as a Custom Api ID
func ValidateCustomApiID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCustomApiID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Custom Api ID
func (id CustomApiId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/customApis/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.CustomApiName)
}

// Segments returns a slice of Resource ID Segments which comprise this Custom Api ID
func (id CustomApiId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticCustomApis", "customApis", "customApis"),
		resourceids.UserSpecifiedSegment("customApiName", "customApiValue"),
	}
}

// String returns a human-readable description of this Custom Api ID
func (id CustomApiId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Custom Api Name: %q", id.CustomApiName),
	}
	return fmt.Sprintf("Custom Api (%s)", strings.Join(components, "\n"))
}
