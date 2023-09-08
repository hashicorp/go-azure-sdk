package domaindomainnamereference

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DomainDomainNameReferenceId{}

// DomainDomainNameReferenceId is a struct representing the Resource ID for a Domain Domain Name Reference
type DomainDomainNameReferenceId struct {
	DomainId          string
	DirectoryObjectId string
}

// NewDomainDomainNameReferenceID returns a new DomainDomainNameReferenceId struct
func NewDomainDomainNameReferenceID(domainId string, directoryObjectId string) DomainDomainNameReferenceId {
	return DomainDomainNameReferenceId{
		DomainId:          domainId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseDomainDomainNameReferenceID parses 'input' into a DomainDomainNameReferenceId
func ParseDomainDomainNameReferenceID(input string) (*DomainDomainNameReferenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(DomainDomainNameReferenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DomainDomainNameReferenceId{}

	if id.DomainId, ok = parsed.Parsed["domainId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "domainId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseDomainDomainNameReferenceIDInsensitively parses 'input' case-insensitively into a DomainDomainNameReferenceId
// note: this method should only be used for API response data and not user input
func ParseDomainDomainNameReferenceIDInsensitively(input string) (*DomainDomainNameReferenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(DomainDomainNameReferenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DomainDomainNameReferenceId{}

	if id.DomainId, ok = parsed.Parsed["domainId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "domainId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateDomainDomainNameReferenceID checks that 'input' can be parsed as a Domain Domain Name Reference ID
func ValidateDomainDomainNameReferenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDomainDomainNameReferenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Domain Domain Name Reference ID
func (id DomainDomainNameReferenceId) ID() string {
	fmtString := "/domains/%s/domainNameReferences/%s"
	return fmt.Sprintf(fmtString, id.DomainId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Domain Domain Name Reference ID
func (id DomainDomainNameReferenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticDomains", "domains", "domains"),
		resourceids.UserSpecifiedSegment("domainId", "domainIdValue"),
		resourceids.StaticSegment("staticDomainNameReferences", "domainNameReferences", "domainNameReferences"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Domain Domain Name Reference ID
func (id DomainDomainNameReferenceId) String() string {
	components := []string{
		fmt.Sprintf("Domain: %q", id.DomainId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Domain Domain Name Reference (%s)", strings.Join(components, "\n"))
}
