package mailfolderchildfolderuserconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailFolderChildFolderUserConfigurationClient struct {
	Client *msgraph.Client
}

func NewMailFolderChildFolderUserConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*MailFolderChildFolderUserConfigurationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "mailfolderchildfolderuserconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MailFolderChildFolderUserConfigurationClient: %+v", err)
	}

	return &MailFolderChildFolderUserConfigurationClient{
		Client: client,
	}, nil
}
