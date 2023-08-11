package subscriptionfeatureregistrations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SubscriptionFeatureRegistrationId{}

// SubscriptionFeatureRegistrationId is a struct representing the Resource ID for a Subscription Feature Registration
type SubscriptionFeatureRegistrationId struct {
	SubscriptionId                      string
	FeatureProviderName                 string
	SubscriptionFeatureRegistrationName string
}

// NewSubscriptionFeatureRegistrationID returns a new SubscriptionFeatureRegistrationId struct
func NewSubscriptionFeatureRegistrationID(subscriptionId string, featureProviderName string, subscriptionFeatureRegistrationName string) SubscriptionFeatureRegistrationId {
	return SubscriptionFeatureRegistrationId{
		SubscriptionId:                      subscriptionId,
		FeatureProviderName:                 featureProviderName,
		SubscriptionFeatureRegistrationName: subscriptionFeatureRegistrationName,
	}
}

// ParseSubscriptionFeatureRegistrationID parses 'input' into a SubscriptionFeatureRegistrationId
func ParseSubscriptionFeatureRegistrationID(input string) (*SubscriptionFeatureRegistrationId, error) {
	parser := resourceids.NewParserFromResourceIdType(SubscriptionFeatureRegistrationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SubscriptionFeatureRegistrationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.FeatureProviderName, ok = parsed.Parsed["featureProviderName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "featureProviderName", *parsed)
	}

	if id.SubscriptionFeatureRegistrationName, ok = parsed.Parsed["subscriptionFeatureRegistrationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionFeatureRegistrationName", *parsed)
	}

	return &id, nil
}

// ParseSubscriptionFeatureRegistrationIDInsensitively parses 'input' case-insensitively into a SubscriptionFeatureRegistrationId
// note: this method should only be used for API response data and not user input
func ParseSubscriptionFeatureRegistrationIDInsensitively(input string) (*SubscriptionFeatureRegistrationId, error) {
	parser := resourceids.NewParserFromResourceIdType(SubscriptionFeatureRegistrationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SubscriptionFeatureRegistrationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.FeatureProviderName, ok = parsed.Parsed["featureProviderName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "featureProviderName", *parsed)
	}

	if id.SubscriptionFeatureRegistrationName, ok = parsed.Parsed["subscriptionFeatureRegistrationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionFeatureRegistrationName", *parsed)
	}

	return &id, nil
}

// ValidateSubscriptionFeatureRegistrationID checks that 'input' can be parsed as a Subscription Feature Registration ID
func ValidateSubscriptionFeatureRegistrationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSubscriptionFeatureRegistrationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Subscription Feature Registration ID
func (id SubscriptionFeatureRegistrationId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Features/featureProviders/%s/subscriptionFeatureRegistrations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.FeatureProviderName, id.SubscriptionFeatureRegistrationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Subscription Feature Registration ID
func (id SubscriptionFeatureRegistrationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftFeatures", "Microsoft.Features", "Microsoft.Features"),
		resourceids.StaticSegment("staticFeatureProviders", "featureProviders", "featureProviders"),
		resourceids.UserSpecifiedSegment("featureProviderName", "featureProviderValue"),
		resourceids.StaticSegment("staticSubscriptionFeatureRegistrations", "subscriptionFeatureRegistrations", "subscriptionFeatureRegistrations"),
		resourceids.UserSpecifiedSegment("subscriptionFeatureRegistrationName", "subscriptionFeatureRegistrationValue"),
	}
}

// String returns a human-readable description of this Subscription Feature Registration ID
func (id SubscriptionFeatureRegistrationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Feature Provider Name: %q", id.FeatureProviderName),
		fmt.Sprintf("Subscription Feature Registration Name: %q", id.SubscriptionFeatureRegistrationName),
	}
	return fmt.Sprintf("Subscription Feature Registration (%s)", strings.Join(components, "\n"))
}
