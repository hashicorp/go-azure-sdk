package virtualendpointreport

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetVirtualEndpointReportsCloudPCRecommendationReportsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetVirtualEndpointReportsCloudPCRecommendationReportsOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultGetVirtualEndpointReportsCloudPCRecommendationReportsOperationOptions() GetVirtualEndpointReportsCloudPCRecommendationReportsOperationOptions {
	return GetVirtualEndpointReportsCloudPCRecommendationReportsOperationOptions{}
}

func (o GetVirtualEndpointReportsCloudPCRecommendationReportsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetVirtualEndpointReportsCloudPCRecommendationReportsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetVirtualEndpointReportsCloudPCRecommendationReportsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetVirtualEndpointReportsCloudPCRecommendationReports - Invoke action getCloudPcRecommendationReports. Get the device
// recommendation reports for Cloud PCs, such as the usage category report. The usage category report categorizes a
// Cloud PC as Undersized, Oversized, Rightsized, or Underutilized, and also provides the recommended SKU when the Cloud
// PC isn't Rightsized.
func (c VirtualEndpointReportClient) GetVirtualEndpointReportsCloudPCRecommendationReports(ctx context.Context, input GetVirtualEndpointReportsCloudPCRecommendationReportsRequest, options GetVirtualEndpointReportsCloudPCRecommendationReportsOperationOptions) (result GetVirtualEndpointReportsCloudPCRecommendationReportsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/virtualEndpoint/reports/getCloudPcRecommendationReports",
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
