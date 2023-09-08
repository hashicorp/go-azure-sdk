package usermailfoldermessagemention

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserMailFolderMessageMentionClient struct {
	Client *msgraph.Client
}

func NewUserMailFolderMessageMentionClientWithBaseURI(api sdkEnv.Api) (*UserMailFolderMessageMentionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermailfoldermessagemention", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserMailFolderMessageMentionClient: %+v", err)
	}

	return &UserMailFolderMessageMentionClient{
		Client: client,
	}, nil
}
