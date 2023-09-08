package userinformationprotectionsensitivitylabelsublabel

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserInformationProtectionSensitivityLabelId{}

// UserInformationProtectionSensitivityLabelId is a struct representing the Resource ID for a User Information Protection Sensitivity Label
type UserInformationProtectionSensitivityLabelId struct {
	UserId             string
	SensitivityLabelId string
}

// NewUserInformationProtectionSensitivityLabelID returns a new UserInformationProtectionSensitivityLabelId struct
func NewUserInformationProtectionSensitivityLabelID(userId string, sensitivityLabelId string) UserInformationProtectionSensitivityLabelId {
	return UserInformationProtectionSensitivityLabelId{
		UserId:             userId,
		SensitivityLabelId: sensitivityLabelId,
	}
}

// ParseUserInformationProtectionSensitivityLabelID parses 'input' into a UserInformationProtectionSensitivityLabelId
func ParseUserInformationProtectionSensitivityLabelID(input string) (*UserInformationProtectionSensitivityLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserInformationProtectionSensitivityLabelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserInformationProtectionSensitivityLabelId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.SensitivityLabelId, ok = parsed.Parsed["sensitivityLabelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId", *parsed)
	}

	return &id, nil
}

// ParseUserInformationProtectionSensitivityLabelIDInsensitively parses 'input' case-insensitively into a UserInformationProtectionSensitivityLabelId
// note: this method should only be used for API response data and not user input
func ParseUserInformationProtectionSensitivityLabelIDInsensitively(input string) (*UserInformationProtectionSensitivityLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserInformationProtectionSensitivityLabelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserInformationProtectionSensitivityLabelId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.SensitivityLabelId, ok = parsed.Parsed["sensitivityLabelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId", *parsed)
	}

	return &id, nil
}

// ValidateUserInformationProtectionSensitivityLabelID checks that 'input' can be parsed as a User Information Protection Sensitivity Label ID
func ValidateUserInformationProtectionSensitivityLabelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserInformationProtectionSensitivityLabelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Information Protection Sensitivity Label ID
func (id UserInformationProtectionSensitivityLabelId) ID() string {
	fmtString := "/users/%s/informationProtection/sensitivityLabels/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.SensitivityLabelId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Information Protection Sensitivity Label ID
func (id UserInformationProtectionSensitivityLabelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticInformationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("staticSensitivityLabels", "sensitivityLabels", "sensitivityLabels"),
		resourceids.UserSpecifiedSegment("sensitivityLabelId", "sensitivityLabelIdValue"),
	}
}

// String returns a human-readable description of this User Information Protection Sensitivity Label ID
func (id UserInformationProtectionSensitivityLabelId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Sensitivity Label: %q", id.SensitivityLabelId),
	}
	return fmt.Sprintf("User Information Protection Sensitivity Label (%s)", strings.Join(components, "\n"))
}
