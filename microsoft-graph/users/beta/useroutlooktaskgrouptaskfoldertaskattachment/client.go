package useroutlooktaskgrouptaskfoldertaskattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOutlookTaskGroupTaskFolderTaskAttachmentClient struct {
	Client *msgraph.Client
}

func NewUserOutlookTaskGroupTaskFolderTaskAttachmentClientWithBaseURI(api sdkEnv.Api) (*UserOutlookTaskGroupTaskFolderTaskAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useroutlooktaskgrouptaskfoldertaskattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOutlookTaskGroupTaskFolderTaskAttachmentClient: %+v", err)
	}

	return &UserOutlookTaskGroupTaskFolderTaskAttachmentClient{
		Client: client,
	}, nil
}
