package useragreementacceptance

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserAgreementAcceptanceId{}

// UserAgreementAcceptanceId is a struct representing the Resource ID for a User Agreement Acceptance
type UserAgreementAcceptanceId struct {
	UserId                string
	AgreementAcceptanceId string
}

// NewUserAgreementAcceptanceID returns a new UserAgreementAcceptanceId struct
func NewUserAgreementAcceptanceID(userId string, agreementAcceptanceId string) UserAgreementAcceptanceId {
	return UserAgreementAcceptanceId{
		UserId:                userId,
		AgreementAcceptanceId: agreementAcceptanceId,
	}
}

// ParseUserAgreementAcceptanceID parses 'input' into a UserAgreementAcceptanceId
func ParseUserAgreementAcceptanceID(input string) (*UserAgreementAcceptanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAgreementAcceptanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAgreementAcceptanceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.AgreementAcceptanceId, ok = parsed.Parsed["agreementAcceptanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "agreementAcceptanceId", *parsed)
	}

	return &id, nil
}

// ParseUserAgreementAcceptanceIDInsensitively parses 'input' case-insensitively into a UserAgreementAcceptanceId
// note: this method should only be used for API response data and not user input
func ParseUserAgreementAcceptanceIDInsensitively(input string) (*UserAgreementAcceptanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAgreementAcceptanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAgreementAcceptanceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.AgreementAcceptanceId, ok = parsed.Parsed["agreementAcceptanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "agreementAcceptanceId", *parsed)
	}

	return &id, nil
}

// ValidateUserAgreementAcceptanceID checks that 'input' can be parsed as a User Agreement Acceptance ID
func ValidateUserAgreementAcceptanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserAgreementAcceptanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Agreement Acceptance ID
func (id UserAgreementAcceptanceId) ID() string {
	fmtString := "/users/%s/agreementAcceptances/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AgreementAcceptanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Agreement Acceptance ID
func (id UserAgreementAcceptanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticAgreementAcceptances", "agreementAcceptances", "agreementAcceptances"),
		resourceids.UserSpecifiedSegment("agreementAcceptanceId", "agreementAcceptanceIdValue"),
	}
}

// String returns a human-readable description of this User Agreement Acceptance ID
func (id UserAgreementAcceptanceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Agreement Acceptance: %q", id.AgreementAcceptanceId),
	}
	return fmt.Sprintf("User Agreement Acceptance (%s)", strings.Join(components, "\n"))
}
