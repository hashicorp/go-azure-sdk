package regulatorycompliance

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RegulatoryComplianceControlId{}

// RegulatoryComplianceControlId is a struct representing the Resource ID for a Regulatory Compliance Control
type RegulatoryComplianceControlId struct {
	SubscriptionId                   string
	RegulatoryComplianceStandardName string
	RegulatoryComplianceControlName  string
}

// NewRegulatoryComplianceControlID returns a new RegulatoryComplianceControlId struct
func NewRegulatoryComplianceControlID(subscriptionId string, regulatoryComplianceStandardName string, regulatoryComplianceControlName string) RegulatoryComplianceControlId {
	return RegulatoryComplianceControlId{
		SubscriptionId:                   subscriptionId,
		RegulatoryComplianceStandardName: regulatoryComplianceStandardName,
		RegulatoryComplianceControlName:  regulatoryComplianceControlName,
	}
}

// ParseRegulatoryComplianceControlID parses 'input' into a RegulatoryComplianceControlId
func ParseRegulatoryComplianceControlID(input string) (*RegulatoryComplianceControlId, error) {
	parser := resourceids.NewParserFromResourceIdType(RegulatoryComplianceControlId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RegulatoryComplianceControlId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.RegulatoryComplianceStandardName, ok = parsed.Parsed["regulatoryComplianceStandardName"]; !ok {
		return nil, fmt.Errorf("the segment 'regulatoryComplianceStandardName' was not found in the resource id %q", input)
	}

	if id.RegulatoryComplianceControlName, ok = parsed.Parsed["regulatoryComplianceControlName"]; !ok {
		return nil, fmt.Errorf("the segment 'regulatoryComplianceControlName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseRegulatoryComplianceControlIDInsensitively parses 'input' case-insensitively into a RegulatoryComplianceControlId
// note: this method should only be used for API response data and not user input
func ParseRegulatoryComplianceControlIDInsensitively(input string) (*RegulatoryComplianceControlId, error) {
	parser := resourceids.NewParserFromResourceIdType(RegulatoryComplianceControlId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RegulatoryComplianceControlId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.RegulatoryComplianceStandardName, ok = parsed.Parsed["regulatoryComplianceStandardName"]; !ok {
		return nil, fmt.Errorf("the segment 'regulatoryComplianceStandardName' was not found in the resource id %q", input)
	}

	if id.RegulatoryComplianceControlName, ok = parsed.Parsed["regulatoryComplianceControlName"]; !ok {
		return nil, fmt.Errorf("the segment 'regulatoryComplianceControlName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateRegulatoryComplianceControlID checks that 'input' can be parsed as a Regulatory Compliance Control ID
func ValidateRegulatoryComplianceControlID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRegulatoryComplianceControlID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Regulatory Compliance Control ID
func (id RegulatoryComplianceControlId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Security/regulatoryComplianceStandards/%s/regulatoryComplianceControls/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.RegulatoryComplianceStandardName, id.RegulatoryComplianceControlName)
}

// Segments returns a slice of Resource ID Segments which comprise this Regulatory Compliance Control ID
func (id RegulatoryComplianceControlId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticRegulatoryComplianceStandards", "regulatoryComplianceStandards", "regulatoryComplianceStandards"),
		resourceids.UserSpecifiedSegment("regulatoryComplianceStandardName", "regulatoryComplianceStandardValue"),
		resourceids.StaticSegment("staticRegulatoryComplianceControls", "regulatoryComplianceControls", "regulatoryComplianceControls"),
		resourceids.UserSpecifiedSegment("regulatoryComplianceControlName", "regulatoryComplianceControlValue"),
	}
}

// String returns a human-readable description of this Regulatory Compliance Control ID
func (id RegulatoryComplianceControlId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Regulatory Compliance Standard Name: %q", id.RegulatoryComplianceStandardName),
		fmt.Sprintf("Regulatory Compliance Control Name: %q", id.RegulatoryComplianceControlName),
	}
	return fmt.Sprintf("Regulatory Compliance Control (%s)", strings.Join(components, "\n"))
}
