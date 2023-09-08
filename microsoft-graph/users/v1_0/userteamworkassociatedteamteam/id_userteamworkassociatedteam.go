package userteamworkassociatedteamteam

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserTeamworkAssociatedTeamId{}

// UserTeamworkAssociatedTeamId is a struct representing the Resource ID for a User Teamwork Associated Team
type UserTeamworkAssociatedTeamId struct {
	UserId               string
	AssociatedTeamInfoId string
}

// NewUserTeamworkAssociatedTeamID returns a new UserTeamworkAssociatedTeamId struct
func NewUserTeamworkAssociatedTeamID(userId string, associatedTeamInfoId string) UserTeamworkAssociatedTeamId {
	return UserTeamworkAssociatedTeamId{
		UserId:               userId,
		AssociatedTeamInfoId: associatedTeamInfoId,
	}
}

// ParseUserTeamworkAssociatedTeamID parses 'input' into a UserTeamworkAssociatedTeamId
func ParseUserTeamworkAssociatedTeamID(input string) (*UserTeamworkAssociatedTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserTeamworkAssociatedTeamId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserTeamworkAssociatedTeamId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.AssociatedTeamInfoId, ok = parsed.Parsed["associatedTeamInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "associatedTeamInfoId", *parsed)
	}

	return &id, nil
}

// ParseUserTeamworkAssociatedTeamIDInsensitively parses 'input' case-insensitively into a UserTeamworkAssociatedTeamId
// note: this method should only be used for API response data and not user input
func ParseUserTeamworkAssociatedTeamIDInsensitively(input string) (*UserTeamworkAssociatedTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserTeamworkAssociatedTeamId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserTeamworkAssociatedTeamId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.AssociatedTeamInfoId, ok = parsed.Parsed["associatedTeamInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "associatedTeamInfoId", *parsed)
	}

	return &id, nil
}

// ValidateUserTeamworkAssociatedTeamID checks that 'input' can be parsed as a User Teamwork Associated Team ID
func ValidateUserTeamworkAssociatedTeamID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserTeamworkAssociatedTeamID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Teamwork Associated Team ID
func (id UserTeamworkAssociatedTeamId) ID() string {
	fmtString := "/users/%s/teamwork/associatedTeams/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AssociatedTeamInfoId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Teamwork Associated Team ID
func (id UserTeamworkAssociatedTeamId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticTeamwork", "teamwork", "teamwork"),
		resourceids.StaticSegment("staticAssociatedTeams", "associatedTeams", "associatedTeams"),
		resourceids.UserSpecifiedSegment("associatedTeamInfoId", "associatedTeamInfoIdValue"),
	}
}

// String returns a human-readable description of this User Teamwork Associated Team ID
func (id UserTeamworkAssociatedTeamId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Associated Team Info: %q", id.AssociatedTeamInfoId),
	}
	return fmt.Sprintf("User Teamwork Associated Team (%s)", strings.Join(components, "\n"))
}
