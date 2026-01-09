package privilegedaccessgroupassignmentscheduleinstancegroupserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleInstanceGroupServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccessGroupAssignmentScheduleInstanceGroupServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccessGroupAssignmentScheduleInstanceGroupServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccessgroupassignmentscheduleinstancegroupserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccessGroupAssignmentScheduleInstanceGroupServiceProvisioningErrorClient: %+v", err)
	}

	return &PrivilegedAccessGroupAssignmentScheduleInstanceGroupServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
