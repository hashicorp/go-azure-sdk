package mailfoldermessagemention

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailFolderMessageMentionClient struct {
	Client *msgraph.Client
}

func NewMailFolderMessageMentionClientWithBaseURI(sdkApi sdkEnv.Api) (*MailFolderMessageMentionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "mailfoldermessagemention", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MailFolderMessageMentionClient: %+v", err)
	}

	return &MailFolderMessageMentionClient{
		Client: client,
	}, nil
}
