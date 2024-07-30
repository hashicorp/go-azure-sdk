package privilegedaccesgroupassignmentschedulerequestprincipal

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestId{}

// IdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestId is a struct representing the Resource ID for a Identity Governance Privileged Acces Group Assignment Schedule Request
type IdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestId struct {
	PrivilegedAccessGroupAssignmentScheduleRequestId string
}

// NewIdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestID returns a new IdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestId struct
func NewIdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestID(privilegedAccessGroupAssignmentScheduleRequestId string) IdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestId {
	return IdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestId{
		PrivilegedAccessGroupAssignmentScheduleRequestId: privilegedAccessGroupAssignmentScheduleRequestId,
	}
}

// ParseIdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestID parses 'input' into a IdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestId
func ParseIdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestID(input string) (*IdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestIDInsensitively parses 'input' case-insensitively into a IdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestIDInsensitively(input string) (*IdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PrivilegedAccessGroupAssignmentScheduleRequestId, ok = input.Parsed["privilegedAccessGroupAssignmentScheduleRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "privilegedAccessGroupAssignmentScheduleRequestId", input)
	}

	return nil
}

// ValidateIdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestID checks that 'input' can be parsed as a Identity Governance Privileged Acces Group Assignment Schedule Request ID
func ValidateIdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Privileged Acces Group Assignment Schedule Request ID
func (id IdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestId) ID() string {
	fmtString := "/identityGovernance/privilegedAccess/group/assignmentScheduleRequests/%s"
	return fmt.Sprintf(fmtString, id.PrivilegedAccessGroupAssignmentScheduleRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Privileged Acces Group Assignment Schedule Request ID
func (id IdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("privilegedAccess", "privilegedAccess", "privilegedAccess"),
		resourceids.StaticSegment("group", "group", "group"),
		resourceids.StaticSegment("assignmentScheduleRequests", "assignmentScheduleRequests", "assignmentScheduleRequests"),
		resourceids.UserSpecifiedSegment("privilegedAccessGroupAssignmentScheduleRequestId", "privilegedAccessGroupAssignmentScheduleRequestIdValue"),
	}
}

// String returns a human-readable description of this Identity Governance Privileged Acces Group Assignment Schedule Request ID
func (id IdentityGovernancePrivilegedAccesGroupAssignmentScheduleRequestId) String() string {
	components := []string{
		fmt.Sprintf("Privileged Access Group Assignment Schedule Request: %q", id.PrivilegedAccessGroupAssignmentScheduleRequestId),
	}
	return fmt.Sprintf("Identity Governance Privileged Acces Group Assignment Schedule Request (%s)", strings.Join(components, "\n"))
}
