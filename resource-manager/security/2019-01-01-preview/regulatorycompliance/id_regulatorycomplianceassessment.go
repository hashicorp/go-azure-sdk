package regulatorycompliance

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RegulatoryComplianceAssessmentId{}

// RegulatoryComplianceAssessmentId is a struct representing the Resource ID for a Regulatory Compliance Assessment
type RegulatoryComplianceAssessmentId struct {
	SubscriptionId                     string
	RegulatoryComplianceStandardName   string
	RegulatoryComplianceControlName    string
	RegulatoryComplianceAssessmentName string
}

// NewRegulatoryComplianceAssessmentID returns a new RegulatoryComplianceAssessmentId struct
func NewRegulatoryComplianceAssessmentID(subscriptionId string, regulatoryComplianceStandardName string, regulatoryComplianceControlName string, regulatoryComplianceAssessmentName string) RegulatoryComplianceAssessmentId {
	return RegulatoryComplianceAssessmentId{
		SubscriptionId:                     subscriptionId,
		RegulatoryComplianceStandardName:   regulatoryComplianceStandardName,
		RegulatoryComplianceControlName:    regulatoryComplianceControlName,
		RegulatoryComplianceAssessmentName: regulatoryComplianceAssessmentName,
	}
}

// ParseRegulatoryComplianceAssessmentID parses 'input' into a RegulatoryComplianceAssessmentId
func ParseRegulatoryComplianceAssessmentID(input string) (*RegulatoryComplianceAssessmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(RegulatoryComplianceAssessmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RegulatoryComplianceAssessmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.RegulatoryComplianceStandardName, ok = parsed.Parsed["regulatoryComplianceStandardName"]; !ok {
		return nil, fmt.Errorf("the segment 'regulatoryComplianceStandardName' was not found in the resource id %q", input)
	}

	if id.RegulatoryComplianceControlName, ok = parsed.Parsed["regulatoryComplianceControlName"]; !ok {
		return nil, fmt.Errorf("the segment 'regulatoryComplianceControlName' was not found in the resource id %q", input)
	}

	if id.RegulatoryComplianceAssessmentName, ok = parsed.Parsed["regulatoryComplianceAssessmentName"]; !ok {
		return nil, fmt.Errorf("the segment 'regulatoryComplianceAssessmentName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseRegulatoryComplianceAssessmentIDInsensitively parses 'input' case-insensitively into a RegulatoryComplianceAssessmentId
// note: this method should only be used for API response data and not user input
func ParseRegulatoryComplianceAssessmentIDInsensitively(input string) (*RegulatoryComplianceAssessmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(RegulatoryComplianceAssessmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RegulatoryComplianceAssessmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.RegulatoryComplianceStandardName, ok = parsed.Parsed["regulatoryComplianceStandardName"]; !ok {
		return nil, fmt.Errorf("the segment 'regulatoryComplianceStandardName' was not found in the resource id %q", input)
	}

	if id.RegulatoryComplianceControlName, ok = parsed.Parsed["regulatoryComplianceControlName"]; !ok {
		return nil, fmt.Errorf("the segment 'regulatoryComplianceControlName' was not found in the resource id %q", input)
	}

	if id.RegulatoryComplianceAssessmentName, ok = parsed.Parsed["regulatoryComplianceAssessmentName"]; !ok {
		return nil, fmt.Errorf("the segment 'regulatoryComplianceAssessmentName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateRegulatoryComplianceAssessmentID checks that 'input' can be parsed as a Regulatory Compliance Assessment ID
func ValidateRegulatoryComplianceAssessmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRegulatoryComplianceAssessmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Regulatory Compliance Assessment ID
func (id RegulatoryComplianceAssessmentId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Security/regulatoryComplianceStandards/%s/regulatoryComplianceControls/%s/regulatoryComplianceAssessments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.RegulatoryComplianceStandardName, id.RegulatoryComplianceControlName, id.RegulatoryComplianceAssessmentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Regulatory Compliance Assessment ID
func (id RegulatoryComplianceAssessmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticRegulatoryComplianceStandards", "regulatoryComplianceStandards", "regulatoryComplianceStandards"),
		resourceids.UserSpecifiedSegment("regulatoryComplianceStandardName", "regulatoryComplianceStandardValue"),
		resourceids.StaticSegment("staticRegulatoryComplianceControls", "regulatoryComplianceControls", "regulatoryComplianceControls"),
		resourceids.UserSpecifiedSegment("regulatoryComplianceControlName", "regulatoryComplianceControlValue"),
		resourceids.StaticSegment("staticRegulatoryComplianceAssessments", "regulatoryComplianceAssessments", "regulatoryComplianceAssessments"),
		resourceids.UserSpecifiedSegment("regulatoryComplianceAssessmentName", "regulatoryComplianceAssessmentValue"),
	}
}

// String returns a human-readable description of this Regulatory Compliance Assessment ID
func (id RegulatoryComplianceAssessmentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Regulatory Compliance Standard Name: %q", id.RegulatoryComplianceStandardName),
		fmt.Sprintf("Regulatory Compliance Control Name: %q", id.RegulatoryComplianceControlName),
		fmt.Sprintf("Regulatory Compliance Assessment Name: %q", id.RegulatoryComplianceAssessmentName),
	}
	return fmt.Sprintf("Regulatory Compliance Assessment (%s)", strings.Join(components, "\n"))
}
