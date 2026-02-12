package iotcentrals

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserId{}

// UserId is a struct representing the Resource ID for a User
type UserId struct {
	BaseURI string
	UserId  string
}

// NewUserID returns a new UserId struct
func NewUserID(baseURI string, userId string) UserId {
	return UserId{
		BaseURI: strings.TrimSuffix(baseURI, "/"),
		UserId:  userId,
	}
}

// ParseUserID parses 'input' into a UserId
func ParseUserID(input string) (*UserId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIDInsensitively parses 'input' case-insensitively into a UserId
// note: this method should only be used for API response data and not user input
func ParseUserIDInsensitively(input string) (*UserId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	return nil
}

// ValidateUserID checks that 'input' can be parsed as a User ID
func ValidateUserID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User ID
func (id UserId) ID() string {
	fmtString := "%s/users/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.UserId)
}

// Path returns the formatted User ID without the BaseURI
func (id UserId) Path() string {
	fmtString := "/users/%s"
	return fmt.Sprintf(fmtString, id.UserId)
}

// PathElements returns the values of User ID Segments without the BaseURI
func (id UserId) PathElements() []any {
	return []any{id.UserId}
}

// Segments returns a slice of Resource ID Segments which comprise this User ID
func (id UserId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
	}
}

// String returns a human-readable description of this User ID
func (id UserId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("User: %q", id.UserId),
	}
	return fmt.Sprintf("User (%s)", strings.Join(components, "\n"))
}
