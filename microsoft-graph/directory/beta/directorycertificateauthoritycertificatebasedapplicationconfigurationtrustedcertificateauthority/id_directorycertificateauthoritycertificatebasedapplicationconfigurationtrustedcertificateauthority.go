package directorycertificateauthoritycertificatebasedapplicationconfigurationtrustedcertificateauthority

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityId{}

// DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityId is a struct representing the Resource ID for a Directory Certificate Authority Certificate Based Application Configuration Trusted Certificate Authority
type DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityId struct {
	CertificateBasedApplicationConfigurationId string
	CertificateAuthorityAsEntityId             string
}

// NewDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityID returns a new DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityId struct
func NewDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityID(certificateBasedApplicationConfigurationId string, certificateAuthorityAsEntityId string) DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityId {
	return DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityId{
		CertificateBasedApplicationConfigurationId: certificateBasedApplicationConfigurationId,
		CertificateAuthorityAsEntityId:             certificateAuthorityAsEntityId,
	}
}

// ParseDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityID parses 'input' into a DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityId
func ParseDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityID(input string) (*DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityId, error) {
	parser := resourceids.NewParserFromResourceIdType(DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityId{}

	if id.CertificateBasedApplicationConfigurationId, ok = parsed.Parsed["certificateBasedApplicationConfigurationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "certificateBasedApplicationConfigurationId", *parsed)
	}

	if id.CertificateAuthorityAsEntityId, ok = parsed.Parsed["certificateAuthorityAsEntityId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "certificateAuthorityAsEntityId", *parsed)
	}

	return &id, nil
}

// ParseDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityIDInsensitively parses 'input' case-insensitively into a DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityId
// note: this method should only be used for API response data and not user input
func ParseDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityIDInsensitively(input string) (*DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityId, error) {
	parser := resourceids.NewParserFromResourceIdType(DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityId{}

	if id.CertificateBasedApplicationConfigurationId, ok = parsed.Parsed["certificateBasedApplicationConfigurationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "certificateBasedApplicationConfigurationId", *parsed)
	}

	if id.CertificateAuthorityAsEntityId, ok = parsed.Parsed["certificateAuthorityAsEntityId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "certificateAuthorityAsEntityId", *parsed)
	}

	return &id, nil
}

// ValidateDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityID checks that 'input' can be parsed as a Directory Certificate Authority Certificate Based Application Configuration Trusted Certificate Authority ID
func ValidateDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Certificate Authority Certificate Based Application Configuration Trusted Certificate Authority ID
func (id DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityId) ID() string {
	fmtString := "/directory/certificateAuthorities/certificateBasedApplicationConfigurations/%s/trustedCertificateAuthorities/%s"
	return fmt.Sprintf(fmtString, id.CertificateBasedApplicationConfigurationId, id.CertificateAuthorityAsEntityId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Certificate Authority Certificate Based Application Configuration Trusted Certificate Authority ID
func (id DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticDirectory", "directory", "directory"),
		resourceids.StaticSegment("staticCertificateAuthorities", "certificateAuthorities", "certificateAuthorities"),
		resourceids.StaticSegment("staticCertificateBasedApplicationConfigurations", "certificateBasedApplicationConfigurations", "certificateBasedApplicationConfigurations"),
		resourceids.UserSpecifiedSegment("certificateBasedApplicationConfigurationId", "certificateBasedApplicationConfigurationIdValue"),
		resourceids.StaticSegment("staticTrustedCertificateAuthorities", "trustedCertificateAuthorities", "trustedCertificateAuthorities"),
		resourceids.UserSpecifiedSegment("certificateAuthorityAsEntityId", "certificateAuthorityAsEntityIdValue"),
	}
}

// String returns a human-readable description of this Directory Certificate Authority Certificate Based Application Configuration Trusted Certificate Authority ID
func (id DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityId) String() string {
	components := []string{
		fmt.Sprintf("Certificate Based Application Configuration: %q", id.CertificateBasedApplicationConfigurationId),
		fmt.Sprintf("Certificate Authority As Entity: %q", id.CertificateAuthorityAsEntityId),
	}
	return fmt.Sprintf("Directory Certificate Authority Certificate Based Application Configuration Trusted Certificate Authority (%s)", strings.Join(components, "\n"))
}
