package mailfolderuserconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailFolderUserConfigurationClient struct {
	Client *msgraph.Client
}

func NewMailFolderUserConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*MailFolderUserConfigurationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "mailfolderuserconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MailFolderUserConfigurationClient: %+v", err)
	}

	return &MailFolderUserConfigurationClient{
		Client: client,
	}, nil
}
