package emailtemplate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailTemplateClient struct {
	Client *resourcemanager.Client
}

func NewEmailTemplateClientWithBaseURI(api environments.Api) (*EmailTemplateClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "emailtemplate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EmailTemplateClient: %+v", err)
	}

	return &EmailTemplateClient{
		Client: client,
	}, nil
}
