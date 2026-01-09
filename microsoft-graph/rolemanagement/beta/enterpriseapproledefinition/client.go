package enterpriseapproledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleDefinitionClient: %+v", err)
	}

	return &EnterpriseAppRoleDefinitionClient{
		Client: client,
	}, nil
}
