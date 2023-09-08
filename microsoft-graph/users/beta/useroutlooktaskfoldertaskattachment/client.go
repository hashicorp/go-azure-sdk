package useroutlooktaskfoldertaskattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOutlookTaskFolderTaskAttachmentClient struct {
	Client *msgraph.Client
}

func NewUserOutlookTaskFolderTaskAttachmentClientWithBaseURI(api sdkEnv.Api) (*UserOutlookTaskFolderTaskAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useroutlooktaskfoldertaskattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOutlookTaskFolderTaskAttachmentClient: %+v", err)
	}

	return &UserOutlookTaskFolderTaskAttachmentClient{
		Client: client,
	}, nil
}
