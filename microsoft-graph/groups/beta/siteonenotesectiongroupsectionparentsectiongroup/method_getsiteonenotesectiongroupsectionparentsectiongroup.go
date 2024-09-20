package siteonenotesectiongroupsectionparentsectiongroup

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

type GetSiteOnenoteSectionGroupSectionParentSectionGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.SectionGroup
}

type GetSiteOnenoteSectionGroupSectionParentSectionGroupOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetSiteOnenoteSectionGroupSectionParentSectionGroupOperationOptions() GetSiteOnenoteSectionGroupSectionParentSectionGroupOperationOptions {
	return GetSiteOnenoteSectionGroupSectionParentSectionGroupOperationOptions{}
}

func (o GetSiteOnenoteSectionGroupSectionParentSectionGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetSiteOnenoteSectionGroupSectionParentSectionGroupOperationOptions) ToOData() *odata.Query {
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

func (o GetSiteOnenoteSectionGroupSectionParentSectionGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetSiteOnenoteSectionGroupSectionParentSectionGroup - Get parentSectionGroup from groups. The section group that
// contains the section. Read-only.
func (c SiteOnenoteSectionGroupSectionParentSectionGroupClient) GetSiteOnenoteSectionGroupSectionParentSectionGroup(ctx context.Context, id beta.GroupIdSiteIdOnenoteSectionGroupIdSectionId, options GetSiteOnenoteSectionGroupSectionParentSectionGroupOperationOptions) (result GetSiteOnenoteSectionGroupSectionParentSectionGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/parentSectionGroup", id.ID()),
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

	var model beta.SectionGroup
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
