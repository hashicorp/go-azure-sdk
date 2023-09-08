package usermailfolderchildfoldermessagemention

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserMailFolderChildFolderMessageMentionClient struct {
	Client *msgraph.Client
}

func NewUserMailFolderChildFolderMessageMentionClientWithBaseURI(api sdkEnv.Api) (*UserMailFolderChildFolderMessageMentionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermailfolderchildfoldermessagemention", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserMailFolderChildFolderMessageMentionClient: %+v", err)
	}

	return &UserMailFolderChildFolderMessageMentionClient{
		Client: client,
	}, nil
}
