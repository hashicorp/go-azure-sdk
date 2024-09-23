package enterpriseapp

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppClient: %+v", err)
	}

	return &EnterpriseAppClient{
		Client: client,
	}, nil
}
