package v2022_02_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sqlvirtualmachine/2022-02-01/availabilitygrouplisteners"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sqlvirtualmachine/2022-02-01/sqlvirtualmachinegroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sqlvirtualmachine/2022-02-01/sqlvirtualmachines"
)

type Client struct {
	AvailabilityGroupListeners *availabilitygrouplisteners.AvailabilityGroupListenersClient
	SqlVirtualMachineGroups    *sqlvirtualmachinegroups.SqlVirtualMachineGroupsClient
	SqlVirtualMachines         *sqlvirtualmachines.SqlVirtualMachinesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	availabilityGroupListenersClient := availabilitygrouplisteners.NewAvailabilityGroupListenersClientWithBaseURI(endpoint)
	configureAuthFunc(&availabilityGroupListenersClient.Client)

	sqlVirtualMachineGroupsClient := sqlvirtualmachinegroups.NewSqlVirtualMachineGroupsClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlVirtualMachineGroupsClient.Client)

	sqlVirtualMachinesClient := sqlvirtualmachines.NewSqlVirtualMachinesClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlVirtualMachinesClient.Client)

	return Client{
		AvailabilityGroupListeners: &availabilityGroupListenersClient,
		SqlVirtualMachineGroups:    &sqlVirtualMachineGroupsClient,
		SqlVirtualMachines:         &sqlVirtualMachinesClient,
	}
}
