package authorizationaccesspolicy

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = AccessPolicyId{}

// AccessPolicyId is a struct representing the Resource ID for a Access Policy
type AccessPolicyId struct {
	SubscriptionId              string
	ResourceGroupName           string
	ServiceName                 string
	AuthorizationProviderId     string
	AuthorizationId             string
	AuthorizationAccessPolicyId string
}

// NewAccessPolicyID returns a new AccessPolicyId struct
func NewAccessPolicyID(subscriptionId string, resourceGroupName string, serviceName string, authorizationProviderId string, authorizationId string, authorizationAccessPolicyId string) AccessPolicyId {
	return AccessPolicyId{
		SubscriptionId:              subscriptionId,
		ResourceGroupName:           resourceGroupName,
		ServiceName:                 serviceName,
		AuthorizationProviderId:     authorizationProviderId,
		AuthorizationId:             authorizationId,
		AuthorizationAccessPolicyId: authorizationAccessPolicyId,
	}
}

// ParseAccessPolicyID parses 'input' into a AccessPolicyId
func ParseAccessPolicyID(input string) (*AccessPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(AccessPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AccessPolicyId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.AuthorizationProviderId, ok = parsed.Parsed["authorizationProviderId"]; !ok {
		return nil, fmt.Errorf("the segment 'authorizationProviderId' was not found in the resource id %q", input)
	}

	if id.AuthorizationId, ok = parsed.Parsed["authorizationId"]; !ok {
		return nil, fmt.Errorf("the segment 'authorizationId' was not found in the resource id %q", input)
	}

	if id.AuthorizationAccessPolicyId, ok = parsed.Parsed["authorizationAccessPolicyId"]; !ok {
		return nil, fmt.Errorf("the segment 'authorizationAccessPolicyId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseAccessPolicyIDInsensitively parses 'input' case-insensitively into a AccessPolicyId
// note: this method should only be used for API response data and not user input
func ParseAccessPolicyIDInsensitively(input string) (*AccessPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(AccessPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AccessPolicyId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.AuthorizationProviderId, ok = parsed.Parsed["authorizationProviderId"]; !ok {
		return nil, fmt.Errorf("the segment 'authorizationProviderId' was not found in the resource id %q", input)
	}

	if id.AuthorizationId, ok = parsed.Parsed["authorizationId"]; !ok {
		return nil, fmt.Errorf("the segment 'authorizationId' was not found in the resource id %q", input)
	}

	if id.AuthorizationAccessPolicyId, ok = parsed.Parsed["authorizationAccessPolicyId"]; !ok {
		return nil, fmt.Errorf("the segment 'authorizationAccessPolicyId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateAccessPolicyID checks that 'input' can be parsed as a Access Policy ID
func ValidateAccessPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAccessPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Access Policy ID
func (id AccessPolicyId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/authorizationProviders/%s/authorizations/%s/accessPolicies/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.AuthorizationProviderId, id.AuthorizationId, id.AuthorizationAccessPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Access Policy ID
func (id AccessPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApiManagement", "Microsoft.ApiManagement", "Microsoft.ApiManagement"),
		resourceids.StaticSegment("staticService", "service", "service"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
		resourceids.StaticSegment("staticAuthorizationProviders", "authorizationProviders", "authorizationProviders"),
		resourceids.UserSpecifiedSegment("authorizationProviderId", "authorizationProviderIdValue"),
		resourceids.StaticSegment("staticAuthorizations", "authorizations", "authorizations"),
		resourceids.UserSpecifiedSegment("authorizationId", "authorizationIdValue"),
		resourceids.StaticSegment("staticAccessPolicies", "accessPolicies", "accessPolicies"),
		resourceids.UserSpecifiedSegment("authorizationAccessPolicyId", "authorizationAccessPolicyIdValue"),
	}
}

// String returns a human-readable description of this Access Policy ID
func (id AccessPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Authorization Provider: %q", id.AuthorizationProviderId),
		fmt.Sprintf("Authorization: %q", id.AuthorizationId),
		fmt.Sprintf("Authorization Access Policy: %q", id.AuthorizationAccessPolicyId),
	}
	return fmt.Sprintf("Access Policy (%s)", strings.Join(components, "\n"))
}
