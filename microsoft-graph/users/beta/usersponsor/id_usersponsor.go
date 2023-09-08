package usersponsor

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserSponsorId{}

// UserSponsorId is a struct representing the Resource ID for a User Sponsor
type UserSponsorId struct {
	UserId            string
	DirectoryObjectId string
}

// NewUserSponsorID returns a new UserSponsorId struct
func NewUserSponsorID(userId string, directoryObjectId string) UserSponsorId {
	return UserSponsorId{
		UserId:            userId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseUserSponsorID parses 'input' into a UserSponsorId
func ParseUserSponsorID(input string) (*UserSponsorId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserSponsorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserSponsorId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseUserSponsorIDInsensitively parses 'input' case-insensitively into a UserSponsorId
// note: this method should only be used for API response data and not user input
func ParseUserSponsorIDInsensitively(input string) (*UserSponsorId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserSponsorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserSponsorId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateUserSponsorID checks that 'input' can be parsed as a User Sponsor ID
func ValidateUserSponsorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserSponsorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Sponsor ID
func (id UserSponsorId) ID() string {
	fmtString := "/users/%s/sponsors/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Sponsor ID
func (id UserSponsorId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticSponsors", "sponsors", "sponsors"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this User Sponsor ID
func (id UserSponsorId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("User Sponsor (%s)", strings.Join(components, "\n"))
}
