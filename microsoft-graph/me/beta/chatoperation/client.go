package chatoperation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatOperationClient struct {
	Client *msgraph.Client
}

func NewChatOperationClientWithBaseURI(sdkApi sdkEnv.Api) (*ChatOperationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "chatoperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ChatOperationClient: %+v", err)
	}

	return &ChatOperationClient{
		Client: client,
	}, nil
}
