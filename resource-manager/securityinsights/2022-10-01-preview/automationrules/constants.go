package automationrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActionType string

const (
	ActionTypeModifyProperties ActionType = "ModifyProperties"
	ActionTypeRunPlaybook      ActionType = "RunPlaybook"
)

func PossibleValuesForActionType() []string {
	return []string{
		string(ActionTypeModifyProperties),
		string(ActionTypeRunPlaybook),
	}
}

type AutomationRuleBooleanConditionSupportedOperator string

const (
	AutomationRuleBooleanConditionSupportedOperatorAnd AutomationRuleBooleanConditionSupportedOperator = "And"
	AutomationRuleBooleanConditionSupportedOperatorOr  AutomationRuleBooleanConditionSupportedOperator = "Or"
)

func PossibleValuesForAutomationRuleBooleanConditionSupportedOperator() []string {
	return []string{
		string(AutomationRuleBooleanConditionSupportedOperatorAnd),
		string(AutomationRuleBooleanConditionSupportedOperatorOr),
	}
}

type AutomationRulePropertyArrayChangedConditionSupportedArrayType string

const (
	AutomationRulePropertyArrayChangedConditionSupportedArrayTypeAlerts   AutomationRulePropertyArrayChangedConditionSupportedArrayType = "Alerts"
	AutomationRulePropertyArrayChangedConditionSupportedArrayTypeComments AutomationRulePropertyArrayChangedConditionSupportedArrayType = "Comments"
	AutomationRulePropertyArrayChangedConditionSupportedArrayTypeLabels   AutomationRulePropertyArrayChangedConditionSupportedArrayType = "Labels"
	AutomationRulePropertyArrayChangedConditionSupportedArrayTypeTactics  AutomationRulePropertyArrayChangedConditionSupportedArrayType = "Tactics"
)

func PossibleValuesForAutomationRulePropertyArrayChangedConditionSupportedArrayType() []string {
	return []string{
		string(AutomationRulePropertyArrayChangedConditionSupportedArrayTypeAlerts),
		string(AutomationRulePropertyArrayChangedConditionSupportedArrayTypeComments),
		string(AutomationRulePropertyArrayChangedConditionSupportedArrayTypeLabels),
		string(AutomationRulePropertyArrayChangedConditionSupportedArrayTypeTactics),
	}
}

type AutomationRulePropertyArrayChangedConditionSupportedChangeType string

const (
	AutomationRulePropertyArrayChangedConditionSupportedChangeTypeAdded AutomationRulePropertyArrayChangedConditionSupportedChangeType = "Added"
)

func PossibleValuesForAutomationRulePropertyArrayChangedConditionSupportedChangeType() []string {
	return []string{
		string(AutomationRulePropertyArrayChangedConditionSupportedChangeTypeAdded),
	}
}

type AutomationRulePropertyArrayConditionSupportedArrayConditionType string

const (
	AutomationRulePropertyArrayConditionSupportedArrayConditionTypeAnyItem AutomationRulePropertyArrayConditionSupportedArrayConditionType = "AnyItem"
)

func PossibleValuesForAutomationRulePropertyArrayConditionSupportedArrayConditionType() []string {
	return []string{
		string(AutomationRulePropertyArrayConditionSupportedArrayConditionTypeAnyItem),
	}
}

type AutomationRulePropertyArrayConditionSupportedArrayType string

const (
	AutomationRulePropertyArrayConditionSupportedArrayTypeCustomDetailValues AutomationRulePropertyArrayConditionSupportedArrayType = "CustomDetailValues"
	AutomationRulePropertyArrayConditionSupportedArrayTypeCustomDetails      AutomationRulePropertyArrayConditionSupportedArrayType = "CustomDetails"
)

func PossibleValuesForAutomationRulePropertyArrayConditionSupportedArrayType() []string {
	return []string{
		string(AutomationRulePropertyArrayConditionSupportedArrayTypeCustomDetailValues),
		string(AutomationRulePropertyArrayConditionSupportedArrayTypeCustomDetails),
	}
}

type AutomationRulePropertyChangedConditionSupportedChangedType string

const (
	AutomationRulePropertyChangedConditionSupportedChangedTypeChangedFrom AutomationRulePropertyChangedConditionSupportedChangedType = "ChangedFrom"
	AutomationRulePropertyChangedConditionSupportedChangedTypeChangedTo   AutomationRulePropertyChangedConditionSupportedChangedType = "ChangedTo"
)

func PossibleValuesForAutomationRulePropertyChangedConditionSupportedChangedType() []string {
	return []string{
		string(AutomationRulePropertyChangedConditionSupportedChangedTypeChangedFrom),
		string(AutomationRulePropertyChangedConditionSupportedChangedTypeChangedTo),
	}
}

type AutomationRulePropertyChangedConditionSupportedPropertyType string

const (
	AutomationRulePropertyChangedConditionSupportedPropertyTypeIncidentOwner    AutomationRulePropertyChangedConditionSupportedPropertyType = "IncidentOwner"
	AutomationRulePropertyChangedConditionSupportedPropertyTypeIncidentSeverity AutomationRulePropertyChangedConditionSupportedPropertyType = "IncidentSeverity"
	AutomationRulePropertyChangedConditionSupportedPropertyTypeIncidentStatus   AutomationRulePropertyChangedConditionSupportedPropertyType = "IncidentStatus"
)

func PossibleValuesForAutomationRulePropertyChangedConditionSupportedPropertyType() []string {
	return []string{
		string(AutomationRulePropertyChangedConditionSupportedPropertyTypeIncidentOwner),
		string(AutomationRulePropertyChangedConditionSupportedPropertyTypeIncidentSeverity),
		string(AutomationRulePropertyChangedConditionSupportedPropertyTypeIncidentStatus),
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
	AutomationRulePropertyConditionSupportedPropertyAlertAnalyticRuleIds           AutomationRulePropertyConditionSupportedProperty = "AlertAnalyticRuleIds"
	AutomationRulePropertyConditionSupportedPropertyAlertProductNames              AutomationRulePropertyConditionSupportedProperty = "AlertProductNames"
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
	AutomationRulePropertyConditionSupportedPropertyIncidentCustomDetailsKey       AutomationRulePropertyConditionSupportedProperty = "IncidentCustomDetailsKey"
	AutomationRulePropertyConditionSupportedPropertyIncidentCustomDetailsValue     AutomationRulePropertyConditionSupportedProperty = "IncidentCustomDetailsValue"
	AutomationRulePropertyConditionSupportedPropertyIncidentDescription            AutomationRulePropertyConditionSupportedProperty = "IncidentDescription"
	AutomationRulePropertyConditionSupportedPropertyIncidentLabel                  AutomationRulePropertyConditionSupportedProperty = "IncidentLabel"
	AutomationRulePropertyConditionSupportedPropertyIncidentProviderName           AutomationRulePropertyConditionSupportedProperty = "IncidentProviderName"
	AutomationRulePropertyConditionSupportedPropertyIncidentRelatedAnalyticRuleIds AutomationRulePropertyConditionSupportedProperty = "IncidentRelatedAnalyticRuleIds"
	AutomationRulePropertyConditionSupportedPropertyIncidentSeverity               AutomationRulePropertyConditionSupportedProperty = "IncidentSeverity"
	AutomationRulePropertyConditionSupportedPropertyIncidentStatus                 AutomationRulePropertyConditionSupportedProperty = "IncidentStatus"
	AutomationRulePropertyConditionSupportedPropertyIncidentTactics                AutomationRulePropertyConditionSupportedProperty = "IncidentTactics"
	AutomationRulePropertyConditionSupportedPropertyIncidentTitle                  AutomationRulePropertyConditionSupportedProperty = "IncidentTitle"
	AutomationRulePropertyConditionSupportedPropertyIncidentUpdatedBySource        AutomationRulePropertyConditionSupportedProperty = "IncidentUpdatedBySource"
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
		string(AutomationRulePropertyConditionSupportedPropertyAlertAnalyticRuleIds),
		string(AutomationRulePropertyConditionSupportedPropertyAlertProductNames),
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
		string(AutomationRulePropertyConditionSupportedPropertyIncidentCustomDetailsKey),
		string(AutomationRulePropertyConditionSupportedPropertyIncidentCustomDetailsValue),
		string(AutomationRulePropertyConditionSupportedPropertyIncidentDescription),
		string(AutomationRulePropertyConditionSupportedPropertyIncidentLabel),
		string(AutomationRulePropertyConditionSupportedPropertyIncidentProviderName),
		string(AutomationRulePropertyConditionSupportedPropertyIncidentRelatedAnalyticRuleIds),
		string(AutomationRulePropertyConditionSupportedPropertyIncidentSeverity),
		string(AutomationRulePropertyConditionSupportedPropertyIncidentStatus),
		string(AutomationRulePropertyConditionSupportedPropertyIncidentTactics),
		string(AutomationRulePropertyConditionSupportedPropertyIncidentTitle),
		string(AutomationRulePropertyConditionSupportedPropertyIncidentUpdatedBySource),
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

type ConditionType string

const (
	ConditionTypeBoolean              ConditionType = "Boolean"
	ConditionTypeProperty             ConditionType = "Property"
	ConditionTypePropertyArray        ConditionType = "PropertyArray"
	ConditionTypePropertyArrayChanged ConditionType = "PropertyArrayChanged"
	ConditionTypePropertyChanged      ConditionType = "PropertyChanged"
)

func PossibleValuesForConditionType() []string {
	return []string{
		string(ConditionTypeBoolean),
		string(ConditionTypeProperty),
		string(ConditionTypePropertyArray),
		string(ConditionTypePropertyArrayChanged),
		string(ConditionTypePropertyChanged),
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
	IncidentLabelTypeAutoAssigned IncidentLabelType = "AutoAssigned"
	IncidentLabelTypeUser         IncidentLabelType = "User"
)

func PossibleValuesForIncidentLabelType() []string {
	return []string{
		string(IncidentLabelTypeAutoAssigned),
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
	TriggersOnAlerts    TriggersOn = "Alerts"
	TriggersOnIncidents TriggersOn = "Incidents"
)

func PossibleValuesForTriggersOn() []string {
	return []string{
		string(TriggersOnAlerts),
		string(TriggersOnIncidents),
	}
}

type TriggersWhen string

const (
	TriggersWhenCreated TriggersWhen = "Created"
	TriggersWhenUpdated TriggersWhen = "Updated"
)

func PossibleValuesForTriggersWhen() []string {
	return []string{
		string(TriggersWhenCreated),
		string(TriggersWhenUpdated),
	}
}
