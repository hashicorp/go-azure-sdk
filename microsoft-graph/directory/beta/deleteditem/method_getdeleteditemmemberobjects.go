package deleteditem

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

type GetDeletedItemMemberObjectsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]string
}

type GetDeletedItemMemberObjectsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []string
}

type GetDeletedItemMemberObjectsOperationOptions struct {
	Metadata *odata.Metadata
	Skip     *int64
	Top      *int64
}

func DefaultGetDeletedItemMemberObjectsOperationOptions() GetDeletedItemMemberObjectsOperationOptions {
	return GetDeletedItemMemberObjectsOperationOptions{}
}

func (o GetDeletedItemMemberObjectsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDeletedItemMemberObjectsOperationOptions) ToOData() *odata.Query {
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

func (o GetDeletedItemMemberObjectsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type GetDeletedItemMemberObjectsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetDeletedItemMemberObjectsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetDeletedItemMemberObjects - Invoke action getMemberObjects
func (c DeletedItemClient) GetDeletedItemMemberObjects(ctx context.Context, id beta.DirectoryDeletedItemId, input GetDeletedItemMemberObjectsRequest, options GetDeletedItemMemberObjectsOperationOptions) (result GetDeletedItemMemberObjectsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &GetDeletedItemMemberObjectsCustomPager{},
		Path:          fmt.Sprintf("%s/getMemberObjects", id.ID()),
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

// GetDeletedItemMemberObjectsComplete retrieves all the results into a single object
func (c DeletedItemClient) GetDeletedItemMemberObjectsComplete(ctx context.Context, id beta.DirectoryDeletedItemId, input GetDeletedItemMemberObjectsRequest, options GetDeletedItemMemberObjectsOperationOptions) (result GetDeletedItemMemberObjectsCompleteResult, err error) {
	items := make([]string, 0)

	resp, err := c.GetDeletedItemMemberObjects(ctx, id, input, options)
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

	result = GetDeletedItemMemberObjectsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
