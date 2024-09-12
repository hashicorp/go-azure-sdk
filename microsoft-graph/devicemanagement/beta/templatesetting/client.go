package templatesetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TemplateSettingClient struct {
	Client *msgraph.Client
}

func NewTemplateSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*TemplateSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "templatesetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TemplateSettingClient: %+v", err)
	}

	return &TemplateSettingClient{
		Client: client,
	}, nil
}
