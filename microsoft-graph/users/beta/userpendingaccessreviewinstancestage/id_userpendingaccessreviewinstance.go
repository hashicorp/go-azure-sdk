package userpendingaccessreviewinstancestage

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPendingAccessReviewInstanceId{}

// UserPendingAccessReviewInstanceId is a struct representing the Resource ID for a User Pending Access Review Instance
type UserPendingAccessReviewInstanceId struct {
	UserId                 string
	AccessReviewInstanceId string
}

// NewUserPendingAccessReviewInstanceID returns a new UserPendingAccessReviewInstanceId struct
func NewUserPendingAccessReviewInstanceID(userId string, accessReviewInstanceId string) UserPendingAccessReviewInstanceId {
	return UserPendingAccessReviewInstanceId{
		UserId:                 userId,
		AccessReviewInstanceId: accessReviewInstanceId,
	}
}

// ParseUserPendingAccessReviewInstanceID parses 'input' into a UserPendingAccessReviewInstanceId
func ParseUserPendingAccessReviewInstanceID(input string) (*UserPendingAccessReviewInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPendingAccessReviewInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPendingAccessReviewInstanceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	return &id, nil
}

// ParseUserPendingAccessReviewInstanceIDInsensitively parses 'input' case-insensitively into a UserPendingAccessReviewInstanceId
// note: this method should only be used for API response data and not user input
func ParseUserPendingAccessReviewInstanceIDInsensitively(input string) (*UserPendingAccessReviewInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPendingAccessReviewInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPendingAccessReviewInstanceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	return &id, nil
}

// ValidateUserPendingAccessReviewInstanceID checks that 'input' can be parsed as a User Pending Access Review Instance ID
func ValidateUserPendingAccessReviewInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserPendingAccessReviewInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Pending Access Review Instance ID
func (id UserPendingAccessReviewInstanceId) ID() string {
	fmtString := "/users/%s/pendingAccessReviewInstances/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AccessReviewInstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Pending Access Review Instance ID
func (id UserPendingAccessReviewInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticPendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceIdValue"),
	}
}

// String returns a human-readable description of this User Pending Access Review Instance ID
func (id UserPendingAccessReviewInstanceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
	}
	return fmt.Sprintf("User Pending Access Review Instance (%s)", strings.Join(components, "\n"))
}
