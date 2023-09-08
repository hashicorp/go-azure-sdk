package meoutlooktaskfoldertaskattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOutlookTaskFolderTaskAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeOutlookTaskFolderTaskAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeOutlookTaskFolderTaskAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meoutlooktaskfoldertaskattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOutlookTaskFolderTaskAttachmentClient: %+v", err)
	}

	return &MeOutlookTaskFolderTaskAttachmentClient{
		Client: client,
	}, nil
}
