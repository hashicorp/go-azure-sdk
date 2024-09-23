package mailfolderchildfoldermessagemention

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailFolderChildFolderMessageMentionClient struct {
	Client *msgraph.Client
}

func NewMailFolderChildFolderMessageMentionClientWithBaseURI(sdkApi sdkEnv.Api) (*MailFolderChildFolderMessageMentionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "mailfolderchildfoldermessagemention", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MailFolderChildFolderMessageMentionClient: %+v", err)
	}

	return &MailFolderChildFolderMessageMentionClient{
		Client: client,
	}, nil
}
