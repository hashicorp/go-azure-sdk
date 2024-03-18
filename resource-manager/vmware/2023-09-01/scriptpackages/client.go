package scriptpackages

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScriptPackagesClient struct {
	Client *resourcemanager.Client
}

func NewScriptPackagesClientWithBaseURI(sdkApi sdkEnv.Api) (*ScriptPackagesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "scriptpackages", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ScriptPackagesClient: %+v", err)
	}

	return &ScriptPackagesClient{
		Client: client,
	}, nil
}
