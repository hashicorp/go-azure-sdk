package user

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserByIdTranslateExchangeIdRequestSourceIdType string

const (
	CreateUserByIdTranslateExchangeIdRequestSourceIdTypeentryId              CreateUserByIdTranslateExchangeIdRequestSourceIdType = "EntryId"
	CreateUserByIdTranslateExchangeIdRequestSourceIdTypeewsId                CreateUserByIdTranslateExchangeIdRequestSourceIdType = "EwsId"
	CreateUserByIdTranslateExchangeIdRequestSourceIdTypeimmutableEntryId     CreateUserByIdTranslateExchangeIdRequestSourceIdType = "ImmutableEntryId"
	CreateUserByIdTranslateExchangeIdRequestSourceIdTyperestId               CreateUserByIdTranslateExchangeIdRequestSourceIdType = "RestId"
	CreateUserByIdTranslateExchangeIdRequestSourceIdTyperestImmutableEntryId CreateUserByIdTranslateExchangeIdRequestSourceIdType = "RestImmutableEntryId"
)

func PossibleValuesForCreateUserByIdTranslateExchangeIdRequestSourceIdType() []string {
	return []string{
		string(CreateUserByIdTranslateExchangeIdRequestSourceIdTypeentryId),
		string(CreateUserByIdTranslateExchangeIdRequestSourceIdTypeewsId),
		string(CreateUserByIdTranslateExchangeIdRequestSourceIdTypeimmutableEntryId),
		string(CreateUserByIdTranslateExchangeIdRequestSourceIdTyperestId),
		string(CreateUserByIdTranslateExchangeIdRequestSourceIdTyperestImmutableEntryId),
	}
}

func (s *CreateUserByIdTranslateExchangeIdRequestSourceIdType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateUserByIdTranslateExchangeIdRequestSourceIdType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateUserByIdTranslateExchangeIdRequestSourceIdType(input string) (*CreateUserByIdTranslateExchangeIdRequestSourceIdType, error) {
	vals := map[string]CreateUserByIdTranslateExchangeIdRequestSourceIdType{
		"entryid":              CreateUserByIdTranslateExchangeIdRequestSourceIdTypeentryId,
		"ewsid":                CreateUserByIdTranslateExchangeIdRequestSourceIdTypeewsId,
		"immutableentryid":     CreateUserByIdTranslateExchangeIdRequestSourceIdTypeimmutableEntryId,
		"restid":               CreateUserByIdTranslateExchangeIdRequestSourceIdTyperestId,
		"restimmutableentryid": CreateUserByIdTranslateExchangeIdRequestSourceIdTyperestImmutableEntryId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateUserByIdTranslateExchangeIdRequestSourceIdType(input)
	return &out, nil
}

type CreateUserByIdTranslateExchangeIdRequestTargetIdType string

const (
	CreateUserByIdTranslateExchangeIdRequestTargetIdTypeentryId              CreateUserByIdTranslateExchangeIdRequestTargetIdType = "EntryId"
	CreateUserByIdTranslateExchangeIdRequestTargetIdTypeewsId                CreateUserByIdTranslateExchangeIdRequestTargetIdType = "EwsId"
	CreateUserByIdTranslateExchangeIdRequestTargetIdTypeimmutableEntryId     CreateUserByIdTranslateExchangeIdRequestTargetIdType = "ImmutableEntryId"
	CreateUserByIdTranslateExchangeIdRequestTargetIdTyperestId               CreateUserByIdTranslateExchangeIdRequestTargetIdType = "RestId"
	CreateUserByIdTranslateExchangeIdRequestTargetIdTyperestImmutableEntryId CreateUserByIdTranslateExchangeIdRequestTargetIdType = "RestImmutableEntryId"
)

func PossibleValuesForCreateUserByIdTranslateExchangeIdRequestTargetIdType() []string {
	return []string{
		string(CreateUserByIdTranslateExchangeIdRequestTargetIdTypeentryId),
		string(CreateUserByIdTranslateExchangeIdRequestTargetIdTypeewsId),
		string(CreateUserByIdTranslateExchangeIdRequestTargetIdTypeimmutableEntryId),
		string(CreateUserByIdTranslateExchangeIdRequestTargetIdTyperestId),
		string(CreateUserByIdTranslateExchangeIdRequestTargetIdTyperestImmutableEntryId),
	}
}

func (s *CreateUserByIdTranslateExchangeIdRequestTargetIdType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateUserByIdTranslateExchangeIdRequestTargetIdType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateUserByIdTranslateExchangeIdRequestTargetIdType(input string) (*CreateUserByIdTranslateExchangeIdRequestTargetIdType, error) {
	vals := map[string]CreateUserByIdTranslateExchangeIdRequestTargetIdType{
		"entryid":              CreateUserByIdTranslateExchangeIdRequestTargetIdTypeentryId,
		"ewsid":                CreateUserByIdTranslateExchangeIdRequestTargetIdTypeewsId,
		"immutableentryid":     CreateUserByIdTranslateExchangeIdRequestTargetIdTypeimmutableEntryId,
		"restid":               CreateUserByIdTranslateExchangeIdRequestTargetIdTyperestId,
		"restimmutableentryid": CreateUserByIdTranslateExchangeIdRequestTargetIdTyperestImmutableEntryId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateUserByIdTranslateExchangeIdRequestTargetIdType(input)
	return &out, nil
}

type GetUserByIdMailTipRequestMailTipsOptions string

const (
	GetUserByIdMailTipRequestMailTipsOptionsautomaticReplies     GetUserByIdMailTipRequestMailTipsOptions = "AutomaticReplies"
	GetUserByIdMailTipRequestMailTipsOptionscustomMailTip        GetUserByIdMailTipRequestMailTipsOptions = "CustomMailTip"
	GetUserByIdMailTipRequestMailTipsOptionsdeliveryRestriction  GetUserByIdMailTipRequestMailTipsOptions = "DeliveryRestriction"
	GetUserByIdMailTipRequestMailTipsOptionsexternalMemberCount  GetUserByIdMailTipRequestMailTipsOptions = "ExternalMemberCount"
	GetUserByIdMailTipRequestMailTipsOptionsmailboxFullStatus    GetUserByIdMailTipRequestMailTipsOptions = "MailboxFullStatus"
	GetUserByIdMailTipRequestMailTipsOptionsmaxMessageSize       GetUserByIdMailTipRequestMailTipsOptions = "MaxMessageSize"
	GetUserByIdMailTipRequestMailTipsOptionsmoderationStatus     GetUserByIdMailTipRequestMailTipsOptions = "ModerationStatus"
	GetUserByIdMailTipRequestMailTipsOptionsrecipientScope       GetUserByIdMailTipRequestMailTipsOptions = "RecipientScope"
	GetUserByIdMailTipRequestMailTipsOptionsrecipientSuggestions GetUserByIdMailTipRequestMailTipsOptions = "RecipientSuggestions"
	GetUserByIdMailTipRequestMailTipsOptionstotalMemberCount     GetUserByIdMailTipRequestMailTipsOptions = "TotalMemberCount"
)

func PossibleValuesForGetUserByIdMailTipRequestMailTipsOptions() []string {
	return []string{
		string(GetUserByIdMailTipRequestMailTipsOptionsautomaticReplies),
		string(GetUserByIdMailTipRequestMailTipsOptionscustomMailTip),
		string(GetUserByIdMailTipRequestMailTipsOptionsdeliveryRestriction),
		string(GetUserByIdMailTipRequestMailTipsOptionsexternalMemberCount),
		string(GetUserByIdMailTipRequestMailTipsOptionsmailboxFullStatus),
		string(GetUserByIdMailTipRequestMailTipsOptionsmaxMessageSize),
		string(GetUserByIdMailTipRequestMailTipsOptionsmoderationStatus),
		string(GetUserByIdMailTipRequestMailTipsOptionsrecipientScope),
		string(GetUserByIdMailTipRequestMailTipsOptionsrecipientSuggestions),
		string(GetUserByIdMailTipRequestMailTipsOptionstotalMemberCount),
	}
}

func (s *GetUserByIdMailTipRequestMailTipsOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGetUserByIdMailTipRequestMailTipsOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGetUserByIdMailTipRequestMailTipsOptions(input string) (*GetUserByIdMailTipRequestMailTipsOptions, error) {
	vals := map[string]GetUserByIdMailTipRequestMailTipsOptions{
		"automaticreplies":     GetUserByIdMailTipRequestMailTipsOptionsautomaticReplies,
		"custommailtip":        GetUserByIdMailTipRequestMailTipsOptionscustomMailTip,
		"deliveryrestriction":  GetUserByIdMailTipRequestMailTipsOptionsdeliveryRestriction,
		"externalmembercount":  GetUserByIdMailTipRequestMailTipsOptionsexternalMemberCount,
		"mailboxfullstatus":    GetUserByIdMailTipRequestMailTipsOptionsmailboxFullStatus,
		"maxmessagesize":       GetUserByIdMailTipRequestMailTipsOptionsmaxMessageSize,
		"moderationstatus":     GetUserByIdMailTipRequestMailTipsOptionsmoderationStatus,
		"recipientscope":       GetUserByIdMailTipRequestMailTipsOptionsrecipientScope,
		"recipientsuggestions": GetUserByIdMailTipRequestMailTipsOptionsrecipientSuggestions,
		"totalmembercount":     GetUserByIdMailTipRequestMailTipsOptionstotalMemberCount,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GetUserByIdMailTipRequestMailTipsOptions(input)
	return &out, nil
}
