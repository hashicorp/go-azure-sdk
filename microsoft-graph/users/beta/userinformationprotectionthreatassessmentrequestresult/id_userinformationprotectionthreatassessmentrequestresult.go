package userinformationprotectionthreatassessmentrequestresult

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserInformationProtectionThreatAssessmentRequestResultId{}

// UserInformationProtectionThreatAssessmentRequestResultId is a struct representing the Resource ID for a User Information Protection Threat Assessment Request Result
type UserInformationProtectionThreatAssessmentRequestResultId struct {
	UserId                    string
	ThreatAssessmentRequestId string
	ThreatAssessmentResultId  string
}

// NewUserInformationProtectionThreatAssessmentRequestResultID returns a new UserInformationProtectionThreatAssessmentRequestResultId struct
func NewUserInformationProtectionThreatAssessmentRequestResultID(userId string, threatAssessmentRequestId string, threatAssessmentResultId string) UserInformationProtectionThreatAssessmentRequestResultId {
	return UserInformationProtectionThreatAssessmentRequestResultId{
		UserId:                    userId,
		ThreatAssessmentRequestId: threatAssessmentRequestId,
		ThreatAssessmentResultId:  threatAssessmentResultId,
	}
}

// ParseUserInformationProtectionThreatAssessmentRequestResultID parses 'input' into a UserInformationProtectionThreatAssessmentRequestResultId
func ParseUserInformationProtectionThreatAssessmentRequestResultID(input string) (*UserInformationProtectionThreatAssessmentRequestResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserInformationProtectionThreatAssessmentRequestResultId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserInformationProtectionThreatAssessmentRequestResultId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ThreatAssessmentRequestId, ok = parsed.Parsed["threatAssessmentRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "threatAssessmentRequestId", *parsed)
	}

	if id.ThreatAssessmentResultId, ok = parsed.Parsed["threatAssessmentResultId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "threatAssessmentResultId", *parsed)
	}

	return &id, nil
}

// ParseUserInformationProtectionThreatAssessmentRequestResultIDInsensitively parses 'input' case-insensitively into a UserInformationProtectionThreatAssessmentRequestResultId
// note: this method should only be used for API response data and not user input
func ParseUserInformationProtectionThreatAssessmentRequestResultIDInsensitively(input string) (*UserInformationProtectionThreatAssessmentRequestResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserInformationProtectionThreatAssessmentRequestResultId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserInformationProtectionThreatAssessmentRequestResultId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ThreatAssessmentRequestId, ok = parsed.Parsed["threatAssessmentRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "threatAssessmentRequestId", *parsed)
	}

	if id.ThreatAssessmentResultId, ok = parsed.Parsed["threatAssessmentResultId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "threatAssessmentResultId", *parsed)
	}

	return &id, nil
}

// ValidateUserInformationProtectionThreatAssessmentRequestResultID checks that 'input' can be parsed as a User Information Protection Threat Assessment Request Result ID
func ValidateUserInformationProtectionThreatAssessmentRequestResultID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserInformationProtectionThreatAssessmentRequestResultID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Information Protection Threat Assessment Request Result ID
func (id UserInformationProtectionThreatAssessmentRequestResultId) ID() string {
	fmtString := "/users/%s/informationProtection/threatAssessmentRequests/%s/results/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ThreatAssessmentRequestId, id.ThreatAssessmentResultId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Information Protection Threat Assessment Request Result ID
func (id UserInformationProtectionThreatAssessmentRequestResultId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticInformationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("staticThreatAssessmentRequests", "threatAssessmentRequests", "threatAssessmentRequests"),
		resourceids.UserSpecifiedSegment("threatAssessmentRequestId", "threatAssessmentRequestIdValue"),
		resourceids.StaticSegment("staticResults", "results", "results"),
		resourceids.UserSpecifiedSegment("threatAssessmentResultId", "threatAssessmentResultIdValue"),
	}
}

// String returns a human-readable description of this User Information Protection Threat Assessment Request Result ID
func (id UserInformationProtectionThreatAssessmentRequestResultId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Threat Assessment Request: %q", id.ThreatAssessmentRequestId),
		fmt.Sprintf("Threat Assessment Result: %q", id.ThreatAssessmentResultId),
	}
	return fmt.Sprintf("User Information Protection Threat Assessment Request Result (%s)", strings.Join(components, "\n"))
}
