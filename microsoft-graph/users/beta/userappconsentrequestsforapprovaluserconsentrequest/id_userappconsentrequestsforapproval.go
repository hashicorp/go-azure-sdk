package userappconsentrequestsforapprovaluserconsentrequest

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserAppConsentRequestsForApprovalId{}

// UserAppConsentRequestsForApprovalId is a struct representing the Resource ID for a User App Consent Requests For Approval
type UserAppConsentRequestsForApprovalId struct {
	UserId              string
	AppConsentRequestId string
}

// NewUserAppConsentRequestsForApprovalID returns a new UserAppConsentRequestsForApprovalId struct
func NewUserAppConsentRequestsForApprovalID(userId string, appConsentRequestId string) UserAppConsentRequestsForApprovalId {
	return UserAppConsentRequestsForApprovalId{
		UserId:              userId,
		AppConsentRequestId: appConsentRequestId,
	}
}

// ParseUserAppConsentRequestsForApprovalID parses 'input' into a UserAppConsentRequestsForApprovalId
func ParseUserAppConsentRequestsForApprovalID(input string) (*UserAppConsentRequestsForApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAppConsentRequestsForApprovalId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAppConsentRequestsForApprovalId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.AppConsentRequestId, ok = parsed.Parsed["appConsentRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appConsentRequestId", *parsed)
	}

	return &id, nil
}

// ParseUserAppConsentRequestsForApprovalIDInsensitively parses 'input' case-insensitively into a UserAppConsentRequestsForApprovalId
// note: this method should only be used for API response data and not user input
func ParseUserAppConsentRequestsForApprovalIDInsensitively(input string) (*UserAppConsentRequestsForApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAppConsentRequestsForApprovalId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAppConsentRequestsForApprovalId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.AppConsentRequestId, ok = parsed.Parsed["appConsentRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appConsentRequestId", *parsed)
	}

	return &id, nil
}

// ValidateUserAppConsentRequestsForApprovalID checks that 'input' can be parsed as a User App Consent Requests For Approval ID
func ValidateUserAppConsentRequestsForApprovalID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserAppConsentRequestsForApprovalID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User App Consent Requests For Approval ID
func (id UserAppConsentRequestsForApprovalId) ID() string {
	fmtString := "/users/%s/appConsentRequestsForApproval/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AppConsentRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this User App Consent Requests For Approval ID
func (id UserAppConsentRequestsForApprovalId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticAppConsentRequestsForApproval", "appConsentRequestsForApproval", "appConsentRequestsForApproval"),
		resourceids.UserSpecifiedSegment("appConsentRequestId", "appConsentRequestIdValue"),
	}
}

// String returns a human-readable description of this User App Consent Requests For Approval ID
func (id UserAppConsentRequestsForApprovalId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("App Consent Request: %q", id.AppConsentRequestId),
	}
	return fmt.Sprintf("User App Consent Requests For Approval (%s)", strings.Join(components, "\n"))
}
