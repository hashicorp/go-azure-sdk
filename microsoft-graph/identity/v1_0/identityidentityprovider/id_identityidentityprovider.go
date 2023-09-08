package identityidentityprovider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = IdentityIdentityProviderId{}

// IdentityIdentityProviderId is a struct representing the Resource ID for a Identity Identity Provider
type IdentityIdentityProviderId struct {
	IdentityProviderBaseId string
}

// NewIdentityIdentityProviderID returns a new IdentityIdentityProviderId struct
func NewIdentityIdentityProviderID(identityProviderBaseId string) IdentityIdentityProviderId {
	return IdentityIdentityProviderId{
		IdentityProviderBaseId: identityProviderBaseId,
	}
}

// ParseIdentityIdentityProviderID parses 'input' into a IdentityIdentityProviderId
func ParseIdentityIdentityProviderID(input string) (*IdentityIdentityProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(IdentityIdentityProviderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IdentityIdentityProviderId{}

	if id.IdentityProviderBaseId, ok = parsed.Parsed["identityProviderBaseId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "identityProviderBaseId", *parsed)
	}

	return &id, nil
}

// ParseIdentityIdentityProviderIDInsensitively parses 'input' case-insensitively into a IdentityIdentityProviderId
// note: this method should only be used for API response data and not user input
func ParseIdentityIdentityProviderIDInsensitively(input string) (*IdentityIdentityProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(IdentityIdentityProviderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IdentityIdentityProviderId{}

	if id.IdentityProviderBaseId, ok = parsed.Parsed["identityProviderBaseId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "identityProviderBaseId", *parsed)
	}

	return &id, nil
}

// ValidateIdentityIdentityProviderID checks that 'input' can be parsed as a Identity Identity Provider ID
func ValidateIdentityIdentityProviderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityIdentityProviderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Identity Provider ID
func (id IdentityIdentityProviderId) ID() string {
	fmtString := "/identity/identityProviders/%s"
	return fmt.Sprintf(fmtString, id.IdentityProviderBaseId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Identity Provider ID
func (id IdentityIdentityProviderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticIdentity", "identity", "identity"),
		resourceids.StaticSegment("staticIdentityProviders", "identityProviders", "identityProviders"),
		resourceids.UserSpecifiedSegment("identityProviderBaseId", "identityProviderBaseIdValue"),
	}
}

// String returns a human-readable description of this Identity Identity Provider ID
func (id IdentityIdentityProviderId) String() string {
	components := []string{
		fmt.Sprintf("Identity Provider Base: %q", id.IdentityProviderBaseId),
	}
	return fmt.Sprintf("Identity Identity Provider (%s)", strings.Join(components, "\n"))
}
