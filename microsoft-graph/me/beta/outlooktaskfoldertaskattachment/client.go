package outlooktaskfoldertaskattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutlookTaskFolderTaskAttachmentClient struct {
	Client *msgraph.Client
}

func NewOutlookTaskFolderTaskAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*OutlookTaskFolderTaskAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "outlooktaskfoldertaskattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OutlookTaskFolderTaskAttachmentClient: %+v", err)
	}

	return &OutlookTaskFolderTaskAttachmentClient{
		Client: client,
	}, nil
}
