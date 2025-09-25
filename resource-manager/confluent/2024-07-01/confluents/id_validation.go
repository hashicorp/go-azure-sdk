package confluents

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ValidationId{})
}

var _ resourceids.ResourceId = &ValidationId{}

// ValidationId is a struct representing the Resource ID for a Validation
type ValidationId struct {
	SubscriptionId    string
	ResourceGroupName string
	ValidationName    string
}

// NewValidationID returns a new ValidationId struct
func NewValidationID(subscriptionId string, resourceGroupName string, validationName string) ValidationId {
	return ValidationId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ValidationName:    validationName,
	}
}

// ParseValidationID parses 'input' into a ValidationId
func ParseValidationID(input string) (*ValidationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ValidationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ValidationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseValidationIDInsensitively parses 'input' case-insensitively into a ValidationId
// note: this method should only be used for API response data and not user input
func ParseValidationIDInsensitively(input string) (*ValidationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ValidationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ValidationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ValidationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ValidationName, ok = input.Parsed["validationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "validationName", input)
	}

	return nil
}

// ValidateValidationID checks that 'input' can be parsed as a Validation ID
func ValidateValidationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseValidationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Validation ID
func (id ValidationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Confluent/validations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ValidationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Validation ID
func (id ValidationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftConfluent", "Microsoft.Confluent", "Microsoft.Confluent"),
		resourceids.StaticSegment("staticValidations", "validations", "validations"),
		resourceids.UserSpecifiedSegment("validationName", "validationName"),
	}
}

// String returns a human-readable description of this Validation ID
func (id ValidationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Validation Name: %q", id.ValidationName),
	}
	return fmt.Sprintf("Validation (%s)", strings.Join(components, "\n"))
}
