package privilegedaccessgroupeligibilityschedulerequestgroupserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupEligibilityScheduleRequestGroupServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccessGroupEligibilityScheduleRequestGroupServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccessGroupEligibilityScheduleRequestGroupServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccessgroupeligibilityschedulerequestgroupserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccessGroupEligibilityScheduleRequestGroupServiceProvisioningErrorClient: %+v", err)
	}

	return &PrivilegedAccessGroupEligibilityScheduleRequestGroupServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
