package memessagemention

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeMessageMentionClient struct {
	Client *msgraph.Client
}

func NewMeMessageMentionClientWithBaseURI(api sdkEnv.Api) (*MeMessageMentionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "memessagemention", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeMessageMentionClient: %+v", err)
	}

	return &MeMessageMentionClient{
		Client: client,
	}, nil
}
