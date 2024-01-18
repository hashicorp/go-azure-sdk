package customimages

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &CustomImageId{}

// CustomImageId is a struct representing the Resource ID for a Custom Image
type CustomImageId struct {
	SubscriptionId    string
	ResourceGroupName string
	LabName           string
	CustomImageName   string
}

// NewCustomImageID returns a new CustomImageId struct
func NewCustomImageID(subscriptionId string, resourceGroupName string, labName string, customImageName string) CustomImageId {
	return CustomImageId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		LabName:           labName,
		CustomImageName:   customImageName,
	}
}

// ParseCustomImageID parses 'input' into a CustomImageId
func ParseCustomImageID(input string) (*CustomImageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CustomImageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CustomImageId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseCustomImageIDInsensitively parses 'input' case-insensitively into a CustomImageId
// note: this method should only be used for API response data and not user input
func ParseCustomImageIDInsensitively(input string) (*CustomImageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CustomImageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CustomImageId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *CustomImageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.LabName, ok = input.Parsed["labName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "labName", input)
	}

	if id.CustomImageName, ok = input.Parsed["customImageName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "customImageName", input)
	}

	return nil
}

// ValidateCustomImageID checks that 'input' can be parsed as a Custom Image ID
func ValidateCustomImageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCustomImageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Custom Image ID
func (id CustomImageId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DevTestLab/labs/%s/customImages/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LabName, id.CustomImageName)
}

// Segments returns a slice of Resource ID Segments which comprise this Custom Image ID
func (id CustomImageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDevTestLab", "Microsoft.DevTestLab", "Microsoft.DevTestLab"),
		resourceids.StaticSegment("staticLabs", "labs", "labs"),
		resourceids.UserSpecifiedSegment("labName", "labValue"),
		resourceids.StaticSegment("staticCustomImages", "customImages", "customImages"),
		resourceids.UserSpecifiedSegment("customImageName", "customImageValue"),
	}
}

// String returns a human-readable description of this Custom Image ID
func (id CustomImageId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Lab Name: %q", id.LabName),
		fmt.Sprintf("Custom Image Name: %q", id.CustomImageName),
	}
	return fmt.Sprintf("Custom Image (%s)", strings.Join(components, "\n"))
}
