package apiissueattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiIssueAttachmentClient struct {
	Client *resourcemanager.Client
}

func NewApiIssueAttachmentClientWithBaseURI(api environments.Api) (*ApiIssueAttachmentClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "apiissueattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApiIssueAttachmentClient: %+v", err)
	}

	return &ApiIssueAttachmentClient{
		Client: client,
	}, nil
}
