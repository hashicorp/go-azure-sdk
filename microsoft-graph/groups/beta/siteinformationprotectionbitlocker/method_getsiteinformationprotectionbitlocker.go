package siteinformationprotectionbitlocker

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

type GetSiteInformationProtectionBitlockerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.Bitlocker
}

type GetSiteInformationProtectionBitlockerOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetSiteInformationProtectionBitlockerOperationOptions() GetSiteInformationProtectionBitlockerOperationOptions {
	return GetSiteInformationProtectionBitlockerOperationOptions{}
}

func (o GetSiteInformationProtectionBitlockerOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetSiteInformationProtectionBitlockerOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetSiteInformationProtectionBitlockerOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetSiteInformationProtectionBitlocker - Get bitlocker from groups
func (c SiteInformationProtectionBitlockerClient) GetSiteInformationProtectionBitlocker(ctx context.Context, id beta.GroupIdSiteId, options GetSiteInformationProtectionBitlockerOperationOptions) (result GetSiteInformationProtectionBitlockerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/informationProtection/bitlocker", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	var model beta.Bitlocker
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
