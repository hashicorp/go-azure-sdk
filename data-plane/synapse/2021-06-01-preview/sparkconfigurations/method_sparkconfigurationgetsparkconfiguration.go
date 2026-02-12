package sparkconfigurations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SparkConfigurationGetSparkConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *SparkConfigurationResource
}

type SparkConfigurationGetSparkConfigurationOperationOptions struct {
	IfNoneMatch *string
}

func DefaultSparkConfigurationGetSparkConfigurationOperationOptions() SparkConfigurationGetSparkConfigurationOperationOptions {
	return SparkConfigurationGetSparkConfigurationOperationOptions{}
}

func (o SparkConfigurationGetSparkConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfNoneMatch != nil {
		out.Append("If-None-Match", fmt.Sprintf("%v", *o.IfNoneMatch))
	}
	return &out
}

func (o SparkConfigurationGetSparkConfigurationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o SparkConfigurationGetSparkConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SparkConfigurationGetSparkConfiguration ...
func (c SparkConfigurationsClient) SparkConfigurationGetSparkConfiguration(ctx context.Context, id SparkConfigurationId, options SparkConfigurationGetSparkConfigurationOperationOptions) (result SparkConfigurationGetSparkConfigurationOperationResponse, err error) {
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

	var model SparkConfigurationResource
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
