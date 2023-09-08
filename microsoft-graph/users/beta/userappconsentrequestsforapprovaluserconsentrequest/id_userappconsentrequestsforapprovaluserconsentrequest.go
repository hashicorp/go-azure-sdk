package userappconsentrequestsforapprovaluserconsentrequest

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserAppConsentRequestsForApprovalUserConsentRequestId{}

// UserAppConsentRequestsForApprovalUserConsentRequestId is a struct representing the Resource ID for a User App Consent Requests For Approval User Consent Request
type UserAppConsentRequestsForApprovalUserConsentRequestId struct {
	UserId               string
	AppConsentRequestId  string
	UserConsentRequestId string
}

// NewUserAppConsentRequestsForApprovalUserConsentRequestID returns a new UserAppConsentRequestsForApprovalUserConsentRequestId struct
func NewUserAppConsentRequestsForApprovalUserConsentRequestID(userId string, appConsentRequestId string, userConsentRequestId string) UserAppConsentRequestsForApprovalUserConsentRequestId {
	return UserAppConsentRequestsForApprovalUserConsentRequestId{
		UserId:               userId,
		AppConsentRequestId:  appConsentRequestId,
		UserConsentRequestId: userConsentRequestId,
	}
}

// ParseUserAppConsentRequestsForApprovalUserConsentRequestID parses 'input' into a UserAppConsentRequestsForApprovalUserConsentRequestId
func ParseUserAppConsentRequestsForApprovalUserConsentRequestID(input string) (*UserAppConsentRequestsForApprovalUserConsentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAppConsentRequestsForApprovalUserConsentRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAppConsentRequestsForApprovalUserConsentRequestId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.AppConsentRequestId, ok = parsed.Parsed["appConsentRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appConsentRequestId", *parsed)
	}

	if id.UserConsentRequestId, ok = parsed.Parsed["userConsentRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userConsentRequestId", *parsed)
	}

	return &id, nil
}

// ParseUserAppConsentRequestsForApprovalUserConsentRequestIDInsensitively parses 'input' case-insensitively into a UserAppConsentRequestsForApprovalUserConsentRequestId
// note: this method should only be used for API response data and not user input
func ParseUserAppConsentRequestsForApprovalUserConsentRequestIDInsensitively(input string) (*UserAppConsentRequestsForApprovalUserConsentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAppConsentRequestsForApprovalUserConsentRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAppConsentRequestsForApprovalUserConsentRequestId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.AppConsentRequestId, ok = parsed.Parsed["appConsentRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appConsentRequestId", *parsed)
	}

	if id.UserConsentRequestId, ok = parsed.Parsed["userConsentRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userConsentRequestId", *parsed)
	}

	return &id, nil
}

// ValidateUserAppConsentRequestsForApprovalUserConsentRequestID checks that 'input' can be parsed as a User App Consent Requests For Approval User Consent Request ID
func ValidateUserAppConsentRequestsForApprovalUserConsentRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserAppConsentRequestsForApprovalUserConsentRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User App Consent Requests For Approval User Consent Request ID
func (id UserAppConsentRequestsForApprovalUserConsentRequestId) ID() string {
	fmtString := "/users/%s/appConsentRequestsForApproval/%s/userConsentRequests/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AppConsentRequestId, id.UserConsentRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this User App Consent Requests For Approval User Consent Request ID
func (id UserAppConsentRequestsForApprovalUserConsentRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticAppConsentRequestsForApproval", "appConsentRequestsForApproval", "appConsentRequestsForApproval"),
		resourceids.UserSpecifiedSegment("appConsentRequestId", "appConsentRequestIdValue"),
		resourceids.StaticSegment("staticUserConsentRequests", "userConsentRequests", "userConsentRequests"),
		resourceids.UserSpecifiedSegment("userConsentRequestId", "userConsentRequestIdValue"),
	}
}

// String returns a human-readable description of this User App Consent Requests For Approval User Consent Request ID
func (id UserAppConsentRequestsForApprovalUserConsentRequestId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("App Consent Request: %q", id.AppConsentRequestId),
		fmt.Sprintf("User Consent Request: %q", id.UserConsentRequestId),
	}
	return fmt.Sprintf("User App Consent Requests For Approval User Consent Request (%s)", strings.Join(components, "\n"))
}
