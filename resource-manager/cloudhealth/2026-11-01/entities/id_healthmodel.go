package entities

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&HealthmodelId{})
}

var _ resourceids.ResourceId = &HealthmodelId{}

// HealthmodelId is a struct representing the Resource ID for a Healthmodel
type HealthmodelId struct {
	SubscriptionId    string
	ResourceGroupName string
	HealthmodelName   string
}

// NewHealthmodelID returns a new HealthmodelId struct
func NewHealthmodelID(subscriptionId string, resourceGroupName string, healthmodelName string) HealthmodelId {
	return HealthmodelId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		HealthmodelName:   healthmodelName,
	}
}

// ParseHealthmodelID parses 'input' into a HealthmodelId
func ParseHealthmodelID(input string) (*HealthmodelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&HealthmodelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := HealthmodelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseHealthmodelIDInsensitively parses 'input' case-insensitively into a HealthmodelId
// note: this method should only be used for API response data and not user input
func ParseHealthmodelIDInsensitively(input string) (*HealthmodelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&HealthmodelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := HealthmodelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *HealthmodelId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateHealthmodelID checks that 'input' can be parsed as a Healthmodel ID
func ValidateHealthmodelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseHealthmodelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Healthmodel ID
func (id HealthmodelId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.CloudHealth/healthmodels/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.HealthmodelName)
}

// Segments returns a slice of Resource ID Segments which comprise this Healthmodel ID
func (id HealthmodelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCloudHealth", "Microsoft.CloudHealth", "Microsoft.CloudHealth"),
		resourceids.StaticSegment("staticHealthmodels", "healthmodels", "healthmodels"),
		resourceids.UserSpecifiedSegment("healthmodelName", "healthmodelName"),
	}
}

// String returns a human-readable description of this Healthmodel ID
func (id HealthmodelId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Healthmodel Name: %q", id.HealthmodelName),
	}
	return fmt.Sprintf("Healthmodel (%s)", strings.Join(components, "\n"))
}
