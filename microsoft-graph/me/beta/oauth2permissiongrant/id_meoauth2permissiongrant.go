package oauth2permissiongrant

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOauth2PermissionGrantId{}

// MeOauth2PermissionGrantId is a struct representing the Resource ID for a Me Oauth 2 Permission Grant
type MeOauth2PermissionGrantId struct {
	OAuth2PermissionGrantId string
}

// NewMeOauth2PermissionGrantID returns a new MeOauth2PermissionGrantId struct
func NewMeOauth2PermissionGrantID(oAuth2PermissionGrantId string) MeOauth2PermissionGrantId {
	return MeOauth2PermissionGrantId{
		OAuth2PermissionGrantId: oAuth2PermissionGrantId,
	}
}

// ParseMeOauth2PermissionGrantID parses 'input' into a MeOauth2PermissionGrantId
func ParseMeOauth2PermissionGrantID(input string) (*MeOauth2PermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOauth2PermissionGrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOauth2PermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOauth2PermissionGrantIDInsensitively parses 'input' case-insensitively into a MeOauth2PermissionGrantId
// note: this method should only be used for API response data and not user input
func ParseMeOauth2PermissionGrantIDInsensitively(input string) (*MeOauth2PermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOauth2PermissionGrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOauth2PermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOauth2PermissionGrantId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OAuth2PermissionGrantId, ok = input.Parsed["oAuth2PermissionGrantId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "oAuth2PermissionGrantId", input)
	}

	return nil
}

// ValidateMeOauth2PermissionGrantID checks that 'input' can be parsed as a Me Oauth 2 Permission Grant ID
func ValidateMeOauth2PermissionGrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOauth2PermissionGrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Oauth 2 Permission Grant ID
func (id MeOauth2PermissionGrantId) ID() string {
	fmtString := "/me/oauth2PermissionGrants/%s"
	return fmt.Sprintf(fmtString, id.OAuth2PermissionGrantId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Oauth 2 Permission Grant ID
func (id MeOauth2PermissionGrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("oauth2PermissionGrants", "oauth2PermissionGrants", "oauth2PermissionGrants"),
		resourceids.UserSpecifiedSegment("oAuth2PermissionGrantId", "oAuth2PermissionGrantIdValue"),
	}
}

// String returns a human-readable description of this Me Oauth 2 Permission Grant ID
func (id MeOauth2PermissionGrantId) String() string {
	components := []string{
		fmt.Sprintf("O Auth 2 Permission Grant: %q", id.OAuth2PermissionGrantId),
	}
	return fmt.Sprintf("Me Oauth 2 Permission Grant (%s)", strings.Join(components, "\n"))
}
