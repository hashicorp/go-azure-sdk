package shares

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureContainerDataFormat string

const (
	AzureContainerDataFormatAzureFile AzureContainerDataFormat = "AzureFile"
	AzureContainerDataFormatBlockBlob AzureContainerDataFormat = "BlockBlob"
	AzureContainerDataFormatPageBlob  AzureContainerDataFormat = "PageBlob"
)

func PossibleValuesForAzureContainerDataFormat() []string {
	return []string{
		string(AzureContainerDataFormatAzureFile),
		string(AzureContainerDataFormatBlockBlob),
		string(AzureContainerDataFormatPageBlob),
	}
}

type ClientPermissionType string

const (
	ClientPermissionTypeNoAccess  ClientPermissionType = "NoAccess"
	ClientPermissionTypeReadOnly  ClientPermissionType = "ReadOnly"
	ClientPermissionTypeReadWrite ClientPermissionType = "ReadWrite"
)

func PossibleValuesForClientPermissionType() []string {
	return []string{
		string(ClientPermissionTypeNoAccess),
		string(ClientPermissionTypeReadOnly),
		string(ClientPermissionTypeReadWrite),
	}
}

type DataPolicy string

const (
	DataPolicyCloud DataPolicy = "Cloud"
	DataPolicyLocal DataPolicy = "Local"
)

func PossibleValuesForDataPolicy() []string {
	return []string{
		string(DataPolicyCloud),
		string(DataPolicyLocal),
	}
}

type MonitoringStatus string

const (
	MonitoringStatusDisabled MonitoringStatus = "Disabled"
	MonitoringStatusEnabled  MonitoringStatus = "Enabled"
)

func PossibleValuesForMonitoringStatus() []string {
	return []string{
		string(MonitoringStatusDisabled),
		string(MonitoringStatusEnabled),
	}
}

type MountType string

const (
	MountTypeHostPath MountType = "HostPath"
	MountTypeVolume   MountType = "Volume"
)

func PossibleValuesForMountType() []string {
	return []string{
		string(MountTypeHostPath),
		string(MountTypeVolume),
	}
}

type RoleTypes string

const (
	RoleTypesASA                 RoleTypes = "ASA"
	RoleTypesCloudEdgeManagement RoleTypes = "CloudEdgeManagement"
	RoleTypesCognitive           RoleTypes = "Cognitive"
	RoleTypesFunctions           RoleTypes = "Functions"
	RoleTypesIOT                 RoleTypes = "IOT"
	RoleTypesKubernetes          RoleTypes = "Kubernetes"
	RoleTypesMEC                 RoleTypes = "MEC"
)

func PossibleValuesForRoleTypes() []string {
	return []string{
		string(RoleTypesASA),
		string(RoleTypesCloudEdgeManagement),
		string(RoleTypesCognitive),
		string(RoleTypesFunctions),
		string(RoleTypesIOT),
		string(RoleTypesKubernetes),
		string(RoleTypesMEC),
	}
}

type ShareAccessProtocol string

const (
	ShareAccessProtocolNFS ShareAccessProtocol = "NFS"
	ShareAccessProtocolSMB ShareAccessProtocol = "SMB"
)

func PossibleValuesForShareAccessProtocol() []string {
	return []string{
		string(ShareAccessProtocolNFS),
		string(ShareAccessProtocolSMB),
	}
}

type ShareAccessType string

const (
	ShareAccessTypeChange ShareAccessType = "Change"
	ShareAccessTypeCustom ShareAccessType = "Custom"
	ShareAccessTypeRead   ShareAccessType = "Read"
)

func PossibleValuesForShareAccessType() []string {
	return []string{
		string(ShareAccessTypeChange),
		string(ShareAccessTypeCustom),
		string(ShareAccessTypeRead),
	}
}

type ShareStatus string

const (
	ShareStatusNeedsAttention ShareStatus = "NeedsAttention"
	ShareStatusOK             ShareStatus = "OK"
	ShareStatusOffline        ShareStatus = "Offline"
	ShareStatusUnknown        ShareStatus = "Unknown"
	ShareStatusUpdating       ShareStatus = "Updating"
)

func PossibleValuesForShareStatus() []string {
	return []string{
		string(ShareStatusNeedsAttention),
		string(ShareStatusOK),
		string(ShareStatusOffline),
		string(ShareStatusUnknown),
		string(ShareStatusUpdating),
	}
}
