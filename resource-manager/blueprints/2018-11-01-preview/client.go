package v2018_11_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/blueprints/2018-11-01-preview/artifact"
	"github.com/hashicorp/go-azure-sdk/resource-manager/blueprints/2018-11-01-preview/assignment"
	"github.com/hashicorp/go-azure-sdk/resource-manager/blueprints/2018-11-01-preview/assignmentoperations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/blueprints/2018-11-01-preview/blueprint"
	"github.com/hashicorp/go-azure-sdk/resource-manager/blueprints/2018-11-01-preview/blueprintassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/blueprints/2018-11-01-preview/publishedartifact"
	"github.com/hashicorp/go-azure-sdk/resource-manager/blueprints/2018-11-01-preview/publishedblueprint"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Artifact             *artifact.ArtifactClient
	Assignment           *assignment.AssignmentClient
	AssignmentOperations *assignmentoperations.AssignmentOperationsClient
	Blueprint            *blueprint.BlueprintClient
	BlueprintAssignments *blueprintassignments.BlueprintAssignmentsClient
	PublishedArtifact    *publishedartifact.PublishedArtifactClient
	PublishedBlueprint   *publishedblueprint.PublishedBlueprintClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	artifactClient, err := artifact.NewArtifactClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Artifact client: %+v", err)
	}
	configureFunc(artifactClient.Client)

	assignmentClient, err := assignment.NewAssignmentClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Assignment client: %+v", err)
	}
	configureFunc(assignmentClient.Client)

	assignmentOperationsClient, err := assignmentoperations.NewAssignmentOperationsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building AssignmentOperations client: %+v", err)
	}
	configureFunc(assignmentOperationsClient.Client)

	blueprintAssignmentsClient, err := blueprintassignments.NewBlueprintAssignmentsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building BlueprintAssignments client: %+v", err)
	}
	configureFunc(blueprintAssignmentsClient.Client)

	blueprintClient, err := blueprint.NewBlueprintClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Blueprint client: %+v", err)
	}
	configureFunc(blueprintClient.Client)

	publishedArtifactClient, err := publishedartifact.NewPublishedArtifactClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PublishedArtifact client: %+v", err)
	}
	configureFunc(publishedArtifactClient.Client)

	publishedBlueprintClient, err := publishedblueprint.NewPublishedBlueprintClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PublishedBlueprint client: %+v", err)
	}
	configureFunc(publishedBlueprintClient.Client)

	return &Client{
		Artifact:             artifactClient,
		Assignment:           assignmentClient,
		AssignmentOperations: assignmentOperationsClient,
		Blueprint:            blueprintClient,
		BlueprintAssignments: blueprintAssignmentsClient,
		PublishedArtifact:    publishedArtifactClient,
		PublishedBlueprint:   publishedBlueprintClient,
	}, nil
}
