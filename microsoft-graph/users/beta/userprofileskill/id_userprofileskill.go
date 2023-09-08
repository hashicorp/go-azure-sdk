package userprofileskill

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfileSkillId{}

// UserProfileSkillId is a struct representing the Resource ID for a User Profile Skill
type UserProfileSkillId struct {
	UserId             string
	SkillProficiencyId string
}

// NewUserProfileSkillID returns a new UserProfileSkillId struct
func NewUserProfileSkillID(userId string, skillProficiencyId string) UserProfileSkillId {
	return UserProfileSkillId{
		UserId:             userId,
		SkillProficiencyId: skillProficiencyId,
	}
}

// ParseUserProfileSkillID parses 'input' into a UserProfileSkillId
func ParseUserProfileSkillID(input string) (*UserProfileSkillId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileSkillId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileSkillId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.SkillProficiencyId, ok = parsed.Parsed["skillProficiencyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "skillProficiencyId", *parsed)
	}

	return &id, nil
}

// ParseUserProfileSkillIDInsensitively parses 'input' case-insensitively into a UserProfileSkillId
// note: this method should only be used for API response data and not user input
func ParseUserProfileSkillIDInsensitively(input string) (*UserProfileSkillId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileSkillId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileSkillId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.SkillProficiencyId, ok = parsed.Parsed["skillProficiencyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "skillProficiencyId", *parsed)
	}

	return &id, nil
}

// ValidateUserProfileSkillID checks that 'input' can be parsed as a User Profile Skill ID
func ValidateUserProfileSkillID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserProfileSkillID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Profile Skill ID
func (id UserProfileSkillId) ID() string {
	fmtString := "/users/%s/profile/skills/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.SkillProficiencyId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Profile Skill ID
func (id UserProfileSkillId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticProfile", "profile", "profile"),
		resourceids.StaticSegment("staticSkills", "skills", "skills"),
		resourceids.UserSpecifiedSegment("skillProficiencyId", "skillProficiencyIdValue"),
	}
}

// String returns a human-readable description of this User Profile Skill ID
func (id UserProfileSkillId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Skill Proficiency: %q", id.SkillProficiencyId),
	}
	return fmt.Sprintf("User Profile Skill (%s)", strings.Join(components, "\n"))
}
