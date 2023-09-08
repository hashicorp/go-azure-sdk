package userprofilecertification

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfileCertificationId{}

// UserProfileCertificationId is a struct representing the Resource ID for a User Profile Certification
type UserProfileCertificationId struct {
	UserId                string
	PersonCertificationId string
}

// NewUserProfileCertificationID returns a new UserProfileCertificationId struct
func NewUserProfileCertificationID(userId string, personCertificationId string) UserProfileCertificationId {
	return UserProfileCertificationId{
		UserId:                userId,
		PersonCertificationId: personCertificationId,
	}
}

// ParseUserProfileCertificationID parses 'input' into a UserProfileCertificationId
func ParseUserProfileCertificationID(input string) (*UserProfileCertificationId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileCertificationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileCertificationId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PersonCertificationId, ok = parsed.Parsed["personCertificationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "personCertificationId", *parsed)
	}

	return &id, nil
}

// ParseUserProfileCertificationIDInsensitively parses 'input' case-insensitively into a UserProfileCertificationId
// note: this method should only be used for API response data and not user input
func ParseUserProfileCertificationIDInsensitively(input string) (*UserProfileCertificationId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileCertificationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileCertificationId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PersonCertificationId, ok = parsed.Parsed["personCertificationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "personCertificationId", *parsed)
	}

	return &id, nil
}

// ValidateUserProfileCertificationID checks that 'input' can be parsed as a User Profile Certification ID
func ValidateUserProfileCertificationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserProfileCertificationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Profile Certification ID
func (id UserProfileCertificationId) ID() string {
	fmtString := "/users/%s/profile/certifications/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PersonCertificationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Profile Certification ID
func (id UserProfileCertificationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticProfile", "profile", "profile"),
		resourceids.StaticSegment("staticCertifications", "certifications", "certifications"),
		resourceids.UserSpecifiedSegment("personCertificationId", "personCertificationIdValue"),
	}
}

// String returns a human-readable description of this User Profile Certification ID
func (id UserProfileCertificationId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Person Certification: %q", id.PersonCertificationId),
	}
	return fmt.Sprintf("User Profile Certification (%s)", strings.Join(components, "\n"))
}
