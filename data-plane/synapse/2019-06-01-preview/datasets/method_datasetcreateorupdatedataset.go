package datasets

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatasetCreateOrUpdateDatasetOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *DatasetResource
}

type DatasetCreateOrUpdateDatasetOperationOptions struct {
	IfMatch *string
}

func DefaultDatasetCreateOrUpdateDatasetOperationOptions() DatasetCreateOrUpdateDatasetOperationOptions {
	return DatasetCreateOrUpdateDatasetOperationOptions{}
}

func (o DatasetCreateOrUpdateDatasetOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DatasetCreateOrUpdateDatasetOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DatasetCreateOrUpdateDatasetOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DatasetCreateOrUpdateDataset ...
func (c DatasetsClient) DatasetCreateOrUpdateDataset(ctx context.Context, id DatasetId, input DatasetResource, options DatasetCreateOrUpdateDatasetOperationOptions) (result DatasetCreateOrUpdateDatasetOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          id.Path(),
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

	result.Poller, err = dataplane.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// DatasetCreateOrUpdateDatasetThenPoll performs DatasetCreateOrUpdateDataset then polls until it's completed
func (c DatasetsClient) DatasetCreateOrUpdateDatasetThenPoll(ctx context.Context, id DatasetId, input DatasetResource, options DatasetCreateOrUpdateDatasetOperationOptions) error {
	result, err := c.DatasetCreateOrUpdateDataset(ctx, id, input, options)
	if err != nil {
		return fmt.Errorf("performing DatasetCreateOrUpdateDataset: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after DatasetCreateOrUpdateDataset: %+v", err)
	}

	return nil
}
