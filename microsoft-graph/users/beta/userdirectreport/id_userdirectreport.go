package userdirectreport

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserDirectReportId{}

// UserDirectReportId is a struct representing the Resource ID for a User Direct Report
type UserDirectReportId struct {
	UserId            string
	DirectoryObjectId string
}

// NewUserDirectReportID returns a new UserDirectReportId struct
func NewUserDirectReportID(userId string, directoryObjectId string) UserDirectReportId {
	return UserDirectReportId{
		UserId:            userId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseUserDirectReportID parses 'input' into a UserDirectReportId
func ParseUserDirectReportID(input string) (*UserDirectReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserDirectReportId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserDirectReportId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseUserDirectReportIDInsensitively parses 'input' case-insensitively into a UserDirectReportId
// note: this method should only be used for API response data and not user input
func ParseUserDirectReportIDInsensitively(input string) (*UserDirectReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserDirectReportId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserDirectReportId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateUserDirectReportID checks that 'input' can be parsed as a User Direct Report ID
func ValidateUserDirectReportID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserDirectReportID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Direct Report ID
func (id UserDirectReportId) ID() string {
	fmtString := "/users/%s/directReports/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Direct Report ID
func (id UserDirectReportId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticDirectReports", "directReports", "directReports"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this User Direct Report ID
func (id UserDirectReportId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("User Direct Report (%s)", strings.Join(components, "\n"))
}
