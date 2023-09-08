package userlicensedetail

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserLicenseDetailId{}

// UserLicenseDetailId is a struct representing the Resource ID for a User License Detail
type UserLicenseDetailId struct {
	UserId           string
	LicenseDetailsId string
}

// NewUserLicenseDetailID returns a new UserLicenseDetailId struct
func NewUserLicenseDetailID(userId string, licenseDetailsId string) UserLicenseDetailId {
	return UserLicenseDetailId{
		UserId:           userId,
		LicenseDetailsId: licenseDetailsId,
	}
}

// ParseUserLicenseDetailID parses 'input' into a UserLicenseDetailId
func ParseUserLicenseDetailID(input string) (*UserLicenseDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserLicenseDetailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserLicenseDetailId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.LicenseDetailsId, ok = parsed.Parsed["licenseDetailsId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "licenseDetailsId", *parsed)
	}

	return &id, nil
}

// ParseUserLicenseDetailIDInsensitively parses 'input' case-insensitively into a UserLicenseDetailId
// note: this method should only be used for API response data and not user input
func ParseUserLicenseDetailIDInsensitively(input string) (*UserLicenseDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserLicenseDetailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserLicenseDetailId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.LicenseDetailsId, ok = parsed.Parsed["licenseDetailsId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "licenseDetailsId", *parsed)
	}

	return &id, nil
}

// ValidateUserLicenseDetailID checks that 'input' can be parsed as a User License Detail ID
func ValidateUserLicenseDetailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserLicenseDetailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User License Detail ID
func (id UserLicenseDetailId) ID() string {
	fmtString := "/users/%s/licenseDetails/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.LicenseDetailsId)
}

// Segments returns a slice of Resource ID Segments which comprise this User License Detail ID
func (id UserLicenseDetailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticLicenseDetails", "licenseDetails", "licenseDetails"),
		resourceids.UserSpecifiedSegment("licenseDetailsId", "licenseDetailsIdValue"),
	}
}

// String returns a human-readable description of this User License Detail ID
func (id UserLicenseDetailId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("License Details: %q", id.LicenseDetailsId),
	}
	return fmt.Sprintf("User License Detail (%s)", strings.Join(components, "\n"))
}
