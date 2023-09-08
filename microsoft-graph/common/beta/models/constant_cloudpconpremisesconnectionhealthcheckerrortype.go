package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCOnPremisesConnectionHealthCheckErrorType string

const (
	CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckAccessDenied                                           CloudPCOnPremisesConnectionHealthCheckErrorType = "AdJoinCheckAccessDenied"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckAccountLockedOrDisabled                                CloudPCOnPremisesConnectionHealthCheckErrorType = "AdJoinCheckAccountLockedOrDisabled"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckAccountQuotaExceeded                                   CloudPCOnPremisesConnectionHealthCheckErrorType = "AdJoinCheckAccountQuotaExceeded"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckComputerObjectAlreadyExists                            CloudPCOnPremisesConnectionHealthCheckErrorType = "AdJoinCheckComputerObjectAlreadyExists"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckCredentialsExpired                                     CloudPCOnPremisesConnectionHealthCheckErrorType = "AdJoinCheckCredentialsExpired"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckFqdnNotFound                                           CloudPCOnPremisesConnectionHealthCheckErrorType = "AdJoinCheckFqdnNotFound"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckIncorrectCredentials                                   CloudPCOnPremisesConnectionHealthCheckErrorType = "AdJoinCheckIncorrectCredentials"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckOrganizationalUnitIncorrectFormat                      CloudPCOnPremisesConnectionHealthCheckErrorType = "AdJoinCheckOrganizationalUnitIncorrectFormat"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckOrganizationalUnitNotFound                             CloudPCOnPremisesConnectionHealthCheckErrorType = "AdJoinCheckOrganizationalUnitNotFound"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckUnknownError                                           CloudPCOnPremisesConnectionHealthCheckErrorType = "AdJoinCheckUnknownError"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeazureAdDeviceSyncCheckConnectDisabled                             CloudPCOnPremisesConnectionHealthCheckErrorType = "AzureAdDeviceSyncCheckConnectDisabled"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeazureAdDeviceSyncCheckDeviceNotFound                              CloudPCOnPremisesConnectionHealthCheckErrorType = "AzureAdDeviceSyncCheckDeviceNotFound"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeazureAdDeviceSyncCheckDurationExceeded                            CloudPCOnPremisesConnectionHealthCheckErrorType = "AzureAdDeviceSyncCheckDurationExceeded"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeazureAdDeviceSyncCheckLongSyncCircle                              CloudPCOnPremisesConnectionHealthCheckErrorType = "AzureAdDeviceSyncCheckLongSyncCircle"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeazureAdDeviceSyncCheckScpNotConfigured                            CloudPCOnPremisesConnectionHealthCheckErrorType = "AzureAdDeviceSyncCheckScpNotConfigured"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeazureAdDeviceSyncCheckTransientServiceError                       CloudPCOnPremisesConnectionHealthCheckErrorType = "AzureAdDeviceSyncCheckTransientServiceError"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeazureAdDeviceSyncCheckUnknownError                                CloudPCOnPremisesConnectionHealthCheckErrorType = "AzureAdDeviceSyncCheckUnknownError"
	CloudPCOnPremisesConnectionHealthCheckErrorTypednsCheckFqdnNotFound                                              CloudPCOnPremisesConnectionHealthCheckErrorType = "DnsCheckFqdnNotFound"
	CloudPCOnPremisesConnectionHealthCheckErrorTypednsCheckNameWithInvalidCharacter                                  CloudPCOnPremisesConnectionHealthCheckErrorType = "DnsCheckNameWithInvalidCharacter"
	CloudPCOnPremisesConnectionHealthCheckErrorTypednsCheckUnknownError                                              CloudPCOnPremisesConnectionHealthCheckErrorType = "DnsCheckUnknownError"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeendpointConnectivityCheckAzureADUrlNotAllowListed                 CloudPCOnPremisesConnectionHealthCheckErrorType = "EndpointConnectivityCheckAzureADUrlNotAllowListed"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeendpointConnectivityCheckCloudPcUrlNotAllowListed                 CloudPCOnPremisesConnectionHealthCheckErrorType = "EndpointConnectivityCheckCloudPCUrlNotAllowListed"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeendpointConnectivityCheckIntuneUrlNotAllowListed                  CloudPCOnPremisesConnectionHealthCheckErrorType = "EndpointConnectivityCheckIntuneUrlNotAllowListed"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeendpointConnectivityCheckLocaleUrlNotAllowListed                  CloudPCOnPremisesConnectionHealthCheckErrorType = "EndpointConnectivityCheckLocaleUrlNotAllowListed"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeendpointConnectivityCheckUnknownError                             CloudPCOnPremisesConnectionHealthCheckErrorType = "EndpointConnectivityCheckUnknownError"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeendpointConnectivityCheckWVDUrlNotAllowListed                     CloudPCOnPremisesConnectionHealthCheckErrorType = "EndpointConnectivityCheckWVDUrlNotAllowListed"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeinternalServerErrorAllocateResourceFailed                         CloudPCOnPremisesConnectionHealthCheckErrorType = "InternalServerErrorAllocateResourceFailed"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeinternalServerErrorDeploymentCanceled                             CloudPCOnPremisesConnectionHealthCheckErrorType = "InternalServerErrorDeploymentCanceled"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeinternalServerErrorUnableToRunDscScript                           CloudPCOnPremisesConnectionHealthCheckErrorType = "InternalServerErrorUnableToRunDscScript"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeinternalServerErrorVMDeploymentTimeout                            CloudPCOnPremisesConnectionHealthCheckErrorType = "InternalServerErrorVMDeploymentTimeout"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeinternalServerUnknownError                                        CloudPCOnPremisesConnectionHealthCheckErrorType = "InternalServerUnknownError"
	CloudPCOnPremisesConnectionHealthCheckErrorTypepermissionCheckNoResourceGroupNetworkContributorRole              CloudPCOnPremisesConnectionHealthCheckErrorType = "PermissionCheckNoResourceGroupNetworkContributorRole"
	CloudPCOnPremisesConnectionHealthCheckErrorTypepermissionCheckNoResourceGroupOwnerRole                           CloudPCOnPremisesConnectionHealthCheckErrorType = "PermissionCheckNoResourceGroupOwnerRole"
	CloudPCOnPremisesConnectionHealthCheckErrorTypepermissionCheckNoSubscriptionReaderRole                           CloudPCOnPremisesConnectionHealthCheckErrorType = "PermissionCheckNoSubscriptionReaderRole"
	CloudPCOnPremisesConnectionHealthCheckErrorTypepermissionCheckNoVNetContributorRole                              CloudPCOnPremisesConnectionHealthCheckErrorType = "PermissionCheckNoVNetContributorRole"
	CloudPCOnPremisesConnectionHealthCheckErrorTypepermissionCheckNoWindows365NetworkInterfaceContributorRole        CloudPCOnPremisesConnectionHealthCheckErrorType = "PermissionCheckNoWindows365NetworkInterfaceContributorRole"
	CloudPCOnPremisesConnectionHealthCheckErrorTypepermissionCheckNoWindows365NetworkUserRole                        CloudPCOnPremisesConnectionHealthCheckErrorType = "PermissionCheckNoWindows365NetworkUserRole"
	CloudPCOnPremisesConnectionHealthCheckErrorTypepermissionCheckTransientServiceError                              CloudPCOnPremisesConnectionHealthCheckErrorType = "PermissionCheckTransientServiceError"
	CloudPCOnPremisesConnectionHealthCheckErrorTypepermissionCheckUnknownError                                       CloudPCOnPremisesConnectionHealthCheckErrorType = "PermissionCheckUnknownError"
	CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckAzurePolicyViolation                     CloudPCOnPremisesConnectionHealthCheckErrorType = "ResourceAvailabilityCheckAzurePolicyViolation"
	CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckGeneralSubscriptionError                 CloudPCOnPremisesConnectionHealthCheckErrorType = "ResourceAvailabilityCheckGeneralSubscriptionError"
	CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckIntuneCustomWindowsRestrictionViolation  CloudPCOnPremisesConnectionHealthCheckErrorType = "ResourceAvailabilityCheckIntuneCustomWindowsRestrictionViolation"
	CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckIntuneDefaultWindowsRestrictionViolation CloudPCOnPremisesConnectionHealthCheckErrorType = "ResourceAvailabilityCheckIntuneDefaultWindowsRestrictionViolation"
	CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckNoIntuneReaderRoleError                  CloudPCOnPremisesConnectionHealthCheckErrorType = "ResourceAvailabilityCheckNoIntuneReaderRoleError"
	CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckNoSubnetIP                               CloudPCOnPremisesConnectionHealthCheckErrorType = "ResourceAvailabilityCheckNoSubnetIP"
	CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckResourceGroupBeingDeleted                CloudPCOnPremisesConnectionHealthCheckErrorType = "ResourceAvailabilityCheckResourceGroupBeingDeleted"
	CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckResourceGroupInvalid                     CloudPCOnPremisesConnectionHealthCheckErrorType = "ResourceAvailabilityCheckResourceGroupInvalid"
	CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckResourceGroupLockedForDelete             CloudPCOnPremisesConnectionHealthCheckErrorType = "ResourceAvailabilityCheckResourceGroupLockedForDelete"
	CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckResourceGroupLockedForReadonly           CloudPCOnPremisesConnectionHealthCheckErrorType = "ResourceAvailabilityCheckResourceGroupLockedForReadonly"
	CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckSubnetDelegationFailed                   CloudPCOnPremisesConnectionHealthCheckErrorType = "ResourceAvailabilityCheckSubnetDelegationFailed"
	CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckSubnetInvalid                            CloudPCOnPremisesConnectionHealthCheckErrorType = "ResourceAvailabilityCheckSubnetInvalid"
	CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckSubnetWithExternalResources              CloudPCOnPremisesConnectionHealthCheckErrorType = "ResourceAvailabilityCheckSubnetWithExternalResources"
	CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckSubscriptionDisabled                     CloudPCOnPremisesConnectionHealthCheckErrorType = "ResourceAvailabilityCheckSubscriptionDisabled"
	CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckSubscriptionNotFound                     CloudPCOnPremisesConnectionHealthCheckErrorType = "ResourceAvailabilityCheckSubscriptionNotFound"
	CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckSubscriptionTransferred                  CloudPCOnPremisesConnectionHealthCheckErrorType = "ResourceAvailabilityCheckSubscriptionTransferred"
	CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckTransientServiceError                    CloudPCOnPremisesConnectionHealthCheckErrorType = "ResourceAvailabilityCheckTransientServiceError"
	CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckUnknownError                             CloudPCOnPremisesConnectionHealthCheckErrorType = "ResourceAvailabilityCheckUnknownError"
	CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckUnsupportedVNetRegion                    CloudPCOnPremisesConnectionHealthCheckErrorType = "ResourceAvailabilityCheckUnsupportedVNetRegion"
	CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckVNetBeingMoved                           CloudPCOnPremisesConnectionHealthCheckErrorType = "ResourceAvailabilityCheckVNetBeingMoved"
	CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckVNetInvalid                              CloudPCOnPremisesConnectionHealthCheckErrorType = "ResourceAvailabilityCheckVNetInvalid"
	CloudPCOnPremisesConnectionHealthCheckErrorTypessoCheckKerberosConfigurationError                                CloudPCOnPremisesConnectionHealthCheckErrorType = "SsoCheckKerberosConfigurationError"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeudpConnectivityCheckStunUrlNotAllowListed                         CloudPCOnPremisesConnectionHealthCheckErrorType = "UdpConnectivityCheckStunUrlNotAllowListed"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeudpConnectivityCheckTurnUrlNotAllowListed                         CloudPCOnPremisesConnectionHealthCheckErrorType = "UdpConnectivityCheckTurnUrlNotAllowListed"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeudpConnectivityCheckUnknownError                                  CloudPCOnPremisesConnectionHealthCheckErrorType = "UdpConnectivityCheckUnknownError"
	CloudPCOnPremisesConnectionHealthCheckErrorTypeudpConnectivityCheckUrlsNotAllowListed                            CloudPCOnPremisesConnectionHealthCheckErrorType = "UdpConnectivityCheckUrlsNotAllowListed"
)

func PossibleValuesForCloudPCOnPremisesConnectionHealthCheckErrorType() []string {
	return []string{
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckAccessDenied),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckAccountLockedOrDisabled),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckAccountQuotaExceeded),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckComputerObjectAlreadyExists),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckCredentialsExpired),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckFqdnNotFound),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckIncorrectCredentials),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckOrganizationalUnitIncorrectFormat),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckOrganizationalUnitNotFound),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckUnknownError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeazureAdDeviceSyncCheckConnectDisabled),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeazureAdDeviceSyncCheckDeviceNotFound),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeazureAdDeviceSyncCheckDurationExceeded),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeazureAdDeviceSyncCheckLongSyncCircle),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeazureAdDeviceSyncCheckScpNotConfigured),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeazureAdDeviceSyncCheckTransientServiceError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeazureAdDeviceSyncCheckUnknownError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypednsCheckFqdnNotFound),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypednsCheckNameWithInvalidCharacter),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypednsCheckUnknownError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeendpointConnectivityCheckAzureADUrlNotAllowListed),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeendpointConnectivityCheckCloudPcUrlNotAllowListed),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeendpointConnectivityCheckIntuneUrlNotAllowListed),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeendpointConnectivityCheckLocaleUrlNotAllowListed),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeendpointConnectivityCheckUnknownError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeendpointConnectivityCheckWVDUrlNotAllowListed),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeinternalServerErrorAllocateResourceFailed),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeinternalServerErrorDeploymentCanceled),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeinternalServerErrorUnableToRunDscScript),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeinternalServerErrorVMDeploymentTimeout),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeinternalServerUnknownError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypepermissionCheckNoResourceGroupNetworkContributorRole),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypepermissionCheckNoResourceGroupOwnerRole),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypepermissionCheckNoSubscriptionReaderRole),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypepermissionCheckNoVNetContributorRole),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypepermissionCheckNoWindows365NetworkInterfaceContributorRole),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypepermissionCheckNoWindows365NetworkUserRole),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypepermissionCheckTransientServiceError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypepermissionCheckUnknownError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckAzurePolicyViolation),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckGeneralSubscriptionError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckIntuneCustomWindowsRestrictionViolation),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckIntuneDefaultWindowsRestrictionViolation),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckNoIntuneReaderRoleError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckNoSubnetIP),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckResourceGroupBeingDeleted),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckResourceGroupInvalid),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckResourceGroupLockedForDelete),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckResourceGroupLockedForReadonly),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckSubnetDelegationFailed),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckSubnetInvalid),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckSubnetWithExternalResources),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckSubscriptionDisabled),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckSubscriptionNotFound),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckSubscriptionTransferred),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckTransientServiceError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckUnknownError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckUnsupportedVNetRegion),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckVNetBeingMoved),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckVNetInvalid),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypessoCheckKerberosConfigurationError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeudpConnectivityCheckStunUrlNotAllowListed),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeudpConnectivityCheckTurnUrlNotAllowListed),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeudpConnectivityCheckUnknownError),
		string(CloudPCOnPremisesConnectionHealthCheckErrorTypeudpConnectivityCheckUrlsNotAllowListed),
	}
}

func parseCloudPCOnPremisesConnectionHealthCheckErrorType(input string) (*CloudPCOnPremisesConnectionHealthCheckErrorType, error) {
	vals := map[string]CloudPCOnPremisesConnectionHealthCheckErrorType{
		"adjoincheckaccessdenied":                                           CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckAccessDenied,
		"adjoincheckaccountlockedordisabled":                                CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckAccountLockedOrDisabled,
		"adjoincheckaccountquotaexceeded":                                   CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckAccountQuotaExceeded,
		"adjoincheckcomputerobjectalreadyexists":                            CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckComputerObjectAlreadyExists,
		"adjoincheckcredentialsexpired":                                     CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckCredentialsExpired,
		"adjoincheckfqdnnotfound":                                           CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckFqdnNotFound,
		"adjoincheckincorrectcredentials":                                   CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckIncorrectCredentials,
		"adjoincheckorganizationalunitincorrectformat":                      CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckOrganizationalUnitIncorrectFormat,
		"adjoincheckorganizationalunitnotfound":                             CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckOrganizationalUnitNotFound,
		"adjoincheckunknownerror":                                           CloudPCOnPremisesConnectionHealthCheckErrorTypeadJoinCheckUnknownError,
		"azureaddevicesynccheckconnectdisabled":                             CloudPCOnPremisesConnectionHealthCheckErrorTypeazureAdDeviceSyncCheckConnectDisabled,
		"azureaddevicesynccheckdevicenotfound":                              CloudPCOnPremisesConnectionHealthCheckErrorTypeazureAdDeviceSyncCheckDeviceNotFound,
		"azureaddevicesynccheckdurationexceeded":                            CloudPCOnPremisesConnectionHealthCheckErrorTypeazureAdDeviceSyncCheckDurationExceeded,
		"azureaddevicesyncchecklongsynccircle":                              CloudPCOnPremisesConnectionHealthCheckErrorTypeazureAdDeviceSyncCheckLongSyncCircle,
		"azureaddevicesynccheckscpnotconfigured":                            CloudPCOnPremisesConnectionHealthCheckErrorTypeazureAdDeviceSyncCheckScpNotConfigured,
		"azureaddevicesyncchecktransientserviceerror":                       CloudPCOnPremisesConnectionHealthCheckErrorTypeazureAdDeviceSyncCheckTransientServiceError,
		"azureaddevicesynccheckunknownerror":                                CloudPCOnPremisesConnectionHealthCheckErrorTypeazureAdDeviceSyncCheckUnknownError,
		"dnscheckfqdnnotfound":                                              CloudPCOnPremisesConnectionHealthCheckErrorTypednsCheckFqdnNotFound,
		"dnschecknamewithinvalidcharacter":                                  CloudPCOnPremisesConnectionHealthCheckErrorTypednsCheckNameWithInvalidCharacter,
		"dnscheckunknownerror":                                              CloudPCOnPremisesConnectionHealthCheckErrorTypednsCheckUnknownError,
		"endpointconnectivitycheckazureadurlnotallowlisted":                 CloudPCOnPremisesConnectionHealthCheckErrorTypeendpointConnectivityCheckAzureADUrlNotAllowListed,
		"endpointconnectivitycheckcloudpcurlnotallowlisted":                 CloudPCOnPremisesConnectionHealthCheckErrorTypeendpointConnectivityCheckCloudPcUrlNotAllowListed,
		"endpointconnectivitycheckintuneurlnotallowlisted":                  CloudPCOnPremisesConnectionHealthCheckErrorTypeendpointConnectivityCheckIntuneUrlNotAllowListed,
		"endpointconnectivitychecklocaleurlnotallowlisted":                  CloudPCOnPremisesConnectionHealthCheckErrorTypeendpointConnectivityCheckLocaleUrlNotAllowListed,
		"endpointconnectivitycheckunknownerror":                             CloudPCOnPremisesConnectionHealthCheckErrorTypeendpointConnectivityCheckUnknownError,
		"endpointconnectivitycheckwvdurlnotallowlisted":                     CloudPCOnPremisesConnectionHealthCheckErrorTypeendpointConnectivityCheckWVDUrlNotAllowListed,
		"internalservererrorallocateresourcefailed":                         CloudPCOnPremisesConnectionHealthCheckErrorTypeinternalServerErrorAllocateResourceFailed,
		"internalservererrordeploymentcanceled":                             CloudPCOnPremisesConnectionHealthCheckErrorTypeinternalServerErrorDeploymentCanceled,
		"internalservererrorunabletorundscscript":                           CloudPCOnPremisesConnectionHealthCheckErrorTypeinternalServerErrorUnableToRunDscScript,
		"internalservererrorvmdeploymenttimeout":                            CloudPCOnPremisesConnectionHealthCheckErrorTypeinternalServerErrorVMDeploymentTimeout,
		"internalserverunknownerror":                                        CloudPCOnPremisesConnectionHealthCheckErrorTypeinternalServerUnknownError,
		"permissionchecknoresourcegroupnetworkcontributorrole":              CloudPCOnPremisesConnectionHealthCheckErrorTypepermissionCheckNoResourceGroupNetworkContributorRole,
		"permissionchecknoresourcegroupownerrole":                           CloudPCOnPremisesConnectionHealthCheckErrorTypepermissionCheckNoResourceGroupOwnerRole,
		"permissionchecknosubscriptionreaderrole":                           CloudPCOnPremisesConnectionHealthCheckErrorTypepermissionCheckNoSubscriptionReaderRole,
		"permissionchecknovnetcontributorrole":                              CloudPCOnPremisesConnectionHealthCheckErrorTypepermissionCheckNoVNetContributorRole,
		"permissionchecknowindows365networkinterfacecontributorrole":        CloudPCOnPremisesConnectionHealthCheckErrorTypepermissionCheckNoWindows365NetworkInterfaceContributorRole,
		"permissionchecknowindows365networkuserrole":                        CloudPCOnPremisesConnectionHealthCheckErrorTypepermissionCheckNoWindows365NetworkUserRole,
		"permissionchecktransientserviceerror":                              CloudPCOnPremisesConnectionHealthCheckErrorTypepermissionCheckTransientServiceError,
		"permissioncheckunknownerror":                                       CloudPCOnPremisesConnectionHealthCheckErrorTypepermissionCheckUnknownError,
		"resourceavailabilitycheckazurepolicyviolation":                     CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckAzurePolicyViolation,
		"resourceavailabilitycheckgeneralsubscriptionerror":                 CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckGeneralSubscriptionError,
		"resourceavailabilitycheckintunecustomwindowsrestrictionviolation":  CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckIntuneCustomWindowsRestrictionViolation,
		"resourceavailabilitycheckintunedefaultwindowsrestrictionviolation": CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckIntuneDefaultWindowsRestrictionViolation,
		"resourceavailabilitychecknointunereaderroleerror":                  CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckNoIntuneReaderRoleError,
		"resourceavailabilitychecknosubnetip":                               CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckNoSubnetIP,
		"resourceavailabilitycheckresourcegroupbeingdeleted":                CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckResourceGroupBeingDeleted,
		"resourceavailabilitycheckresourcegroupinvalid":                     CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckResourceGroupInvalid,
		"resourceavailabilitycheckresourcegrouplockedfordelete":             CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckResourceGroupLockedForDelete,
		"resourceavailabilitycheckresourcegrouplockedforreadonly":           CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckResourceGroupLockedForReadonly,
		"resourceavailabilitychecksubnetdelegationfailed":                   CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckSubnetDelegationFailed,
		"resourceavailabilitychecksubnetinvalid":                            CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckSubnetInvalid,
		"resourceavailabilitychecksubnetwithexternalresources":              CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckSubnetWithExternalResources,
		"resourceavailabilitychecksubscriptiondisabled":                     CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckSubscriptionDisabled,
		"resourceavailabilitychecksubscriptionnotfound":                     CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckSubscriptionNotFound,
		"resourceavailabilitychecksubscriptiontransferred":                  CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckSubscriptionTransferred,
		"resourceavailabilitychecktransientserviceerror":                    CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckTransientServiceError,
		"resourceavailabilitycheckunknownerror":                             CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckUnknownError,
		"resourceavailabilitycheckunsupportedvnetregion":                    CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckUnsupportedVNetRegion,
		"resourceavailabilitycheckvnetbeingmoved":                           CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckVNetBeingMoved,
		"resourceavailabilitycheckvnetinvalid":                              CloudPCOnPremisesConnectionHealthCheckErrorTyperesourceAvailabilityCheckVNetInvalid,
		"ssocheckkerberosconfigurationerror":                                CloudPCOnPremisesConnectionHealthCheckErrorTypessoCheckKerberosConfigurationError,
		"udpconnectivitycheckstunurlnotallowlisted":                         CloudPCOnPremisesConnectionHealthCheckErrorTypeudpConnectivityCheckStunUrlNotAllowListed,
		"udpconnectivitycheckturnurlnotallowlisted":                         CloudPCOnPremisesConnectionHealthCheckErrorTypeudpConnectivityCheckTurnUrlNotAllowListed,
		"udpconnectivitycheckunknownerror":                                  CloudPCOnPremisesConnectionHealthCheckErrorTypeudpConnectivityCheckUnknownError,
		"udpconnectivitycheckurlsnotallowlisted":                            CloudPCOnPremisesConnectionHealthCheckErrorTypeudpConnectivityCheckUrlsNotAllowListed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCOnPremisesConnectionHealthCheckErrorType(input)
	return &out, nil
}
