package deleteditem

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckDeletedItemMemberObjectsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]string
}

type CheckDeletedItemMemberObjectsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []string
}

type CheckDeletedItemMemberObjectsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultCheckDeletedItemMemberObjectsOperationOptions() CheckDeletedItemMemberObjectsOperationOptions {
	return CheckDeletedItemMemberObjectsOperationOptions{}
}

func (o CheckDeletedItemMemberObjectsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CheckDeletedItemMemberObjectsOperationOptions) ToOData() *odata.Query {
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

func (o CheckDeletedItemMemberObjectsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type CheckDeletedItemMemberObjectsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *CheckDeletedItemMemberObjectsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// CheckDeletedItemMemberObjects - Invoke action checkMemberObjects
func (c DeletedItemClient) CheckDeletedItemMemberObjects(ctx context.Context, id beta.DirectoryDeletedItemId, input CheckDeletedItemMemberObjectsRequest, options CheckDeletedItemMemberObjectsOperationOptions) (result CheckDeletedItemMemberObjectsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &CheckDeletedItemMemberObjectsCustomPager{},
		Path:          fmt.Sprintf("%s/checkMemberObjects", id.ID()),
		RetryFunc:     options.RetryFunc,
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

// CheckDeletedItemMemberObjectsComplete retrieves all the results into a single object
func (c DeletedItemClient) CheckDeletedItemMemberObjectsComplete(ctx context.Context, id beta.DirectoryDeletedItemId, input CheckDeletedItemMemberObjectsRequest, options CheckDeletedItemMemberObjectsOperationOptions) (result CheckDeletedItemMemberObjectsCompleteResult, err error) {
	items := make([]string, 0)

	resp, err := c.CheckDeletedItemMemberObjects(ctx, id, input, options)
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

	result = CheckDeletedItemMemberObjectsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
