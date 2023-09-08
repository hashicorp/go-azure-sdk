package meprofileskill

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeProfileSkillId{}

// MeProfileSkillId is a struct representing the Resource ID for a Me Profile Skill
type MeProfileSkillId struct {
	SkillProficiencyId string
}

// NewMeProfileSkillID returns a new MeProfileSkillId struct
func NewMeProfileSkillID(skillProficiencyId string) MeProfileSkillId {
	return MeProfileSkillId{
		SkillProficiencyId: skillProficiencyId,
	}
}

// ParseMeProfileSkillID parses 'input' into a MeProfileSkillId
func ParseMeProfileSkillID(input string) (*MeProfileSkillId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeProfileSkillId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeProfileSkillId{}

	if id.SkillProficiencyId, ok = parsed.Parsed["skillProficiencyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "skillProficiencyId", *parsed)
	}

	return &id, nil
}

// ParseMeProfileSkillIDInsensitively parses 'input' case-insensitively into a MeProfileSkillId
// note: this method should only be used for API response data and not user input
func ParseMeProfileSkillIDInsensitively(input string) (*MeProfileSkillId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeProfileSkillId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeProfileSkillId{}

	if id.SkillProficiencyId, ok = parsed.Parsed["skillProficiencyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "skillProficiencyId", *parsed)
	}

	return &id, nil
}

// ValidateMeProfileSkillID checks that 'input' can be parsed as a Me Profile Skill ID
func ValidateMeProfileSkillID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeProfileSkillID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Profile Skill ID
func (id MeProfileSkillId) ID() string {
	fmtString := "/me/profile/skills/%s"
	return fmt.Sprintf(fmtString, id.SkillProficiencyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Profile Skill ID
func (id MeProfileSkillId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticProfile", "profile", "profile"),
		resourceids.StaticSegment("staticSkills", "skills", "skills"),
		resourceids.UserSpecifiedSegment("skillProficiencyId", "skillProficiencyIdValue"),
	}
}

// String returns a human-readable description of this Me Profile Skill ID
func (id MeProfileSkillId) String() string {
	components := []string{
		fmt.Sprintf("Skill Proficiency: %q", id.SkillProficiencyId),
	}
	return fmt.Sprintf("Me Profile Skill (%s)", strings.Join(components, "\n"))
}
