package mechatmessage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeChatMessageClient struct {
	Client *msgraph.Client
}

func NewMeChatMessageClientWithBaseURI(api sdkEnv.Api) (*MeChatMessageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mechatmessage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeChatMessageClient: %+v", err)
	}

	return &MeChatMessageClient{
		Client: client,
	}, nil
}
