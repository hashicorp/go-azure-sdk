package usermailfolderchildfoldermessageattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserMailFolderChildFolderMessageAttachmentClient struct {
	Client *msgraph.Client
}

func NewUserMailFolderChildFolderMessageAttachmentClientWithBaseURI(api sdkEnv.Api) (*UserMailFolderChildFolderMessageAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermailfolderchildfoldermessageattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserMailFolderChildFolderMessageAttachmentClient: %+v", err)
	}

	return &UserMailFolderChildFolderMessageAttachmentClient{
		Client: client,
	}, nil
}
