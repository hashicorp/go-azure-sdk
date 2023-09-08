package groupthreadpostmention

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupThreadPostMentionClient struct {
	Client *msgraph.Client
}

func NewGroupThreadPostMentionClientWithBaseURI(api sdkEnv.Api) (*GroupThreadPostMentionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupthreadpostmention", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupThreadPostMentionClient: %+v", err)
	}

	return &GroupThreadPostMentionClient{
		Client: client,
	}, nil
}
