package userauthenticationoperation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAuthenticationOperationClient struct {
	Client *msgraph.Client
}

func NewUserAuthenticationOperationClientWithBaseURI(api sdkEnv.Api) (*UserAuthenticationOperationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userauthenticationoperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserAuthenticationOperationClient: %+v", err)
	}

	return &UserAuthenticationOperationClient{
		Client: client,
	}, nil
}
