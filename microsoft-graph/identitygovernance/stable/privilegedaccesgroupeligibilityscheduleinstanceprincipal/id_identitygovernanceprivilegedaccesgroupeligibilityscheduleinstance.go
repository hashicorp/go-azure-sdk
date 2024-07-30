package privilegedaccesgroupeligibilityscheduleinstanceprincipal

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceId{}

// IdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceId is a struct representing the Resource ID for a Identity Governance Privileged Acces Group Eligibility Schedule Instance
type IdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceId struct {
	PrivilegedAccessGroupEligibilityScheduleInstanceId string
}

// NewIdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceID returns a new IdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceId struct
func NewIdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceID(privilegedAccessGroupEligibilityScheduleInstanceId string) IdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceId {
	return IdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceId{
		PrivilegedAccessGroupEligibilityScheduleInstanceId: privilegedAccessGroupEligibilityScheduleInstanceId,
	}
}

// ParseIdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceID parses 'input' into a IdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceId
func ParseIdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceID(input string) (*IdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceIDInsensitively parses 'input' case-insensitively into a IdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceIDInsensitively(input string) (*IdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PrivilegedAccessGroupEligibilityScheduleInstanceId, ok = input.Parsed["privilegedAccessGroupEligibilityScheduleInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "privilegedAccessGroupEligibilityScheduleInstanceId", input)
	}

	return nil
}

// ValidateIdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceID checks that 'input' can be parsed as a Identity Governance Privileged Acces Group Eligibility Schedule Instance ID
func ValidateIdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Privileged Acces Group Eligibility Schedule Instance ID
func (id IdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceId) ID() string {
	fmtString := "/identityGovernance/privilegedAccess/group/eligibilityScheduleInstances/%s"
	return fmt.Sprintf(fmtString, id.PrivilegedAccessGroupEligibilityScheduleInstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Privileged Acces Group Eligibility Schedule Instance ID
func (id IdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("privilegedAccess", "privilegedAccess", "privilegedAccess"),
		resourceids.StaticSegment("group", "group", "group"),
		resourceids.StaticSegment("eligibilityScheduleInstances", "eligibilityScheduleInstances", "eligibilityScheduleInstances"),
		resourceids.UserSpecifiedSegment("privilegedAccessGroupEligibilityScheduleInstanceId", "privilegedAccessGroupEligibilityScheduleInstanceIdValue"),
	}
}

// String returns a human-readable description of this Identity Governance Privileged Acces Group Eligibility Schedule Instance ID
func (id IdentityGovernancePrivilegedAccesGroupEligibilityScheduleInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Privileged Access Group Eligibility Schedule Instance: %q", id.PrivilegedAccessGroupEligibilityScheduleInstanceId),
	}
	return fmt.Sprintf("Identity Governance Privileged Acces Group Eligibility Schedule Instance (%s)", strings.Join(components, "\n"))
}
