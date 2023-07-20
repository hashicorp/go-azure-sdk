package v2021_05_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/videoanalyzer/2021-05-01-preview/edgemodules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/videoanalyzer/2021-05-01-preview/videoanalyzers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/videoanalyzer/2021-05-01-preview/videos"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	EdgeModules    *edgemodules.EdgeModulesClient
	VideoAnalyzers *videoanalyzers.VideoAnalyzersClient
	Videos         *videos.VideosClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	edgeModulesClient, err := edgemodules.NewEdgeModulesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building EdgeModules client: %+v", err)
	}
	configureFunc(edgeModulesClient.Client)

	videoAnalyzersClient, err := videoanalyzers.NewVideoAnalyzersClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building VideoAnalyzers client: %+v", err)
	}
	configureFunc(videoAnalyzersClient.Client)

	videosClient, err := videos.NewVideosClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Videos client: %+v", err)
	}
	configureFunc(videosClient.Client)

	return &Client{
		EdgeModules:    edgeModulesClient,
		VideoAnalyzers: videoAnalyzersClient,
		Videos:         videosClient,
	}, nil
}
