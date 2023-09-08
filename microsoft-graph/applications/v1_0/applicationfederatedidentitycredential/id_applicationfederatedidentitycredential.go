package applicationfederatedidentitycredential

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ApplicationFederatedIdentityCredentialId{}

// ApplicationFederatedIdentityCredentialId is a struct representing the Resource ID for a Application Federated Identity Credential
type ApplicationFederatedIdentityCredentialId struct {
	ApplicationId                 string
	FederatedIdentityCredentialId string
}

// NewApplicationFederatedIdentityCredentialID returns a new ApplicationFederatedIdentityCredentialId struct
func NewApplicationFederatedIdentityCredentialID(applicationId string, federatedIdentityCredentialId string) ApplicationFederatedIdentityCredentialId {
	return ApplicationFederatedIdentityCredentialId{
		ApplicationId:                 applicationId,
		FederatedIdentityCredentialId: federatedIdentityCredentialId,
	}
}

// ParseApplicationFederatedIdentityCredentialID parses 'input' into a ApplicationFederatedIdentityCredentialId
func ParseApplicationFederatedIdentityCredentialID(input string) (*ApplicationFederatedIdentityCredentialId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApplicationFederatedIdentityCredentialId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApplicationFederatedIdentityCredentialId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.FederatedIdentityCredentialId, ok = parsed.Parsed["federatedIdentityCredentialId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "federatedIdentityCredentialId", *parsed)
	}

	return &id, nil
}

// ParseApplicationFederatedIdentityCredentialIDInsensitively parses 'input' case-insensitively into a ApplicationFederatedIdentityCredentialId
// note: this method should only be used for API response data and not user input
func ParseApplicationFederatedIdentityCredentialIDInsensitively(input string) (*ApplicationFederatedIdentityCredentialId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApplicationFederatedIdentityCredentialId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApplicationFederatedIdentityCredentialId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.FederatedIdentityCredentialId, ok = parsed.Parsed["federatedIdentityCredentialId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "federatedIdentityCredentialId", *parsed)
	}

	return &id, nil
}

// ValidateApplicationFederatedIdentityCredentialID checks that 'input' can be parsed as a Application Federated Identity Credential ID
func ValidateApplicationFederatedIdentityCredentialID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApplicationFederatedIdentityCredentialID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Application Federated Identity Credential ID
func (id ApplicationFederatedIdentityCredentialId) ID() string {
	fmtString := "/applications/%s/federatedIdentityCredentials/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.FederatedIdentityCredentialId)
}

// Segments returns a slice of Resource ID Segments which comprise this Application Federated Identity Credential ID
func (id ApplicationFederatedIdentityCredentialId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticApplications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "applicationIdValue"),
		resourceids.StaticSegment("staticFederatedIdentityCredentials", "federatedIdentityCredentials", "federatedIdentityCredentials"),
		resourceids.UserSpecifiedSegment("federatedIdentityCredentialId", "federatedIdentityCredentialIdValue"),
	}
}

// String returns a human-readable description of this Application Federated Identity Credential ID
func (id ApplicationFederatedIdentityCredentialId) String() string {
	components := []string{
		fmt.Sprintf("Application: %q", id.ApplicationId),
		fmt.Sprintf("Federated Identity Credential: %q", id.FederatedIdentityCredentialId),
	}
	return fmt.Sprintf("Application Federated Identity Credential (%s)", strings.Join(components, "\n"))
}
