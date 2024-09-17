package scriptcmdlets

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScriptCmdletsClient struct {
	Client *resourcemanager.Client
}

func NewScriptCmdletsClientWithBaseURI(sdkApi sdkEnv.Api) (*ScriptCmdletsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "scriptcmdlets", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ScriptCmdletsClient: %+v", err)
	}

	return &ScriptCmdletsClient{
		Client: client,
	}, nil
}
