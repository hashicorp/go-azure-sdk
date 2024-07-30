package privilegedaccesgroupeligibilityschedulerequesttargetschedule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccesGroupEligibilityScheduleRequestTargetScheduleClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccesGroupEligibilityScheduleRequestTargetScheduleClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccesGroupEligibilityScheduleRequestTargetScheduleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccesgroupeligibilityschedulerequesttargetschedule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccesGroupEligibilityScheduleRequestTargetScheduleClient: %+v", err)
	}

	return &PrivilegedAccesGroupEligibilityScheduleRequestTargetScheduleClient{
		Client: client,
	}, nil
}
