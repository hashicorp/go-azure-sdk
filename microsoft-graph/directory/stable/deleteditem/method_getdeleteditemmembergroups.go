package deleteditem

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

type GetDeletedItemMemberGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]string
}

type GetDeletedItemMemberGroupsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []string
}

type GetDeletedItemMemberGroupsOperationOptions struct {
	Metadata *odata.Metadata
	Skip     *int64
	Top      *int64
}

func DefaultGetDeletedItemMemberGroupsOperationOptions() GetDeletedItemMemberGroupsOperationOptions {
	return GetDeletedItemMemberGroupsOperationOptions{}
}

func (o GetDeletedItemMemberGroupsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDeletedItemMemberGroupsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o GetDeletedItemMemberGroupsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type GetDeletedItemMemberGroupsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetDeletedItemMemberGroupsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetDeletedItemMemberGroups - Invoke action getMemberGroups. Return all the group IDs for the groups that the
// specified user, group, service principal, organizational contact, device, or directory object is a member of. This
// function is transitive. This API returns up to 11,000 group IDs. If more than 11,000 results are available, it
// returns a 400 Bad Request error with the DirectoryResultSizeLimitExceeded error code. If you get the
// DirectoryResultSizeLimitExceeded error code, use the List group transitive memberOf API instead.
func (c DeletedItemClient) GetDeletedItemMemberGroups(ctx context.Context, id stable.DirectoryDeletedItemId, input GetDeletedItemMemberGroupsRequest, options GetDeletedItemMemberGroupsOperationOptions) (result GetDeletedItemMemberGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &GetDeletedItemMemberGroupsCustomPager{},
		Path:          fmt.Sprintf("%s/getMemberGroups", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]string `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetDeletedItemMemberGroupsComplete retrieves all the results into a single object
func (c DeletedItemClient) GetDeletedItemMemberGroupsComplete(ctx context.Context, id stable.DirectoryDeletedItemId, input GetDeletedItemMemberGroupsRequest, options GetDeletedItemMemberGroupsOperationOptions) (result GetDeletedItemMemberGroupsCompleteResult, err error) {
	items := make([]string, 0)

	resp, err := c.GetDeletedItemMemberGroups(ctx, id, input, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			items = append(items, v)
		}
	}

	result = GetDeletedItemMemberGroupsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
