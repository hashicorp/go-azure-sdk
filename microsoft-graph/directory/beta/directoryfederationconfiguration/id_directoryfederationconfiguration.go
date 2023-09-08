package directoryfederationconfiguration

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DirectoryFederationConfigurationId{}

// DirectoryFederationConfigurationId is a struct representing the Resource ID for a Directory Federation Configuration
type DirectoryFederationConfigurationId struct {
	IdentityProviderBaseId string
}

// NewDirectoryFederationConfigurationID returns a new DirectoryFederationConfigurationId struct
func NewDirectoryFederationConfigurationID(identityProviderBaseId string) DirectoryFederationConfigurationId {
	return DirectoryFederationConfigurationId{
		IdentityProviderBaseId: identityProviderBaseId,
	}
}

// ParseDirectoryFederationConfigurationID parses 'input' into a DirectoryFederationConfigurationId
func ParseDirectoryFederationConfigurationID(input string) (*DirectoryFederationConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(DirectoryFederationConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DirectoryFederationConfigurationId{}

	if id.IdentityProviderBaseId, ok = parsed.Parsed["identityProviderBaseId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "identityProviderBaseId", *parsed)
	}

	return &id, nil
}

// ParseDirectoryFederationConfigurationIDInsensitively parses 'input' case-insensitively into a DirectoryFederationConfigurationId
// note: this method should only be used for API response data and not user input
func ParseDirectoryFederationConfigurationIDInsensitively(input string) (*DirectoryFederationConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(DirectoryFederationConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DirectoryFederationConfigurationId{}

	if id.IdentityProviderBaseId, ok = parsed.Parsed["identityProviderBaseId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "identityProviderBaseId", *parsed)
	}

	return &id, nil
}

// ValidateDirectoryFederationConfigurationID checks that 'input' can be parsed as a Directory Federation Configuration ID
func ValidateDirectoryFederationConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryFederationConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Federation Configuration ID
func (id DirectoryFederationConfigurationId) ID() string {
	fmtString := "/directory/federationConfigurations/%s"
	return fmt.Sprintf(fmtString, id.IdentityProviderBaseId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Federation Configuration ID
func (id DirectoryFederationConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticDirectory", "directory", "directory"),
		resourceids.StaticSegment("staticFederationConfigurations", "federationConfigurations", "federationConfigurations"),
		resourceids.UserSpecifiedSegment("identityProviderBaseId", "identityProviderBaseIdValue"),
	}
}

// String returns a human-readable description of this Directory Federation Configuration ID
func (id DirectoryFederationConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Identity Provider Base: %q", id.IdentityProviderBaseId),
	}
	return fmt.Sprintf("Directory Federation Configuration (%s)", strings.Join(components, "\n"))
}
