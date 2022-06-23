package automationrules

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomationRuleActionType string

const (
	AutomationRuleActionTypeModifyProperties AutomationRuleActionType = "ModifyProperties"
	AutomationRuleActionTypeRunPlaybook      AutomationRuleActionType = "RunPlaybook"
)

func PossibleValuesForAutomationRuleActionType() []string {
	return []string{
		string(AutomationRuleActionTypeModifyProperties),
		string(AutomationRuleActionTypeRunPlaybook),
	}
}

func parseAutomationRuleActionType(input string) (*AutomationRuleActionType, error) {
	vals := map[string]AutomationRuleActionType{
		"modifyproperties": AutomationRuleActionTypeModifyProperties,
		"runplaybook":      AutomationRuleActionTypeRunPlaybook,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutomationRuleActionType(input)
	return &out, nil
}

type AutomationRuleConditionType string

const (
	AutomationRuleConditionTypeProperty AutomationRuleConditionType = "Property"
)

func PossibleValuesForAutomationRuleConditionType() []string {
	return []string{
		string(AutomationRuleConditionTypeProperty),
	}
}

func parseAutomationRuleConditionType(input string) (*AutomationRuleConditionType, error) {
	vals := map[string]AutomationRuleConditionType{
		"property": AutomationRuleConditionTypeProperty,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutomationRuleConditionType(input)
	return &out, nil
}

type AutomationRulePropertyConditionSupportedOperator string

const (
	AutomationRulePropertyConditionSupportedOperatorContains      AutomationRulePropertyConditionSupportedOperator = "Contains"
	AutomationRulePropertyConditionSupportedOperatorEndsWith      AutomationRulePropertyConditionSupportedOperator = "EndsWith"
	AutomationRulePropertyConditionSupportedOperatorEquals        AutomationRulePropertyConditionSupportedOperator = "Equals"
	AutomationRulePropertyConditionSupportedOperatorNotContains   AutomationRulePropertyConditionSupportedOperator = "NotContains"
	AutomationRulePropertyConditionSupportedOperatorNotEndsWith   AutomationRulePropertyConditionSupportedOperator = "NotEndsWith"
	AutomationRulePropertyConditionSupportedOperatorNotEquals     AutomationRulePropertyConditionSupportedOperator = "NotEquals"
	AutomationRulePropertyConditionSupportedOperatorNotStartsWith AutomationRulePropertyConditionSupportedOperator = "NotStartsWith"
	AutomationRulePropertyConditionSupportedOperatorStartsWith    AutomationRulePropertyConditionSupportedOperator = "StartsWith"
)

func PossibleValuesForAutomationRulePropertyConditionSupportedOperator() []string {
	return []string{
		string(AutomationRulePropertyConditionSupportedOperatorContains),
		string(AutomationRulePropertyConditionSupportedOperatorEndsWith),
		string(AutomationRulePropertyConditionSupportedOperatorEquals),
		string(AutomationRulePropertyConditionSupportedOperatorNotContains),
		string(AutomationRulePropertyConditionSupportedOperatorNotEndsWith),
		string(AutomationRulePropertyConditionSupportedOperatorNotEquals),
		string(AutomationRulePropertyConditionSupportedOperatorNotStartsWith),
		string(AutomationRulePropertyConditionSupportedOperatorStartsWith),
	}
}

func parseAutomationRulePropertyConditionSupportedOperator(input string) (*AutomationRulePropertyConditionSupportedOperator, error) {
	vals := map[string]AutomationRulePropertyConditionSupportedOperator{
		"contains":      AutomationRulePropertyConditionSupportedOperatorContains,
		"endswith":      AutomationRulePropertyConditionSupportedOperatorEndsWith,
		"equals":        AutomationRulePropertyConditionSupportedOperatorEquals,
		"notcontains":   AutomationRulePropertyConditionSupportedOperatorNotContains,
		"notendswith":   AutomationRulePropertyConditionSupportedOperatorNotEndsWith,
		"notequals":     AutomationRulePropertyConditionSupportedOperatorNotEquals,
		"notstartswith": AutomationRulePropertyConditionSupportedOperatorNotStartsWith,
		"startswith":    AutomationRulePropertyConditionSupportedOperatorStartsWith,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutomationRulePropertyConditionSupportedOperator(input)
	return &out, nil
}

type AutomationRulePropertyConditionSupportedProperty string

const (
	AutomationRulePropertyConditionSupportedPropertyAccountAadTenantId             AutomationRulePropertyConditionSupportedProperty = "AccountAadTenantId"
	AutomationRulePropertyConditionSupportedPropertyAccountAadUserId               AutomationRulePropertyConditionSupportedProperty = "AccountAadUserId"
	AutomationRulePropertyConditionSupportedPropertyAccountNTDomain                AutomationRulePropertyConditionSupportedProperty = "AccountNTDomain"
	AutomationRulePropertyConditionSupportedPropertyAccountName                    AutomationRulePropertyConditionSupportedProperty = "AccountName"
	AutomationRulePropertyConditionSupportedPropertyAccountObjectGuid              AutomationRulePropertyConditionSupportedProperty = "AccountObjectGuid"
	AutomationRulePropertyConditionSupportedPropertyAccountPUID                    AutomationRulePropertyConditionSupportedProperty = "AccountPUID"
	AutomationRulePropertyConditionSupportedPropertyAccountSid                     AutomationRulePropertyConditionSupportedProperty = "AccountSid"
	AutomationRulePropertyConditionSupportedPropertyAccountUPNSuffix               AutomationRulePropertyConditionSupportedProperty = "AccountUPNSuffix"
	AutomationRulePropertyConditionSupportedPropertyAzureResourceResourceId        AutomationRulePropertyConditionSupportedProperty = "AzureResourceResourceId"
	AutomationRulePropertyConditionSupportedPropertyAzureResourceSubscriptionId    AutomationRulePropertyConditionSupportedProperty = "AzureResourceSubscriptionId"
	AutomationRulePropertyConditionSupportedPropertyCloudApplicationAppId          AutomationRulePropertyConditionSupportedProperty = "CloudApplicationAppId"
	AutomationRulePropertyConditionSupportedPropertyCloudApplicationAppName        AutomationRulePropertyConditionSupportedProperty = "CloudApplicationAppName"
	AutomationRulePropertyConditionSupportedPropertyDNSDomainName                  AutomationRulePropertyConditionSupportedProperty = "DNSDomainName"
	AutomationRulePropertyConditionSupportedPropertyFileDirectory                  AutomationRulePropertyConditionSupportedProperty = "FileDirectory"
	AutomationRulePropertyConditionSupportedPropertyFileHashValue                  AutomationRulePropertyConditionSupportedProperty = "FileHashValue"
	AutomationRulePropertyConditionSupportedPropertyFileName                       AutomationRulePropertyConditionSupportedProperty = "FileName"
	AutomationRulePropertyConditionSupportedPropertyHostAzureID                    AutomationRulePropertyConditionSupportedProperty = "HostAzureID"
	AutomationRulePropertyConditionSupportedPropertyHostNTDomain                   AutomationRulePropertyConditionSupportedProperty = "HostNTDomain"
	AutomationRulePropertyConditionSupportedPropertyHostName                       AutomationRulePropertyConditionSupportedProperty = "HostName"
	AutomationRulePropertyConditionSupportedPropertyHostNetBiosName                AutomationRulePropertyConditionSupportedProperty = "HostNetBiosName"
	AutomationRulePropertyConditionSupportedPropertyHostOSVersion                  AutomationRulePropertyConditionSupportedProperty = "HostOSVersion"
	AutomationRulePropertyConditionSupportedPropertyIPAddress                      AutomationRulePropertyConditionSupportedProperty = "IPAddress"
	AutomationRulePropertyConditionSupportedPropertyIncidentDescription            AutomationRulePropertyConditionSupportedProperty = "IncidentDescription"
	AutomationRulePropertyConditionSupportedPropertyIncidentProviderName           AutomationRulePropertyConditionSupportedProperty = "IncidentProviderName"
	AutomationRulePropertyConditionSupportedPropertyIncidentRelatedAnalyticRuleIds AutomationRulePropertyConditionSupportedProperty = "IncidentRelatedAnalyticRuleIds"
	AutomationRulePropertyConditionSupportedPropertyIncidentSeverity               AutomationRulePropertyConditionSupportedProperty = "IncidentSeverity"
	AutomationRulePropertyConditionSupportedPropertyIncidentStatus                 AutomationRulePropertyConditionSupportedProperty = "IncidentStatus"
	AutomationRulePropertyConditionSupportedPropertyIncidentTactics                AutomationRulePropertyConditionSupportedProperty = "IncidentTactics"
	AutomationRulePropertyConditionSupportedPropertyIncidentTitle                  AutomationRulePropertyConditionSupportedProperty = "IncidentTitle"
	AutomationRulePropertyConditionSupportedPropertyIoTDeviceId                    AutomationRulePropertyConditionSupportedProperty = "IoTDeviceId"
	AutomationRulePropertyConditionSupportedPropertyIoTDeviceModel                 AutomationRulePropertyConditionSupportedProperty = "IoTDeviceModel"
	AutomationRulePropertyConditionSupportedPropertyIoTDeviceName                  AutomationRulePropertyConditionSupportedProperty = "IoTDeviceName"
	AutomationRulePropertyConditionSupportedPropertyIoTDeviceOperatingSystem       AutomationRulePropertyConditionSupportedProperty = "IoTDeviceOperatingSystem"
	AutomationRulePropertyConditionSupportedPropertyIoTDeviceType                  AutomationRulePropertyConditionSupportedProperty = "IoTDeviceType"
	AutomationRulePropertyConditionSupportedPropertyIoTDeviceVendor                AutomationRulePropertyConditionSupportedProperty = "IoTDeviceVendor"
	AutomationRulePropertyConditionSupportedPropertyMailMessageDeliveryAction      AutomationRulePropertyConditionSupportedProperty = "MailMessageDeliveryAction"
	AutomationRulePropertyConditionSupportedPropertyMailMessageDeliveryLocation    AutomationRulePropertyConditionSupportedProperty = "MailMessageDeliveryLocation"
	AutomationRulePropertyConditionSupportedPropertyMailMessagePOneSender          AutomationRulePropertyConditionSupportedProperty = "MailMessageP1Sender"
	AutomationRulePropertyConditionSupportedPropertyMailMessagePTwoSender          AutomationRulePropertyConditionSupportedProperty = "MailMessageP2Sender"
	AutomationRulePropertyConditionSupportedPropertyMailMessageRecipient           AutomationRulePropertyConditionSupportedProperty = "MailMessageRecipient"
	AutomationRulePropertyConditionSupportedPropertyMailMessageSenderIP            AutomationRulePropertyConditionSupportedProperty = "MailMessageSenderIP"
	AutomationRulePropertyConditionSupportedPropertyMailMessageSubject             AutomationRulePropertyConditionSupportedProperty = "MailMessageSubject"
	AutomationRulePropertyConditionSupportedPropertyMailboxDisplayName             AutomationRulePropertyConditionSupportedProperty = "MailboxDisplayName"
	AutomationRulePropertyConditionSupportedPropertyMailboxPrimaryAddress          AutomationRulePropertyConditionSupportedProperty = "MailboxPrimaryAddress"
	AutomationRulePropertyConditionSupportedPropertyMailboxUPN                     AutomationRulePropertyConditionSupportedProperty = "MailboxUPN"
	AutomationRulePropertyConditionSupportedPropertyMalwareCategory                AutomationRulePropertyConditionSupportedProperty = "MalwareCategory"
	AutomationRulePropertyConditionSupportedPropertyMalwareName                    AutomationRulePropertyConditionSupportedProperty = "MalwareName"
	AutomationRulePropertyConditionSupportedPropertyProcessCommandLine             AutomationRulePropertyConditionSupportedProperty = "ProcessCommandLine"
	AutomationRulePropertyConditionSupportedPropertyProcessId                      AutomationRulePropertyConditionSupportedProperty = "ProcessId"
	AutomationRulePropertyConditionSupportedPropertyRegistryKey                    AutomationRulePropertyConditionSupportedProperty = "RegistryKey"
	AutomationRulePropertyConditionSupportedPropertyRegistryValueData              AutomationRulePropertyConditionSupportedProperty = "RegistryValueData"
	AutomationRulePropertyConditionSupportedPropertyUrl                            AutomationRulePropertyConditionSupportedProperty = "Url"
)

func PossibleValuesForAutomationRulePropertyConditionSupportedProperty() []string {
	return []string{
		string(AutomationRulePropertyConditionSupportedPropertyAccountAadTenantId),
		string(AutomationRulePropertyConditionSupportedPropertyAccountAadUserId),
		string(AutomationRulePropertyConditionSupportedPropertyAccountNTDomain),
		string(AutomationRulePropertyConditionSupportedPropertyAccountName),
		string(AutomationRulePropertyConditionSupportedPropertyAccountObjectGuid),
		string(AutomationRulePropertyConditionSupportedPropertyAccountPUID),
		string(AutomationRulePropertyConditionSupportedPropertyAccountSid),
		string(AutomationRulePropertyConditionSupportedPropertyAccountUPNSuffix),
		string(AutomationRulePropertyConditionSupportedPropertyAzureResourceResourceId),
		string(AutomationRulePropertyConditionSupportedPropertyAzureResourceSubscriptionId),
		string(AutomationRulePropertyConditionSupportedPropertyCloudApplicationAppId),
		string(AutomationRulePropertyConditionSupportedPropertyCloudApplicationAppName),
		string(AutomationRulePropertyConditionSupportedPropertyDNSDomainName),
		string(AutomationRulePropertyConditionSupportedPropertyFileDirectory),
		string(AutomationRulePropertyConditionSupportedPropertyFileHashValue),
		string(AutomationRulePropertyConditionSupportedPropertyFileName),
		string(AutomationRulePropertyConditionSupportedPropertyHostAzureID),
		string(AutomationRulePropertyConditionSupportedPropertyHostNTDomain),
		string(AutomationRulePropertyConditionSupportedPropertyHostName),
		string(AutomationRulePropertyConditionSupportedPropertyHostNetBiosName),
		string(AutomationRulePropertyConditionSupportedPropertyHostOSVersion),
		string(AutomationRulePropertyConditionSupportedPropertyIPAddress),
		string(AutomationRulePropertyConditionSupportedPropertyIncidentDescription),
		string(AutomationRulePropertyConditionSupportedPropertyIncidentProviderName),
		string(AutomationRulePropertyConditionSupportedPropertyIncidentRelatedAnalyticRuleIds),
		string(AutomationRulePropertyConditionSupportedPropertyIncidentSeverity),
		string(AutomationRulePropertyConditionSupportedPropertyIncidentStatus),
		string(AutomationRulePropertyConditionSupportedPropertyIncidentTactics),
		string(AutomationRulePropertyConditionSupportedPropertyIncidentTitle),
		string(AutomationRulePropertyConditionSupportedPropertyIoTDeviceId),
		string(AutomationRulePropertyConditionSupportedPropertyIoTDeviceModel),
		string(AutomationRulePropertyConditionSupportedPropertyIoTDeviceName),
		string(AutomationRulePropertyConditionSupportedPropertyIoTDeviceOperatingSystem),
		string(AutomationRulePropertyConditionSupportedPropertyIoTDeviceType),
		string(AutomationRulePropertyConditionSupportedPropertyIoTDeviceVendor),
		string(AutomationRulePropertyConditionSupportedPropertyMailMessageDeliveryAction),
		string(AutomationRulePropertyConditionSupportedPropertyMailMessageDeliveryLocation),
		string(AutomationRulePropertyConditionSupportedPropertyMailMessagePOneSender),
		string(AutomationRulePropertyConditionSupportedPropertyMailMessagePTwoSender),
		string(AutomationRulePropertyConditionSupportedPropertyMailMessageRecipient),
		string(AutomationRulePropertyConditionSupportedPropertyMailMessageSenderIP),
		string(AutomationRulePropertyConditionSupportedPropertyMailMessageSubject),
		string(AutomationRulePropertyConditionSupportedPropertyMailboxDisplayName),
		string(AutomationRulePropertyConditionSupportedPropertyMailboxPrimaryAddress),
		string(AutomationRulePropertyConditionSupportedPropertyMailboxUPN),
		string(AutomationRulePropertyConditionSupportedPropertyMalwareCategory),
		string(AutomationRulePropertyConditionSupportedPropertyMalwareName),
		string(AutomationRulePropertyConditionSupportedPropertyProcessCommandLine),
		string(AutomationRulePropertyConditionSupportedPropertyProcessId),
		string(AutomationRulePropertyConditionSupportedPropertyRegistryKey),
		string(AutomationRulePropertyConditionSupportedPropertyRegistryValueData),
		string(AutomationRulePropertyConditionSupportedPropertyUrl),
	}
}

func parseAutomationRulePropertyConditionSupportedProperty(input string) (*AutomationRulePropertyConditionSupportedProperty, error) {
	vals := map[string]AutomationRulePropertyConditionSupportedProperty{
		"accountaadtenantid":             AutomationRulePropertyConditionSupportedPropertyAccountAadTenantId,
		"accountaaduserid":               AutomationRulePropertyConditionSupportedPropertyAccountAadUserId,
		"accountntdomain":                AutomationRulePropertyConditionSupportedPropertyAccountNTDomain,
		"accountname":                    AutomationRulePropertyConditionSupportedPropertyAccountName,
		"accountobjectguid":              AutomationRulePropertyConditionSupportedPropertyAccountObjectGuid,
		"accountpuid":                    AutomationRulePropertyConditionSupportedPropertyAccountPUID,
		"accountsid":                     AutomationRulePropertyConditionSupportedPropertyAccountSid,
		"accountupnsuffix":               AutomationRulePropertyConditionSupportedPropertyAccountUPNSuffix,
		"azureresourceresourceid":        AutomationRulePropertyConditionSupportedPropertyAzureResourceResourceId,
		"azureresourcesubscriptionid":    AutomationRulePropertyConditionSupportedPropertyAzureResourceSubscriptionId,
		"cloudapplicationappid":          AutomationRulePropertyConditionSupportedPropertyCloudApplicationAppId,
		"cloudapplicationappname":        AutomationRulePropertyConditionSupportedPropertyCloudApplicationAppName,
		"dnsdomainname":                  AutomationRulePropertyConditionSupportedPropertyDNSDomainName,
		"filedirectory":                  AutomationRulePropertyConditionSupportedPropertyFileDirectory,
		"filehashvalue":                  AutomationRulePropertyConditionSupportedPropertyFileHashValue,
		"filename":                       AutomationRulePropertyConditionSupportedPropertyFileName,
		"hostazureid":                    AutomationRulePropertyConditionSupportedPropertyHostAzureID,
		"hostntdomain":                   AutomationRulePropertyConditionSupportedPropertyHostNTDomain,
		"hostname":                       AutomationRulePropertyConditionSupportedPropertyHostName,
		"hostnetbiosname":                AutomationRulePropertyConditionSupportedPropertyHostNetBiosName,
		"hostosversion":                  AutomationRulePropertyConditionSupportedPropertyHostOSVersion,
		"ipaddress":                      AutomationRulePropertyConditionSupportedPropertyIPAddress,
		"incidentdescription":            AutomationRulePropertyConditionSupportedPropertyIncidentDescription,
		"incidentprovidername":           AutomationRulePropertyConditionSupportedPropertyIncidentProviderName,
		"incidentrelatedanalyticruleids": AutomationRulePropertyConditionSupportedPropertyIncidentRelatedAnalyticRuleIds,
		"incidentseverity":               AutomationRulePropertyConditionSupportedPropertyIncidentSeverity,
		"incidentstatus":                 AutomationRulePropertyConditionSupportedPropertyIncidentStatus,
		"incidenttactics":                AutomationRulePropertyConditionSupportedPropertyIncidentTactics,
		"incidenttitle":                  AutomationRulePropertyConditionSupportedPropertyIncidentTitle,
		"iotdeviceid":                    AutomationRulePropertyConditionSupportedPropertyIoTDeviceId,
		"iotdevicemodel":                 AutomationRulePropertyConditionSupportedPropertyIoTDeviceModel,
		"iotdevicename":                  AutomationRulePropertyConditionSupportedPropertyIoTDeviceName,
		"iotdeviceoperatingsystem":       AutomationRulePropertyConditionSupportedPropertyIoTDeviceOperatingSystem,
		"iotdevicetype":                  AutomationRulePropertyConditionSupportedPropertyIoTDeviceType,
		"iotdevicevendor":                AutomationRulePropertyConditionSupportedPropertyIoTDeviceVendor,
		"mailmessagedeliveryaction":      AutomationRulePropertyConditionSupportedPropertyMailMessageDeliveryAction,
		"mailmessagedeliverylocation":    AutomationRulePropertyConditionSupportedPropertyMailMessageDeliveryLocation,
		"mailmessagep1sender":            AutomationRulePropertyConditionSupportedPropertyMailMessagePOneSender,
		"mailmessagep2sender":            AutomationRulePropertyConditionSupportedPropertyMailMessagePTwoSender,
		"mailmessagerecipient":           AutomationRulePropertyConditionSupportedPropertyMailMessageRecipient,
		"mailmessagesenderip":            AutomationRulePropertyConditionSupportedPropertyMailMessageSenderIP,
		"mailmessagesubject":             AutomationRulePropertyConditionSupportedPropertyMailMessageSubject,
		"mailboxdisplayname":             AutomationRulePropertyConditionSupportedPropertyMailboxDisplayName,
		"mailboxprimaryaddress":          AutomationRulePropertyConditionSupportedPropertyMailboxPrimaryAddress,
		"mailboxupn":                     AutomationRulePropertyConditionSupportedPropertyMailboxUPN,
		"malwarecategory":                AutomationRulePropertyConditionSupportedPropertyMalwareCategory,
		"malwarename":                    AutomationRulePropertyConditionSupportedPropertyMalwareName,
		"processcommandline":             AutomationRulePropertyConditionSupportedPropertyProcessCommandLine,
		"processid":                      AutomationRulePropertyConditionSupportedPropertyProcessId,
		"registrykey":                    AutomationRulePropertyConditionSupportedPropertyRegistryKey,
		"registryvaluedata":              AutomationRulePropertyConditionSupportedPropertyRegistryValueData,
		"url":                            AutomationRulePropertyConditionSupportedPropertyUrl,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutomationRulePropertyConditionSupportedProperty(input)
	return &out, nil
}

type IncidentClassification string

const (
	IncidentClassificationBenignPositive IncidentClassification = "BenignPositive"
	IncidentClassificationFalsePositive  IncidentClassification = "FalsePositive"
	IncidentClassificationTruePositive   IncidentClassification = "TruePositive"
	IncidentClassificationUndetermined   IncidentClassification = "Undetermined"
)

func PossibleValuesForIncidentClassification() []string {
	return []string{
		string(IncidentClassificationBenignPositive),
		string(IncidentClassificationFalsePositive),
		string(IncidentClassificationTruePositive),
		string(IncidentClassificationUndetermined),
	}
}

func parseIncidentClassification(input string) (*IncidentClassification, error) {
	vals := map[string]IncidentClassification{
		"benignpositive": IncidentClassificationBenignPositive,
		"falsepositive":  IncidentClassificationFalsePositive,
		"truepositive":   IncidentClassificationTruePositive,
		"undetermined":   IncidentClassificationUndetermined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IncidentClassification(input)
	return &out, nil
}

type IncidentClassificationReason string

const (
	IncidentClassificationReasonInaccurateData        IncidentClassificationReason = "InaccurateData"
	IncidentClassificationReasonIncorrectAlertLogic   IncidentClassificationReason = "IncorrectAlertLogic"
	IncidentClassificationReasonSuspiciousActivity    IncidentClassificationReason = "SuspiciousActivity"
	IncidentClassificationReasonSuspiciousButExpected IncidentClassificationReason = "SuspiciousButExpected"
)

func PossibleValuesForIncidentClassificationReason() []string {
	return []string{
		string(IncidentClassificationReasonInaccurateData),
		string(IncidentClassificationReasonIncorrectAlertLogic),
		string(IncidentClassificationReasonSuspiciousActivity),
		string(IncidentClassificationReasonSuspiciousButExpected),
	}
}

func parseIncidentClassificationReason(input string) (*IncidentClassificationReason, error) {
	vals := map[string]IncidentClassificationReason{
		"inaccuratedata":        IncidentClassificationReasonInaccurateData,
		"incorrectalertlogic":   IncidentClassificationReasonIncorrectAlertLogic,
		"suspiciousactivity":    IncidentClassificationReasonSuspiciousActivity,
		"suspiciousbutexpected": IncidentClassificationReasonSuspiciousButExpected,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IncidentClassificationReason(input)
	return &out, nil
}

type IncidentLabelType string

const (
	IncidentLabelTypeSystem IncidentLabelType = "System"
	IncidentLabelTypeUser   IncidentLabelType = "User"
)

func PossibleValuesForIncidentLabelType() []string {
	return []string{
		string(IncidentLabelTypeSystem),
		string(IncidentLabelTypeUser),
	}
}

func parseIncidentLabelType(input string) (*IncidentLabelType, error) {
	vals := map[string]IncidentLabelType{
		"system": IncidentLabelTypeSystem,
		"user":   IncidentLabelTypeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IncidentLabelType(input)
	return &out, nil
}

type IncidentSeverity string

const (
	IncidentSeverityHigh          IncidentSeverity = "High"
	IncidentSeverityInformational IncidentSeverity = "Informational"
	IncidentSeverityLow           IncidentSeverity = "Low"
	IncidentSeverityMedium        IncidentSeverity = "Medium"
)

func PossibleValuesForIncidentSeverity() []string {
	return []string{
		string(IncidentSeverityHigh),
		string(IncidentSeverityInformational),
		string(IncidentSeverityLow),
		string(IncidentSeverityMedium),
	}
}

func parseIncidentSeverity(input string) (*IncidentSeverity, error) {
	vals := map[string]IncidentSeverity{
		"high":          IncidentSeverityHigh,
		"informational": IncidentSeverityInformational,
		"low":           IncidentSeverityLow,
		"medium":        IncidentSeverityMedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IncidentSeverity(input)
	return &out, nil
}

type IncidentStatus string

const (
	IncidentStatusActive IncidentStatus = "Active"
	IncidentStatusClosed IncidentStatus = "Closed"
	IncidentStatusNew    IncidentStatus = "New"
)

func PossibleValuesForIncidentStatus() []string {
	return []string{
		string(IncidentStatusActive),
		string(IncidentStatusClosed),
		string(IncidentStatusNew),
	}
}

func parseIncidentStatus(input string) (*IncidentStatus, error) {
	vals := map[string]IncidentStatus{
		"active": IncidentStatusActive,
		"closed": IncidentStatusClosed,
		"new":    IncidentStatusNew,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IncidentStatus(input)
	return &out, nil
}

type OwnerType string

const (
	OwnerTypeGroup   OwnerType = "Group"
	OwnerTypeUnknown OwnerType = "Unknown"
	OwnerTypeUser    OwnerType = "User"
)

func PossibleValuesForOwnerType() []string {
	return []string{
		string(OwnerTypeGroup),
		string(OwnerTypeUnknown),
		string(OwnerTypeUser),
	}
}

func parseOwnerType(input string) (*OwnerType, error) {
	vals := map[string]OwnerType{
		"group":   OwnerTypeGroup,
		"unknown": OwnerTypeUnknown,
		"user":    OwnerTypeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OwnerType(input)
	return &out, nil
}

type TriggersOn string

const (
	TriggersOnIncidents TriggersOn = "Incidents"
)

func PossibleValuesForTriggersOn() []string {
	return []string{
		string(TriggersOnIncidents),
	}
}

func parseTriggersOn(input string) (*TriggersOn, error) {
	vals := map[string]TriggersOn{
		"incidents": TriggersOnIncidents,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TriggersOn(input)
	return &out, nil
}

type TriggersWhen string

const (
	TriggersWhenCreated TriggersWhen = "Created"
)

func PossibleValuesForTriggersWhen() []string {
	return []string{
		string(TriggersWhenCreated),
	}
}

func parseTriggersWhen(input string) (*TriggersWhen, error) {
	vals := map[string]TriggersWhen{
		"created": TriggersWhenCreated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TriggersWhen(input)
	return &out, nil
}
