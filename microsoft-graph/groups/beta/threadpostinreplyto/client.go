package threadpostinreplyto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreadPostInReplyToClient struct {
	Client *msgraph.Client
}

func NewThreadPostInReplyToClientWithBaseURI(sdkApi sdkEnv.Api) (*ThreadPostInReplyToClient, error) {
	client, err := msgraph.NewClient(sdkApi, "threadpostinreplyto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ThreadPostInReplyToClient: %+v", err)
	}

	return &ThreadPostInReplyToClient{
		Client: client,
	}, nil
}
