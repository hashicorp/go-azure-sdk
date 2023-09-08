package me

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMeTranslateExchangeIdRequestSourceIdType string

const (
	CreateMeTranslateExchangeIdRequestSourceIdTypeentryId              CreateMeTranslateExchangeIdRequestSourceIdType = "EntryId"
	CreateMeTranslateExchangeIdRequestSourceIdTypeewsId                CreateMeTranslateExchangeIdRequestSourceIdType = "EwsId"
	CreateMeTranslateExchangeIdRequestSourceIdTypeimmutableEntryId     CreateMeTranslateExchangeIdRequestSourceIdType = "ImmutableEntryId"
	CreateMeTranslateExchangeIdRequestSourceIdTyperestId               CreateMeTranslateExchangeIdRequestSourceIdType = "RestId"
	CreateMeTranslateExchangeIdRequestSourceIdTyperestImmutableEntryId CreateMeTranslateExchangeIdRequestSourceIdType = "RestImmutableEntryId"
)

func PossibleValuesForCreateMeTranslateExchangeIdRequestSourceIdType() []string {
	return []string{
		string(CreateMeTranslateExchangeIdRequestSourceIdTypeentryId),
		string(CreateMeTranslateExchangeIdRequestSourceIdTypeewsId),
		string(CreateMeTranslateExchangeIdRequestSourceIdTypeimmutableEntryId),
		string(CreateMeTranslateExchangeIdRequestSourceIdTyperestId),
		string(CreateMeTranslateExchangeIdRequestSourceIdTyperestImmutableEntryId),
	}
}

func (s *CreateMeTranslateExchangeIdRequestSourceIdType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateMeTranslateExchangeIdRequestSourceIdType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateMeTranslateExchangeIdRequestSourceIdType(input string) (*CreateMeTranslateExchangeIdRequestSourceIdType, error) {
	vals := map[string]CreateMeTranslateExchangeIdRequestSourceIdType{
		"entryid":              CreateMeTranslateExchangeIdRequestSourceIdTypeentryId,
		"ewsid":                CreateMeTranslateExchangeIdRequestSourceIdTypeewsId,
		"immutableentryid":     CreateMeTranslateExchangeIdRequestSourceIdTypeimmutableEntryId,
		"restid":               CreateMeTranslateExchangeIdRequestSourceIdTyperestId,
		"restimmutableentryid": CreateMeTranslateExchangeIdRequestSourceIdTyperestImmutableEntryId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateMeTranslateExchangeIdRequestSourceIdType(input)
	return &out, nil
}

type CreateMeTranslateExchangeIdRequestTargetIdType string

const (
	CreateMeTranslateExchangeIdRequestTargetIdTypeentryId              CreateMeTranslateExchangeIdRequestTargetIdType = "EntryId"
	CreateMeTranslateExchangeIdRequestTargetIdTypeewsId                CreateMeTranslateExchangeIdRequestTargetIdType = "EwsId"
	CreateMeTranslateExchangeIdRequestTargetIdTypeimmutableEntryId     CreateMeTranslateExchangeIdRequestTargetIdType = "ImmutableEntryId"
	CreateMeTranslateExchangeIdRequestTargetIdTyperestId               CreateMeTranslateExchangeIdRequestTargetIdType = "RestId"
	CreateMeTranslateExchangeIdRequestTargetIdTyperestImmutableEntryId CreateMeTranslateExchangeIdRequestTargetIdType = "RestImmutableEntryId"
)

func PossibleValuesForCreateMeTranslateExchangeIdRequestTargetIdType() []string {
	return []string{
		string(CreateMeTranslateExchangeIdRequestTargetIdTypeentryId),
		string(CreateMeTranslateExchangeIdRequestTargetIdTypeewsId),
		string(CreateMeTranslateExchangeIdRequestTargetIdTypeimmutableEntryId),
		string(CreateMeTranslateExchangeIdRequestTargetIdTyperestId),
		string(CreateMeTranslateExchangeIdRequestTargetIdTyperestImmutableEntryId),
	}
}

func (s *CreateMeTranslateExchangeIdRequestTargetIdType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateMeTranslateExchangeIdRequestTargetIdType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateMeTranslateExchangeIdRequestTargetIdType(input string) (*CreateMeTranslateExchangeIdRequestTargetIdType, error) {
	vals := map[string]CreateMeTranslateExchangeIdRequestTargetIdType{
		"entryid":              CreateMeTranslateExchangeIdRequestTargetIdTypeentryId,
		"ewsid":                CreateMeTranslateExchangeIdRequestTargetIdTypeewsId,
		"immutableentryid":     CreateMeTranslateExchangeIdRequestTargetIdTypeimmutableEntryId,
		"restid":               CreateMeTranslateExchangeIdRequestTargetIdTyperestId,
		"restimmutableentryid": CreateMeTranslateExchangeIdRequestTargetIdTyperestImmutableEntryId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateMeTranslateExchangeIdRequestTargetIdType(input)
	return &out, nil
}

type GetMeMailTipRequestMailTipsOptions string

const (
	GetMeMailTipRequestMailTipsOptionsautomaticReplies     GetMeMailTipRequestMailTipsOptions = "AutomaticReplies"
	GetMeMailTipRequestMailTipsOptionscustomMailTip        GetMeMailTipRequestMailTipsOptions = "CustomMailTip"
	GetMeMailTipRequestMailTipsOptionsdeliveryRestriction  GetMeMailTipRequestMailTipsOptions = "DeliveryRestriction"
	GetMeMailTipRequestMailTipsOptionsexternalMemberCount  GetMeMailTipRequestMailTipsOptions = "ExternalMemberCount"
	GetMeMailTipRequestMailTipsOptionsmailboxFullStatus    GetMeMailTipRequestMailTipsOptions = "MailboxFullStatus"
	GetMeMailTipRequestMailTipsOptionsmaxMessageSize       GetMeMailTipRequestMailTipsOptions = "MaxMessageSize"
	GetMeMailTipRequestMailTipsOptionsmoderationStatus     GetMeMailTipRequestMailTipsOptions = "ModerationStatus"
	GetMeMailTipRequestMailTipsOptionsrecipientScope       GetMeMailTipRequestMailTipsOptions = "RecipientScope"
	GetMeMailTipRequestMailTipsOptionsrecipientSuggestions GetMeMailTipRequestMailTipsOptions = "RecipientSuggestions"
	GetMeMailTipRequestMailTipsOptionstotalMemberCount     GetMeMailTipRequestMailTipsOptions = "TotalMemberCount"
)

func PossibleValuesForGetMeMailTipRequestMailTipsOptions() []string {
	return []string{
		string(GetMeMailTipRequestMailTipsOptionsautomaticReplies),
		string(GetMeMailTipRequestMailTipsOptionscustomMailTip),
		string(GetMeMailTipRequestMailTipsOptionsdeliveryRestriction),
		string(GetMeMailTipRequestMailTipsOptionsexternalMemberCount),
		string(GetMeMailTipRequestMailTipsOptionsmailboxFullStatus),
		string(GetMeMailTipRequestMailTipsOptionsmaxMessageSize),
		string(GetMeMailTipRequestMailTipsOptionsmoderationStatus),
		string(GetMeMailTipRequestMailTipsOptionsrecipientScope),
		string(GetMeMailTipRequestMailTipsOptionsrecipientSuggestions),
		string(GetMeMailTipRequestMailTipsOptionstotalMemberCount),
	}
}

func (s *GetMeMailTipRequestMailTipsOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGetMeMailTipRequestMailTipsOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGetMeMailTipRequestMailTipsOptions(input string) (*GetMeMailTipRequestMailTipsOptions, error) {
	vals := map[string]GetMeMailTipRequestMailTipsOptions{
		"automaticreplies":     GetMeMailTipRequestMailTipsOptionsautomaticReplies,
		"custommailtip":        GetMeMailTipRequestMailTipsOptionscustomMailTip,
		"deliveryrestriction":  GetMeMailTipRequestMailTipsOptionsdeliveryRestriction,
		"externalmembercount":  GetMeMailTipRequestMailTipsOptionsexternalMemberCount,
		"mailboxfullstatus":    GetMeMailTipRequestMailTipsOptionsmailboxFullStatus,
		"maxmessagesize":       GetMeMailTipRequestMailTipsOptionsmaxMessageSize,
		"moderationstatus":     GetMeMailTipRequestMailTipsOptionsmoderationStatus,
		"recipientscope":       GetMeMailTipRequestMailTipsOptionsrecipientScope,
		"recipientsuggestions": GetMeMailTipRequestMailTipsOptionsrecipientSuggestions,
		"totalmembercount":     GetMeMailTipRequestMailTipsOptionstotalMemberCount,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GetMeMailTipRequestMailTipsOptions(input)
	return &out, nil
}
