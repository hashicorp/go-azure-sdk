package templatemigratableto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TemplateMigratableToClient struct {
	Client *msgraph.Client
}

func NewTemplateMigratableToClientWithBaseURI(sdkApi sdkEnv.Api) (*TemplateMigratableToClient, error) {
	client, err := msgraph.NewClient(sdkApi, "templatemigratableto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TemplateMigratableToClient: %+v", err)
	}

	return &TemplateMigratableToClient{
		Client: client,
	}, nil
}
