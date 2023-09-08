package meauthenticationoperation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeAuthenticationOperationClient struct {
	Client *msgraph.Client
}

func NewMeAuthenticationOperationClientWithBaseURI(api sdkEnv.Api) (*MeAuthenticationOperationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meauthenticationoperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeAuthenticationOperationClient: %+v", err)
	}

	return &MeAuthenticationOperationClient{
		Client: client,
	}, nil
}
