package privilegedaccessgroupeligibilityschedulerequestgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupEligibilityScheduleRequestGroupClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccessGroupEligibilityScheduleRequestGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccessGroupEligibilityScheduleRequestGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccessgroupeligibilityschedulerequestgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccessGroupEligibilityScheduleRequestGroupClient: %+v", err)
	}

	return &PrivilegedAccessGroupEligibilityScheduleRequestGroupClient{
		Client: client,
	}, nil
}
