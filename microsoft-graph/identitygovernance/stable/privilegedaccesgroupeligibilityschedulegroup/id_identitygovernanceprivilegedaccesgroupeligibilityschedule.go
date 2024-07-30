package privilegedaccesgroupeligibilityschedulegroup

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccesGroupEligibilityScheduleId{}

// IdentityGovernancePrivilegedAccesGroupEligibilityScheduleId is a struct representing the Resource ID for a Identity Governance Privileged Acces Group Eligibility Schedule
type IdentityGovernancePrivilegedAccesGroupEligibilityScheduleId struct {
	PrivilegedAccessGroupEligibilityScheduleId string
}

// NewIdentityGovernancePrivilegedAccesGroupEligibilityScheduleID returns a new IdentityGovernancePrivilegedAccesGroupEligibilityScheduleId struct
func NewIdentityGovernancePrivilegedAccesGroupEligibilityScheduleID(privilegedAccessGroupEligibilityScheduleId string) IdentityGovernancePrivilegedAccesGroupEligibilityScheduleId {
	return IdentityGovernancePrivilegedAccesGroupEligibilityScheduleId{
		PrivilegedAccessGroupEligibilityScheduleId: privilegedAccessGroupEligibilityScheduleId,
	}
}

// ParseIdentityGovernancePrivilegedAccesGroupEligibilityScheduleID parses 'input' into a IdentityGovernancePrivilegedAccesGroupEligibilityScheduleId
func ParseIdentityGovernancePrivilegedAccesGroupEligibilityScheduleID(input string) (*IdentityGovernancePrivilegedAccesGroupEligibilityScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccesGroupEligibilityScheduleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccesGroupEligibilityScheduleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernancePrivilegedAccesGroupEligibilityScheduleIDInsensitively parses 'input' case-insensitively into a IdentityGovernancePrivilegedAccesGroupEligibilityScheduleId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernancePrivilegedAccesGroupEligibilityScheduleIDInsensitively(input string) (*IdentityGovernancePrivilegedAccesGroupEligibilityScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccesGroupEligibilityScheduleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccesGroupEligibilityScheduleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernancePrivilegedAccesGroupEligibilityScheduleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PrivilegedAccessGroupEligibilityScheduleId, ok = input.Parsed["privilegedAccessGroupEligibilityScheduleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "privilegedAccessGroupEligibilityScheduleId", input)
	}

	return nil
}

// ValidateIdentityGovernancePrivilegedAccesGroupEligibilityScheduleID checks that 'input' can be parsed as a Identity Governance Privileged Acces Group Eligibility Schedule ID
func ValidateIdentityGovernancePrivilegedAccesGroupEligibilityScheduleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernancePrivilegedAccesGroupEligibilityScheduleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Privileged Acces Group Eligibility Schedule ID
func (id IdentityGovernancePrivilegedAccesGroupEligibilityScheduleId) ID() string {
	fmtString := "/identityGovernance/privilegedAccess/group/eligibilitySchedules/%s"
	return fmt.Sprintf(fmtString, id.PrivilegedAccessGroupEligibilityScheduleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Privileged Acces Group Eligibility Schedule ID
func (id IdentityGovernancePrivilegedAccesGroupEligibilityScheduleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("privilegedAccess", "privilegedAccess", "privilegedAccess"),
		resourceids.StaticSegment("group", "group", "group"),
		resourceids.StaticSegment("eligibilitySchedules", "eligibilitySchedules", "eligibilitySchedules"),
		resourceids.UserSpecifiedSegment("privilegedAccessGroupEligibilityScheduleId", "privilegedAccessGroupEligibilityScheduleIdValue"),
	}
}

// String returns a human-readable description of this Identity Governance Privileged Acces Group Eligibility Schedule ID
func (id IdentityGovernancePrivilegedAccesGroupEligibilityScheduleId) String() string {
	components := []string{
		fmt.Sprintf("Privileged Access Group Eligibility Schedule: %q", id.PrivilegedAccessGroupEligibilityScheduleId),
	}
	return fmt.Sprintf("Identity Governance Privileged Acces Group Eligibility Schedule (%s)", strings.Join(components, "\n"))
}
