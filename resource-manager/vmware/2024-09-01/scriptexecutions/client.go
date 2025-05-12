package scriptexecutions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScriptExecutionsClient struct {
	Client *resourcemanager.Client
}

func NewScriptExecutionsClientWithBaseURI(sdkApi sdkEnv.Api) (*ScriptExecutionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "scriptexecutions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ScriptExecutionsClient: %+v", err)
	}

	return &ScriptExecutionsClient{
		Client: client,
	}, nil
}
