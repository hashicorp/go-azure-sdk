package userprofilelanguage

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfileLanguageId{}

// UserProfileLanguageId is a struct representing the Resource ID for a User Profile Language
type UserProfileLanguageId struct {
	UserId                string
	LanguageProficiencyId string
}

// NewUserProfileLanguageID returns a new UserProfileLanguageId struct
func NewUserProfileLanguageID(userId string, languageProficiencyId string) UserProfileLanguageId {
	return UserProfileLanguageId{
		UserId:                userId,
		LanguageProficiencyId: languageProficiencyId,
	}
}

// ParseUserProfileLanguageID parses 'input' into a UserProfileLanguageId
func ParseUserProfileLanguageID(input string) (*UserProfileLanguageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileLanguageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileLanguageId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.LanguageProficiencyId, ok = parsed.Parsed["languageProficiencyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "languageProficiencyId", *parsed)
	}

	return &id, nil
}

// ParseUserProfileLanguageIDInsensitively parses 'input' case-insensitively into a UserProfileLanguageId
// note: this method should only be used for API response data and not user input
func ParseUserProfileLanguageIDInsensitively(input string) (*UserProfileLanguageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileLanguageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileLanguageId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.LanguageProficiencyId, ok = parsed.Parsed["languageProficiencyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "languageProficiencyId", *parsed)
	}

	return &id, nil
}

// ValidateUserProfileLanguageID checks that 'input' can be parsed as a User Profile Language ID
func ValidateUserProfileLanguageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserProfileLanguageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Profile Language ID
func (id UserProfileLanguageId) ID() string {
	fmtString := "/users/%s/profile/languages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.LanguageProficiencyId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Profile Language ID
func (id UserProfileLanguageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticProfile", "profile", "profile"),
		resourceids.StaticSegment("staticLanguages", "languages", "languages"),
		resourceids.UserSpecifiedSegment("languageProficiencyId", "languageProficiencyIdValue"),
	}
}

// String returns a human-readable description of this User Profile Language ID
func (id UserProfileLanguageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Language Proficiency: %q", id.LanguageProficiencyId),
	}
	return fmt.Sprintf("User Profile Language (%s)", strings.Join(components, "\n"))
}
