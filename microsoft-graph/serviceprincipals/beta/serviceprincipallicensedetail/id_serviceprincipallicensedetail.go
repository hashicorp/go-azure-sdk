package serviceprincipallicensedetail

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalLicenseDetailId{}

// ServicePrincipalLicenseDetailId is a struct representing the Resource ID for a Service Principal License Detail
type ServicePrincipalLicenseDetailId struct {
	ServicePrincipalId string
	LicenseDetailsId   string
}

// NewServicePrincipalLicenseDetailID returns a new ServicePrincipalLicenseDetailId struct
func NewServicePrincipalLicenseDetailID(servicePrincipalId string, licenseDetailsId string) ServicePrincipalLicenseDetailId {
	return ServicePrincipalLicenseDetailId{
		ServicePrincipalId: servicePrincipalId,
		LicenseDetailsId:   licenseDetailsId,
	}
}

// ParseServicePrincipalLicenseDetailID parses 'input' into a ServicePrincipalLicenseDetailId
func ParseServicePrincipalLicenseDetailID(input string) (*ServicePrincipalLicenseDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalLicenseDetailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalLicenseDetailId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.LicenseDetailsId, ok = parsed.Parsed["licenseDetailsId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "licenseDetailsId", *parsed)
	}

	return &id, nil
}

// ParseServicePrincipalLicenseDetailIDInsensitively parses 'input' case-insensitively into a ServicePrincipalLicenseDetailId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalLicenseDetailIDInsensitively(input string) (*ServicePrincipalLicenseDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalLicenseDetailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalLicenseDetailId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.LicenseDetailsId, ok = parsed.Parsed["licenseDetailsId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "licenseDetailsId", *parsed)
	}

	return &id, nil
}

// ValidateServicePrincipalLicenseDetailID checks that 'input' can be parsed as a Service Principal License Detail ID
func ValidateServicePrincipalLicenseDetailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalLicenseDetailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal License Detail ID
func (id ServicePrincipalLicenseDetailId) ID() string {
	fmtString := "/servicePrincipals/%s/licenseDetails/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.LicenseDetailsId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal License Detail ID
func (id ServicePrincipalLicenseDetailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticServicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalIdValue"),
		resourceids.StaticSegment("staticLicenseDetails", "licenseDetails", "licenseDetails"),
		resourceids.UserSpecifiedSegment("licenseDetailsId", "licenseDetailsIdValue"),
	}
}

// String returns a human-readable description of this Service Principal License Detail ID
func (id ServicePrincipalLicenseDetailId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("License Details: %q", id.LicenseDetailsId),
	}
	return fmt.Sprintf("Service Principal License Detail (%s)", strings.Join(components, "\n"))
}
