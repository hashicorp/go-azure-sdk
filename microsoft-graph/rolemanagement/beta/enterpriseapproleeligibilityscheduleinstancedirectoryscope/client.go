package enterpriseapproleeligibilityscheduleinstancedirectoryscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleEligibilityScheduleInstanceDirectoryScopeClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleEligibilityScheduleInstanceDirectoryScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleEligibilityScheduleInstanceDirectoryScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleeligibilityscheduleinstancedirectoryscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleEligibilityScheduleInstanceDirectoryScopeClient: %+v", err)
	}

	return &EnterpriseAppRoleEligibilityScheduleInstanceDirectoryScopeClient{
		Client: client,
	}, nil
}
