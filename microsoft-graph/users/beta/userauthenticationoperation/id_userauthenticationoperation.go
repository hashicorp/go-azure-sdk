package userauthenticationoperation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserAuthenticationOperationId{}

// UserAuthenticationOperationId is a struct representing the Resource ID for a User Authentication Operation
type UserAuthenticationOperationId struct {
	UserId                 string
	LongRunningOperationId string
}

// NewUserAuthenticationOperationID returns a new UserAuthenticationOperationId struct
func NewUserAuthenticationOperationID(userId string, longRunningOperationId string) UserAuthenticationOperationId {
	return UserAuthenticationOperationId{
		UserId:                 userId,
		LongRunningOperationId: longRunningOperationId,
	}
}

// ParseUserAuthenticationOperationID parses 'input' into a UserAuthenticationOperationId
func ParseUserAuthenticationOperationID(input string) (*UserAuthenticationOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAuthenticationOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAuthenticationOperationId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.LongRunningOperationId, ok = parsed.Parsed["longRunningOperationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "longRunningOperationId", *parsed)
	}

	return &id, nil
}

// ParseUserAuthenticationOperationIDInsensitively parses 'input' case-insensitively into a UserAuthenticationOperationId
// note: this method should only be used for API response data and not user input
func ParseUserAuthenticationOperationIDInsensitively(input string) (*UserAuthenticationOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAuthenticationOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAuthenticationOperationId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.LongRunningOperationId, ok = parsed.Parsed["longRunningOperationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "longRunningOperationId", *parsed)
	}

	return &id, nil
}

// ValidateUserAuthenticationOperationID checks that 'input' can be parsed as a User Authentication Operation ID
func ValidateUserAuthenticationOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserAuthenticationOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Authentication Operation ID
func (id UserAuthenticationOperationId) ID() string {
	fmtString := "/users/%s/authentication/operations/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.LongRunningOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Authentication Operation ID
func (id UserAuthenticationOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticAuthentication", "authentication", "authentication"),
		resourceids.StaticSegment("staticOperations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("longRunningOperationId", "longRunningOperationIdValue"),
	}
}

// String returns a human-readable description of this User Authentication Operation ID
func (id UserAuthenticationOperationId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Long Running Operation: %q", id.LongRunningOperationId),
	}
	return fmt.Sprintf("User Authentication Operation (%s)", strings.Join(components, "\n"))
}
