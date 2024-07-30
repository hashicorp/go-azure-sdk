package conditionalaccesauthenticationcontextclassreference

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityConditionalAccesAuthenticationContextClassReferenceId{}

// IdentityConditionalAccesAuthenticationContextClassReferenceId is a struct representing the Resource ID for a Identity Conditional Acces Authentication Context Class Reference
type IdentityConditionalAccesAuthenticationContextClassReferenceId struct {
	AuthenticationContextClassReferenceId string
}

// NewIdentityConditionalAccesAuthenticationContextClassReferenceID returns a new IdentityConditionalAccesAuthenticationContextClassReferenceId struct
func NewIdentityConditionalAccesAuthenticationContextClassReferenceID(authenticationContextClassReferenceId string) IdentityConditionalAccesAuthenticationContextClassReferenceId {
	return IdentityConditionalAccesAuthenticationContextClassReferenceId{
		AuthenticationContextClassReferenceId: authenticationContextClassReferenceId,
	}
}

// ParseIdentityConditionalAccesAuthenticationContextClassReferenceID parses 'input' into a IdentityConditionalAccesAuthenticationContextClassReferenceId
func ParseIdentityConditionalAccesAuthenticationContextClassReferenceID(input string) (*IdentityConditionalAccesAuthenticationContextClassReferenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityConditionalAccesAuthenticationContextClassReferenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityConditionalAccesAuthenticationContextClassReferenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityConditionalAccesAuthenticationContextClassReferenceIDInsensitively parses 'input' case-insensitively into a IdentityConditionalAccesAuthenticationContextClassReferenceId
// note: this method should only be used for API response data and not user input
func ParseIdentityConditionalAccesAuthenticationContextClassReferenceIDInsensitively(input string) (*IdentityConditionalAccesAuthenticationContextClassReferenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityConditionalAccesAuthenticationContextClassReferenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityConditionalAccesAuthenticationContextClassReferenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityConditionalAccesAuthenticationContextClassReferenceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AuthenticationContextClassReferenceId, ok = input.Parsed["authenticationContextClassReferenceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authenticationContextClassReferenceId", input)
	}

	return nil
}

// ValidateIdentityConditionalAccesAuthenticationContextClassReferenceID checks that 'input' can be parsed as a Identity Conditional Acces Authentication Context Class Reference ID
func ValidateIdentityConditionalAccesAuthenticationContextClassReferenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityConditionalAccesAuthenticationContextClassReferenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Conditional Acces Authentication Context Class Reference ID
func (id IdentityConditionalAccesAuthenticationContextClassReferenceId) ID() string {
	fmtString := "/identity/conditionalAccess/authenticationContextClassReferences/%s"
	return fmt.Sprintf(fmtString, id.AuthenticationContextClassReferenceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Conditional Acces Authentication Context Class Reference ID
func (id IdentityConditionalAccesAuthenticationContextClassReferenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("conditionalAccess", "conditionalAccess", "conditionalAccess"),
		resourceids.StaticSegment("authenticationContextClassReferences", "authenticationContextClassReferences", "authenticationContextClassReferences"),
		resourceids.UserSpecifiedSegment("authenticationContextClassReferenceId", "authenticationContextClassReferenceIdValue"),
	}
}

// String returns a human-readable description of this Identity Conditional Acces Authentication Context Class Reference ID
func (id IdentityConditionalAccesAuthenticationContextClassReferenceId) String() string {
	components := []string{
		fmt.Sprintf("Authentication Context Class Reference: %q", id.AuthenticationContextClassReferenceId),
	}
	return fmt.Sprintf("Identity Conditional Acces Authentication Context Class Reference (%s)", strings.Join(components, "\n"))
}
