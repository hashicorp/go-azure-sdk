package drivelistcontenttype

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddDriveListContentTypesCopyFromHubOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.ContentType
}

type AddDriveListContentTypesCopyFromHubOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultAddDriveListContentTypesCopyFromHubOperationOptions() AddDriveListContentTypesCopyFromHubOperationOptions {
	return AddDriveListContentTypesCopyFromHubOperationOptions{}
}

func (o AddDriveListContentTypesCopyFromHubOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AddDriveListContentTypesCopyFromHubOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AddDriveListContentTypesCopyFromHubOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AddDriveListContentTypesCopyFromHub - Invoke action addCopyFromContentTypeHub. Add or sync a copy of a published
// content type from the content type hub to a target site or a list. This method is part of the content type publishing
// changes to optimize the syncing of published content types to sites and lists, effectively switching from a 'push
// everywhere' to 'pull as needed' approach. The method allows users to pull content types directly from the content
// type hub to a site or list. For more information, see contentType: getCompatibleHubContentTypes and the blog post
// Syntex Product Updates – August 2021.
func (c DriveListContentTypeClient) AddDriveListContentTypesCopyFromHub(ctx context.Context, id stable.MeDriveId, input AddDriveListContentTypesCopyFromHubRequest, options AddDriveListContentTypesCopyFromHubOperationOptions) (result AddDriveListContentTypesCopyFromHubOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/list/contentTypes/addCopyFromContentTypeHub", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	var model stable.ContentType
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
