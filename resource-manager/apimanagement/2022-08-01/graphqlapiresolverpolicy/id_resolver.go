package graphqlapiresolverpolicy

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ResolverId{}

// ResolverId is a struct representing the Resource ID for a Resolver
type ResolverId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	ApiId             string
	ResolverId        string
}

// NewResolverID returns a new ResolverId struct
func NewResolverID(subscriptionId string, resourceGroupName string, serviceName string, apiId string, resolverId string) ResolverId {
	return ResolverId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		ApiId:             apiId,
		ResolverId:        resolverId,
	}
}

// ParseResolverID parses 'input' into a ResolverId
func ParseResolverID(input string) (*ResolverId, error) {
	parser := resourceids.NewParserFromResourceIdType(ResolverId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ResolverId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.ApiId, ok = parsed.Parsed["apiId"]; !ok {
		return nil, fmt.Errorf("the segment 'apiId' was not found in the resource id %q", input)
	}

	if id.ResolverId, ok = parsed.Parsed["resolverId"]; !ok {
		return nil, fmt.Errorf("the segment 'resolverId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseResolverIDInsensitively parses 'input' case-insensitively into a ResolverId
// note: this method should only be used for API response data and not user input
func ParseResolverIDInsensitively(input string) (*ResolverId, error) {
	parser := resourceids.NewParserFromResourceIdType(ResolverId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ResolverId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.ApiId, ok = parsed.Parsed["apiId"]; !ok {
		return nil, fmt.Errorf("the segment 'apiId' was not found in the resource id %q", input)
	}

	if id.ResolverId, ok = parsed.Parsed["resolverId"]; !ok {
		return nil, fmt.Errorf("the segment 'resolverId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateResolverID checks that 'input' can be parsed as a Resolver ID
func ValidateResolverID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseResolverID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Resolver ID
func (id ResolverId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/apis/%s/resolvers/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.ApiId, id.ResolverId)
}

// Segments returns a slice of Resource ID Segments which comprise this Resolver ID
func (id ResolverId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApiManagement", "Microsoft.ApiManagement", "Microsoft.ApiManagement"),
		resourceids.StaticSegment("staticService", "service", "service"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
		resourceids.StaticSegment("staticApis", "apis", "apis"),
		resourceids.UserSpecifiedSegment("apiId", "apiIdValue"),
		resourceids.StaticSegment("staticResolvers", "resolvers", "resolvers"),
		resourceids.UserSpecifiedSegment("resolverId", "resolverIdValue"),
	}
}

// String returns a human-readable description of this Resolver ID
func (id ResolverId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Api: %q", id.ApiId),
		fmt.Sprintf("Resolver: %q", id.ResolverId),
	}
	return fmt.Sprintf("Resolver (%s)", strings.Join(components, "\n"))
}
