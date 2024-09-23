package enterpriseapproleassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleAssignmentClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleAssignmentClient: %+v", err)
	}

	return &EnterpriseAppRoleAssignmentClient{
		Client: client,
	}, nil
}
