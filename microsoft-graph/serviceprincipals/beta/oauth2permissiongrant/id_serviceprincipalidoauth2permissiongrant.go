package oauth2permissiongrant

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdOauth2PermissionGrantId{}

// ServicePrincipalIdOauth2PermissionGrantId is a struct representing the Resource ID for a Service Principal Id Oauth 2 Permission Grant
type ServicePrincipalIdOauth2PermissionGrantId struct {
	ServicePrincipalId      string
	OAuth2PermissionGrantId string
}

// NewServicePrincipalIdOauth2PermissionGrantID returns a new ServicePrincipalIdOauth2PermissionGrantId struct
func NewServicePrincipalIdOauth2PermissionGrantID(servicePrincipalId string, oAuth2PermissionGrantId string) ServicePrincipalIdOauth2PermissionGrantId {
	return ServicePrincipalIdOauth2PermissionGrantId{
		ServicePrincipalId:      servicePrincipalId,
		OAuth2PermissionGrantId: oAuth2PermissionGrantId,
	}
}

// ParseServicePrincipalIdOauth2PermissionGrantID parses 'input' into a ServicePrincipalIdOauth2PermissionGrantId
func ParseServicePrincipalIdOauth2PermissionGrantID(input string) (*ServicePrincipalIdOauth2PermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdOauth2PermissionGrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdOauth2PermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServicePrincipalIdOauth2PermissionGrantIDInsensitively parses 'input' case-insensitively into a ServicePrincipalIdOauth2PermissionGrantId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalIdOauth2PermissionGrantIDInsensitively(input string) (*ServicePrincipalIdOauth2PermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalIdOauth2PermissionGrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServicePrincipalIdOauth2PermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServicePrincipalIdOauth2PermissionGrantId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	if id.OAuth2PermissionGrantId, ok = input.Parsed["oAuth2PermissionGrantId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "oAuth2PermissionGrantId", input)
	}

	return nil
}

// ValidateServicePrincipalIdOauth2PermissionGrantID checks that 'input' can be parsed as a Service Principal Id Oauth 2 Permission Grant ID
func ValidateServicePrincipalIdOauth2PermissionGrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalIdOauth2PermissionGrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Id Oauth 2 Permission Grant ID
func (id ServicePrincipalIdOauth2PermissionGrantId) ID() string {
	fmtString := "/servicePrincipals/%s/oauth2PermissionGrants/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.OAuth2PermissionGrantId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Id Oauth 2 Permission Grant ID
func (id ServicePrincipalIdOauth2PermissionGrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("servicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalIdValue"),
		resourceids.StaticSegment("oauth2PermissionGrants", "oauth2PermissionGrants", "oauth2PermissionGrants"),
		resourceids.UserSpecifiedSegment("oAuth2PermissionGrantId", "oAuth2PermissionGrantIdValue"),
	}
}

// String returns a human-readable description of this Service Principal Id Oauth 2 Permission Grant ID
func (id ServicePrincipalIdOauth2PermissionGrantId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("O Auth 2 Permission Grant: %q", id.OAuth2PermissionGrantId),
	}
	return fmt.Sprintf("Service Principal Id Oauth 2 Permission Grant (%s)", strings.Join(components, "\n"))
}
