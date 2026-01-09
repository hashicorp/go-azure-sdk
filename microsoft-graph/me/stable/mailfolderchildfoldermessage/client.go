package mailfolderchildfoldermessage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailFolderChildFolderMessageClient struct {
	Client *msgraph.Client
}

func NewMailFolderChildFolderMessageClientWithBaseURI(sdkApi sdkEnv.Api) (*MailFolderChildFolderMessageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "mailfolderchildfoldermessage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MailFolderChildFolderMessageClient: %+v", err)
	}

	return &MailFolderChildFolderMessageClient{
		Client: client,
	}, nil
}
