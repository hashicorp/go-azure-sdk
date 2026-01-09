package messagemention

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageMentionClient struct {
	Client *msgraph.Client
}

func NewMessageMentionClientWithBaseURI(sdkApi sdkEnv.Api) (*MessageMentionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "messagemention", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MessageMentionClient: %+v", err)
	}

	return &MessageMentionClient{
		Client: client,
	}, nil
}
