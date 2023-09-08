package userchatoperation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserChatOperationClient struct {
	Client *msgraph.Client
}

func NewUserChatOperationClientWithBaseURI(api sdkEnv.Api) (*UserChatOperationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userchatoperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserChatOperationClient: %+v", err)
	}

	return &UserChatOperationClient{
		Client: client,
	}, nil
}
