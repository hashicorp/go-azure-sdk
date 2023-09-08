package serviceprincipaloauth2permissiongrant

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalOauth2PermissionGrantId{}

// ServicePrincipalOauth2PermissionGrantId is a struct representing the Resource ID for a Service Principal Oauth 2 Permission Grant
type ServicePrincipalOauth2PermissionGrantId struct {
	ServicePrincipalId      string
	OAuth2PermissionGrantId string
}

// NewServicePrincipalOauth2PermissionGrantID returns a new ServicePrincipalOauth2PermissionGrantId struct
func NewServicePrincipalOauth2PermissionGrantID(servicePrincipalId string, oAuth2PermissionGrantId string) ServicePrincipalOauth2PermissionGrantId {
	return ServicePrincipalOauth2PermissionGrantId{
		ServicePrincipalId:      servicePrincipalId,
		OAuth2PermissionGrantId: oAuth2PermissionGrantId,
	}
}

// ParseServicePrincipalOauth2PermissionGrantID parses 'input' into a ServicePrincipalOauth2PermissionGrantId
func ParseServicePrincipalOauth2PermissionGrantID(input string) (*ServicePrincipalOauth2PermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalOauth2PermissionGrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalOauth2PermissionGrantId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.OAuth2PermissionGrantId, ok = parsed.Parsed["oAuth2PermissionGrantId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "oAuth2PermissionGrantId", *parsed)
	}

	return &id, nil
}

// ParseServicePrincipalOauth2PermissionGrantIDInsensitively parses 'input' case-insensitively into a ServicePrincipalOauth2PermissionGrantId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalOauth2PermissionGrantIDInsensitively(input string) (*ServicePrincipalOauth2PermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalOauth2PermissionGrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalOauth2PermissionGrantId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.OAuth2PermissionGrantId, ok = parsed.Parsed["oAuth2PermissionGrantId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "oAuth2PermissionGrantId", *parsed)
	}

	return &id, nil
}

// ValidateServicePrincipalOauth2PermissionGrantID checks that 'input' can be parsed as a Service Principal Oauth 2 Permission Grant ID
func ValidateServicePrincipalOauth2PermissionGrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalOauth2PermissionGrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Oauth 2 Permission Grant ID
func (id ServicePrincipalOauth2PermissionGrantId) ID() string {
	fmtString := "/servicePrincipals/%s/oauth2PermissionGrants/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.OAuth2PermissionGrantId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Oauth 2 Permission Grant ID
func (id ServicePrincipalOauth2PermissionGrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticServicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalIdValue"),
		resourceids.StaticSegment("staticOauth2PermissionGrants", "oauth2PermissionGrants", "oauth2PermissionGrants"),
		resourceids.UserSpecifiedSegment("oAuth2PermissionGrantId", "oAuth2PermissionGrantIdValue"),
	}
}

// String returns a human-readable description of this Service Principal Oauth 2 Permission Grant ID
func (id ServicePrincipalOauth2PermissionGrantId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("O Auth 2 Permission Grant: %q", id.OAuth2PermissionGrantId),
	}
	return fmt.Sprintf("Service Principal Oauth 2 Permission Grant (%s)", strings.Join(components, "\n"))
}
