package meappconsentrequestsforapprovaluserconsentrequestapproval

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeAppConsentRequestsForApprovalUserConsentRequestId{}

// MeAppConsentRequestsForApprovalUserConsentRequestId is a struct representing the Resource ID for a Me App Consent Requests For Approval User Consent Request
type MeAppConsentRequestsForApprovalUserConsentRequestId struct {
	AppConsentRequestId  string
	UserConsentRequestId string
}

// NewMeAppConsentRequestsForApprovalUserConsentRequestID returns a new MeAppConsentRequestsForApprovalUserConsentRequestId struct
func NewMeAppConsentRequestsForApprovalUserConsentRequestID(appConsentRequestId string, userConsentRequestId string) MeAppConsentRequestsForApprovalUserConsentRequestId {
	return MeAppConsentRequestsForApprovalUserConsentRequestId{
		AppConsentRequestId:  appConsentRequestId,
		UserConsentRequestId: userConsentRequestId,
	}
}

// ParseMeAppConsentRequestsForApprovalUserConsentRequestID parses 'input' into a MeAppConsentRequestsForApprovalUserConsentRequestId
func ParseMeAppConsentRequestsForApprovalUserConsentRequestID(input string) (*MeAppConsentRequestsForApprovalUserConsentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeAppConsentRequestsForApprovalUserConsentRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeAppConsentRequestsForApprovalUserConsentRequestId{}

	if id.AppConsentRequestId, ok = parsed.Parsed["appConsentRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appConsentRequestId", *parsed)
	}

	if id.UserConsentRequestId, ok = parsed.Parsed["userConsentRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userConsentRequestId", *parsed)
	}

	return &id, nil
}

// ParseMeAppConsentRequestsForApprovalUserConsentRequestIDInsensitively parses 'input' case-insensitively into a MeAppConsentRequestsForApprovalUserConsentRequestId
// note: this method should only be used for API response data and not user input
func ParseMeAppConsentRequestsForApprovalUserConsentRequestIDInsensitively(input string) (*MeAppConsentRequestsForApprovalUserConsentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeAppConsentRequestsForApprovalUserConsentRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeAppConsentRequestsForApprovalUserConsentRequestId{}

	if id.AppConsentRequestId, ok = parsed.Parsed["appConsentRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appConsentRequestId", *parsed)
	}

	if id.UserConsentRequestId, ok = parsed.Parsed["userConsentRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userConsentRequestId", *parsed)
	}

	return &id, nil
}

// ValidateMeAppConsentRequestsForApprovalUserConsentRequestID checks that 'input' can be parsed as a Me App Consent Requests For Approval User Consent Request ID
func ValidateMeAppConsentRequestsForApprovalUserConsentRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeAppConsentRequestsForApprovalUserConsentRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me App Consent Requests For Approval User Consent Request ID
func (id MeAppConsentRequestsForApprovalUserConsentRequestId) ID() string {
	fmtString := "/me/appConsentRequestsForApproval/%s/userConsentRequests/%s"
	return fmt.Sprintf(fmtString, id.AppConsentRequestId, id.UserConsentRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me App Consent Requests For Approval User Consent Request ID
func (id MeAppConsentRequestsForApprovalUserConsentRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticAppConsentRequestsForApproval", "appConsentRequestsForApproval", "appConsentRequestsForApproval"),
		resourceids.UserSpecifiedSegment("appConsentRequestId", "appConsentRequestIdValue"),
		resourceids.StaticSegment("staticUserConsentRequests", "userConsentRequests", "userConsentRequests"),
		resourceids.UserSpecifiedSegment("userConsentRequestId", "userConsentRequestIdValue"),
	}
}

// String returns a human-readable description of this Me App Consent Requests For Approval User Consent Request ID
func (id MeAppConsentRequestsForApprovalUserConsentRequestId) String() string {
	components := []string{
		fmt.Sprintf("App Consent Request: %q", id.AppConsentRequestId),
		fmt.Sprintf("User Consent Request: %q", id.UserConsentRequestId),
	}
	return fmt.Sprintf("Me App Consent Requests For Approval User Consent Request (%s)", strings.Join(components, "\n"))
}
