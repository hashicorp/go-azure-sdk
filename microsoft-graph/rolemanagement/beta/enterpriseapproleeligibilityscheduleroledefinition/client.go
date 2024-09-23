package enterpriseapproleeligibilityscheduleroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleEligibilityScheduleRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleEligibilityScheduleRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleEligibilityScheduleRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleeligibilityscheduleroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleEligibilityScheduleRoleDefinitionClient: %+v", err)
	}

	return &EnterpriseAppRoleEligibilityScheduleRoleDefinitionClient{
		Client: client,
	}, nil
}
