package regulatorycompliance

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RegulatoryComplianceStandardId{}

// RegulatoryComplianceStandardId is a struct representing the Resource ID for a Regulatory Compliance Standard
type RegulatoryComplianceStandardId struct {
	SubscriptionId                   string
	RegulatoryComplianceStandardName string
}

// NewRegulatoryComplianceStandardID returns a new RegulatoryComplianceStandardId struct
func NewRegulatoryComplianceStandardID(subscriptionId string, regulatoryComplianceStandardName string) RegulatoryComplianceStandardId {
	return RegulatoryComplianceStandardId{
		SubscriptionId:                   subscriptionId,
		RegulatoryComplianceStandardName: regulatoryComplianceStandardName,
	}
}

// ParseRegulatoryComplianceStandardID parses 'input' into a RegulatoryComplianceStandardId
func ParseRegulatoryComplianceStandardID(input string) (*RegulatoryComplianceStandardId, error) {
	parser := resourceids.NewParserFromResourceIdType(RegulatoryComplianceStandardId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RegulatoryComplianceStandardId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.RegulatoryComplianceStandardName, ok = parsed.Parsed["regulatoryComplianceStandardName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "regulatoryComplianceStandardName", *parsed)
	}

	return &id, nil
}

// ParseRegulatoryComplianceStandardIDInsensitively parses 'input' case-insensitively into a RegulatoryComplianceStandardId
// note: this method should only be used for API response data and not user input
func ParseRegulatoryComplianceStandardIDInsensitively(input string) (*RegulatoryComplianceStandardId, error) {
	parser := resourceids.NewParserFromResourceIdType(RegulatoryComplianceStandardId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RegulatoryComplianceStandardId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.RegulatoryComplianceStandardName, ok = parsed.Parsed["regulatoryComplianceStandardName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "regulatoryComplianceStandardName", *parsed)
	}

	return &id, nil
}

// ValidateRegulatoryComplianceStandardID checks that 'input' can be parsed as a Regulatory Compliance Standard ID
func ValidateRegulatoryComplianceStandardID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRegulatoryComplianceStandardID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Regulatory Compliance Standard ID
func (id RegulatoryComplianceStandardId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Security/regulatoryComplianceStandards/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.RegulatoryComplianceStandardName)
}

// Segments returns a slice of Resource ID Segments which comprise this Regulatory Compliance Standard ID
func (id RegulatoryComplianceStandardId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticRegulatoryComplianceStandards", "regulatoryComplianceStandards", "regulatoryComplianceStandards"),
		resourceids.UserSpecifiedSegment("regulatoryComplianceStandardName", "regulatoryComplianceStandardValue"),
	}
}

// String returns a human-readable description of this Regulatory Compliance Standard ID
func (id RegulatoryComplianceStandardId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Regulatory Compliance Standard Name: %q", id.RegulatoryComplianceStandardName),
	}
	return fmt.Sprintf("Regulatory Compliance Standard (%s)", strings.Join(components, "\n"))
}
