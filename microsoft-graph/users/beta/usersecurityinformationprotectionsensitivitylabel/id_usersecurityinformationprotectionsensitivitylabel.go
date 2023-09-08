package usersecurityinformationprotectionsensitivitylabel

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserSecurityInformationProtectionSensitivityLabelId{}

// UserSecurityInformationProtectionSensitivityLabelId is a struct representing the Resource ID for a User Security Information Protection Sensitivity Label
type UserSecurityInformationProtectionSensitivityLabelId struct {
	UserId             string
	SensitivityLabelId string
}

// NewUserSecurityInformationProtectionSensitivityLabelID returns a new UserSecurityInformationProtectionSensitivityLabelId struct
func NewUserSecurityInformationProtectionSensitivityLabelID(userId string, sensitivityLabelId string) UserSecurityInformationProtectionSensitivityLabelId {
	return UserSecurityInformationProtectionSensitivityLabelId{
		UserId:             userId,
		SensitivityLabelId: sensitivityLabelId,
	}
}

// ParseUserSecurityInformationProtectionSensitivityLabelID parses 'input' into a UserSecurityInformationProtectionSensitivityLabelId
func ParseUserSecurityInformationProtectionSensitivityLabelID(input string) (*UserSecurityInformationProtectionSensitivityLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserSecurityInformationProtectionSensitivityLabelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserSecurityInformationProtectionSensitivityLabelId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.SensitivityLabelId, ok = parsed.Parsed["sensitivityLabelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId", *parsed)
	}

	return &id, nil
}

// ParseUserSecurityInformationProtectionSensitivityLabelIDInsensitively parses 'input' case-insensitively into a UserSecurityInformationProtectionSensitivityLabelId
// note: this method should only be used for API response data and not user input
func ParseUserSecurityInformationProtectionSensitivityLabelIDInsensitively(input string) (*UserSecurityInformationProtectionSensitivityLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserSecurityInformationProtectionSensitivityLabelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserSecurityInformationProtectionSensitivityLabelId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.SensitivityLabelId, ok = parsed.Parsed["sensitivityLabelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId", *parsed)
	}

	return &id, nil
}

// ValidateUserSecurityInformationProtectionSensitivityLabelID checks that 'input' can be parsed as a User Security Information Protection Sensitivity Label ID
func ValidateUserSecurityInformationProtectionSensitivityLabelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserSecurityInformationProtectionSensitivityLabelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Security Information Protection Sensitivity Label ID
func (id UserSecurityInformationProtectionSensitivityLabelId) ID() string {
	fmtString := "/users/%s/security/informationProtection/sensitivityLabels/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.SensitivityLabelId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Security Information Protection Sensitivity Label ID
func (id UserSecurityInformationProtectionSensitivityLabelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticSecurity", "security", "security"),
		resourceids.StaticSegment("staticInformationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("staticSensitivityLabels", "sensitivityLabels", "sensitivityLabels"),
		resourceids.UserSpecifiedSegment("sensitivityLabelId", "sensitivityLabelIdValue"),
	}
}

// String returns a human-readable description of this User Security Information Protection Sensitivity Label ID
func (id UserSecurityInformationProtectionSensitivityLabelId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Sensitivity Label: %q", id.SensitivityLabelId),
	}
	return fmt.Sprintf("User Security Information Protection Sensitivity Label (%s)", strings.Join(components, "\n"))
}
