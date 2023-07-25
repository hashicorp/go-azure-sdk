package domains

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureResourceType string

const (
	AzureResourceTypeTrafficManager AzureResourceType = "TrafficManager"
	AzureResourceTypeWebsite        AzureResourceType = "Website"
)

func PossibleValuesForAzureResourceType() []string {
	return []string{
		string(AzureResourceTypeTrafficManager),
		string(AzureResourceTypeWebsite),
	}
}

func parseAzureResourceType(input string) (*AzureResourceType, error) {
	vals := map[string]AzureResourceType{
		"trafficmanager": AzureResourceTypeTrafficManager,
		"website":        AzureResourceTypeWebsite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureResourceType(input)
	return &out, nil
}

type CustomHostNameDnsRecordType string

const (
	CustomHostNameDnsRecordTypeA     CustomHostNameDnsRecordType = "A"
	CustomHostNameDnsRecordTypeCName CustomHostNameDnsRecordType = "CName"
)

func PossibleValuesForCustomHostNameDnsRecordType() []string {
	return []string{
		string(CustomHostNameDnsRecordTypeA),
		string(CustomHostNameDnsRecordTypeCName),
	}
}

func parseCustomHostNameDnsRecordType(input string) (*CustomHostNameDnsRecordType, error) {
	vals := map[string]CustomHostNameDnsRecordType{
		"a":     CustomHostNameDnsRecordTypeA,
		"cname": CustomHostNameDnsRecordTypeCName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomHostNameDnsRecordType(input)
	return &out, nil
}

type DnsType string

const (
	DnsTypeAzureDns                  DnsType = "AzureDns"
	DnsTypeDefaultDomainRegistrarDns DnsType = "DefaultDomainRegistrarDns"
)

func PossibleValuesForDnsType() []string {
	return []string{
		string(DnsTypeAzureDns),
		string(DnsTypeDefaultDomainRegistrarDns),
	}
}

func parseDnsType(input string) (*DnsType, error) {
	vals := map[string]DnsType{
		"azuredns":                  DnsTypeAzureDns,
		"defaultdomainregistrardns": DnsTypeDefaultDomainRegistrarDns,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DnsType(input)
	return &out, nil
}

type DomainStatus string

const (
	DomainStatusActive              DomainStatus = "Active"
	DomainStatusAwaiting            DomainStatus = "Awaiting"
	DomainStatusCancelled           DomainStatus = "Cancelled"
	DomainStatusConfiscated         DomainStatus = "Confiscated"
	DomainStatusDisabled            DomainStatus = "Disabled"
	DomainStatusExcluded            DomainStatus = "Excluded"
	DomainStatusExpired             DomainStatus = "Expired"
	DomainStatusFailed              DomainStatus = "Failed"
	DomainStatusHeld                DomainStatus = "Held"
	DomainStatusJsonConverterFailed DomainStatus = "JsonConverterFailed"
	DomainStatusLocked              DomainStatus = "Locked"
	DomainStatusParked              DomainStatus = "Parked"
	DomainStatusPending             DomainStatus = "Pending"
	DomainStatusReserved            DomainStatus = "Reserved"
	DomainStatusReverted            DomainStatus = "Reverted"
	DomainStatusSuspended           DomainStatus = "Suspended"
	DomainStatusTransferred         DomainStatus = "Transferred"
	DomainStatusUnknown             DomainStatus = "Unknown"
	DomainStatusUnlocked            DomainStatus = "Unlocked"
	DomainStatusUnparked            DomainStatus = "Unparked"
	DomainStatusUpdated             DomainStatus = "Updated"
)

func PossibleValuesForDomainStatus() []string {
	return []string{
		string(DomainStatusActive),
		string(DomainStatusAwaiting),
		string(DomainStatusCancelled),
		string(DomainStatusConfiscated),
		string(DomainStatusDisabled),
		string(DomainStatusExcluded),
		string(DomainStatusExpired),
		string(DomainStatusFailed),
		string(DomainStatusHeld),
		string(DomainStatusJsonConverterFailed),
		string(DomainStatusLocked),
		string(DomainStatusParked),
		string(DomainStatusPending),
		string(DomainStatusReserved),
		string(DomainStatusReverted),
		string(DomainStatusSuspended),
		string(DomainStatusTransferred),
		string(DomainStatusUnknown),
		string(DomainStatusUnlocked),
		string(DomainStatusUnparked),
		string(DomainStatusUpdated),
	}
}

func parseDomainStatus(input string) (*DomainStatus, error) {
	vals := map[string]DomainStatus{
		"active":              DomainStatusActive,
		"awaiting":            DomainStatusAwaiting,
		"cancelled":           DomainStatusCancelled,
		"confiscated":         DomainStatusConfiscated,
		"disabled":            DomainStatusDisabled,
		"excluded":            DomainStatusExcluded,
		"expired":             DomainStatusExpired,
		"failed":              DomainStatusFailed,
		"held":                DomainStatusHeld,
		"jsonconverterfailed": DomainStatusJsonConverterFailed,
		"locked":              DomainStatusLocked,
		"parked":              DomainStatusParked,
		"pending":             DomainStatusPending,
		"reserved":            DomainStatusReserved,
		"reverted":            DomainStatusReverted,
		"suspended":           DomainStatusSuspended,
		"transferred":         DomainStatusTransferred,
		"unknown":             DomainStatusUnknown,
		"unlocked":            DomainStatusUnlocked,
		"unparked":            DomainStatusUnparked,
		"updated":             DomainStatusUpdated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DomainStatus(input)
	return &out, nil
}

type DomainType string

const (
	DomainTypeRegular     DomainType = "Regular"
	DomainTypeSoftDeleted DomainType = "SoftDeleted"
)

func PossibleValuesForDomainType() []string {
	return []string{
		string(DomainTypeRegular),
		string(DomainTypeSoftDeleted),
	}
}

func parseDomainType(input string) (*DomainType, error) {
	vals := map[string]DomainType{
		"regular":     DomainTypeRegular,
		"softdeleted": DomainTypeSoftDeleted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DomainType(input)
	return &out, nil
}

type HostNameType string

const (
	HostNameTypeManaged  HostNameType = "Managed"
	HostNameTypeVerified HostNameType = "Verified"
)

func PossibleValuesForHostNameType() []string {
	return []string{
		string(HostNameTypeManaged),
		string(HostNameTypeVerified),
	}
}

func parseHostNameType(input string) (*HostNameType, error) {
	vals := map[string]HostNameType{
		"managed":  HostNameTypeManaged,
		"verified": HostNameTypeVerified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HostNameType(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateCanceled   ProvisioningState = "Canceled"
	ProvisioningStateDeleting   ProvisioningState = "Deleting"
	ProvisioningStateFailed     ProvisioningState = "Failed"
	ProvisioningStateInProgress ProvisioningState = "InProgress"
	ProvisioningStateSucceeded  ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateInProgress),
		string(ProvisioningStateSucceeded),
	}
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"canceled":   ProvisioningStateCanceled,
		"deleting":   ProvisioningStateDeleting,
		"failed":     ProvisioningStateFailed,
		"inprogress": ProvisioningStateInProgress,
		"succeeded":  ProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type ResourceNotRenewableReason string

const (
	ResourceNotRenewableReasonExpirationNotInRenewalTimeRange          ResourceNotRenewableReason = "ExpirationNotInRenewalTimeRange"
	ResourceNotRenewableReasonRegistrationStatusNotSupportedForRenewal ResourceNotRenewableReason = "RegistrationStatusNotSupportedForRenewal"
	ResourceNotRenewableReasonSubscriptionNotActive                    ResourceNotRenewableReason = "SubscriptionNotActive"
)

func PossibleValuesForResourceNotRenewableReason() []string {
	return []string{
		string(ResourceNotRenewableReasonExpirationNotInRenewalTimeRange),
		string(ResourceNotRenewableReasonRegistrationStatusNotSupportedForRenewal),
		string(ResourceNotRenewableReasonSubscriptionNotActive),
	}
}

func parseResourceNotRenewableReason(input string) (*ResourceNotRenewableReason, error) {
	vals := map[string]ResourceNotRenewableReason{
		"expirationnotinrenewaltimerange":          ResourceNotRenewableReasonExpirationNotInRenewalTimeRange,
		"registrationstatusnotsupportedforrenewal": ResourceNotRenewableReasonRegistrationStatusNotSupportedForRenewal,
		"subscriptionnotactive":                    ResourceNotRenewableReasonSubscriptionNotActive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceNotRenewableReason(input)
	return &out, nil
}
