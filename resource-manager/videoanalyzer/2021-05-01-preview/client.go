package v2021_05_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/videoanalyzer/2021-05-01-preview/edgemodules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/videoanalyzer/2021-05-01-preview/videoanalyzers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/videoanalyzer/2021-05-01-preview/videos"
)

type Client struct {
	EdgeModules    *edgemodules.EdgeModulesClient
	VideoAnalyzers *videoanalyzers.VideoAnalyzersClient
	Videos         *videos.VideosClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	edgeModulesClient := edgemodules.NewEdgeModulesClientWithBaseURI(endpoint)
	configureAuthFunc(&edgeModulesClient.Client)

	videoAnalyzersClient := videoanalyzers.NewVideoAnalyzersClientWithBaseURI(endpoint)
	configureAuthFunc(&videoAnalyzersClient.Client)

	videosClient := videos.NewVideosClientWithBaseURI(endpoint)
	configureAuthFunc(&videosClient.Client)

	return Client{
		EdgeModules:    &edgeModulesClient,
		VideoAnalyzers: &videoAnalyzersClient,
		Videos:         &videosClient,
	}
}
