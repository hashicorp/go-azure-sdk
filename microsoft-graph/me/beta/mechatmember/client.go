package mechatmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeChatMemberClient struct {
	Client *msgraph.Client
}

func NewMeChatMemberClientWithBaseURI(api sdkEnv.Api) (*MeChatMemberClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mechatmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeChatMemberClient: %+v", err)
	}

	return &MeChatMemberClient{
		Client: client,
	}, nil
}
