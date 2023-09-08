package meoutlooktaskfolder

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOutlookTaskFolderClient struct {
	Client *msgraph.Client
}

func NewMeOutlookTaskFolderClientWithBaseURI(api sdkEnv.Api) (*MeOutlookTaskFolderClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meoutlooktaskfolder", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOutlookTaskFolderClient: %+v", err)
	}

	return &MeOutlookTaskFolderClient{
		Client: client,
	}, nil
}
