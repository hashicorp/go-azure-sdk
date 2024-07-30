package appmanagementpolicyappliesto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppManagementPolicyAppliesToClient struct {
	Client *msgraph.Client
}

func NewAppManagementPolicyAppliesToClientWithBaseURI(sdkApi sdkEnv.Api) (*AppManagementPolicyAppliesToClient, error) {
	client, err := msgraph.NewClient(sdkApi, "appmanagementpolicyappliesto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppManagementPolicyAppliesToClient: %+v", err)
	}

	return &AppManagementPolicyAppliesToClient{
		Client: client,
	}, nil
}
