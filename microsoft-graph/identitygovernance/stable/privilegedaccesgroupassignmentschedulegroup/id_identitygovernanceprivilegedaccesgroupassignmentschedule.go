package privilegedaccesgroupassignmentschedulegroup

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId{}

// IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId is a struct representing the Resource ID for a Identity Governance Privileged Acces Group Assignment Schedule
type IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId struct {
	PrivilegedAccessGroupAssignmentScheduleId string
}

// NewIdentityGovernancePrivilegedAccesGroupAssignmentScheduleID returns a new IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId struct
func NewIdentityGovernancePrivilegedAccesGroupAssignmentScheduleID(privilegedAccessGroupAssignmentScheduleId string) IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId {
	return IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId{
		PrivilegedAccessGroupAssignmentScheduleId: privilegedAccessGroupAssignmentScheduleId,
	}
}

// ParseIdentityGovernancePrivilegedAccesGroupAssignmentScheduleID parses 'input' into a IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId
func ParseIdentityGovernancePrivilegedAccesGroupAssignmentScheduleID(input string) (*IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernancePrivilegedAccesGroupAssignmentScheduleIDInsensitively parses 'input' case-insensitively into a IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernancePrivilegedAccesGroupAssignmentScheduleIDInsensitively(input string) (*IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PrivilegedAccessGroupAssignmentScheduleId, ok = input.Parsed["privilegedAccessGroupAssignmentScheduleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "privilegedAccessGroupAssignmentScheduleId", input)
	}

	return nil
}

// ValidateIdentityGovernancePrivilegedAccesGroupAssignmentScheduleID checks that 'input' can be parsed as a Identity Governance Privileged Acces Group Assignment Schedule ID
func ValidateIdentityGovernancePrivilegedAccesGroupAssignmentScheduleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernancePrivilegedAccesGroupAssignmentScheduleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Privileged Acces Group Assignment Schedule ID
func (id IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId) ID() string {
	fmtString := "/identityGovernance/privilegedAccess/group/assignmentSchedules/%s"
	return fmt.Sprintf(fmtString, id.PrivilegedAccessGroupAssignmentScheduleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Privileged Acces Group Assignment Schedule ID
func (id IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("privilegedAccess", "privilegedAccess", "privilegedAccess"),
		resourceids.StaticSegment("group", "group", "group"),
		resourceids.StaticSegment("assignmentSchedules", "assignmentSchedules", "assignmentSchedules"),
		resourceids.UserSpecifiedSegment("privilegedAccessGroupAssignmentScheduleId", "privilegedAccessGroupAssignmentScheduleIdValue"),
	}
}

// String returns a human-readable description of this Identity Governance Privileged Acces Group Assignment Schedule ID
func (id IdentityGovernancePrivilegedAccesGroupAssignmentScheduleId) String() string {
	components := []string{
		fmt.Sprintf("Privileged Access Group Assignment Schedule: %q", id.PrivilegedAccessGroupAssignmentScheduleId),
	}
	return fmt.Sprintf("Identity Governance Privileged Acces Group Assignment Schedule (%s)", strings.Join(components, "\n"))
}
