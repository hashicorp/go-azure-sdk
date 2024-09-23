package enterpriseapproleeligibilityscheduledirectoryscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleEligibilityScheduleDirectoryScopeClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleEligibilityScheduleDirectoryScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleEligibilityScheduleDirectoryScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleeligibilityscheduledirectoryscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleEligibilityScheduleDirectoryScopeClient: %+v", err)
	}

	return &EnterpriseAppRoleEligibilityScheduleDirectoryScopeClient{
		Client: client,
	}, nil
}
