package oauth2permissiongrant

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &Oauth2PermissionGrantId{}

// Oauth2PermissionGrantId is a struct representing the Resource ID for a Oauth 2 Permission Grant
type Oauth2PermissionGrantId struct {
	OAuth2PermissionGrantId string
}

// NewOauth2PermissionGrantID returns a new Oauth2PermissionGrantId struct
func NewOauth2PermissionGrantID(oAuth2PermissionGrantId string) Oauth2PermissionGrantId {
	return Oauth2PermissionGrantId{
		OAuth2PermissionGrantId: oAuth2PermissionGrantId,
	}
}

// ParseOauth2PermissionGrantID parses 'input' into a Oauth2PermissionGrantId
func ParseOauth2PermissionGrantID(input string) (*Oauth2PermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&Oauth2PermissionGrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := Oauth2PermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseOauth2PermissionGrantIDInsensitively parses 'input' case-insensitively into a Oauth2PermissionGrantId
// note: this method should only be used for API response data and not user input
func ParseOauth2PermissionGrantIDInsensitively(input string) (*Oauth2PermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&Oauth2PermissionGrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := Oauth2PermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *Oauth2PermissionGrantId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OAuth2PermissionGrantId, ok = input.Parsed["oAuth2PermissionGrantId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "oAuth2PermissionGrantId", input)
	}

	return nil
}

// ValidateOauth2PermissionGrantID checks that 'input' can be parsed as a Oauth 2 Permission Grant ID
func ValidateOauth2PermissionGrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseOauth2PermissionGrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Oauth 2 Permission Grant ID
func (id Oauth2PermissionGrantId) ID() string {
	fmtString := "/oauth2PermissionGrants/%s"
	return fmt.Sprintf(fmtString, id.OAuth2PermissionGrantId)
}

// Segments returns a slice of Resource ID Segments which comprise this Oauth 2 Permission Grant ID
func (id Oauth2PermissionGrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("oauth2PermissionGrants", "oauth2PermissionGrants", "oauth2PermissionGrants"),
		resourceids.UserSpecifiedSegment("oAuth2PermissionGrantId", "oAuth2PermissionGrantIdValue"),
	}
}

// String returns a human-readable description of this Oauth 2 Permission Grant ID
func (id Oauth2PermissionGrantId) String() string {
	components := []string{
		fmt.Sprintf("O Auth 2 Permission Grant: %q", id.OAuth2PermissionGrantId),
	}
	return fmt.Sprintf("Oauth 2 Permission Grant (%s)", strings.Join(components, "\n"))
}
