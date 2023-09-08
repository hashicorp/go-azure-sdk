package groupthreadpostinreplytomention

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupThreadPostInReplyToMentionClient struct {
	Client *msgraph.Client
}

func NewGroupThreadPostInReplyToMentionClientWithBaseURI(api sdkEnv.Api) (*GroupThreadPostInReplyToMentionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupthreadpostinreplytomention", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupThreadPostInReplyToMentionClient: %+v", err)
	}

	return &GroupThreadPostInReplyToMentionClient{
		Client: client,
	}, nil
}
