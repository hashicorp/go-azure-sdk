package createresource

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&SubscriptionStatusId{})
}

var _ resourceids.ResourceId = &SubscriptionStatusId{}

// SubscriptionStatusId is a struct representing the Resource ID for a Subscription Status
type SubscriptionStatusId struct {
	SubscriptionId         string
	DynatraceEnvironmentId string
}

// NewSubscriptionStatusID returns a new SubscriptionStatusId struct
func NewSubscriptionStatusID(subscriptionId string, dynatraceEnvironmentId string) SubscriptionStatusId {
	return SubscriptionStatusId{
		SubscriptionId:         subscriptionId,
		DynatraceEnvironmentId: dynatraceEnvironmentId,
	}
}

// ParseSubscriptionStatusID parses 'input' into a SubscriptionStatusId
func ParseSubscriptionStatusID(input string) (*SubscriptionStatusId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SubscriptionStatusId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SubscriptionStatusId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSubscriptionStatusIDInsensitively parses 'input' case-insensitively into a SubscriptionStatusId
// note: this method should only be used for API response data and not user input
func ParseSubscriptionStatusIDInsensitively(input string) (*SubscriptionStatusId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SubscriptionStatusId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SubscriptionStatusId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SubscriptionStatusId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.DynatraceEnvironmentId, ok = input.Parsed["dynatraceEnvironmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dynatraceEnvironmentId", input)
	}

	return nil
}

// ValidateSubscriptionStatusID checks that 'input' can be parsed as a Subscription Status ID
func ValidateSubscriptionStatusID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSubscriptionStatusID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Subscription Status ID
func (id SubscriptionStatusId) ID() string {
	fmtString := "/subscriptions/%s/providers/Dynatrace.Observability/subscriptionStatuses/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.DynatraceEnvironmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Subscription Status ID
func (id SubscriptionStatusId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticDynatraceObservability", "Dynatrace.Observability", "Dynatrace.Observability"),
		resourceids.StaticSegment("staticSubscriptionStatuses", "subscriptionStatuses", "subscriptionStatuses"),
		resourceids.UserSpecifiedSegment("dynatraceEnvironmentId", "dynatraceEnvironmentId"),
	}
}

// String returns a human-readable description of this Subscription Status ID
func (id SubscriptionStatusId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Dynatrace Environment: %q", id.DynatraceEnvironmentId),
	}
	return fmt.Sprintf("Subscription Status (%s)", strings.Join(components, "\n"))
}
