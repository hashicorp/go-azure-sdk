package usermailfolderuserconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserMailFolderUserConfigurationClient struct {
	Client *msgraph.Client
}

func NewUserMailFolderUserConfigurationClientWithBaseURI(api sdkEnv.Api) (*UserMailFolderUserConfigurationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermailfolderuserconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserMailFolderUserConfigurationClient: %+v", err)
	}

	return &UserMailFolderUserConfigurationClient{
		Client: client,
	}, nil
}
