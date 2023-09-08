package memailfoldermessagemention

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeMailFolderMessageMentionClient struct {
	Client *msgraph.Client
}

func NewMeMailFolderMessageMentionClientWithBaseURI(api sdkEnv.Api) (*MeMailFolderMessageMentionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "memailfoldermessagemention", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeMailFolderMessageMentionClient: %+v", err)
	}

	return &MeMailFolderMessageMentionClient{
		Client: client,
	}, nil
}
