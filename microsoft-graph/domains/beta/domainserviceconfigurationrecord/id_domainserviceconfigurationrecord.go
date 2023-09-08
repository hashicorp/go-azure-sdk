package domainserviceconfigurationrecord

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DomainServiceConfigurationRecordId{}

// DomainServiceConfigurationRecordId is a struct representing the Resource ID for a Domain Service Configuration Record
type DomainServiceConfigurationRecordId struct {
	DomainId          string
	DomainDnsRecordId string
}

// NewDomainServiceConfigurationRecordID returns a new DomainServiceConfigurationRecordId struct
func NewDomainServiceConfigurationRecordID(domainId string, domainDnsRecordId string) DomainServiceConfigurationRecordId {
	return DomainServiceConfigurationRecordId{
		DomainId:          domainId,
		DomainDnsRecordId: domainDnsRecordId,
	}
}

// ParseDomainServiceConfigurationRecordID parses 'input' into a DomainServiceConfigurationRecordId
func ParseDomainServiceConfigurationRecordID(input string) (*DomainServiceConfigurationRecordId, error) {
	parser := resourceids.NewParserFromResourceIdType(DomainServiceConfigurationRecordId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DomainServiceConfigurationRecordId{}

	if id.DomainId, ok = parsed.Parsed["domainId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "domainId", *parsed)
	}

	if id.DomainDnsRecordId, ok = parsed.Parsed["domainDnsRecordId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "domainDnsRecordId", *parsed)
	}

	return &id, nil
}

// ParseDomainServiceConfigurationRecordIDInsensitively parses 'input' case-insensitively into a DomainServiceConfigurationRecordId
// note: this method should only be used for API response data and not user input
func ParseDomainServiceConfigurationRecordIDInsensitively(input string) (*DomainServiceConfigurationRecordId, error) {
	parser := resourceids.NewParserFromResourceIdType(DomainServiceConfigurationRecordId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DomainServiceConfigurationRecordId{}

	if id.DomainId, ok = parsed.Parsed["domainId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "domainId", *parsed)
	}

	if id.DomainDnsRecordId, ok = parsed.Parsed["domainDnsRecordId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "domainDnsRecordId", *parsed)
	}

	return &id, nil
}

// ValidateDomainServiceConfigurationRecordID checks that 'input' can be parsed as a Domain Service Configuration Record ID
func ValidateDomainServiceConfigurationRecordID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDomainServiceConfigurationRecordID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Domain Service Configuration Record ID
func (id DomainServiceConfigurationRecordId) ID() string {
	fmtString := "/domains/%s/serviceConfigurationRecords/%s"
	return fmt.Sprintf(fmtString, id.DomainId, id.DomainDnsRecordId)
}

// Segments returns a slice of Resource ID Segments which comprise this Domain Service Configuration Record ID
func (id DomainServiceConfigurationRecordId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticDomains", "domains", "domains"),
		resourceids.UserSpecifiedSegment("domainId", "domainIdValue"),
		resourceids.StaticSegment("staticServiceConfigurationRecords", "serviceConfigurationRecords", "serviceConfigurationRecords"),
		resourceids.UserSpecifiedSegment("domainDnsRecordId", "domainDnsRecordIdValue"),
	}
}

// String returns a human-readable description of this Domain Service Configuration Record ID
func (id DomainServiceConfigurationRecordId) String() string {
	components := []string{
		fmt.Sprintf("Domain: %q", id.DomainId),
		fmt.Sprintf("Domain Dns Record: %q", id.DomainDnsRecordId),
	}
	return fmt.Sprintf("Domain Service Configuration Record (%s)", strings.Join(components, "\n"))
}
