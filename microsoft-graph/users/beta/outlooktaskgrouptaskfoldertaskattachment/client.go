package outlooktaskgrouptaskfoldertaskattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutlookTaskGroupTaskFolderTaskAttachmentClient struct {
	Client *msgraph.Client
}

func NewOutlookTaskGroupTaskFolderTaskAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*OutlookTaskGroupTaskFolderTaskAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "outlooktaskgrouptaskfoldertaskattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OutlookTaskGroupTaskFolderTaskAttachmentClient: %+v", err)
	}

	return &OutlookTaskGroupTaskFolderTaskAttachmentClient{
		Client: client,
	}, nil
}
