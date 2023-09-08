package domainverificationdnsrecord

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DomainVerificationDnsRecordId{}

// DomainVerificationDnsRecordId is a struct representing the Resource ID for a Domain Verification Dns Record
type DomainVerificationDnsRecordId struct {
	DomainId          string
	DomainDnsRecordId string
}

// NewDomainVerificationDnsRecordID returns a new DomainVerificationDnsRecordId struct
func NewDomainVerificationDnsRecordID(domainId string, domainDnsRecordId string) DomainVerificationDnsRecordId {
	return DomainVerificationDnsRecordId{
		DomainId:          domainId,
		DomainDnsRecordId: domainDnsRecordId,
	}
}

// ParseDomainVerificationDnsRecordID parses 'input' into a DomainVerificationDnsRecordId
func ParseDomainVerificationDnsRecordID(input string) (*DomainVerificationDnsRecordId, error) {
	parser := resourceids.NewParserFromResourceIdType(DomainVerificationDnsRecordId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DomainVerificationDnsRecordId{}

	if id.DomainId, ok = parsed.Parsed["domainId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "domainId", *parsed)
	}

	if id.DomainDnsRecordId, ok = parsed.Parsed["domainDnsRecordId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "domainDnsRecordId", *parsed)
	}

	return &id, nil
}

// ParseDomainVerificationDnsRecordIDInsensitively parses 'input' case-insensitively into a DomainVerificationDnsRecordId
// note: this method should only be used for API response data and not user input
func ParseDomainVerificationDnsRecordIDInsensitively(input string) (*DomainVerificationDnsRecordId, error) {
	parser := resourceids.NewParserFromResourceIdType(DomainVerificationDnsRecordId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DomainVerificationDnsRecordId{}

	if id.DomainId, ok = parsed.Parsed["domainId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "domainId", *parsed)
	}

	if id.DomainDnsRecordId, ok = parsed.Parsed["domainDnsRecordId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "domainDnsRecordId", *parsed)
	}

	return &id, nil
}

// ValidateDomainVerificationDnsRecordID checks that 'input' can be parsed as a Domain Verification Dns Record ID
func ValidateDomainVerificationDnsRecordID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDomainVerificationDnsRecordID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Domain Verification Dns Record ID
func (id DomainVerificationDnsRecordId) ID() string {
	fmtString := "/domains/%s/verificationDnsRecords/%s"
	return fmt.Sprintf(fmtString, id.DomainId, id.DomainDnsRecordId)
}

// Segments returns a slice of Resource ID Segments which comprise this Domain Verification Dns Record ID
func (id DomainVerificationDnsRecordId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticDomains", "domains", "domains"),
		resourceids.UserSpecifiedSegment("domainId", "domainIdValue"),
		resourceids.StaticSegment("staticVerificationDnsRecords", "verificationDnsRecords", "verificationDnsRecords"),
		resourceids.UserSpecifiedSegment("domainDnsRecordId", "domainDnsRecordIdValue"),
	}
}

// String returns a human-readable description of this Domain Verification Dns Record ID
func (id DomainVerificationDnsRecordId) String() string {
	components := []string{
		fmt.Sprintf("Domain: %q", id.DomainId),
		fmt.Sprintf("Domain Dns Record: %q", id.DomainDnsRecordId),
	}
	return fmt.Sprintf("Domain Verification Dns Record (%s)", strings.Join(components, "\n"))
}
