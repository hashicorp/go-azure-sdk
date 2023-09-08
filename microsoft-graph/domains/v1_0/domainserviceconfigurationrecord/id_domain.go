package domainserviceconfigurationrecord

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DomainId{}

// DomainId is a struct representing the Resource ID for a Domain
type DomainId struct {
	DomainId string
}

// NewDomainID returns a new DomainId struct
func NewDomainID(domainId string) DomainId {
	return DomainId{
		DomainId: domainId,
	}
}

// ParseDomainID parses 'input' into a DomainId
func ParseDomainID(input string) (*DomainId, error) {
	parser := resourceids.NewParserFromResourceIdType(DomainId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DomainId{}

	if id.DomainId, ok = parsed.Parsed["domainId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "domainId", *parsed)
	}

	return &id, nil
}

// ParseDomainIDInsensitively parses 'input' case-insensitively into a DomainId
// note: this method should only be used for API response data and not user input
func ParseDomainIDInsensitively(input string) (*DomainId, error) {
	parser := resourceids.NewParserFromResourceIdType(DomainId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DomainId{}

	if id.DomainId, ok = parsed.Parsed["domainId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "domainId", *parsed)
	}

	return &id, nil
}

// ValidateDomainID checks that 'input' can be parsed as a Domain ID
func ValidateDomainID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDomainID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Domain ID
func (id DomainId) ID() string {
	fmtString := "/domains/%s"
	return fmt.Sprintf(fmtString, id.DomainId)
}

// Segments returns a slice of Resource ID Segments which comprise this Domain ID
func (id DomainId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticDomains", "domains", "domains"),
		resourceids.UserSpecifiedSegment("domainId", "domainIdValue"),
	}
}

// String returns a human-readable description of this Domain ID
func (id DomainId) String() string {
	components := []string{
		fmt.Sprintf("Domain: %q", id.DomainId),
	}
	return fmt.Sprintf("Domain (%s)", strings.Join(components, "\n"))
}
