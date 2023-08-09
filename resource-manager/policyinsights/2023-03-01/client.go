package v2023_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2023-03-01/checkpolicyrestrictions"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	CheckPolicyRestrictions *checkpolicyrestrictions.CheckPolicyRestrictionsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	checkPolicyRestrictionsClient, err := checkpolicyrestrictions.NewCheckPolicyRestrictionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CheckPolicyRestrictions client: %+v", err)
	}
	configureFunc(checkPolicyRestrictionsClient.Client)

	return &Client{
		CheckPolicyRestrictions: checkPolicyRestrictionsClient,
	}, nil
}
