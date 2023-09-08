package memanageddeviceassignmentfilterevaluationstatusdetail

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeManagedDeviceAssignmentFilterEvaluationStatusDetailId{}

// MeManagedDeviceAssignmentFilterEvaluationStatusDetailId is a struct representing the Resource ID for a Me Managed Device Assignment Filter Evaluation Status Detail
type MeManagedDeviceAssignmentFilterEvaluationStatusDetailId struct {
	ManagedDeviceId                           string
	AssignmentFilterEvaluationStatusDetailsId string
}

// NewMeManagedDeviceAssignmentFilterEvaluationStatusDetailID returns a new MeManagedDeviceAssignmentFilterEvaluationStatusDetailId struct
func NewMeManagedDeviceAssignmentFilterEvaluationStatusDetailID(managedDeviceId string, assignmentFilterEvaluationStatusDetailsId string) MeManagedDeviceAssignmentFilterEvaluationStatusDetailId {
	return MeManagedDeviceAssignmentFilterEvaluationStatusDetailId{
		ManagedDeviceId: managedDeviceId,
		AssignmentFilterEvaluationStatusDetailsId: assignmentFilterEvaluationStatusDetailsId,
	}
}

// ParseMeManagedDeviceAssignmentFilterEvaluationStatusDetailID parses 'input' into a MeManagedDeviceAssignmentFilterEvaluationStatusDetailId
func ParseMeManagedDeviceAssignmentFilterEvaluationStatusDetailID(input string) (*MeManagedDeviceAssignmentFilterEvaluationStatusDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeManagedDeviceAssignmentFilterEvaluationStatusDetailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeManagedDeviceAssignmentFilterEvaluationStatusDetailId{}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.AssignmentFilterEvaluationStatusDetailsId, ok = parsed.Parsed["assignmentFilterEvaluationStatusDetailsId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "assignmentFilterEvaluationStatusDetailsId", *parsed)
	}

	return &id, nil
}

// ParseMeManagedDeviceAssignmentFilterEvaluationStatusDetailIDInsensitively parses 'input' case-insensitively into a MeManagedDeviceAssignmentFilterEvaluationStatusDetailId
// note: this method should only be used for API response data and not user input
func ParseMeManagedDeviceAssignmentFilterEvaluationStatusDetailIDInsensitively(input string) (*MeManagedDeviceAssignmentFilterEvaluationStatusDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeManagedDeviceAssignmentFilterEvaluationStatusDetailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeManagedDeviceAssignmentFilterEvaluationStatusDetailId{}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.AssignmentFilterEvaluationStatusDetailsId, ok = parsed.Parsed["assignmentFilterEvaluationStatusDetailsId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "assignmentFilterEvaluationStatusDetailsId", *parsed)
	}

	return &id, nil
}

// ValidateMeManagedDeviceAssignmentFilterEvaluationStatusDetailID checks that 'input' can be parsed as a Me Managed Device Assignment Filter Evaluation Status Detail ID
func ValidateMeManagedDeviceAssignmentFilterEvaluationStatusDetailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeManagedDeviceAssignmentFilterEvaluationStatusDetailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Managed Device Assignment Filter Evaluation Status Detail ID
func (id MeManagedDeviceAssignmentFilterEvaluationStatusDetailId) ID() string {
	fmtString := "/me/managedDevices/%s/assignmentFilterEvaluationStatusDetails/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.AssignmentFilterEvaluationStatusDetailsId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Managed Device Assignment Filter Evaluation Status Detail ID
func (id MeManagedDeviceAssignmentFilterEvaluationStatusDetailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticManagedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceIdValue"),
		resourceids.StaticSegment("staticAssignmentFilterEvaluationStatusDetails", "assignmentFilterEvaluationStatusDetails", "assignmentFilterEvaluationStatusDetails"),
		resourceids.UserSpecifiedSegment("assignmentFilterEvaluationStatusDetailsId", "assignmentFilterEvaluationStatusDetailsIdValue"),
	}
}

// String returns a human-readable description of this Me Managed Device Assignment Filter Evaluation Status Detail ID
func (id MeManagedDeviceAssignmentFilterEvaluationStatusDetailId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Assignment Filter Evaluation Status Details: %q", id.AssignmentFilterEvaluationStatusDetailsId),
	}
	return fmt.Sprintf("Me Managed Device Assignment Filter Evaluation Status Detail (%s)", strings.Join(components, "\n"))
}
