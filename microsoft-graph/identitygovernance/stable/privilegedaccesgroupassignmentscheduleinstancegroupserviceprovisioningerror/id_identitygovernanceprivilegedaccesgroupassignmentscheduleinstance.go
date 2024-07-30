package privilegedaccesgroupassignmentscheduleinstancegroupserviceprovisioningerror

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId{}

// IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId is a struct representing the Resource ID for a Identity Governance Privileged Acces Group Assignment Schedule Instance
type IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId struct {
	PrivilegedAccessGroupAssignmentScheduleInstanceId string
}

// NewIdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceID returns a new IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId struct
func NewIdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceID(privilegedAccessGroupAssignmentScheduleInstanceId string) IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId {
	return IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId{
		PrivilegedAccessGroupAssignmentScheduleInstanceId: privilegedAccessGroupAssignmentScheduleInstanceId,
	}
}

// ParseIdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceID parses 'input' into a IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId
func ParseIdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceID(input string) (*IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceIDInsensitively parses 'input' case-insensitively into a IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceIDInsensitively(input string) (*IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PrivilegedAccessGroupAssignmentScheduleInstanceId, ok = input.Parsed["privilegedAccessGroupAssignmentScheduleInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "privilegedAccessGroupAssignmentScheduleInstanceId", input)
	}

	return nil
}

// ValidateIdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceID checks that 'input' can be parsed as a Identity Governance Privileged Acces Group Assignment Schedule Instance ID
func ValidateIdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Privileged Acces Group Assignment Schedule Instance ID
func (id IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId) ID() string {
	fmtString := "/identityGovernance/privilegedAccess/group/assignmentScheduleInstances/%s"
	return fmt.Sprintf(fmtString, id.PrivilegedAccessGroupAssignmentScheduleInstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Privileged Acces Group Assignment Schedule Instance ID
func (id IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("privilegedAccess", "privilegedAccess", "privilegedAccess"),
		resourceids.StaticSegment("group", "group", "group"),
		resourceids.StaticSegment("assignmentScheduleInstances", "assignmentScheduleInstances", "assignmentScheduleInstances"),
		resourceids.UserSpecifiedSegment("privilegedAccessGroupAssignmentScheduleInstanceId", "privilegedAccessGroupAssignmentScheduleInstanceIdValue"),
	}
}

// String returns a human-readable description of this Identity Governance Privileged Acces Group Assignment Schedule Instance ID
func (id IdentityGovernancePrivilegedAccesGroupAssignmentScheduleInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Privileged Access Group Assignment Schedule Instance: %q", id.PrivilegedAccessGroupAssignmentScheduleInstanceId),
	}
	return fmt.Sprintf("Identity Governance Privileged Acces Group Assignment Schedule Instance (%s)", strings.Join(components, "\n"))
}
