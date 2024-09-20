package siteonenotesectiongroupsection

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateSiteOnenoteSectionGroupSectionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateSiteOnenoteSectionGroupSectionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateSiteOnenoteSectionGroupSectionOperationOptions() UpdateSiteOnenoteSectionGroupSectionOperationOptions {
	return UpdateSiteOnenoteSectionGroupSectionOperationOptions{}
}

func (o UpdateSiteOnenoteSectionGroupSectionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateSiteOnenoteSectionGroupSectionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateSiteOnenoteSectionGroupSectionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateSiteOnenoteSectionGroupSection - Update the navigation property sections in groups
func (c SiteOnenoteSectionGroupSectionClient) UpdateSiteOnenoteSectionGroupSection(ctx context.Context, id beta.GroupIdSiteIdOnenoteSectionGroupIdSectionId, input beta.OnenoteSection, options UpdateSiteOnenoteSectionGroupSectionOperationOptions) (result UpdateSiteOnenoteSectionGroupSectionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}
