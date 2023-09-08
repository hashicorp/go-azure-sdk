package userpendingaccessreviewinstancecontactedreviewer

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPendingAccessReviewInstanceContactedReviewerId{}

// UserPendingAccessReviewInstanceContactedReviewerId is a struct representing the Resource ID for a User Pending Access Review Instance Contacted Reviewer
type UserPendingAccessReviewInstanceContactedReviewerId struct {
	UserId                 string
	AccessReviewInstanceId string
	AccessReviewReviewerId string
}

// NewUserPendingAccessReviewInstanceContactedReviewerID returns a new UserPendingAccessReviewInstanceContactedReviewerId struct
func NewUserPendingAccessReviewInstanceContactedReviewerID(userId string, accessReviewInstanceId string, accessReviewReviewerId string) UserPendingAccessReviewInstanceContactedReviewerId {
	return UserPendingAccessReviewInstanceContactedReviewerId{
		UserId:                 userId,
		AccessReviewInstanceId: accessReviewInstanceId,
		AccessReviewReviewerId: accessReviewReviewerId,
	}
}

// ParseUserPendingAccessReviewInstanceContactedReviewerID parses 'input' into a UserPendingAccessReviewInstanceContactedReviewerId
func ParseUserPendingAccessReviewInstanceContactedReviewerID(input string) (*UserPendingAccessReviewInstanceContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPendingAccessReviewInstanceContactedReviewerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPendingAccessReviewInstanceContactedReviewerId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewReviewerId, ok = parsed.Parsed["accessReviewReviewerId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewReviewerId", *parsed)
	}

	return &id, nil
}

// ParseUserPendingAccessReviewInstanceContactedReviewerIDInsensitively parses 'input' case-insensitively into a UserPendingAccessReviewInstanceContactedReviewerId
// note: this method should only be used for API response data and not user input
func ParseUserPendingAccessReviewInstanceContactedReviewerIDInsensitively(input string) (*UserPendingAccessReviewInstanceContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPendingAccessReviewInstanceContactedReviewerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPendingAccessReviewInstanceContactedReviewerId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewReviewerId, ok = parsed.Parsed["accessReviewReviewerId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewReviewerId", *parsed)
	}

	return &id, nil
}

// ValidateUserPendingAccessReviewInstanceContactedReviewerID checks that 'input' can be parsed as a User Pending Access Review Instance Contacted Reviewer ID
func ValidateUserPendingAccessReviewInstanceContactedReviewerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserPendingAccessReviewInstanceContactedReviewerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Pending Access Review Instance Contacted Reviewer ID
func (id UserPendingAccessReviewInstanceContactedReviewerId) ID() string {
	fmtString := "/users/%s/pendingAccessReviewInstances/%s/contactedReviewers/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AccessReviewInstanceId, id.AccessReviewReviewerId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Pending Access Review Instance Contacted Reviewer ID
func (id UserPendingAccessReviewInstanceContactedReviewerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticPendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceIdValue"),
		resourceids.StaticSegment("staticContactedReviewers", "contactedReviewers", "contactedReviewers"),
		resourceids.UserSpecifiedSegment("accessReviewReviewerId", "accessReviewReviewerIdValue"),
	}
}

// String returns a human-readable description of this User Pending Access Review Instance Contacted Reviewer ID
func (id UserPendingAccessReviewInstanceContactedReviewerId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Reviewer: %q", id.AccessReviewReviewerId),
	}
	return fmt.Sprintf("User Pending Access Review Instance Contacted Reviewer (%s)", strings.Join(components, "\n"))
}
