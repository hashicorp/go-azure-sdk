package meinformationprotectionthreatassessmentrequestresult

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeInformationProtectionThreatAssessmentRequestResultId{}

// MeInformationProtectionThreatAssessmentRequestResultId is a struct representing the Resource ID for a Me Information Protection Threat Assessment Request Result
type MeInformationProtectionThreatAssessmentRequestResultId struct {
	ThreatAssessmentRequestId string
	ThreatAssessmentResultId  string
}

// NewMeInformationProtectionThreatAssessmentRequestResultID returns a new MeInformationProtectionThreatAssessmentRequestResultId struct
func NewMeInformationProtectionThreatAssessmentRequestResultID(threatAssessmentRequestId string, threatAssessmentResultId string) MeInformationProtectionThreatAssessmentRequestResultId {
	return MeInformationProtectionThreatAssessmentRequestResultId{
		ThreatAssessmentRequestId: threatAssessmentRequestId,
		ThreatAssessmentResultId:  threatAssessmentResultId,
	}
}

// ParseMeInformationProtectionThreatAssessmentRequestResultID parses 'input' into a MeInformationProtectionThreatAssessmentRequestResultId
func ParseMeInformationProtectionThreatAssessmentRequestResultID(input string) (*MeInformationProtectionThreatAssessmentRequestResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeInformationProtectionThreatAssessmentRequestResultId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeInformationProtectionThreatAssessmentRequestResultId{}

	if id.ThreatAssessmentRequestId, ok = parsed.Parsed["threatAssessmentRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "threatAssessmentRequestId", *parsed)
	}

	if id.ThreatAssessmentResultId, ok = parsed.Parsed["threatAssessmentResultId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "threatAssessmentResultId", *parsed)
	}

	return &id, nil
}

// ParseMeInformationProtectionThreatAssessmentRequestResultIDInsensitively parses 'input' case-insensitively into a MeInformationProtectionThreatAssessmentRequestResultId
// note: this method should only be used for API response data and not user input
func ParseMeInformationProtectionThreatAssessmentRequestResultIDInsensitively(input string) (*MeInformationProtectionThreatAssessmentRequestResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeInformationProtectionThreatAssessmentRequestResultId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeInformationProtectionThreatAssessmentRequestResultId{}

	if id.ThreatAssessmentRequestId, ok = parsed.Parsed["threatAssessmentRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "threatAssessmentRequestId", *parsed)
	}

	if id.ThreatAssessmentResultId, ok = parsed.Parsed["threatAssessmentResultId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "threatAssessmentResultId", *parsed)
	}

	return &id, nil
}

// ValidateMeInformationProtectionThreatAssessmentRequestResultID checks that 'input' can be parsed as a Me Information Protection Threat Assessment Request Result ID
func ValidateMeInformationProtectionThreatAssessmentRequestResultID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeInformationProtectionThreatAssessmentRequestResultID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Information Protection Threat Assessment Request Result ID
func (id MeInformationProtectionThreatAssessmentRequestResultId) ID() string {
	fmtString := "/me/informationProtection/threatAssessmentRequests/%s/results/%s"
	return fmt.Sprintf(fmtString, id.ThreatAssessmentRequestId, id.ThreatAssessmentResultId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Information Protection Threat Assessment Request Result ID
func (id MeInformationProtectionThreatAssessmentRequestResultId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticInformationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("staticThreatAssessmentRequests", "threatAssessmentRequests", "threatAssessmentRequests"),
		resourceids.UserSpecifiedSegment("threatAssessmentRequestId", "threatAssessmentRequestIdValue"),
		resourceids.StaticSegment("staticResults", "results", "results"),
		resourceids.UserSpecifiedSegment("threatAssessmentResultId", "threatAssessmentResultIdValue"),
	}
}

// String returns a human-readable description of this Me Information Protection Threat Assessment Request Result ID
func (id MeInformationProtectionThreatAssessmentRequestResultId) String() string {
	components := []string{
		fmt.Sprintf("Threat Assessment Request: %q", id.ThreatAssessmentRequestId),
		fmt.Sprintf("Threat Assessment Result: %q", id.ThreatAssessmentResultId),
	}
	return fmt.Sprintf("Me Information Protection Threat Assessment Request Result (%s)", strings.Join(components, "\n"))
}
