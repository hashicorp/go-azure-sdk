package usercloudpc

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCloudPCId{}

// UserCloudPCId is a struct representing the Resource ID for a User Cloud P C
type UserCloudPCId struct {
	UserId    string
	CloudPCId string
}

// NewUserCloudPCID returns a new UserCloudPCId struct
func NewUserCloudPCID(userId string, cloudPCId string) UserCloudPCId {
	return UserCloudPCId{
		UserId:    userId,
		CloudPCId: cloudPCId,
	}
}

// ParseUserCloudPCID parses 'input' into a UserCloudPCId
func ParseUserCloudPCID(input string) (*UserCloudPCId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCloudPCId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCloudPCId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.CloudPCId, ok = parsed.Parsed["cloudPCId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "cloudPCId", *parsed)
	}

	return &id, nil
}

// ParseUserCloudPCIDInsensitively parses 'input' case-insensitively into a UserCloudPCId
// note: this method should only be used for API response data and not user input
func ParseUserCloudPCIDInsensitively(input string) (*UserCloudPCId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCloudPCId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCloudPCId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.CloudPCId, ok = parsed.Parsed["cloudPCId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "cloudPCId", *parsed)
	}

	return &id, nil
}

// ValidateUserCloudPCID checks that 'input' can be parsed as a User Cloud P C ID
func ValidateUserCloudPCID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCloudPCID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Cloud P C ID
func (id UserCloudPCId) ID() string {
	fmtString := "/users/%s/cloudPCs/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CloudPCId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Cloud P C ID
func (id UserCloudPCId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCloudPCs", "cloudPCs", "cloudPCs"),
		resourceids.UserSpecifiedSegment("cloudPCId", "cloudPCIdValue"),
	}
}

// String returns a human-readable description of this User Cloud P C ID
func (id UserCloudPCId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Cloud P C: %q", id.CloudPCId),
	}
	return fmt.Sprintf("User Cloud P C (%s)", strings.Join(components, "\n"))
}
