package userprofilenote

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfileNoteId{}

// UserProfileNoteId is a struct representing the Resource ID for a User Profile Note
type UserProfileNoteId struct {
	UserId             string
	PersonAnnotationId string
}

// NewUserProfileNoteID returns a new UserProfileNoteId struct
func NewUserProfileNoteID(userId string, personAnnotationId string) UserProfileNoteId {
	return UserProfileNoteId{
		UserId:             userId,
		PersonAnnotationId: personAnnotationId,
	}
}

// ParseUserProfileNoteID parses 'input' into a UserProfileNoteId
func ParseUserProfileNoteID(input string) (*UserProfileNoteId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileNoteId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileNoteId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PersonAnnotationId, ok = parsed.Parsed["personAnnotationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "personAnnotationId", *parsed)
	}

	return &id, nil
}

// ParseUserProfileNoteIDInsensitively parses 'input' case-insensitively into a UserProfileNoteId
// note: this method should only be used for API response data and not user input
func ParseUserProfileNoteIDInsensitively(input string) (*UserProfileNoteId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileNoteId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileNoteId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PersonAnnotationId, ok = parsed.Parsed["personAnnotationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "personAnnotationId", *parsed)
	}

	return &id, nil
}

// ValidateUserProfileNoteID checks that 'input' can be parsed as a User Profile Note ID
func ValidateUserProfileNoteID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserProfileNoteID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Profile Note ID
func (id UserProfileNoteId) ID() string {
	fmtString := "/users/%s/profile/notes/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PersonAnnotationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Profile Note ID
func (id UserProfileNoteId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticProfile", "profile", "profile"),
		resourceids.StaticSegment("staticNotes", "notes", "notes"),
		resourceids.UserSpecifiedSegment("personAnnotationId", "personAnnotationIdValue"),
	}
}

// String returns a human-readable description of this User Profile Note ID
func (id UserProfileNoteId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Person Annotation: %q", id.PersonAnnotationId),
	}
	return fmt.Sprintf("User Profile Note (%s)", strings.Join(components, "\n"))
}
