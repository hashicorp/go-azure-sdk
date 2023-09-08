package usertransitivereport

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserTransitiveReportId{}

// UserTransitiveReportId is a struct representing the Resource ID for a User Transitive Report
type UserTransitiveReportId struct {
	UserId            string
	DirectoryObjectId string
}

// NewUserTransitiveReportID returns a new UserTransitiveReportId struct
func NewUserTransitiveReportID(userId string, directoryObjectId string) UserTransitiveReportId {
	return UserTransitiveReportId{
		UserId:            userId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseUserTransitiveReportID parses 'input' into a UserTransitiveReportId
func ParseUserTransitiveReportID(input string) (*UserTransitiveReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserTransitiveReportId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserTransitiveReportId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseUserTransitiveReportIDInsensitively parses 'input' case-insensitively into a UserTransitiveReportId
// note: this method should only be used for API response data and not user input
func ParseUserTransitiveReportIDInsensitively(input string) (*UserTransitiveReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserTransitiveReportId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserTransitiveReportId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateUserTransitiveReportID checks that 'input' can be parsed as a User Transitive Report ID
func ValidateUserTransitiveReportID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserTransitiveReportID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Transitive Report ID
func (id UserTransitiveReportId) ID() string {
	fmtString := "/users/%s/transitiveReports/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Transitive Report ID
func (id UserTransitiveReportId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticTransitiveReports", "transitiveReports", "transitiveReports"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this User Transitive Report ID
func (id UserTransitiveReportId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("User Transitive Report (%s)", strings.Join(components, "\n"))
}
