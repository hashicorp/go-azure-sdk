package oauth2permissiongrant

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOauth2PermissionGrantId{}

// UserIdOauth2PermissionGrantId is a struct representing the Resource ID for a User Id Oauth 2 Permission Grant
type UserIdOauth2PermissionGrantId struct {
	UserId                  string
	OAuth2PermissionGrantId string
}

// NewUserIdOauth2PermissionGrantID returns a new UserIdOauth2PermissionGrantId struct
func NewUserIdOauth2PermissionGrantID(userId string, oAuth2PermissionGrantId string) UserIdOauth2PermissionGrantId {
	return UserIdOauth2PermissionGrantId{
		UserId:                  userId,
		OAuth2PermissionGrantId: oAuth2PermissionGrantId,
	}
}

// ParseUserIdOauth2PermissionGrantID parses 'input' into a UserIdOauth2PermissionGrantId
func ParseUserIdOauth2PermissionGrantID(input string) (*UserIdOauth2PermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOauth2PermissionGrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOauth2PermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOauth2PermissionGrantIDInsensitively parses 'input' case-insensitively into a UserIdOauth2PermissionGrantId
// note: this method should only be used for API response data and not user input
func ParseUserIdOauth2PermissionGrantIDInsensitively(input string) (*UserIdOauth2PermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOauth2PermissionGrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOauth2PermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOauth2PermissionGrantId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.OAuth2PermissionGrantId, ok = input.Parsed["oAuth2PermissionGrantId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "oAuth2PermissionGrantId", input)
	}

	return nil
}

// ValidateUserIdOauth2PermissionGrantID checks that 'input' can be parsed as a User Id Oauth 2 Permission Grant ID
func ValidateUserIdOauth2PermissionGrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOauth2PermissionGrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Oauth 2 Permission Grant ID
func (id UserIdOauth2PermissionGrantId) ID() string {
	fmtString := "/users/%s/oauth2PermissionGrants/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OAuth2PermissionGrantId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Oauth 2 Permission Grant ID
func (id UserIdOauth2PermissionGrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("oauth2PermissionGrants", "oauth2PermissionGrants", "oauth2PermissionGrants"),
		resourceids.UserSpecifiedSegment("oAuth2PermissionGrantId", "oAuth2PermissionGrantIdValue"),
	}
}

// String returns a human-readable description of this User Id Oauth 2 Permission Grant ID
func (id UserIdOauth2PermissionGrantId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("O Auth 2 Permission Grant: %q", id.OAuth2PermissionGrantId),
	}
	return fmt.Sprintf("User Id Oauth 2 Permission Grant (%s)", strings.Join(components, "\n"))
}
