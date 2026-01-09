package enterpriseapproleeligibilityschedulerequestroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleEligibilityScheduleRequestRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleEligibilityScheduleRequestRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleEligibilityScheduleRequestRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleeligibilityschedulerequestroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleEligibilityScheduleRequestRoleDefinitionClient: %+v", err)
	}

	return &EnterpriseAppRoleEligibilityScheduleRequestRoleDefinitionClient{
		Client: client,
	}, nil
}
