package usermanageddeviceassignmentfilterevaluationstatusdetail

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserManagedDeviceAssignmentFilterEvaluationStatusDetailId{}

// UserManagedDeviceAssignmentFilterEvaluationStatusDetailId is a struct representing the Resource ID for a User Managed Device Assignment Filter Evaluation Status Detail
type UserManagedDeviceAssignmentFilterEvaluationStatusDetailId struct {
	UserId                                    string
	ManagedDeviceId                           string
	AssignmentFilterEvaluationStatusDetailsId string
}

// NewUserManagedDeviceAssignmentFilterEvaluationStatusDetailID returns a new UserManagedDeviceAssignmentFilterEvaluationStatusDetailId struct
func NewUserManagedDeviceAssignmentFilterEvaluationStatusDetailID(userId string, managedDeviceId string, assignmentFilterEvaluationStatusDetailsId string) UserManagedDeviceAssignmentFilterEvaluationStatusDetailId {
	return UserManagedDeviceAssignmentFilterEvaluationStatusDetailId{
		UserId:          userId,
		ManagedDeviceId: managedDeviceId,
		AssignmentFilterEvaluationStatusDetailsId: assignmentFilterEvaluationStatusDetailsId,
	}
}

// ParseUserManagedDeviceAssignmentFilterEvaluationStatusDetailID parses 'input' into a UserManagedDeviceAssignmentFilterEvaluationStatusDetailId
func ParseUserManagedDeviceAssignmentFilterEvaluationStatusDetailID(input string) (*UserManagedDeviceAssignmentFilterEvaluationStatusDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserManagedDeviceAssignmentFilterEvaluationStatusDetailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserManagedDeviceAssignmentFilterEvaluationStatusDetailId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.AssignmentFilterEvaluationStatusDetailsId, ok = parsed.Parsed["assignmentFilterEvaluationStatusDetailsId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "assignmentFilterEvaluationStatusDetailsId", *parsed)
	}

	return &id, nil
}

// ParseUserManagedDeviceAssignmentFilterEvaluationStatusDetailIDInsensitively parses 'input' case-insensitively into a UserManagedDeviceAssignmentFilterEvaluationStatusDetailId
// note: this method should only be used for API response data and not user input
func ParseUserManagedDeviceAssignmentFilterEvaluationStatusDetailIDInsensitively(input string) (*UserManagedDeviceAssignmentFilterEvaluationStatusDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserManagedDeviceAssignmentFilterEvaluationStatusDetailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserManagedDeviceAssignmentFilterEvaluationStatusDetailId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.AssignmentFilterEvaluationStatusDetailsId, ok = parsed.Parsed["assignmentFilterEvaluationStatusDetailsId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "assignmentFilterEvaluationStatusDetailsId", *parsed)
	}

	return &id, nil
}

// ValidateUserManagedDeviceAssignmentFilterEvaluationStatusDetailID checks that 'input' can be parsed as a User Managed Device Assignment Filter Evaluation Status Detail ID
func ValidateUserManagedDeviceAssignmentFilterEvaluationStatusDetailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserManagedDeviceAssignmentFilterEvaluationStatusDetailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Managed Device Assignment Filter Evaluation Status Detail ID
func (id UserManagedDeviceAssignmentFilterEvaluationStatusDetailId) ID() string {
	fmtString := "/users/%s/managedDevices/%s/assignmentFilterEvaluationStatusDetails/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ManagedDeviceId, id.AssignmentFilterEvaluationStatusDetailsId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Managed Device Assignment Filter Evaluation Status Detail ID
func (id UserManagedDeviceAssignmentFilterEvaluationStatusDetailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticManagedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceIdValue"),
		resourceids.StaticSegment("staticAssignmentFilterEvaluationStatusDetails", "assignmentFilterEvaluationStatusDetails", "assignmentFilterEvaluationStatusDetails"),
		resourceids.UserSpecifiedSegment("assignmentFilterEvaluationStatusDetailsId", "assignmentFilterEvaluationStatusDetailsIdValue"),
	}
}

// String returns a human-readable description of this User Managed Device Assignment Filter Evaluation Status Detail ID
func (id UserManagedDeviceAssignmentFilterEvaluationStatusDetailId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Assignment Filter Evaluation Status Details: %q", id.AssignmentFilterEvaluationStatusDetailsId),
	}
	return fmt.Sprintf("User Managed Device Assignment Filter Evaluation Status Detail (%s)", strings.Join(components, "\n"))
}
