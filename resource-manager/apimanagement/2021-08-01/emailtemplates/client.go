package emailtemplates

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailTemplatesClient struct {
	Client *resourcemanager.Client
}

func NewEmailTemplatesClientWithBaseURI(api sdkEnv.Api) (*EmailTemplatesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "emailtemplates", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EmailTemplatesClient: %+v", err)
	}

	return &EmailTemplatesClient{
		Client: client,
	}, nil
}
