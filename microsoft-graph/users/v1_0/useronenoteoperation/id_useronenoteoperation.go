package useronenoteoperation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnenoteOperationId{}

// UserOnenoteOperationId is a struct representing the Resource ID for a User Onenote Operation
type UserOnenoteOperationId struct {
	UserId             string
	OnenoteOperationId string
}

// NewUserOnenoteOperationID returns a new UserOnenoteOperationId struct
func NewUserOnenoteOperationID(userId string, onenoteOperationId string) UserOnenoteOperationId {
	return UserOnenoteOperationId{
		UserId:             userId,
		OnenoteOperationId: onenoteOperationId,
	}
}

// ParseUserOnenoteOperationID parses 'input' into a UserOnenoteOperationId
func ParseUserOnenoteOperationID(input string) (*UserOnenoteOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteOperationId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OnenoteOperationId, ok = parsed.Parsed["onenoteOperationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteOperationId", *parsed)
	}

	return &id, nil
}

// ParseUserOnenoteOperationIDInsensitively parses 'input' case-insensitively into a UserOnenoteOperationId
// note: this method should only be used for API response data and not user input
func ParseUserOnenoteOperationIDInsensitively(input string) (*UserOnenoteOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteOperationId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OnenoteOperationId, ok = parsed.Parsed["onenoteOperationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteOperationId", *parsed)
	}

	return &id, nil
}

// ValidateUserOnenoteOperationID checks that 'input' can be parsed as a User Onenote Operation ID
func ValidateUserOnenoteOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOnenoteOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Onenote Operation ID
func (id UserOnenoteOperationId) ID() string {
	fmtString := "/users/%s/onenote/operations/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OnenoteOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Onenote Operation ID
func (id UserOnenoteOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticOperations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("onenoteOperationId", "onenoteOperationIdValue"),
	}
}

// String returns a human-readable description of this User Onenote Operation ID
func (id UserOnenoteOperationId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Onenote Operation: %q", id.OnenoteOperationId),
	}
	return fmt.Sprintf("User Onenote Operation (%s)", strings.Join(components, "\n"))
}
