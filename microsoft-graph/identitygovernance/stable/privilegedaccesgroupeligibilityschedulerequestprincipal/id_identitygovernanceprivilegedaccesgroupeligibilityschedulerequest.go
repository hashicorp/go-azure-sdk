package privilegedaccesgroupeligibilityschedulerequestprincipal

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId{}

// IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId is a struct representing the Resource ID for a Identity Governance Privileged Acces Group Eligibility Schedule Request
type IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId struct {
	PrivilegedAccessGroupEligibilityScheduleRequestId string
}

// NewIdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestID returns a new IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId struct
func NewIdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestID(privilegedAccessGroupEligibilityScheduleRequestId string) IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId {
	return IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId{
		PrivilegedAccessGroupEligibilityScheduleRequestId: privilegedAccessGroupEligibilityScheduleRequestId,
	}
}

// ParseIdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestID parses 'input' into a IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId
func ParseIdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestID(input string) (*IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestIDInsensitively parses 'input' case-insensitively into a IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestIDInsensitively(input string) (*IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PrivilegedAccessGroupEligibilityScheduleRequestId, ok = input.Parsed["privilegedAccessGroupEligibilityScheduleRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "privilegedAccessGroupEligibilityScheduleRequestId", input)
	}

	return nil
}

// ValidateIdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestID checks that 'input' can be parsed as a Identity Governance Privileged Acces Group Eligibility Schedule Request ID
func ValidateIdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Privileged Acces Group Eligibility Schedule Request ID
func (id IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId) ID() string {
	fmtString := "/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests/%s"
	return fmt.Sprintf(fmtString, id.PrivilegedAccessGroupEligibilityScheduleRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Privileged Acces Group Eligibility Schedule Request ID
func (id IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("privilegedAccess", "privilegedAccess", "privilegedAccess"),
		resourceids.StaticSegment("group", "group", "group"),
		resourceids.StaticSegment("eligibilityScheduleRequests", "eligibilityScheduleRequests", "eligibilityScheduleRequests"),
		resourceids.UserSpecifiedSegment("privilegedAccessGroupEligibilityScheduleRequestId", "privilegedAccessGroupEligibilityScheduleRequestIdValue"),
	}
}

// String returns a human-readable description of this Identity Governance Privileged Acces Group Eligibility Schedule Request ID
func (id IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId) String() string {
	components := []string{
		fmt.Sprintf("Privileged Access Group Eligibility Schedule Request: %q", id.PrivilegedAccessGroupEligibilityScheduleRequestId),
	}
	return fmt.Sprintf("Identity Governance Privileged Acces Group Eligibility Schedule Request (%s)", strings.Join(components, "\n"))
}
