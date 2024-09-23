package enterpriseapproleeligibilityscheduleappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleEligibilityScheduleAppScopeClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleEligibilityScheduleAppScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleEligibilityScheduleAppScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleeligibilityscheduleappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleEligibilityScheduleAppScopeClient: %+v", err)
	}

	return &EnterpriseAppRoleEligibilityScheduleAppScopeClient{
		Client: client,
	}, nil
}
