package mechatoperation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeChatOperationClient struct {
	Client *msgraph.Client
}

func NewMeChatOperationClientWithBaseURI(api sdkEnv.Api) (*MeChatOperationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mechatoperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeChatOperationClient: %+v", err)
	}

	return &MeChatOperationClient{
		Client: client,
	}, nil
}
