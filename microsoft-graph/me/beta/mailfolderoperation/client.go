package mailfolderoperation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailFolderOperationClient struct {
	Client *msgraph.Client
}

func NewMailFolderOperationClientWithBaseURI(sdkApi sdkEnv.Api) (*MailFolderOperationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "mailfolderoperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MailFolderOperationClient: %+v", err)
	}

	return &MailFolderOperationClient{
		Client: client,
	}, nil
}
