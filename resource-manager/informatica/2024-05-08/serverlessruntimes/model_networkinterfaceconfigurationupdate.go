package serverlessruntimes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkInterfaceConfigurationUpdate struct {
	SubnetId         *string `json:"subnetId,omitempty"`
	VnetId           *string `json:"vnetId,omitempty"`
	VnetResourceGuid *string `json:"vnetResourceGuid,omitempty"`
}
