package enterpriseapproleeligibilityschedulerequestappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleEligibilityScheduleRequestAppScopeClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleEligibilityScheduleRequestAppScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleEligibilityScheduleRequestAppScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleeligibilityschedulerequestappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleEligibilityScheduleRequestAppScopeClient: %+v", err)
	}

	return &EnterpriseAppRoleEligibilityScheduleRequestAppScopeClient{
		Client: client,
	}, nil
}
