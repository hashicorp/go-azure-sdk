package mailfoldermessageextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailFolderMessageExtensionClient struct {
	Client *msgraph.Client
}

func NewMailFolderMessageExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*MailFolderMessageExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "mailfoldermessageextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MailFolderMessageExtensionClient: %+v", err)
	}

	return &MailFolderMessageExtensionClient{
		Client: client,
	}, nil
}
