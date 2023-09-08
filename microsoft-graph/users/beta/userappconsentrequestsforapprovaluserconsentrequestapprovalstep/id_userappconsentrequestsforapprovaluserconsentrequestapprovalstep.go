package userappconsentrequestsforapprovaluserconsentrequestapprovalstep

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserAppConsentRequestsForApprovalUserConsentRequestApprovalStepId{}

// UserAppConsentRequestsForApprovalUserConsentRequestApprovalStepId is a struct representing the Resource ID for a User App Consent Requests For Approval User Consent Request Approval Step
type UserAppConsentRequestsForApprovalUserConsentRequestApprovalStepId struct {
	UserId               string
	AppConsentRequestId  string
	UserConsentRequestId string
	ApprovalStepId       string
}

// NewUserAppConsentRequestsForApprovalUserConsentRequestApprovalStepID returns a new UserAppConsentRequestsForApprovalUserConsentRequestApprovalStepId struct
func NewUserAppConsentRequestsForApprovalUserConsentRequestApprovalStepID(userId string, appConsentRequestId string, userConsentRequestId string, approvalStepId string) UserAppConsentRequestsForApprovalUserConsentRequestApprovalStepId {
	return UserAppConsentRequestsForApprovalUserConsentRequestApprovalStepId{
		UserId:               userId,
		AppConsentRequestId:  appConsentRequestId,
		UserConsentRequestId: userConsentRequestId,
		ApprovalStepId:       approvalStepId,
	}
}

// ParseUserAppConsentRequestsForApprovalUserConsentRequestApprovalStepID parses 'input' into a UserAppConsentRequestsForApprovalUserConsentRequestApprovalStepId
func ParseUserAppConsentRequestsForApprovalUserConsentRequestApprovalStepID(input string) (*UserAppConsentRequestsForApprovalUserConsentRequestApprovalStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAppConsentRequestsForApprovalUserConsentRequestApprovalStepId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAppConsentRequestsForApprovalUserConsentRequestApprovalStepId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.AppConsentRequestId, ok = parsed.Parsed["appConsentRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appConsentRequestId", *parsed)
	}

	if id.UserConsentRequestId, ok = parsed.Parsed["userConsentRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userConsentRequestId", *parsed)
	}

	if id.ApprovalStepId, ok = parsed.Parsed["approvalStepId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalStepId", *parsed)
	}

	return &id, nil
}

// ParseUserAppConsentRequestsForApprovalUserConsentRequestApprovalStepIDInsensitively parses 'input' case-insensitively into a UserAppConsentRequestsForApprovalUserConsentRequestApprovalStepId
// note: this method should only be used for API response data and not user input
func ParseUserAppConsentRequestsForApprovalUserConsentRequestApprovalStepIDInsensitively(input string) (*UserAppConsentRequestsForApprovalUserConsentRequestApprovalStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAppConsentRequestsForApprovalUserConsentRequestApprovalStepId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAppConsentRequestsForApprovalUserConsentRequestApprovalStepId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.AppConsentRequestId, ok = parsed.Parsed["appConsentRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appConsentRequestId", *parsed)
	}

	if id.UserConsentRequestId, ok = parsed.Parsed["userConsentRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userConsentRequestId", *parsed)
	}

	if id.ApprovalStepId, ok = parsed.Parsed["approvalStepId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalStepId", *parsed)
	}

	return &id, nil
}

// ValidateUserAppConsentRequestsForApprovalUserConsentRequestApprovalStepID checks that 'input' can be parsed as a User App Consent Requests For Approval User Consent Request Approval Step ID
func ValidateUserAppConsentRequestsForApprovalUserConsentRequestApprovalStepID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserAppConsentRequestsForApprovalUserConsentRequestApprovalStepID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User App Consent Requests For Approval User Consent Request Approval Step ID
func (id UserAppConsentRequestsForApprovalUserConsentRequestApprovalStepId) ID() string {
	fmtString := "/users/%s/appConsentRequestsForApproval/%s/userConsentRequests/%s/approval/steps/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AppConsentRequestId, id.UserConsentRequestId, id.ApprovalStepId)
}

// Segments returns a slice of Resource ID Segments which comprise this User App Consent Requests For Approval User Consent Request Approval Step ID
func (id UserAppConsentRequestsForApprovalUserConsentRequestApprovalStepId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticAppConsentRequestsForApproval", "appConsentRequestsForApproval", "appConsentRequestsForApproval"),
		resourceids.UserSpecifiedSegment("appConsentRequestId", "appConsentRequestIdValue"),
		resourceids.StaticSegment("staticUserConsentRequests", "userConsentRequests", "userConsentRequests"),
		resourceids.UserSpecifiedSegment("userConsentRequestId", "userConsentRequestIdValue"),
		resourceids.StaticSegment("staticApproval", "approval", "approval"),
		resourceids.StaticSegment("staticSteps", "steps", "steps"),
		resourceids.UserSpecifiedSegment("approvalStepId", "approvalStepIdValue"),
	}
}

// String returns a human-readable description of this User App Consent Requests For Approval User Consent Request Approval Step ID
func (id UserAppConsentRequestsForApprovalUserConsentRequestApprovalStepId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("App Consent Request: %q", id.AppConsentRequestId),
		fmt.Sprintf("User Consent Request: %q", id.UserConsentRequestId),
		fmt.Sprintf("Approval Step: %q", id.ApprovalStepId),
	}
	return fmt.Sprintf("User App Consent Requests For Approval User Consent Request Approval Step (%s)", strings.Join(components, "\n"))
}
