package sitecontenttype

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddSiteContentTypesCopyFromHubOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ContentType
}

// AddSiteContentTypesCopyFromHub - Invoke action addCopyFromContentTypeHub. Add or sync a copy of a published content
// type from the content type hub to a target site or a list. This method is part of the content type publishing changes
// to optimize the syncing of published content types to sites and lists, effectively switching from a 'push everywhere'
// to 'pull as needed' approach. The method allows users to pull content types directly from the content type hub to a
// site or list. For more information, see getCompatibleHubContentTypes and the blog post Syntex Product Updates –
// August 2021.
func (c SiteContentTypeClient) AddSiteContentTypesCopyFromHub(ctx context.Context, id beta.GroupIdSiteId, input AddSiteContentTypesCopyFromHubRequest) (result AddSiteContentTypesCopyFromHubOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/contentTypes/addCopyFromContentTypeHub", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var model beta.ContentType
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
