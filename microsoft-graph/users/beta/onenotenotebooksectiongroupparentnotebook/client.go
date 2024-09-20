package onenotenotebooksectiongroupparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteNotebookSectionGroupParentNotebookClient struct {
	Client *msgraph.Client
}

func NewOnenoteNotebookSectionGroupParentNotebookClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteNotebookSectionGroupParentNotebookClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotenotebooksectiongroupparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteNotebookSectionGroupParentNotebookClient: %+v", err)
	}

	return &OnenoteNotebookSectionGroupParentNotebookClient{
		Client: client,
	}, nil
}
