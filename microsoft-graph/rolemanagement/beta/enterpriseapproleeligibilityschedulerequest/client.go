package enterpriseapproleeligibilityschedulerequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleEligibilityScheduleRequestClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleEligibilityScheduleRequestClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleEligibilityScheduleRequestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleeligibilityschedulerequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleEligibilityScheduleRequestClient: %+v", err)
	}

	return &EnterpriseAppRoleEligibilityScheduleRequestClient{
		Client: client,
	}, nil
}
