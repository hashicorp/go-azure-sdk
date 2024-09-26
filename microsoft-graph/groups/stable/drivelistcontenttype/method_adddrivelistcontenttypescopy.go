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

type AddDriveListContentTypesCopyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.ContentType
}

type AddDriveListContentTypesCopyOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultAddDriveListContentTypesCopyOperationOptions() AddDriveListContentTypesCopyOperationOptions {
	return AddDriveListContentTypesCopyOperationOptions{}
}

func (o AddDriveListContentTypesCopyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AddDriveListContentTypesCopyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AddDriveListContentTypesCopyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AddDriveListContentTypesCopy - Invoke action addCopy. Add a copy of a content type from a site to a list.
func (c DriveListContentTypeClient) AddDriveListContentTypesCopy(ctx context.Context, id stable.GroupIdDriveId, input AddDriveListContentTypesCopyRequest, options AddDriveListContentTypesCopyOperationOptions) (result AddDriveListContentTypesCopyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/list/contentTypes/addCopy", id.ID()),
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
