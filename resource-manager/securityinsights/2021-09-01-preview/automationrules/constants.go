package automationrules

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

type AutomationRuleConditionType string

const (
	AutomationRuleConditionTypeProperty AutomationRuleConditionType = "Property"
)

func PossibleValuesForAutomationRuleConditionType() []string {
	return []string{
		string(AutomationRuleConditionTypeProperty),
	}
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

type TriggersOn string

const (
	TriggersOnIncidents TriggersOn = "Incidents"
)

func PossibleValuesForTriggersOn() []string {
	return []string{
		string(TriggersOnIncidents),
	}
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
