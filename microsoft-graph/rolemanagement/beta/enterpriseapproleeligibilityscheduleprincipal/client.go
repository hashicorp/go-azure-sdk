package enterpriseapproleeligibilityscheduleprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleEligibilitySchedulePrincipalClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleEligibilitySchedulePrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleEligibilitySchedulePrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleeligibilityscheduleprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleEligibilitySchedulePrincipalClient: %+v", err)
	}

	return &EnterpriseAppRoleEligibilitySchedulePrincipalClient{
		Client: client,
	}, nil
}
