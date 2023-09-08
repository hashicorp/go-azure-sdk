package userchatpinnedmessagemessage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserChatPinnedMessageMessageClient struct {
	Client *msgraph.Client
}

func NewUserChatPinnedMessageMessageClientWithBaseURI(api sdkEnv.Api) (*UserChatPinnedMessageMessageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userchatpinnedmessagemessage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserChatPinnedMessageMessageClient: %+v", err)
	}

	return &UserChatPinnedMessageMessageClient{
		Client: client,
	}, nil
}
