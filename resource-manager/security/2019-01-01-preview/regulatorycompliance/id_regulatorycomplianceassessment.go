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

	id := RegulatoryComplianceAssessmentId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
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

	id := RegulatoryComplianceAssessmentId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RegulatoryComplianceAssessmentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.RegulatoryComplianceAssessmentName, ok = input.Parsed["regulatoryComplianceAssessmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "regulatoryComplianceAssessmentName", input)
	}

	return nil
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
