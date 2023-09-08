package userprofilepublication

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfilePublicationId{}

// UserProfilePublicationId is a struct representing the Resource ID for a User Profile Publication
type UserProfilePublicationId struct {
	UserId            string
	ItemPublicationId string
}

// NewUserProfilePublicationID returns a new UserProfilePublicationId struct
func NewUserProfilePublicationID(userId string, itemPublicationId string) UserProfilePublicationId {
	return UserProfilePublicationId{
		UserId:            userId,
		ItemPublicationId: itemPublicationId,
	}
}

// ParseUserProfilePublicationID parses 'input' into a UserProfilePublicationId
func ParseUserProfilePublicationID(input string) (*UserProfilePublicationId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfilePublicationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfilePublicationId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ItemPublicationId, ok = parsed.Parsed["itemPublicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "itemPublicationId", *parsed)
	}

	return &id, nil
}

// ParseUserProfilePublicationIDInsensitively parses 'input' case-insensitively into a UserProfilePublicationId
// note: this method should only be used for API response data and not user input
func ParseUserProfilePublicationIDInsensitively(input string) (*UserProfilePublicationId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfilePublicationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfilePublicationId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ItemPublicationId, ok = parsed.Parsed["itemPublicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "itemPublicationId", *parsed)
	}

	return &id, nil
}

// ValidateUserProfilePublicationID checks that 'input' can be parsed as a User Profile Publication ID
func ValidateUserProfilePublicationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserProfilePublicationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Profile Publication ID
func (id UserProfilePublicationId) ID() string {
	fmtString := "/users/%s/profile/publications/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ItemPublicationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Profile Publication ID
func (id UserProfilePublicationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticProfile", "profile", "profile"),
		resourceids.StaticSegment("staticPublications", "publications", "publications"),
		resourceids.UserSpecifiedSegment("itemPublicationId", "itemPublicationIdValue"),
	}
}

// String returns a human-readable description of this User Profile Publication ID
func (id UserProfilePublicationId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Item Publication: %q", id.ItemPublicationId),
	}
	return fmt.Sprintf("User Profile Publication (%s)", strings.Join(components, "\n"))
}
