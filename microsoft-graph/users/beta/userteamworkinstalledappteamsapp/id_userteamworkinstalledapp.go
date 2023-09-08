package userteamworkinstalledappteamsapp

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserTeamworkInstalledAppId{}

// UserTeamworkInstalledAppId is a struct representing the Resource ID for a User Teamwork Installed App
type UserTeamworkInstalledAppId struct {
	UserId                          string
	UserScopeTeamsAppInstallationId string
}

// NewUserTeamworkInstalledAppID returns a new UserTeamworkInstalledAppId struct
func NewUserTeamworkInstalledAppID(userId string, userScopeTeamsAppInstallationId string) UserTeamworkInstalledAppId {
	return UserTeamworkInstalledAppId{
		UserId:                          userId,
		UserScopeTeamsAppInstallationId: userScopeTeamsAppInstallationId,
	}
}

// ParseUserTeamworkInstalledAppID parses 'input' into a UserTeamworkInstalledAppId
func ParseUserTeamworkInstalledAppID(input string) (*UserTeamworkInstalledAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserTeamworkInstalledAppId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserTeamworkInstalledAppId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.UserScopeTeamsAppInstallationId, ok = parsed.Parsed["userScopeTeamsAppInstallationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userScopeTeamsAppInstallationId", *parsed)
	}

	return &id, nil
}

// ParseUserTeamworkInstalledAppIDInsensitively parses 'input' case-insensitively into a UserTeamworkInstalledAppId
// note: this method should only be used for API response data and not user input
func ParseUserTeamworkInstalledAppIDInsensitively(input string) (*UserTeamworkInstalledAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserTeamworkInstalledAppId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserTeamworkInstalledAppId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.UserScopeTeamsAppInstallationId, ok = parsed.Parsed["userScopeTeamsAppInstallationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userScopeTeamsAppInstallationId", *parsed)
	}

	return &id, nil
}

// ValidateUserTeamworkInstalledAppID checks that 'input' can be parsed as a User Teamwork Installed App ID
func ValidateUserTeamworkInstalledAppID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserTeamworkInstalledAppID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Teamwork Installed App ID
func (id UserTeamworkInstalledAppId) ID() string {
	fmtString := "/users/%s/teamwork/installedApps/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.UserScopeTeamsAppInstallationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Teamwork Installed App ID
func (id UserTeamworkInstalledAppId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticTeamwork", "teamwork", "teamwork"),
		resourceids.StaticSegment("staticInstalledApps", "installedApps", "installedApps"),
		resourceids.UserSpecifiedSegment("userScopeTeamsAppInstallationId", "userScopeTeamsAppInstallationIdValue"),
	}
}

// String returns a human-readable description of this User Teamwork Installed App ID
func (id UserTeamworkInstalledAppId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("User Scope Teams App Installation: %q", id.UserScopeTeamsAppInstallationId),
	}
	return fmt.Sprintf("User Teamwork Installed App (%s)", strings.Join(components, "\n"))
}
