package usermailfolderchildfolderuserconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserMailFolderChildFolderUserConfigurationClient struct {
	Client *msgraph.Client
}

func NewUserMailFolderChildFolderUserConfigurationClientWithBaseURI(api sdkEnv.Api) (*UserMailFolderChildFolderUserConfigurationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermailfolderchildfolderuserconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserMailFolderChildFolderUserConfigurationClient: %+v", err)
	}

	return &UserMailFolderChildFolderUserConfigurationClient{
		Client: client,
	}, nil
}
