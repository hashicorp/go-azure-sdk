package useronenoteresource

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnenoteResourceId{}

// UserOnenoteResourceId is a struct representing the Resource ID for a User Onenote Resource
type UserOnenoteResourceId struct {
	UserId            string
	OnenoteResourceId string
}

// NewUserOnenoteResourceID returns a new UserOnenoteResourceId struct
func NewUserOnenoteResourceID(userId string, onenoteResourceId string) UserOnenoteResourceId {
	return UserOnenoteResourceId{
		UserId:            userId,
		OnenoteResourceId: onenoteResourceId,
	}
}

// ParseUserOnenoteResourceID parses 'input' into a UserOnenoteResourceId
func ParseUserOnenoteResourceID(input string) (*UserOnenoteResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteResourceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OnenoteResourceId, ok = parsed.Parsed["onenoteResourceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteResourceId", *parsed)
	}

	return &id, nil
}

// ParseUserOnenoteResourceIDInsensitively parses 'input' case-insensitively into a UserOnenoteResourceId
// note: this method should only be used for API response data and not user input
func ParseUserOnenoteResourceIDInsensitively(input string) (*UserOnenoteResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenoteResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenoteResourceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OnenoteResourceId, ok = parsed.Parsed["onenoteResourceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteResourceId", *parsed)
	}

	return &id, nil
}

// ValidateUserOnenoteResourceID checks that 'input' can be parsed as a User Onenote Resource ID
func ValidateUserOnenoteResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOnenoteResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Onenote Resource ID
func (id UserOnenoteResourceId) ID() string {
	fmtString := "/users/%s/onenote/resources/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OnenoteResourceId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Onenote Resource ID
func (id UserOnenoteResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticResources", "resources", "resources"),
		resourceids.UserSpecifiedSegment("onenoteResourceId", "onenoteResourceIdValue"),
	}
}

// String returns a human-readable description of this User Onenote Resource ID
func (id UserOnenoteResourceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Onenote Resource: %q", id.OnenoteResourceId),
	}
	return fmt.Sprintf("User Onenote Resource (%s)", strings.Join(components, "\n"))
}
