package namespaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessRights string

const (
	AccessRightsListen AccessRights = "Listen"
	AccessRightsManage AccessRights = "Manage"
	AccessRightsSend   AccessRights = "Send"
)

func PossibleValuesForAccessRights() []string {
	return []string{
		string(AccessRightsListen),
		string(AccessRightsManage),
		string(AccessRightsSend),
	}
}

type DefaultAction string

const (
	DefaultActionAllow DefaultAction = "Allow"
	DefaultActionDeny  DefaultAction = "Deny"
)

func PossibleValuesForDefaultAction() []string {
	return []string{
		string(DefaultActionAllow),
		string(DefaultActionDeny),
	}
}

type EndPointProvisioningState string

const (
	EndPointProvisioningStateCanceled  EndPointProvisioningState = "Canceled"
	EndPointProvisioningStateCreating  EndPointProvisioningState = "Creating"
	EndPointProvisioningStateDeleting  EndPointProvisioningState = "Deleting"
	EndPointProvisioningStateFailed    EndPointProvisioningState = "Failed"
	EndPointProvisioningStateSucceeded EndPointProvisioningState = "Succeeded"
	EndPointProvisioningStateUpdating  EndPointProvisioningState = "Updating"
)

func PossibleValuesForEndPointProvisioningState() []string {
	return []string{
		string(EndPointProvisioningStateCanceled),
		string(EndPointProvisioningStateCreating),
		string(EndPointProvisioningStateDeleting),
		string(EndPointProvisioningStateFailed),
		string(EndPointProvisioningStateSucceeded),
		string(EndPointProvisioningStateUpdating),
	}
}

type KeyType string

const (
	KeyTypePrimaryKey   KeyType = "PrimaryKey"
	KeyTypeSecondaryKey KeyType = "SecondaryKey"
)

func PossibleValuesForKeyType() []string {
	return []string{
		string(KeyTypePrimaryKey),
		string(KeyTypeSecondaryKey),
	}
}

type NetworkRuleIPAction string

const (
	NetworkRuleIPActionAllow NetworkRuleIPAction = "Allow"
)

func PossibleValuesForNetworkRuleIPAction() []string {
	return []string{
		string(NetworkRuleIPActionAllow),
	}
}

type PrivateLinkConnectionStatus string

const (
	PrivateLinkConnectionStatusApproved     PrivateLinkConnectionStatus = "Approved"
	PrivateLinkConnectionStatusDisconnected PrivateLinkConnectionStatus = "Disconnected"
	PrivateLinkConnectionStatusPending      PrivateLinkConnectionStatus = "Pending"
	PrivateLinkConnectionStatusRejected     PrivateLinkConnectionStatus = "Rejected"
)

func PossibleValuesForPrivateLinkConnectionStatus() []string {
	return []string{
		string(PrivateLinkConnectionStatusApproved),
		string(PrivateLinkConnectionStatusDisconnected),
		string(PrivateLinkConnectionStatusPending),
		string(PrivateLinkConnectionStatusRejected),
	}
}

type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled           PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled            PublicNetworkAccess = "Enabled"
	PublicNetworkAccessSecuredByPerimeter PublicNetworkAccess = "SecuredByPerimeter"
)

func PossibleValuesForPublicNetworkAccess() []string {
	return []string{
		string(PublicNetworkAccessDisabled),
		string(PublicNetworkAccessEnabled),
		string(PublicNetworkAccessSecuredByPerimeter),
	}
}

type SkuName string

const (
	SkuNameStandard SkuName = "Standard"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNameStandard),
	}
}

type SkuTier string

const (
	SkuTierStandard SkuTier = "Standard"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierStandard),
	}
}

type UnavailableReason string

const (
	UnavailableReasonInvalidName                           UnavailableReason = "InvalidName"
	UnavailableReasonNameInLockdown                        UnavailableReason = "NameInLockdown"
	UnavailableReasonNameInUse                             UnavailableReason = "NameInUse"
	UnavailableReasonNone                                  UnavailableReason = "None"
	UnavailableReasonSubscriptionIsDisabled                UnavailableReason = "SubscriptionIsDisabled"
	UnavailableReasonTooManyNamespaceInCurrentSubscription UnavailableReason = "TooManyNamespaceInCurrentSubscription"
)

func PossibleValuesForUnavailableReason() []string {
	return []string{
		string(UnavailableReasonInvalidName),
		string(UnavailableReasonNameInLockdown),
		string(UnavailableReasonNameInUse),
		string(UnavailableReasonNone),
		string(UnavailableReasonSubscriptionIsDisabled),
		string(UnavailableReasonTooManyNamespaceInCurrentSubscription),
	}
}
