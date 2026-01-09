package mailfoldermessage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailFolderMessageClient struct {
	Client *msgraph.Client
}

func NewMailFolderMessageClientWithBaseURI(sdkApi sdkEnv.Api) (*MailFolderMessageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "mailfoldermessage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MailFolderMessageClient: %+v", err)
	}

	return &MailFolderMessageClient{
		Client: client,
	}, nil
}
