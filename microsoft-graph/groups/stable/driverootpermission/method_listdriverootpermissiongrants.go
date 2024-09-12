package driverootpermission

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

type ListDriveRootPermissionGrantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Permission
}

type ListDriveRootPermissionGrantsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Permission
}

type ListDriveRootPermissionGrantsOperationOptions struct {
	Skip *int64
	Top  *int64
}

func DefaultListDriveRootPermissionGrantsOperationOptions() ListDriveRootPermissionGrantsOperationOptions {
	return ListDriveRootPermissionGrantsOperationOptions{}
}

func (o ListDriveRootPermissionGrantsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveRootPermissionGrantsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListDriveRootPermissionGrantsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveRootPermissionGrantsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveRootPermissionGrantsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveRootPermissionGrants - Invoke action grant. Grant users access to a link represented by a permission.
func (c DriveRootPermissionClient) ListDriveRootPermissionGrants(ctx context.Context, id stable.GroupIdDriveIdRootPermissionId, input ListDriveRootPermissionGrantsRequest, options ListDriveRootPermissionGrantsOperationOptions) (result ListDriveRootPermissionGrantsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListDriveRootPermissionGrantsCustomPager{},
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
		Values *[]stable.Permission `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDriveRootPermissionGrantsComplete retrieves all the results into a single object
func (c DriveRootPermissionClient) ListDriveRootPermissionGrantsComplete(ctx context.Context, id stable.GroupIdDriveIdRootPermissionId, input ListDriveRootPermissionGrantsRequest, options ListDriveRootPermissionGrantsOperationOptions) (ListDriveRootPermissionGrantsCompleteResult, error) {
	return c.ListDriveRootPermissionGrantsCompleteMatchingPredicate(ctx, id, input, options, PermissionOperationPredicate{})
}

// ListDriveRootPermissionGrantsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveRootPermissionClient) ListDriveRootPermissionGrantsCompleteMatchingPredicate(ctx context.Context, id stable.GroupIdDriveIdRootPermissionId, input ListDriveRootPermissionGrantsRequest, options ListDriveRootPermissionGrantsOperationOptions, predicate PermissionOperationPredicate) (result ListDriveRootPermissionGrantsCompleteResult, err error) {
	items := make([]stable.Permission, 0)

	resp, err := c.ListDriveRootPermissionGrants(ctx, id, input, options)
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

	result = ListDriveRootPermissionGrantsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
