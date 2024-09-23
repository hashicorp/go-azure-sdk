package appmanagementpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppManagementPolicyClient struct {
	Client *msgraph.Client
}

func NewAppManagementPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*AppManagementPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "appmanagementpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppManagementPolicyClient: %+v", err)
	}

	return &AppManagementPolicyClient{
		Client: client,
	}, nil
}
