package drivelistitempermission

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

type ListDriveListItemPermissionGrantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Permission
}

type ListDriveListItemPermissionGrantsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Permission
}

type ListDriveListItemPermissionGrantsOperationOptions struct {
	Skip *int64
	Top  *int64
}

func DefaultListDriveListItemPermissionGrantsOperationOptions() ListDriveListItemPermissionGrantsOperationOptions {
	return ListDriveListItemPermissionGrantsOperationOptions{}
}

func (o ListDriveListItemPermissionGrantsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveListItemPermissionGrantsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListDriveListItemPermissionGrantsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveListItemPermissionGrantsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveListItemPermissionGrantsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveListItemPermissionGrants - Invoke action grant. Grant users access to a link represented by a permission.
func (c DriveListItemPermissionClient) ListDriveListItemPermissionGrants(ctx context.Context, id beta.UserIdDriveIdListItemIdPermissionId, input ListDriveListItemPermissionGrantsRequest, options ListDriveListItemPermissionGrantsOperationOptions) (result ListDriveListItemPermissionGrantsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListDriveListItemPermissionGrantsCustomPager{},
		Path:          fmt.Sprintf("%s/grant", id.ID()),
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
		Values *[]beta.Permission `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDriveListItemPermissionGrantsComplete retrieves all the results into a single object
func (c DriveListItemPermissionClient) ListDriveListItemPermissionGrantsComplete(ctx context.Context, id beta.UserIdDriveIdListItemIdPermissionId, input ListDriveListItemPermissionGrantsRequest, options ListDriveListItemPermissionGrantsOperationOptions) (ListDriveListItemPermissionGrantsCompleteResult, error) {
	return c.ListDriveListItemPermissionGrantsCompleteMatchingPredicate(ctx, id, input, options, PermissionOperationPredicate{})
}

// ListDriveListItemPermissionGrantsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveListItemPermissionClient) ListDriveListItemPermissionGrantsCompleteMatchingPredicate(ctx context.Context, id beta.UserIdDriveIdListItemIdPermissionId, input ListDriveListItemPermissionGrantsRequest, options ListDriveListItemPermissionGrantsOperationOptions, predicate PermissionOperationPredicate) (result ListDriveListItemPermissionGrantsCompleteResult, err error) {
	items := make([]beta.Permission, 0)

	resp, err := c.ListDriveListItemPermissionGrants(ctx, id, input, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListDriveListItemPermissionGrantsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
