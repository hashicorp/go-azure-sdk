package userinformationprotectiondatalosspreventionpolicy

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserInformationProtectionDataLossPreventionPolicyId{}

// UserInformationProtectionDataLossPreventionPolicyId is a struct representing the Resource ID for a User Information Protection Data Loss Prevention Policy
type UserInformationProtectionDataLossPreventionPolicyId struct {
	UserId                     string
	DataLossPreventionPolicyId string
}

// NewUserInformationProtectionDataLossPreventionPolicyID returns a new UserInformationProtectionDataLossPreventionPolicyId struct
func NewUserInformationProtectionDataLossPreventionPolicyID(userId string, dataLossPreventionPolicyId string) UserInformationProtectionDataLossPreventionPolicyId {
	return UserInformationProtectionDataLossPreventionPolicyId{
		UserId:                     userId,
		DataLossPreventionPolicyId: dataLossPreventionPolicyId,
	}
}

// ParseUserInformationProtectionDataLossPreventionPolicyID parses 'input' into a UserInformationProtectionDataLossPreventionPolicyId
func ParseUserInformationProtectionDataLossPreventionPolicyID(input string) (*UserInformationProtectionDataLossPreventionPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserInformationProtectionDataLossPreventionPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserInformationProtectionDataLossPreventionPolicyId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DataLossPreventionPolicyId, ok = parsed.Parsed["dataLossPreventionPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "dataLossPreventionPolicyId", *parsed)
	}

	return &id, nil
}

// ParseUserInformationProtectionDataLossPreventionPolicyIDInsensitively parses 'input' case-insensitively into a UserInformationProtectionDataLossPreventionPolicyId
// note: this method should only be used for API response data and not user input
func ParseUserInformationProtectionDataLossPreventionPolicyIDInsensitively(input string) (*UserInformationProtectionDataLossPreventionPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserInformationProtectionDataLossPreventionPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserInformationProtectionDataLossPreventionPolicyId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DataLossPreventionPolicyId, ok = parsed.Parsed["dataLossPreventionPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "dataLossPreventionPolicyId", *parsed)
	}

	return &id, nil
}

// ValidateUserInformationProtectionDataLossPreventionPolicyID checks that 'input' can be parsed as a User Information Protection Data Loss Prevention Policy ID
func ValidateUserInformationProtectionDataLossPreventionPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserInformationProtectionDataLossPreventionPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Information Protection Data Loss Prevention Policy ID
func (id UserInformationProtectionDataLossPreventionPolicyId) ID() string {
	fmtString := "/users/%s/informationProtection/dataLossPreventionPolicies/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DataLossPreventionPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Information Protection Data Loss Prevention Policy ID
func (id UserInformationProtectionDataLossPreventionPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticInformationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("staticDataLossPreventionPolicies", "dataLossPreventionPolicies", "dataLossPreventionPolicies"),
		resourceids.UserSpecifiedSegment("dataLossPreventionPolicyId", "dataLossPreventionPolicyIdValue"),
	}
}

// String returns a human-readable description of this User Information Protection Data Loss Prevention Policy ID
func (id UserInformationProtectionDataLossPreventionPolicyId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Data Loss Prevention Policy: %q", id.DataLossPreventionPolicyId),
	}
	return fmt.Sprintf("User Information Protection Data Loss Prevention Policy (%s)", strings.Join(components, "\n"))
}
