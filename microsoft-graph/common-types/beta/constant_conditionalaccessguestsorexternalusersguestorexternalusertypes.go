package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes string

const (
	ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes_B2bCollaborationGuest  ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes = "b2bCollaborationGuest"
	ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes_B2bCollaborationMember ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes = "b2bCollaborationMember"
	ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes_B2bDirectConnectUser   ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes = "b2bDirectConnectUser"
	ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes_InternalGuest          ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes = "internalGuest"
	ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes_None                   ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes = "none"
	ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes_OtherExternalUser      ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes = "otherExternalUser"
	ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes_ServiceProvider        ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes = "serviceProvider"
)

func PossibleValuesForConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes() []string {
	return []string{
		string(ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes_B2bCollaborationGuest),
		string(ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes_B2bCollaborationMember),
		string(ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes_B2bDirectConnectUser),
		string(ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes_InternalGuest),
		string(ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes_None),
		string(ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes_OtherExternalUser),
		string(ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes_ServiceProvider),
	}
}

func (s *ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes(input string) (*ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes, error) {
	vals := map[string]ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes{
		"b2bcollaborationguest":  ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes_B2bCollaborationGuest,
		"b2bcollaborationmember": ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes_B2bCollaborationMember,
		"b2bdirectconnectuser":   ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes_B2bDirectConnectUser,
		"internalguest":          ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes_InternalGuest,
		"none":                   ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes_None,
		"otherexternaluser":      ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes_OtherExternalUser,
		"serviceprovider":        ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes_ServiceProvider,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes(input)
	return &out, nil
}
