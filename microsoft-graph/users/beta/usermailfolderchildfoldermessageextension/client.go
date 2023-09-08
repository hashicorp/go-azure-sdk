package usermailfolderchildfoldermessageextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserMailFolderChildFolderMessageExtensionClient struct {
	Client *msgraph.Client
}

func NewUserMailFolderChildFolderMessageExtensionClientWithBaseURI(api sdkEnv.Api) (*UserMailFolderChildFolderMessageExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermailfolderchildfoldermessageextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserMailFolderChildFolderMessageExtensionClient: %+v", err)
	}

	return &UserMailFolderChildFolderMessageExtensionClient{
		Client: client,
	}, nil
}
