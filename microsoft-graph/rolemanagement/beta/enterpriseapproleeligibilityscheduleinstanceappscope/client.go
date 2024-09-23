package enterpriseapproleeligibilityscheduleinstanceappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleEligibilityScheduleInstanceAppScopeClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleEligibilityScheduleInstanceAppScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleEligibilityScheduleInstanceAppScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleeligibilityscheduleinstanceappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleEligibilityScheduleInstanceAppScopeClient: %+v", err)
	}

	return &EnterpriseAppRoleEligibilityScheduleInstanceAppScopeClient{
		Client: client,
	}, nil
}
