package authorizations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = AuthorizationProviderId{}

// AuthorizationProviderId is a struct representing the Resource ID for a Authorization Provider
type AuthorizationProviderId struct {
	SubscriptionId          string
	ResourceGroupName       string
	ServiceName             string
	AuthorizationProviderId string
}

// NewAuthorizationProviderID returns a new AuthorizationProviderId struct
func NewAuthorizationProviderID(subscriptionId string, resourceGroupName string, serviceName string, authorizationProviderId string) AuthorizationProviderId {
	return AuthorizationProviderId{
		SubscriptionId:          subscriptionId,
		ResourceGroupName:       resourceGroupName,
		ServiceName:             serviceName,
		AuthorizationProviderId: authorizationProviderId,
	}
}

// ParseAuthorizationProviderID parses 'input' into a AuthorizationProviderId
func ParseAuthorizationProviderID(input string) (*AuthorizationProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(AuthorizationProviderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AuthorizationProviderId{}

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

	return &id, nil
}

// ParseAuthorizationProviderIDInsensitively parses 'input' case-insensitively into a AuthorizationProviderId
// note: this method should only be used for API response data and not user input
func ParseAuthorizationProviderIDInsensitively(input string) (*AuthorizationProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(AuthorizationProviderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AuthorizationProviderId{}

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

	return &id, nil
}

// ValidateAuthorizationProviderID checks that 'input' can be parsed as a Authorization Provider ID
func ValidateAuthorizationProviderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAuthorizationProviderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Authorization Provider ID
func (id AuthorizationProviderId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/authorizationProviders/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.AuthorizationProviderId)
}

// Segments returns a slice of Resource ID Segments which comprise this Authorization Provider ID
func (id AuthorizationProviderId) Segments() []resourceids.Segment {
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
	}
}

// String returns a human-readable description of this Authorization Provider ID
func (id AuthorizationProviderId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Authorization Provider: %q", id.AuthorizationProviderId),
	}
	return fmt.Sprintf("Authorization Provider (%s)", strings.Join(components, "\n"))
}
