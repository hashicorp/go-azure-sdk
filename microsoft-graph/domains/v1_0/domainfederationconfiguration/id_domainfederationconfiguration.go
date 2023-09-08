package domainfederationconfiguration

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DomainFederationConfigurationId{}

// DomainFederationConfigurationId is a struct representing the Resource ID for a Domain Federation Configuration
type DomainFederationConfigurationId struct {
	DomainId                   string
	InternalDomainFederationId string
}

// NewDomainFederationConfigurationID returns a new DomainFederationConfigurationId struct
func NewDomainFederationConfigurationID(domainId string, internalDomainFederationId string) DomainFederationConfigurationId {
	return DomainFederationConfigurationId{
		DomainId:                   domainId,
		InternalDomainFederationId: internalDomainFederationId,
	}
}

// ParseDomainFederationConfigurationID parses 'input' into a DomainFederationConfigurationId
func ParseDomainFederationConfigurationID(input string) (*DomainFederationConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(DomainFederationConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DomainFederationConfigurationId{}

	if id.DomainId, ok = parsed.Parsed["domainId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "domainId", *parsed)
	}

	if id.InternalDomainFederationId, ok = parsed.Parsed["internalDomainFederationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "internalDomainFederationId", *parsed)
	}

	return &id, nil
}

// ParseDomainFederationConfigurationIDInsensitively parses 'input' case-insensitively into a DomainFederationConfigurationId
// note: this method should only be used for API response data and not user input
func ParseDomainFederationConfigurationIDInsensitively(input string) (*DomainFederationConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(DomainFederationConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DomainFederationConfigurationId{}

	if id.DomainId, ok = parsed.Parsed["domainId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "domainId", *parsed)
	}

	if id.InternalDomainFederationId, ok = parsed.Parsed["internalDomainFederationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "internalDomainFederationId", *parsed)
	}

	return &id, nil
}

// ValidateDomainFederationConfigurationID checks that 'input' can be parsed as a Domain Federation Configuration ID
func ValidateDomainFederationConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDomainFederationConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Domain Federation Configuration ID
func (id DomainFederationConfigurationId) ID() string {
	fmtString := "/domains/%s/federationConfiguration/%s"
	return fmt.Sprintf(fmtString, id.DomainId, id.InternalDomainFederationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Domain Federation Configuration ID
func (id DomainFederationConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticDomains", "domains", "domains"),
		resourceids.UserSpecifiedSegment("domainId", "domainIdValue"),
		resourceids.StaticSegment("staticFederationConfiguration", "federationConfiguration", "federationConfiguration"),
		resourceids.UserSpecifiedSegment("internalDomainFederationId", "internalDomainFederationIdValue"),
	}
}

// String returns a human-readable description of this Domain Federation Configuration ID
func (id DomainFederationConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Domain: %q", id.DomainId),
		fmt.Sprintf("Internal Domain Federation: %q", id.InternalDomainFederationId),
	}
	return fmt.Sprintf("Domain Federation Configuration (%s)", strings.Join(components, "\n"))
}
