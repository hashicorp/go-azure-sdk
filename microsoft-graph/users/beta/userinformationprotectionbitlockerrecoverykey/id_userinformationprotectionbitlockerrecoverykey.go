package userinformationprotectionbitlockerrecoverykey

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserInformationProtectionBitlockerRecoveryKeyId{}

// UserInformationProtectionBitlockerRecoveryKeyId is a struct representing the Resource ID for a User Information Protection Bitlocker Recovery Key
type UserInformationProtectionBitlockerRecoveryKeyId struct {
	UserId                 string
	BitlockerRecoveryKeyId string
}

// NewUserInformationProtectionBitlockerRecoveryKeyID returns a new UserInformationProtectionBitlockerRecoveryKeyId struct
func NewUserInformationProtectionBitlockerRecoveryKeyID(userId string, bitlockerRecoveryKeyId string) UserInformationProtectionBitlockerRecoveryKeyId {
	return UserInformationProtectionBitlockerRecoveryKeyId{
		UserId:                 userId,
		BitlockerRecoveryKeyId: bitlockerRecoveryKeyId,
	}
}

// ParseUserInformationProtectionBitlockerRecoveryKeyID parses 'input' into a UserInformationProtectionBitlockerRecoveryKeyId
func ParseUserInformationProtectionBitlockerRecoveryKeyID(input string) (*UserInformationProtectionBitlockerRecoveryKeyId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserInformationProtectionBitlockerRecoveryKeyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserInformationProtectionBitlockerRecoveryKeyId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.BitlockerRecoveryKeyId, ok = parsed.Parsed["bitlockerRecoveryKeyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "bitlockerRecoveryKeyId", *parsed)
	}

	return &id, nil
}

// ParseUserInformationProtectionBitlockerRecoveryKeyIDInsensitively parses 'input' case-insensitively into a UserInformationProtectionBitlockerRecoveryKeyId
// note: this method should only be used for API response data and not user input
func ParseUserInformationProtectionBitlockerRecoveryKeyIDInsensitively(input string) (*UserInformationProtectionBitlockerRecoveryKeyId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserInformationProtectionBitlockerRecoveryKeyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserInformationProtectionBitlockerRecoveryKeyId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.BitlockerRecoveryKeyId, ok = parsed.Parsed["bitlockerRecoveryKeyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "bitlockerRecoveryKeyId", *parsed)
	}

	return &id, nil
}

// ValidateUserInformationProtectionBitlockerRecoveryKeyID checks that 'input' can be parsed as a User Information Protection Bitlocker Recovery Key ID
func ValidateUserInformationProtectionBitlockerRecoveryKeyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserInformationProtectionBitlockerRecoveryKeyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Information Protection Bitlocker Recovery Key ID
func (id UserInformationProtectionBitlockerRecoveryKeyId) ID() string {
	fmtString := "/users/%s/informationProtection/bitlocker/recoveryKeys/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.BitlockerRecoveryKeyId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Information Protection Bitlocker Recovery Key ID
func (id UserInformationProtectionBitlockerRecoveryKeyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticInformationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("staticBitlocker", "bitlocker", "bitlocker"),
		resourceids.StaticSegment("staticRecoveryKeys", "recoveryKeys", "recoveryKeys"),
		resourceids.UserSpecifiedSegment("bitlockerRecoveryKeyId", "bitlockerRecoveryKeyIdValue"),
	}
}

// String returns a human-readable description of this User Information Protection Bitlocker Recovery Key ID
func (id UserInformationProtectionBitlockerRecoveryKeyId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Bitlocker Recovery Key: %q", id.BitlockerRecoveryKeyId),
	}
	return fmt.Sprintf("User Information Protection Bitlocker Recovery Key (%s)", strings.Join(components, "\n"))
}
