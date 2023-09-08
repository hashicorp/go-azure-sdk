package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes string

const (
	ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypesb2bCollaborationGuest  ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes = "B2bCollaborationGuest"
	ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypesb2bCollaborationMember ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes = "B2bCollaborationMember"
	ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypesb2bDirectConnectUser   ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes = "B2bDirectConnectUser"
	ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypesinternalGuest          ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes = "InternalGuest"
	ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypesnone                   ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes = "None"
	ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypesotherExternalUser      ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes = "OtherExternalUser"
	ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypesserviceProvider        ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes = "ServiceProvider"
)

func PossibleValuesForConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes() []string {
	return []string{
		string(ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypesb2bCollaborationGuest),
		string(ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypesb2bCollaborationMember),
		string(ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypesb2bDirectConnectUser),
		string(ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypesinternalGuest),
		string(ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypesnone),
		string(ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypesotherExternalUser),
		string(ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypesserviceProvider),
	}
}

func parseConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes(input string) (*ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes, error) {
	vals := map[string]ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes{
		"b2bcollaborationguest":  ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypesb2bCollaborationGuest,
		"b2bcollaborationmember": ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypesb2bCollaborationMember,
		"b2bdirectconnectuser":   ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypesb2bDirectConnectUser,
		"internalguest":          ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypesinternalGuest,
		"none":                   ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypesnone,
		"otherexternaluser":      ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypesotherExternalUser,
		"serviceprovider":        ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypesserviceProvider,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes(input)
	return &out, nil
}
