package v2023_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/sqlvirtualmachine/2023-10-01/availabilitygrouplisteners"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sqlvirtualmachine/2023-10-01/sqlvirtualmachinegroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sqlvirtualmachine/2023-10-01/sqlvirtualmachines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sqlvirtualmachine/2023-10-01/sqlvirtualmachinetroubleshoot"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AvailabilityGroupListeners    *availabilitygrouplisteners.AvailabilityGroupListenersClient
	SqlVirtualMachineGroups       *sqlvirtualmachinegroups.SqlVirtualMachineGroupsClient
	SqlVirtualMachineTroubleshoot *sqlvirtualmachinetroubleshoot.SqlVirtualMachineTroubleshootClient
	SqlVirtualMachines            *sqlvirtualmachines.SqlVirtualMachinesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	availabilityGroupListenersClient, err := availabilitygrouplisteners.NewAvailabilityGroupListenersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AvailabilityGroupListeners client: %+v", err)
	}
	configureFunc(availabilityGroupListenersClient.Client)

	sqlVirtualMachineGroupsClient, err := sqlvirtualmachinegroups.NewSqlVirtualMachineGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlVirtualMachineGroups client: %+v", err)
	}
	configureFunc(sqlVirtualMachineGroupsClient.Client)

	sqlVirtualMachineTroubleshootClient, err := sqlvirtualmachinetroubleshoot.NewSqlVirtualMachineTroubleshootClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlVirtualMachineTroubleshoot client: %+v", err)
	}
	configureFunc(sqlVirtualMachineTroubleshootClient.Client)

	sqlVirtualMachinesClient, err := sqlvirtualmachines.NewSqlVirtualMachinesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlVirtualMachines client: %+v", err)
	}
	configureFunc(sqlVirtualMachinesClient.Client)

	return &Client{
		AvailabilityGroupListeners:    availabilityGroupListenersClient,
		SqlVirtualMachineGroups:       sqlVirtualMachineGroupsClient,
		SqlVirtualMachineTroubleshoot: sqlVirtualMachineTroubleshootClient,
		SqlVirtualMachines:            sqlVirtualMachinesClient,
	}, nil
}
