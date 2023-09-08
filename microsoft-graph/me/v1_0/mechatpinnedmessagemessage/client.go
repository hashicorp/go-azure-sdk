package mechatpinnedmessagemessage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeChatPinnedMessageMessageClient struct {
	Client *msgraph.Client
}

func NewMeChatPinnedMessageMessageClientWithBaseURI(api sdkEnv.Api) (*MeChatPinnedMessageMessageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mechatpinnedmessagemessage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeChatPinnedMessageMessageClient: %+v", err)
	}

	return &MeChatPinnedMessageMessageClient{
		Client: client,
	}, nil
}
