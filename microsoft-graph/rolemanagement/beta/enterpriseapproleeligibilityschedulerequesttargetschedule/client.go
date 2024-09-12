package enterpriseapproleeligibilityschedulerequesttargetschedule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleEligibilityScheduleRequestTargetScheduleClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleEligibilityScheduleRequestTargetScheduleClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleEligibilityScheduleRequestTargetScheduleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleeligibilityschedulerequesttargetschedule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleEligibilityScheduleRequestTargetScheduleClient: %+v", err)
	}

	return &EnterpriseAppRoleEligibilityScheduleRequestTargetScheduleClient{
		Client: client,
	}, nil
}
