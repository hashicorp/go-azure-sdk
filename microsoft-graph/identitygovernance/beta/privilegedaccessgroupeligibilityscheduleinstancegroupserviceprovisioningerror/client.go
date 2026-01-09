package privilegedaccessgroupeligibilityscheduleinstancegroupserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupEligibilityScheduleInstanceGroupServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccessGroupEligibilityScheduleInstanceGroupServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccessGroupEligibilityScheduleInstanceGroupServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccessgroupeligibilityscheduleinstancegroupserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccessGroupEligibilityScheduleInstanceGroupServiceProvisioningErrorClient: %+v", err)
	}

	return &PrivilegedAccessGroupEligibilityScheduleInstanceGroupServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
