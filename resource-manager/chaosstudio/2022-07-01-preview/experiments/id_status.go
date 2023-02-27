// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package experiments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = StatusId{}

// StatusId is a struct representing the Resource ID for a Status
type StatusId struct {
	SubscriptionId    string
	ResourceGroupName string
	ExperimentName    string
	StatusId          string
}

// NewStatusID returns a new StatusId struct
func NewStatusID(subscriptionId string, resourceGroupName string, experimentName string, statusId string) StatusId {
	return StatusId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ExperimentName:    experimentName,
		StatusId:          statusId,
	}
}

// ParseStatusID parses 'input' into a StatusId
func ParseStatusID(input string) (*StatusId, error) {
	parser := resourceids.NewParserFromResourceIdType(StatusId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := StatusId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ExperimentName, ok = parsed.Parsed["experimentName"]; !ok {
		return nil, fmt.Errorf("the segment 'experimentName' was not found in the resource id %q", input)
	}

	if id.StatusId, ok = parsed.Parsed["statusId"]; !ok {
		return nil, fmt.Errorf("the segment 'statusId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseStatusIDInsensitively parses 'input' case-insensitively into a StatusId
// note: this method should only be used for API response data and not user input
func ParseStatusIDInsensitively(input string) (*StatusId, error) {
	parser := resourceids.NewParserFromResourceIdType(StatusId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := StatusId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ExperimentName, ok = parsed.Parsed["experimentName"]; !ok {
		return nil, fmt.Errorf("the segment 'experimentName' was not found in the resource id %q", input)
	}

	if id.StatusId, ok = parsed.Parsed["statusId"]; !ok {
		return nil, fmt.Errorf("the segment 'statusId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateStatusID checks that 'input' can be parsed as a Status ID
func ValidateStatusID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseStatusID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Status ID
func (id StatusId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Chaos/experiments/%s/statuses/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ExperimentName, id.StatusId)
}

// Segments returns a slice of Resource ID Segments which comprise this Status ID
func (id StatusId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftChaos", "Microsoft.Chaos", "Microsoft.Chaos"),
		resourceids.StaticSegment("staticExperiments", "experiments", "experiments"),
		resourceids.UserSpecifiedSegment("experimentName", "experimentValue"),
		resourceids.StaticSegment("staticStatuses", "statuses", "statuses"),
		resourceids.UserSpecifiedSegment("statusId", "statusIdValue"),
	}
}

// String returns a human-readable description of this Status ID
func (id StatusId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Experiment Name: %q", id.ExperimentName),
		fmt.Sprintf("Status: %q", id.StatusId),
	}
	return fmt.Sprintf("Status (%s)", strings.Join(components, "\n"))
}
