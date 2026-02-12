package v1_0

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/data-plane/appconfiguration/1-0/keys"
	"github.com/hashicorp/go-azure-sdk/data-plane/appconfiguration/1-0/keyvalues"
	"github.com/hashicorp/go-azure-sdk/data-plane/appconfiguration/1-0/labels"
	"github.com/hashicorp/go-azure-sdk/data-plane/appconfiguration/1-0/locks"
	"github.com/hashicorp/go-azure-sdk/data-plane/appconfiguration/1-0/revisions"
	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

type Client struct {
	KeyValues *keyvalues.KeyValuesClient
	Keys      *keys.KeysClient
	Labels    *labels.LabelsClient
	Locks     *locks.LocksClient
	Revisions *revisions.RevisionsClient
}

func NewClient(configureFunc func(c *dataplane.Client)) (*Client, error) {
	keyValuesClient, err := keyvalues.NewKeyValuesClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building KeyValues client: %+v", err)
	}
	configureFunc(keyValuesClient.Client)

	keysClient, err := keys.NewKeysClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Keys client: %+v", err)
	}
	configureFunc(keysClient.Client)

	labelsClient, err := labels.NewLabelsClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Labels client: %+v", err)
	}
	configureFunc(labelsClient.Client)

	locksClient, err := locks.NewLocksClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Locks client: %+v", err)
	}
	configureFunc(locksClient.Client)

	revisionsClient, err := revisions.NewRevisionsClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Revisions client: %+v", err)
	}
	configureFunc(revisionsClient.Client)

	return &Client{
		KeyValues: keyValuesClient,
		Keys:      keysClient,
		Labels:    labelsClient,
		Locks:     locksClient,
		Revisions: revisionsClient,
	}, nil
}
