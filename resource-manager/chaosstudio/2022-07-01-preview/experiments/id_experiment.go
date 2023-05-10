package experiments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ExperimentId{}

// ExperimentId is a struct representing the Resource ID for a Experiment
type ExperimentId struct {
	SubscriptionId    string
	ResourceGroupName string
	ExperimentName    string
}

// NewExperimentID returns a new ExperimentId struct
func NewExperimentID(subscriptionId string, resourceGroupName string, experimentName string) ExperimentId {
	return ExperimentId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ExperimentName:    experimentName,
	}
}

// ParseExperimentID parses 'input' into a ExperimentId
func ParseExperimentID(input string) (*ExperimentId, error) {
	parser := resourceids.NewParserFromResourceIdType(ExperimentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ExperimentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ExperimentName, ok = parsed.Parsed["experimentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "experimentName", *parsed)
	}

	return &id, nil
}

// ParseExperimentIDInsensitively parses 'input' case-insensitively into a ExperimentId
// note: this method should only be used for API response data and not user input
func ParseExperimentIDInsensitively(input string) (*ExperimentId, error) {
	parser := resourceids.NewParserFromResourceIdType(ExperimentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ExperimentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ExperimentName, ok = parsed.Parsed["experimentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "experimentName", *parsed)
	}

	return &id, nil
}

// ValidateExperimentID checks that 'input' can be parsed as a Experiment ID
func ValidateExperimentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseExperimentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Experiment ID
func (id ExperimentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Chaos/experiments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ExperimentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Experiment ID
func (id ExperimentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftChaos", "Microsoft.Chaos", "Microsoft.Chaos"),
		resourceids.StaticSegment("staticExperiments", "experiments", "experiments"),
		resourceids.UserSpecifiedSegment("experimentName", "experimentValue"),
	}
}

// String returns a human-readable description of this Experiment ID
func (id ExperimentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Experiment Name: %q", id.ExperimentName),
	}
	return fmt.Sprintf("Experiment (%s)", strings.Join(components, "\n"))
}
