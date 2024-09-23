package enterpriseapproleeligibilityscheduleinstanceroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleEligibilityScheduleInstanceRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleEligibilityScheduleInstanceRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleEligibilityScheduleInstanceRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleeligibilityscheduleinstanceroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleEligibilityScheduleInstanceRoleDefinitionClient: %+v", err)
	}

	return &EnterpriseAppRoleEligibilityScheduleInstanceRoleDefinitionClient{
		Client: client,
	}, nil
}
