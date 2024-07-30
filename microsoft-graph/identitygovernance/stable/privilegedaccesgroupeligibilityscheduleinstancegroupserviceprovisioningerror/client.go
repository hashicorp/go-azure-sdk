package privilegedaccesgroupeligibilityscheduleinstancegroupserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccesGroupEligibilityScheduleInstanceGroupServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccesGroupEligibilityScheduleInstanceGroupServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccesGroupEligibilityScheduleInstanceGroupServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccesgroupeligibilityscheduleinstancegroupserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccesGroupEligibilityScheduleInstanceGroupServiceProvisioningErrorClient: %+v", err)
	}

	return &PrivilegedAccesGroupEligibilityScheduleInstanceGroupServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
