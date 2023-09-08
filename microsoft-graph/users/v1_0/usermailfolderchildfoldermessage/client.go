package usermailfolderchildfoldermessage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserMailFolderChildFolderMessageClient struct {
	Client *msgraph.Client
}

func NewUserMailFolderChildFolderMessageClientWithBaseURI(api sdkEnv.Api) (*UserMailFolderChildFolderMessageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermailfolderchildfoldermessage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserMailFolderChildFolderMessageClient: %+v", err)
	}

	return &UserMailFolderChildFolderMessageClient{
		Client: client,
	}, nil
}
