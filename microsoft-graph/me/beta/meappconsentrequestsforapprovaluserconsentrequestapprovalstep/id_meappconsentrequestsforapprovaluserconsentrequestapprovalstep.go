package meappconsentrequestsforapprovaluserconsentrequestapprovalstep

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepId{}

// MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepId is a struct representing the Resource ID for a Me App Consent Requests For Approval User Consent Request Approval Step
type MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepId struct {
	AppConsentRequestId  string
	UserConsentRequestId string
	ApprovalStepId       string
}

// NewMeAppConsentRequestsForApprovalUserConsentRequestApprovalStepID returns a new MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepId struct
func NewMeAppConsentRequestsForApprovalUserConsentRequestApprovalStepID(appConsentRequestId string, userConsentRequestId string, approvalStepId string) MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepId {
	return MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepId{
		AppConsentRequestId:  appConsentRequestId,
		UserConsentRequestId: userConsentRequestId,
		ApprovalStepId:       approvalStepId,
	}
}

// ParseMeAppConsentRequestsForApprovalUserConsentRequestApprovalStepID parses 'input' into a MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepId
func ParseMeAppConsentRequestsForApprovalUserConsentRequestApprovalStepID(input string) (*MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepId{}

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

// ParseMeAppConsentRequestsForApprovalUserConsentRequestApprovalStepIDInsensitively parses 'input' case-insensitively into a MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepId
// note: this method should only be used for API response data and not user input
func ParseMeAppConsentRequestsForApprovalUserConsentRequestApprovalStepIDInsensitively(input string) (*MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepId{}

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

// ValidateMeAppConsentRequestsForApprovalUserConsentRequestApprovalStepID checks that 'input' can be parsed as a Me App Consent Requests For Approval User Consent Request Approval Step ID
func ValidateMeAppConsentRequestsForApprovalUserConsentRequestApprovalStepID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeAppConsentRequestsForApprovalUserConsentRequestApprovalStepID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me App Consent Requests For Approval User Consent Request Approval Step ID
func (id MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepId) ID() string {
	fmtString := "/me/appConsentRequestsForApproval/%s/userConsentRequests/%s/approval/steps/%s"
	return fmt.Sprintf(fmtString, id.AppConsentRequestId, id.UserConsentRequestId, id.ApprovalStepId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me App Consent Requests For Approval User Consent Request Approval Step ID
func (id MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticAppConsentRequestsForApproval", "appConsentRequestsForApproval", "appConsentRequestsForApproval"),
		resourceids.UserSpecifiedSegment("appConsentRequestId", "appConsentRequestIdValue"),
		resourceids.StaticSegment("staticUserConsentRequests", "userConsentRequests", "userConsentRequests"),
		resourceids.UserSpecifiedSegment("userConsentRequestId", "userConsentRequestIdValue"),
		resourceids.StaticSegment("staticApproval", "approval", "approval"),
		resourceids.StaticSegment("staticSteps", "steps", "steps"),
		resourceids.UserSpecifiedSegment("approvalStepId", "approvalStepIdValue"),
	}
}

// String returns a human-readable description of this Me App Consent Requests For Approval User Consent Request Approval Step ID
func (id MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepId) String() string {
	components := []string{
		fmt.Sprintf("App Consent Request: %q", id.AppConsentRequestId),
		fmt.Sprintf("User Consent Request: %q", id.UserConsentRequestId),
		fmt.Sprintf("Approval Step: %q", id.ApprovalStepId),
	}
	return fmt.Sprintf("Me App Consent Requests For Approval User Consent Request Approval Step (%s)", strings.Join(components, "\n"))
}
