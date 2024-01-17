package regulatorycompliance

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RegulatoryComplianceControlId{}

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
	parser := resourceids.NewParserFromResourceIdType(&RegulatoryComplianceControlId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RegulatoryComplianceControlId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRegulatoryComplianceControlIDInsensitively parses 'input' case-insensitively into a RegulatoryComplianceControlId
// note: this method should only be used for API response data and not user input
func ParseRegulatoryComplianceControlIDInsensitively(input string) (*RegulatoryComplianceControlId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RegulatoryComplianceControlId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RegulatoryComplianceControlId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RegulatoryComplianceControlId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.RegulatoryComplianceStandardName, ok = input.Parsed["regulatoryComplianceStandardName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "regulatoryComplianceStandardName", input)
	}

	if id.RegulatoryComplianceControlName, ok = input.Parsed["regulatoryComplianceControlName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "regulatoryComplianceControlName", input)
	}

	return nil
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
