package userinformationprotectionsensitivitylabelsublabel

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserInformationProtectionSensitivityLabelSublabelId{}

// UserInformationProtectionSensitivityLabelSublabelId is a struct representing the Resource ID for a User Information Protection Sensitivity Label Sublabel
type UserInformationProtectionSensitivityLabelSublabelId struct {
	UserId              string
	SensitivityLabelId  string
	SensitivityLabelId1 string
}

// NewUserInformationProtectionSensitivityLabelSublabelID returns a new UserInformationProtectionSensitivityLabelSublabelId struct
func NewUserInformationProtectionSensitivityLabelSublabelID(userId string, sensitivityLabelId string, sensitivityLabelId1 string) UserInformationProtectionSensitivityLabelSublabelId {
	return UserInformationProtectionSensitivityLabelSublabelId{
		UserId:              userId,
		SensitivityLabelId:  sensitivityLabelId,
		SensitivityLabelId1: sensitivityLabelId1,
	}
}

// ParseUserInformationProtectionSensitivityLabelSublabelID parses 'input' into a UserInformationProtectionSensitivityLabelSublabelId
func ParseUserInformationProtectionSensitivityLabelSublabelID(input string) (*UserInformationProtectionSensitivityLabelSublabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserInformationProtectionSensitivityLabelSublabelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserInformationProtectionSensitivityLabelSublabelId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.SensitivityLabelId, ok = parsed.Parsed["sensitivityLabelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId", *parsed)
	}

	if id.SensitivityLabelId1, ok = parsed.Parsed["sensitivityLabelId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId1", *parsed)
	}

	return &id, nil
}

// ParseUserInformationProtectionSensitivityLabelSublabelIDInsensitively parses 'input' case-insensitively into a UserInformationProtectionSensitivityLabelSublabelId
// note: this method should only be used for API response data and not user input
func ParseUserInformationProtectionSensitivityLabelSublabelIDInsensitively(input string) (*UserInformationProtectionSensitivityLabelSublabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserInformationProtectionSensitivityLabelSublabelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserInformationProtectionSensitivityLabelSublabelId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.SensitivityLabelId, ok = parsed.Parsed["sensitivityLabelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId", *parsed)
	}

	if id.SensitivityLabelId1, ok = parsed.Parsed["sensitivityLabelId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId1", *parsed)
	}

	return &id, nil
}

// ValidateUserInformationProtectionSensitivityLabelSublabelID checks that 'input' can be parsed as a User Information Protection Sensitivity Label Sublabel ID
func ValidateUserInformationProtectionSensitivityLabelSublabelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserInformationProtectionSensitivityLabelSublabelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Information Protection Sensitivity Label Sublabel ID
func (id UserInformationProtectionSensitivityLabelSublabelId) ID() string {
	fmtString := "/users/%s/informationProtection/sensitivityLabels/%s/sublabels/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.SensitivityLabelId, id.SensitivityLabelId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Information Protection Sensitivity Label Sublabel ID
func (id UserInformationProtectionSensitivityLabelSublabelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticInformationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("staticSensitivityLabels", "sensitivityLabels", "sensitivityLabels"),
		resourceids.UserSpecifiedSegment("sensitivityLabelId", "sensitivityLabelIdValue"),
		resourceids.StaticSegment("staticSublabels", "sublabels", "sublabels"),
		resourceids.UserSpecifiedSegment("sensitivityLabelId1", "sensitivityLabelId1Value"),
	}
}

// String returns a human-readable description of this User Information Protection Sensitivity Label Sublabel ID
func (id UserInformationProtectionSensitivityLabelSublabelId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Sensitivity Label: %q", id.SensitivityLabelId),
		fmt.Sprintf("Sensitivity Label Id 1: %q", id.SensitivityLabelId1),
	}
	return fmt.Sprintf("User Information Protection Sensitivity Label Sublabel (%s)", strings.Join(components, "\n"))
}
