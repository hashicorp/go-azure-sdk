package datasets

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatasetGetDatasetOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *DatasetResource
}

type DatasetGetDatasetOperationOptions struct {
	IfNoneMatch *string
}

func DefaultDatasetGetDatasetOperationOptions() DatasetGetDatasetOperationOptions {
	return DatasetGetDatasetOperationOptions{}
}

func (o DatasetGetDatasetOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfNoneMatch != nil {
		out.Append("If-None-Match", fmt.Sprintf("%v", *o.IfNoneMatch))
	}
	return &out
}

func (o DatasetGetDatasetOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DatasetGetDatasetOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DatasetGetDataset ...
func (c DatasetsClient) DatasetGetDataset(ctx context.Context, id DatasetId, options DatasetGetDatasetOperationOptions) (result DatasetGetDatasetOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.Path(),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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

	var model DatasetResource
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
