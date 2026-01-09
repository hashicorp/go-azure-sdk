package privilegedaccessgroupeligibilityschedulerequesttargetschedule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupEligibilityScheduleRequestTargetScheduleClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccessGroupEligibilityScheduleRequestTargetScheduleClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccessGroupEligibilityScheduleRequestTargetScheduleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccessgroupeligibilityschedulerequesttargetschedule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccessGroupEligibilityScheduleRequestTargetScheduleClient: %+v", err)
	}

	return &PrivilegedAccessGroupEligibilityScheduleRequestTargetScheduleClient{
		Client: client,
	}, nil
}
