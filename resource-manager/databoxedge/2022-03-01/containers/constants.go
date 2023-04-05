package containers

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

type ContainerStatus string

const (
	ContainerStatusNeedsAttention ContainerStatus = "NeedsAttention"
	ContainerStatusOK             ContainerStatus = "OK"
	ContainerStatusOffline        ContainerStatus = "Offline"
	ContainerStatusUnknown        ContainerStatus = "Unknown"
	ContainerStatusUpdating       ContainerStatus = "Updating"
)

func PossibleValuesForContainerStatus() []string {
	return []string{
		string(ContainerStatusNeedsAttention),
		string(ContainerStatusOK),
		string(ContainerStatusOffline),
		string(ContainerStatusUnknown),
		string(ContainerStatusUpdating),
	}
}
