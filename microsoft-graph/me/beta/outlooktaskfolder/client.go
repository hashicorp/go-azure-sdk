package outlooktaskfolder

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutlookTaskFolderClient struct {
	Client *msgraph.Client
}

func NewOutlookTaskFolderClientWithBaseURI(sdkApi sdkEnv.Api) (*OutlookTaskFolderClient, error) {
	client, err := msgraph.NewClient(sdkApi, "outlooktaskfolder", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OutlookTaskFolderClient: %+v", err)
	}

	return &OutlookTaskFolderClient{
		Client: client,
	}, nil
}
