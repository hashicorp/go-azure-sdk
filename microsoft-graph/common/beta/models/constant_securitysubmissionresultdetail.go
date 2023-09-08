package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySubmissionResultDetail string

const (
	SecuritySubmissionResultDetailallowedByConnection                    SecuritySubmissionResultDetail = "AllowedByConnection"
	SecuritySubmissionResultDetailallowedByExchangeTransportRule         SecuritySubmissionResultDetail = "AllowedByExchangeTransportRule"
	SecuritySubmissionResultDetailallowedBySecOps                        SecuritySubmissionResultDetail = "AllowedBySecOps"
	SecuritySubmissionResultDetailallowedByTenant                        SecuritySubmissionResultDetail = "AllowedByTenant"
	SecuritySubmissionResultDetailallowedByTenantAllowBlockList          SecuritySubmissionResultDetail = "AllowedByTenantAllowBlockList"
	SecuritySubmissionResultDetailallowedByThirdPartyFilters             SecuritySubmissionResultDetail = "AllowedByThirdPartyFilters"
	SecuritySubmissionResultDetailallowedByUserSetting                   SecuritySubmissionResultDetail = "AllowedByUserSetting"
	SecuritySubmissionResultDetailallowedFileByTenantAllowBlockList      SecuritySubmissionResultDetail = "AllowedFileByTenantAllowBlockList"
	SecuritySubmissionResultDetailallowedRecipientByTenantAllowBlockList SecuritySubmissionResultDetail = "AllowedRecipientByTenantAllowBlockList"
	SecuritySubmissionResultDetailallowedSenderByTenantAllowBlockList    SecuritySubmissionResultDetail = "AllowedSenderByTenantAllowBlockList"
	SecuritySubmissionResultDetailallowedUrlByTenantAllowBlockList       SecuritySubmissionResultDetail = "AllowedUrlByTenantAllowBlockList"
	SecuritySubmissionResultDetailbadReclassifiedAsBad                   SecuritySubmissionResultDetail = "BadReclassifiedAsBad"
	SecuritySubmissionResultDetailbadReclassifiedAsBulk                  SecuritySubmissionResultDetail = "BadReclassifiedAsBulk"
	SecuritySubmissionResultDetailbadReclassifiedAsCannotMakeDecision    SecuritySubmissionResultDetail = "BadReclassifiedAsCannotMakeDecision"
	SecuritySubmissionResultDetailbadReclassifiedAsGood                  SecuritySubmissionResultDetail = "BadReclassifiedAsGood"
	SecuritySubmissionResultDetailblockedByConnection                    SecuritySubmissionResultDetail = "BlockedByConnection"
	SecuritySubmissionResultDetailblockedByExchangeTransportRule         SecuritySubmissionResultDetail = "BlockedByExchangeTransportRule"
	SecuritySubmissionResultDetailblockedByTenant                        SecuritySubmissionResultDetail = "BlockedByTenant"
	SecuritySubmissionResultDetailblockedByTenantAllowBlockList          SecuritySubmissionResultDetail = "BlockedByTenantAllowBlockList"
	SecuritySubmissionResultDetailblockedByUserSetting                   SecuritySubmissionResultDetail = "BlockedByUserSetting"
	SecuritySubmissionResultDetailblockedFileByTenantAllowBlockList      SecuritySubmissionResultDetail = "BlockedFileByTenantAllowBlockList"
	SecuritySubmissionResultDetailblockedRecipientByTenantAllowBlockList SecuritySubmissionResultDetail = "BlockedRecipientByTenantAllowBlockList"
	SecuritySubmissionResultDetailblockedSenderByTenantAllowBlockList    SecuritySubmissionResultDetail = "BlockedSenderByTenantAllowBlockList"
	SecuritySubmissionResultDetailblockedUrlByTenantAllowBlockList       SecuritySubmissionResultDetail = "BlockedUrlByTenantAllowBlockList"
	SecuritySubmissionResultDetailbrandImpersonation                     SecuritySubmissionResultDetail = "BrandImpersonation"
	SecuritySubmissionResultDetaildomainImpersonation                    SecuritySubmissionResultDetail = "DomainImpersonation"
	SecuritySubmissionResultDetailgoodReclassifiedAsBad                  SecuritySubmissionResultDetail = "GoodReclassifiedAsBad"
	SecuritySubmissionResultDetailgoodReclassifiedAsBulk                 SecuritySubmissionResultDetail = "GoodReclassifiedAsBulk"
	SecuritySubmissionResultDetailgoodReclassifiedAsCannotMakeDecision   SecuritySubmissionResultDetail = "GoodReclassifiedAsCannotMakeDecision"
	SecuritySubmissionResultDetailgoodReclassifiedAsGood                 SecuritySubmissionResultDetail = "GoodReclassifiedAsGood"
	SecuritySubmissionResultDetailinvalidFalseNegative                   SecuritySubmissionResultDetail = "InvalidFalseNegative"
	SecuritySubmissionResultDetailinvalidFalsePositive                   SecuritySubmissionResultDetail = "InvalidFalsePositive"
	SecuritySubmissionResultDetailjunkMailRuleDisabled                   SecuritySubmissionResultDetail = "JunkMailRuleDisabled"
	SecuritySubmissionResultDetailmessageNotFound                        SecuritySubmissionResultDetail = "MessageNotFound"
	SecuritySubmissionResultDetailnone                                   SecuritySubmissionResultDetail = "None"
	SecuritySubmissionResultDetailonPremisesSkip                         SecuritySubmissionResultDetail = "OnPremisesSkip"
	SecuritySubmissionResultDetailoutboundBulk                           SecuritySubmissionResultDetail = "OutboundBulk"
	SecuritySubmissionResultDetailoutboundCannotMakeDecision             SecuritySubmissionResultDetail = "OutboundCannotMakeDecision"
	SecuritySubmissionResultDetailoutboundNotRescanned                   SecuritySubmissionResultDetail = "OutboundNotRescanned"
	SecuritySubmissionResultDetailoutboundShouldBeBlocked                SecuritySubmissionResultDetail = "OutboundShouldBeBlocked"
	SecuritySubmissionResultDetailoutboundShouldNotBeBlocked             SecuritySubmissionResultDetail = "OutboundShouldNotBeBlocked"
	SecuritySubmissionResultDetailquarantineReleased                     SecuritySubmissionResultDetail = "QuarantineReleased"
	SecuritySubmissionResultDetailquarantineReleasedThenBlocked          SecuritySubmissionResultDetail = "QuarantineReleasedThenBlocked"
	SecuritySubmissionResultDetailsimulatedThreat                        SecuritySubmissionResultDetail = "SimulatedThreat"
	SecuritySubmissionResultDetailspoofBlocked                           SecuritySubmissionResultDetail = "SpoofBlocked"
	SecuritySubmissionResultDetailunderInvestigation                     SecuritySubmissionResultDetail = "UnderInvestigation"
	SecuritySubmissionResultDetailurlFileCannotMakeDecision              SecuritySubmissionResultDetail = "UrlFileCannotMakeDecision"
	SecuritySubmissionResultDetailurlFileShouldBeBlocked                 SecuritySubmissionResultDetail = "UrlFileShouldBeBlocked"
	SecuritySubmissionResultDetailurlFileShouldNotBeBlocked              SecuritySubmissionResultDetail = "UrlFileShouldNotBeBlocked"
	SecuritySubmissionResultDetailuserImpersonation                      SecuritySubmissionResultDetail = "UserImpersonation"
	SecuritySubmissionResultDetailzeroHourAutoPurgeAllowed               SecuritySubmissionResultDetail = "ZeroHourAutoPurgeAllowed"
	SecuritySubmissionResultDetailzeroHourAutoPurgeBlocked               SecuritySubmissionResultDetail = "ZeroHourAutoPurgeBlocked"
	SecuritySubmissionResultDetailzeroHourAutoPurgeQuarantineReleased    SecuritySubmissionResultDetail = "ZeroHourAutoPurgeQuarantineReleased"
)

func PossibleValuesForSecuritySubmissionResultDetail() []string {
	return []string{
		string(SecuritySubmissionResultDetailallowedByConnection),
		string(SecuritySubmissionResultDetailallowedByExchangeTransportRule),
		string(SecuritySubmissionResultDetailallowedBySecOps),
		string(SecuritySubmissionResultDetailallowedByTenant),
		string(SecuritySubmissionResultDetailallowedByTenantAllowBlockList),
		string(SecuritySubmissionResultDetailallowedByThirdPartyFilters),
		string(SecuritySubmissionResultDetailallowedByUserSetting),
		string(SecuritySubmissionResultDetailallowedFileByTenantAllowBlockList),
		string(SecuritySubmissionResultDetailallowedRecipientByTenantAllowBlockList),
		string(SecuritySubmissionResultDetailallowedSenderByTenantAllowBlockList),
		string(SecuritySubmissionResultDetailallowedUrlByTenantAllowBlockList),
		string(SecuritySubmissionResultDetailbadReclassifiedAsBad),
		string(SecuritySubmissionResultDetailbadReclassifiedAsBulk),
		string(SecuritySubmissionResultDetailbadReclassifiedAsCannotMakeDecision),
		string(SecuritySubmissionResultDetailbadReclassifiedAsGood),
		string(SecuritySubmissionResultDetailblockedByConnection),
		string(SecuritySubmissionResultDetailblockedByExchangeTransportRule),
		string(SecuritySubmissionResultDetailblockedByTenant),
		string(SecuritySubmissionResultDetailblockedByTenantAllowBlockList),
		string(SecuritySubmissionResultDetailblockedByUserSetting),
		string(SecuritySubmissionResultDetailblockedFileByTenantAllowBlockList),
		string(SecuritySubmissionResultDetailblockedRecipientByTenantAllowBlockList),
		string(SecuritySubmissionResultDetailblockedSenderByTenantAllowBlockList),
		string(SecuritySubmissionResultDetailblockedUrlByTenantAllowBlockList),
		string(SecuritySubmissionResultDetailbrandImpersonation),
		string(SecuritySubmissionResultDetaildomainImpersonation),
		string(SecuritySubmissionResultDetailgoodReclassifiedAsBad),
		string(SecuritySubmissionResultDetailgoodReclassifiedAsBulk),
		string(SecuritySubmissionResultDetailgoodReclassifiedAsCannotMakeDecision),
		string(SecuritySubmissionResultDetailgoodReclassifiedAsGood),
		string(SecuritySubmissionResultDetailinvalidFalseNegative),
		string(SecuritySubmissionResultDetailinvalidFalsePositive),
		string(SecuritySubmissionResultDetailjunkMailRuleDisabled),
		string(SecuritySubmissionResultDetailmessageNotFound),
		string(SecuritySubmissionResultDetailnone),
		string(SecuritySubmissionResultDetailonPremisesSkip),
		string(SecuritySubmissionResultDetailoutboundBulk),
		string(SecuritySubmissionResultDetailoutboundCannotMakeDecision),
		string(SecuritySubmissionResultDetailoutboundNotRescanned),
		string(SecuritySubmissionResultDetailoutboundShouldBeBlocked),
		string(SecuritySubmissionResultDetailoutboundShouldNotBeBlocked),
		string(SecuritySubmissionResultDetailquarantineReleased),
		string(SecuritySubmissionResultDetailquarantineReleasedThenBlocked),
		string(SecuritySubmissionResultDetailsimulatedThreat),
		string(SecuritySubmissionResultDetailspoofBlocked),
		string(SecuritySubmissionResultDetailunderInvestigation),
		string(SecuritySubmissionResultDetailurlFileCannotMakeDecision),
		string(SecuritySubmissionResultDetailurlFileShouldBeBlocked),
		string(SecuritySubmissionResultDetailurlFileShouldNotBeBlocked),
		string(SecuritySubmissionResultDetailuserImpersonation),
		string(SecuritySubmissionResultDetailzeroHourAutoPurgeAllowed),
		string(SecuritySubmissionResultDetailzeroHourAutoPurgeBlocked),
		string(SecuritySubmissionResultDetailzeroHourAutoPurgeQuarantineReleased),
	}
}

func parseSecuritySubmissionResultDetail(input string) (*SecuritySubmissionResultDetail, error) {
	vals := map[string]SecuritySubmissionResultDetail{
		"allowedbyconnection":                    SecuritySubmissionResultDetailallowedByConnection,
		"allowedbyexchangetransportrule":         SecuritySubmissionResultDetailallowedByExchangeTransportRule,
		"allowedbysecops":                        SecuritySubmissionResultDetailallowedBySecOps,
		"allowedbytenant":                        SecuritySubmissionResultDetailallowedByTenant,
		"allowedbytenantallowblocklist":          SecuritySubmissionResultDetailallowedByTenantAllowBlockList,
		"allowedbythirdpartyfilters":             SecuritySubmissionResultDetailallowedByThirdPartyFilters,
		"allowedbyusersetting":                   SecuritySubmissionResultDetailallowedByUserSetting,
		"allowedfilebytenantallowblocklist":      SecuritySubmissionResultDetailallowedFileByTenantAllowBlockList,
		"allowedrecipientbytenantallowblocklist": SecuritySubmissionResultDetailallowedRecipientByTenantAllowBlockList,
		"allowedsenderbytenantallowblocklist":    SecuritySubmissionResultDetailallowedSenderByTenantAllowBlockList,
		"allowedurlbytenantallowblocklist":       SecuritySubmissionResultDetailallowedUrlByTenantAllowBlockList,
		"badreclassifiedasbad":                   SecuritySubmissionResultDetailbadReclassifiedAsBad,
		"badreclassifiedasbulk":                  SecuritySubmissionResultDetailbadReclassifiedAsBulk,
		"badreclassifiedascannotmakedecision":    SecuritySubmissionResultDetailbadReclassifiedAsCannotMakeDecision,
		"badreclassifiedasgood":                  SecuritySubmissionResultDetailbadReclassifiedAsGood,
		"blockedbyconnection":                    SecuritySubmissionResultDetailblockedByConnection,
		"blockedbyexchangetransportrule":         SecuritySubmissionResultDetailblockedByExchangeTransportRule,
		"blockedbytenant":                        SecuritySubmissionResultDetailblockedByTenant,
		"blockedbytenantallowblocklist":          SecuritySubmissionResultDetailblockedByTenantAllowBlockList,
		"blockedbyusersetting":                   SecuritySubmissionResultDetailblockedByUserSetting,
		"blockedfilebytenantallowblocklist":      SecuritySubmissionResultDetailblockedFileByTenantAllowBlockList,
		"blockedrecipientbytenantallowblocklist": SecuritySubmissionResultDetailblockedRecipientByTenantAllowBlockList,
		"blockedsenderbytenantallowblocklist":    SecuritySubmissionResultDetailblockedSenderByTenantAllowBlockList,
		"blockedurlbytenantallowblocklist":       SecuritySubmissionResultDetailblockedUrlByTenantAllowBlockList,
		"brandimpersonation":                     SecuritySubmissionResultDetailbrandImpersonation,
		"domainimpersonation":                    SecuritySubmissionResultDetaildomainImpersonation,
		"goodreclassifiedasbad":                  SecuritySubmissionResultDetailgoodReclassifiedAsBad,
		"goodreclassifiedasbulk":                 SecuritySubmissionResultDetailgoodReclassifiedAsBulk,
		"goodreclassifiedascannotmakedecision":   SecuritySubmissionResultDetailgoodReclassifiedAsCannotMakeDecision,
		"goodreclassifiedasgood":                 SecuritySubmissionResultDetailgoodReclassifiedAsGood,
		"invalidfalsenegative":                   SecuritySubmissionResultDetailinvalidFalseNegative,
		"invalidfalsepositive":                   SecuritySubmissionResultDetailinvalidFalsePositive,
		"junkmailruledisabled":                   SecuritySubmissionResultDetailjunkMailRuleDisabled,
		"messagenotfound":                        SecuritySubmissionResultDetailmessageNotFound,
		"none":                                   SecuritySubmissionResultDetailnone,
		"onpremisesskip":                         SecuritySubmissionResultDetailonPremisesSkip,
		"outboundbulk":                           SecuritySubmissionResultDetailoutboundBulk,
		"outboundcannotmakedecision":             SecuritySubmissionResultDetailoutboundCannotMakeDecision,
		"outboundnotrescanned":                   SecuritySubmissionResultDetailoutboundNotRescanned,
		"outboundshouldbeblocked":                SecuritySubmissionResultDetailoutboundShouldBeBlocked,
		"outboundshouldnotbeblocked":             SecuritySubmissionResultDetailoutboundShouldNotBeBlocked,
		"quarantinereleased":                     SecuritySubmissionResultDetailquarantineReleased,
		"quarantinereleasedthenblocked":          SecuritySubmissionResultDetailquarantineReleasedThenBlocked,
		"simulatedthreat":                        SecuritySubmissionResultDetailsimulatedThreat,
		"spoofblocked":                           SecuritySubmissionResultDetailspoofBlocked,
		"underinvestigation":                     SecuritySubmissionResultDetailunderInvestigation,
		"urlfilecannotmakedecision":              SecuritySubmissionResultDetailurlFileCannotMakeDecision,
		"urlfileshouldbeblocked":                 SecuritySubmissionResultDetailurlFileShouldBeBlocked,
		"urlfileshouldnotbeblocked":              SecuritySubmissionResultDetailurlFileShouldNotBeBlocked,
		"userimpersonation":                      SecuritySubmissionResultDetailuserImpersonation,
		"zerohourautopurgeallowed":               SecuritySubmissionResultDetailzeroHourAutoPurgeAllowed,
		"zerohourautopurgeblocked":               SecuritySubmissionResultDetailzeroHourAutoPurgeBlocked,
		"zerohourautopurgequarantinereleased":    SecuritySubmissionResultDetailzeroHourAutoPurgeQuarantineReleased,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySubmissionResultDetail(input)
	return &out, nil
}
