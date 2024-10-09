package secrets

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&SecretId{})
}

var _ resourceids.ResourceId = &SecretId{}

// SecretId is a struct representing the Resource ID for a Secret
type SecretId struct {
	SubscriptionId    string
	ResourceGroupName string
	LabName           string
	UserName          string
	SecretName        string
}

// NewSecretID returns a new SecretId struct
func NewSecretID(subscriptionId string, resourceGroupName string, labName string, userName string, secretName string) SecretId {
	return SecretId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		LabName:           labName,
		UserName:          userName,
		SecretName:        secretName,
	}
}

// ParseSecretID parses 'input' into a SecretId
func ParseSecretID(input string) (*SecretId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SecretId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SecretId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSecretIDInsensitively parses 'input' case-insensitively into a SecretId
// note: this method should only be used for API response data and not user input
func ParseSecretIDInsensitively(input string) (*SecretId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SecretId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SecretId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SecretId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.LabName, ok = input.Parsed["labName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "labName", input)
	}

	if id.UserName, ok = input.Parsed["userName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userName", input)
	}

	if id.SecretName, ok = input.Parsed["secretName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "secretName", input)
	}

	return nil
}

// ValidateSecretID checks that 'input' can be parsed as a Secret ID
func ValidateSecretID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSecretID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Secret ID
func (id SecretId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DevTestLab/labs/%s/users/%s/secrets/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LabName, id.UserName, id.SecretName)
}

// Segments returns a slice of Resource ID Segments which comprise this Secret ID
func (id SecretId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDevTestLab", "Microsoft.DevTestLab", "Microsoft.DevTestLab"),
		resourceids.StaticSegment("staticLabs", "labs", "labs"),
		resourceids.UserSpecifiedSegment("labName", "labName"),
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userName", "userName"),
		resourceids.StaticSegment("staticSecrets", "secrets", "secrets"),
		resourceids.UserSpecifiedSegment("secretName", "secretName"),
	}
}

// String returns a human-readable description of this Secret ID
func (id SecretId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Lab Name: %q", id.LabName),
		fmt.Sprintf("User Name: %q", id.UserName),
		fmt.Sprintf("Secret Name: %q", id.SecretName),
	}
	return fmt.Sprintf("Secret (%s)", strings.Join(components, "\n"))
}
