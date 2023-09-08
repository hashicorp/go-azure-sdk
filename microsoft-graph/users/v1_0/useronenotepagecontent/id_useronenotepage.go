package useronenotepagecontent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnenotePageId{}

// UserOnenotePageId is a struct representing the Resource ID for a User Onenote Page
type UserOnenotePageId struct {
	UserId        string
	OnenotePageId string
}

// NewUserOnenotePageID returns a new UserOnenotePageId struct
func NewUserOnenotePageID(userId string, onenotePageId string) UserOnenotePageId {
	return UserOnenotePageId{
		UserId:        userId,
		OnenotePageId: onenotePageId,
	}
}

// ParseUserOnenotePageID parses 'input' into a UserOnenotePageId
func ParseUserOnenotePageID(input string) (*UserOnenotePageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenotePageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenotePageId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OnenotePageId, ok = parsed.Parsed["onenotePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", *parsed)
	}

	return &id, nil
}

// ParseUserOnenotePageIDInsensitively parses 'input' case-insensitively into a UserOnenotePageId
// note: this method should only be used for API response data and not user input
func ParseUserOnenotePageIDInsensitively(input string) (*UserOnenotePageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnenotePageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnenotePageId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OnenotePageId, ok = parsed.Parsed["onenotePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", *parsed)
	}

	return &id, nil
}

// ValidateUserOnenotePageID checks that 'input' can be parsed as a User Onenote Page ID
func ValidateUserOnenotePageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOnenotePageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Onenote Page ID
func (id UserOnenotePageId) ID() string {
	fmtString := "/users/%s/onenote/pages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Onenote Page ID
func (id UserOnenotePageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticPages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageIdValue"),
	}
}

// String returns a human-readable description of this User Onenote Page ID
func (id UserOnenotePageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("User Onenote Page (%s)", strings.Join(components, "\n"))
}
