package userinformationprotectionpolicylabel

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserInformationProtectionPolicyLabelId{}

// UserInformationProtectionPolicyLabelId is a struct representing the Resource ID for a User Information Protection Policy Label
type UserInformationProtectionPolicyLabelId struct {
	UserId                       string
	InformationProtectionLabelId string
}

// NewUserInformationProtectionPolicyLabelID returns a new UserInformationProtectionPolicyLabelId struct
func NewUserInformationProtectionPolicyLabelID(userId string, informationProtectionLabelId string) UserInformationProtectionPolicyLabelId {
	return UserInformationProtectionPolicyLabelId{
		UserId:                       userId,
		InformationProtectionLabelId: informationProtectionLabelId,
	}
}

// ParseUserInformationProtectionPolicyLabelID parses 'input' into a UserInformationProtectionPolicyLabelId
func ParseUserInformationProtectionPolicyLabelID(input string) (*UserInformationProtectionPolicyLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserInformationProtectionPolicyLabelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserInformationProtectionPolicyLabelId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.InformationProtectionLabelId, ok = parsed.Parsed["informationProtectionLabelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "informationProtectionLabelId", *parsed)
	}

	return &id, nil
}

// ParseUserInformationProtectionPolicyLabelIDInsensitively parses 'input' case-insensitively into a UserInformationProtectionPolicyLabelId
// note: this method should only be used for API response data and not user input
func ParseUserInformationProtectionPolicyLabelIDInsensitively(input string) (*UserInformationProtectionPolicyLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserInformationProtectionPolicyLabelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserInformationProtectionPolicyLabelId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.InformationProtectionLabelId, ok = parsed.Parsed["informationProtectionLabelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "informationProtectionLabelId", *parsed)
	}

	return &id, nil
}

// ValidateUserInformationProtectionPolicyLabelID checks that 'input' can be parsed as a User Information Protection Policy Label ID
func ValidateUserInformationProtectionPolicyLabelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserInformationProtectionPolicyLabelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Information Protection Policy Label ID
func (id UserInformationProtectionPolicyLabelId) ID() string {
	fmtString := "/users/%s/informationProtection/policy/labels/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.InformationProtectionLabelId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Information Protection Policy Label ID
func (id UserInformationProtectionPolicyLabelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticInformationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("staticPolicy", "policy", "policy"),
		resourceids.StaticSegment("staticLabels", "labels", "labels"),
		resourceids.UserSpecifiedSegment("informationProtectionLabelId", "informationProtectionLabelIdValue"),
	}
}

// String returns a human-readable description of this User Information Protection Policy Label ID
func (id UserInformationProtectionPolicyLabelId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Information Protection Label: %q", id.InformationProtectionLabelId),
	}
	return fmt.Sprintf("User Information Protection Policy Label (%s)", strings.Join(components, "\n"))
}
