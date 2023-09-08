package useroutlookmastercategory

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOutlookMasterCategoryId{}

// UserOutlookMasterCategoryId is a struct representing the Resource ID for a User Outlook Master Category
type UserOutlookMasterCategoryId struct {
	UserId            string
	OutlookCategoryId string
}

// NewUserOutlookMasterCategoryID returns a new UserOutlookMasterCategoryId struct
func NewUserOutlookMasterCategoryID(userId string, outlookCategoryId string) UserOutlookMasterCategoryId {
	return UserOutlookMasterCategoryId{
		UserId:            userId,
		OutlookCategoryId: outlookCategoryId,
	}
}

// ParseUserOutlookMasterCategoryID parses 'input' into a UserOutlookMasterCategoryId
func ParseUserOutlookMasterCategoryID(input string) (*UserOutlookMasterCategoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOutlookMasterCategoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOutlookMasterCategoryId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OutlookCategoryId, ok = parsed.Parsed["outlookCategoryId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookCategoryId", *parsed)
	}

	return &id, nil
}

// ParseUserOutlookMasterCategoryIDInsensitively parses 'input' case-insensitively into a UserOutlookMasterCategoryId
// note: this method should only be used for API response data and not user input
func ParseUserOutlookMasterCategoryIDInsensitively(input string) (*UserOutlookMasterCategoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOutlookMasterCategoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOutlookMasterCategoryId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OutlookCategoryId, ok = parsed.Parsed["outlookCategoryId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookCategoryId", *parsed)
	}

	return &id, nil
}

// ValidateUserOutlookMasterCategoryID checks that 'input' can be parsed as a User Outlook Master Category ID
func ValidateUserOutlookMasterCategoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOutlookMasterCategoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Outlook Master Category ID
func (id UserOutlookMasterCategoryId) ID() string {
	fmtString := "/users/%s/outlook/masterCategories/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OutlookCategoryId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Outlook Master Category ID
func (id UserOutlookMasterCategoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOutlook", "outlook", "outlook"),
		resourceids.StaticSegment("staticMasterCategories", "masterCategories", "masterCategories"),
		resourceids.UserSpecifiedSegment("outlookCategoryId", "outlookCategoryIdValue"),
	}
}

// String returns a human-readable description of this User Outlook Master Category ID
func (id UserOutlookMasterCategoryId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Outlook Category: %q", id.OutlookCategoryId),
	}
	return fmt.Sprintf("User Outlook Master Category (%s)", strings.Join(components, "\n"))
}
