package userfollowedsite

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserFollowedSiteId{}

// UserFollowedSiteId is a struct representing the Resource ID for a User Followed Site
type UserFollowedSiteId struct {
	UserId string
	SiteId string
}

// NewUserFollowedSiteID returns a new UserFollowedSiteId struct
func NewUserFollowedSiteID(userId string, siteId string) UserFollowedSiteId {
	return UserFollowedSiteId{
		UserId: userId,
		SiteId: siteId,
	}
}

// ParseUserFollowedSiteID parses 'input' into a UserFollowedSiteId
func ParseUserFollowedSiteID(input string) (*UserFollowedSiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserFollowedSiteId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserFollowedSiteId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	return &id, nil
}

// ParseUserFollowedSiteIDInsensitively parses 'input' case-insensitively into a UserFollowedSiteId
// note: this method should only be used for API response data and not user input
func ParseUserFollowedSiteIDInsensitively(input string) (*UserFollowedSiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserFollowedSiteId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserFollowedSiteId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.SiteId, ok = parsed.Parsed["siteId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteId", *parsed)
	}

	return &id, nil
}

// ValidateUserFollowedSiteID checks that 'input' can be parsed as a User Followed Site ID
func ValidateUserFollowedSiteID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserFollowedSiteID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Followed Site ID
func (id UserFollowedSiteId) ID() string {
	fmtString := "/users/%s/followedSites/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.SiteId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Followed Site ID
func (id UserFollowedSiteId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticFollowedSites", "followedSites", "followedSites"),
		resourceids.UserSpecifiedSegment("siteId", "siteIdValue"),
	}
}

// String returns a human-readable description of this User Followed Site ID
func (id UserFollowedSiteId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Site: %q", id.SiteId),
	}
	return fmt.Sprintf("User Followed Site (%s)", strings.Join(components, "\n"))
}
