package usermailfoldermessageattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserMailFolderMessageAttachmentClient struct {
	Client *msgraph.Client
}

func NewUserMailFolderMessageAttachmentClientWithBaseURI(api sdkEnv.Api) (*UserMailFolderMessageAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermailfoldermessageattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserMailFolderMessageAttachmentClient: %+v", err)
	}

	return &UserMailFolderMessageAttachmentClient{
		Client: client,
	}, nil
}
