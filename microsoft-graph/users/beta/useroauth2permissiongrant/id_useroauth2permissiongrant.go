package useroauth2permissiongrant

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOauth2PermissionGrantId{}

// UserOauth2PermissionGrantId is a struct representing the Resource ID for a User Oauth 2 Permission Grant
type UserOauth2PermissionGrantId struct {
	UserId                  string
	OAuth2PermissionGrantId string
}

// NewUserOauth2PermissionGrantID returns a new UserOauth2PermissionGrantId struct
func NewUserOauth2PermissionGrantID(userId string, oAuth2PermissionGrantId string) UserOauth2PermissionGrantId {
	return UserOauth2PermissionGrantId{
		UserId:                  userId,
		OAuth2PermissionGrantId: oAuth2PermissionGrantId,
	}
}

// ParseUserOauth2PermissionGrantID parses 'input' into a UserOauth2PermissionGrantId
func ParseUserOauth2PermissionGrantID(input string) (*UserOauth2PermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOauth2PermissionGrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOauth2PermissionGrantId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OAuth2PermissionGrantId, ok = parsed.Parsed["oAuth2PermissionGrantId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "oAuth2PermissionGrantId", *parsed)
	}

	return &id, nil
}

// ParseUserOauth2PermissionGrantIDInsensitively parses 'input' case-insensitively into a UserOauth2PermissionGrantId
// note: this method should only be used for API response data and not user input
func ParseUserOauth2PermissionGrantIDInsensitively(input string) (*UserOauth2PermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOauth2PermissionGrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOauth2PermissionGrantId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OAuth2PermissionGrantId, ok = parsed.Parsed["oAuth2PermissionGrantId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "oAuth2PermissionGrantId", *parsed)
	}

	return &id, nil
}

// ValidateUserOauth2PermissionGrantID checks that 'input' can be parsed as a User Oauth 2 Permission Grant ID
func ValidateUserOauth2PermissionGrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOauth2PermissionGrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Oauth 2 Permission Grant ID
func (id UserOauth2PermissionGrantId) ID() string {
	fmtString := "/users/%s/oauth2PermissionGrants/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OAuth2PermissionGrantId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Oauth 2 Permission Grant ID
func (id UserOauth2PermissionGrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOauth2PermissionGrants", "oauth2PermissionGrants", "oauth2PermissionGrants"),
		resourceids.UserSpecifiedSegment("oAuth2PermissionGrantId", "oAuth2PermissionGrantIdValue"),
	}
}

// String returns a human-readable description of this User Oauth 2 Permission Grant ID
func (id UserOauth2PermissionGrantId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("O Auth 2 Permission Grant: %q", id.OAuth2PermissionGrantId),
	}
	return fmt.Sprintf("User Oauth 2 Permission Grant (%s)", strings.Join(components, "\n"))
}
