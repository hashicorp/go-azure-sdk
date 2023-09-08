package userinformationprotectionthreatassessmentrequestresult

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserInformationProtectionThreatAssessmentRequestId{}

// UserInformationProtectionThreatAssessmentRequestId is a struct representing the Resource ID for a User Information Protection Threat Assessment Request
type UserInformationProtectionThreatAssessmentRequestId struct {
	UserId                    string
	ThreatAssessmentRequestId string
}

// NewUserInformationProtectionThreatAssessmentRequestID returns a new UserInformationProtectionThreatAssessmentRequestId struct
func NewUserInformationProtectionThreatAssessmentRequestID(userId string, threatAssessmentRequestId string) UserInformationProtectionThreatAssessmentRequestId {
	return UserInformationProtectionThreatAssessmentRequestId{
		UserId:                    userId,
		ThreatAssessmentRequestId: threatAssessmentRequestId,
	}
}

// ParseUserInformationProtectionThreatAssessmentRequestID parses 'input' into a UserInformationProtectionThreatAssessmentRequestId
func ParseUserInformationProtectionThreatAssessmentRequestID(input string) (*UserInformationProtectionThreatAssessmentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserInformationProtectionThreatAssessmentRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserInformationProtectionThreatAssessmentRequestId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ThreatAssessmentRequestId, ok = parsed.Parsed["threatAssessmentRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "threatAssessmentRequestId", *parsed)
	}

	return &id, nil
}

// ParseUserInformationProtectionThreatAssessmentRequestIDInsensitively parses 'input' case-insensitively into a UserInformationProtectionThreatAssessmentRequestId
// note: this method should only be used for API response data and not user input
func ParseUserInformationProtectionThreatAssessmentRequestIDInsensitively(input string) (*UserInformationProtectionThreatAssessmentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserInformationProtectionThreatAssessmentRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserInformationProtectionThreatAssessmentRequestId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ThreatAssessmentRequestId, ok = parsed.Parsed["threatAssessmentRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "threatAssessmentRequestId", *parsed)
	}

	return &id, nil
}

// ValidateUserInformationProtectionThreatAssessmentRequestID checks that 'input' can be parsed as a User Information Protection Threat Assessment Request ID
func ValidateUserInformationProtectionThreatAssessmentRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserInformationProtectionThreatAssessmentRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Information Protection Threat Assessment Request ID
func (id UserInformationProtectionThreatAssessmentRequestId) ID() string {
	fmtString := "/users/%s/informationProtection/threatAssessmentRequests/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ThreatAssessmentRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Information Protection Threat Assessment Request ID
func (id UserInformationProtectionThreatAssessmentRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticInformationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("staticThreatAssessmentRequests", "threatAssessmentRequests", "threatAssessmentRequests"),
		resourceids.UserSpecifiedSegment("threatAssessmentRequestId", "threatAssessmentRequestIdValue"),
	}
}

// String returns a human-readable description of this User Information Protection Threat Assessment Request ID
func (id UserInformationProtectionThreatAssessmentRequestId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Threat Assessment Request: %q", id.ThreatAssessmentRequestId),
	}
	return fmt.Sprintf("User Information Protection Threat Assessment Request (%s)", strings.Join(components, "\n"))
}
