package useronenoteoperation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteOperationClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteOperationClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteOperationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenoteoperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteOperationClient: %+v", err)
	}

	return &UserOnenoteOperationClient{
		Client: client,
	}, nil
}
