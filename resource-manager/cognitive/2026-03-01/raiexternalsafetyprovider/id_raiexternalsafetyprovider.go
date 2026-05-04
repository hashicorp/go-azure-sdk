package raiexternalsafetyprovider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&RaiExternalSafetyProviderId{})
}

var _ resourceids.ResourceId = &RaiExternalSafetyProviderId{}

// RaiExternalSafetyProviderId is a struct representing the Resource ID for a Rai External Safety Provider
type RaiExternalSafetyProviderId struct {
	SubscriptionId                string
	RaiExternalSafetyProviderName string
}

// NewRaiExternalSafetyProviderID returns a new RaiExternalSafetyProviderId struct
func NewRaiExternalSafetyProviderID(subscriptionId string, raiExternalSafetyProviderName string) RaiExternalSafetyProviderId {
	return RaiExternalSafetyProviderId{
		SubscriptionId:                subscriptionId,
		RaiExternalSafetyProviderName: raiExternalSafetyProviderName,
	}
}

// ParseRaiExternalSafetyProviderID parses 'input' into a RaiExternalSafetyProviderId
func ParseRaiExternalSafetyProviderID(input string) (*RaiExternalSafetyProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RaiExternalSafetyProviderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RaiExternalSafetyProviderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRaiExternalSafetyProviderIDInsensitively parses 'input' case-insensitively into a RaiExternalSafetyProviderId
// note: this method should only be used for API response data and not user input
func ParseRaiExternalSafetyProviderIDInsensitively(input string) (*RaiExternalSafetyProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RaiExternalSafetyProviderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RaiExternalSafetyProviderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RaiExternalSafetyProviderId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.RaiExternalSafetyProviderName, ok = input.Parsed["raiExternalSafetyProviderName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "raiExternalSafetyProviderName", input)
	}

	return nil
}

// ValidateRaiExternalSafetyProviderID checks that 'input' can be parsed as a Rai External Safety Provider ID
func ValidateRaiExternalSafetyProviderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRaiExternalSafetyProviderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Rai External Safety Provider ID
func (id RaiExternalSafetyProviderId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.CognitiveServices/raiExternalSafetyProviders/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.RaiExternalSafetyProviderName)
}

// Segments returns a slice of Resource ID Segments which comprise this Rai External Safety Provider ID
func (id RaiExternalSafetyProviderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCognitiveServices", "Microsoft.CognitiveServices", "Microsoft.CognitiveServices"),
		resourceids.StaticSegment("staticRaiExternalSafetyProviders", "raiExternalSafetyProviders", "raiExternalSafetyProviders"),
		resourceids.UserSpecifiedSegment("raiExternalSafetyProviderName", "raiExternalSafetyProviderName"),
	}
}

// String returns a human-readable description of this Rai External Safety Provider ID
func (id RaiExternalSafetyProviderId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Rai External Safety Provider Name: %q", id.RaiExternalSafetyProviderName),
	}
	return fmt.Sprintf("Rai External Safety Provider (%s)", strings.Join(components, "\n"))
}
