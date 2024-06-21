package containers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerProperties struct {
	ContainerStatus *ContainerStatus         `json:"containerStatus,omitempty"`
	CreatedDateTime *string                  `json:"createdDateTime,omitempty"`
	DataFormat      AzureContainerDataFormat `json:"dataFormat"`
	RefreshDetails  *RefreshDetails          `json:"refreshDetails,omitempty"`
}
