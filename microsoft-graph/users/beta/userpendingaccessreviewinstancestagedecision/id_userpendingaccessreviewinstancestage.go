package userpendingaccessreviewinstancestagedecision

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPendingAccessReviewInstanceStageId{}

// UserPendingAccessReviewInstanceStageId is a struct representing the Resource ID for a User Pending Access Review Instance Stage
type UserPendingAccessReviewInstanceStageId struct {
	UserId                 string
	AccessReviewInstanceId string
	AccessReviewStageId    string
}

// NewUserPendingAccessReviewInstanceStageID returns a new UserPendingAccessReviewInstanceStageId struct
func NewUserPendingAccessReviewInstanceStageID(userId string, accessReviewInstanceId string, accessReviewStageId string) UserPendingAccessReviewInstanceStageId {
	return UserPendingAccessReviewInstanceStageId{
		UserId:                 userId,
		AccessReviewInstanceId: accessReviewInstanceId,
		AccessReviewStageId:    accessReviewStageId,
	}
}

// ParseUserPendingAccessReviewInstanceStageID parses 'input' into a UserPendingAccessReviewInstanceStageId
func ParseUserPendingAccessReviewInstanceStageID(input string) (*UserPendingAccessReviewInstanceStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPendingAccessReviewInstanceStageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPendingAccessReviewInstanceStageId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewStageId, ok = parsed.Parsed["accessReviewStageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewStageId", *parsed)
	}

	return &id, nil
}

// ParseUserPendingAccessReviewInstanceStageIDInsensitively parses 'input' case-insensitively into a UserPendingAccessReviewInstanceStageId
// note: this method should only be used for API response data and not user input
func ParseUserPendingAccessReviewInstanceStageIDInsensitively(input string) (*UserPendingAccessReviewInstanceStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPendingAccessReviewInstanceStageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPendingAccessReviewInstanceStageId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.AccessReviewInstanceId, ok = parsed.Parsed["accessReviewInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", *parsed)
	}

	if id.AccessReviewStageId, ok = parsed.Parsed["accessReviewStageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accessReviewStageId", *parsed)
	}

	return &id, nil
}

// ValidateUserPendingAccessReviewInstanceStageID checks that 'input' can be parsed as a User Pending Access Review Instance Stage ID
func ValidateUserPendingAccessReviewInstanceStageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserPendingAccessReviewInstanceStageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Pending Access Review Instance Stage ID
func (id UserPendingAccessReviewInstanceStageId) ID() string {
	fmtString := "/users/%s/pendingAccessReviewInstances/%s/stages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AccessReviewInstanceId, id.AccessReviewStageId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Pending Access Review Instance Stage ID
func (id UserPendingAccessReviewInstanceStageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticPendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceIdValue"),
		resourceids.StaticSegment("staticStages", "stages", "stages"),
		resourceids.UserSpecifiedSegment("accessReviewStageId", "accessReviewStageIdValue"),
	}
}

// String returns a human-readable description of this User Pending Access Review Instance Stage ID
func (id UserPendingAccessReviewInstanceStageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
	}
	return fmt.Sprintf("User Pending Access Review Instance Stage (%s)", strings.Join(components, "\n"))
}
