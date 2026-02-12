package iotcentrals

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RolesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Role
}

type RolesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Role
}

type RolesListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *RolesListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// RolesList ...
func (c IotcentralsClient) RolesList(ctx context.Context) (result RolesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &RolesListCustomPager{},
		Path:       "/roles",
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
		Values *[]Role `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// RolesListComplete retrieves all the results into a single object
func (c IotcentralsClient) RolesListComplete(ctx context.Context) (RolesListCompleteResult, error) {
	return c.RolesListCompleteMatchingPredicate(ctx, RoleOperationPredicate{})
}

// RolesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IotcentralsClient) RolesListCompleteMatchingPredicate(ctx context.Context, predicate RoleOperationPredicate) (result RolesListCompleteResult, err error) {
	items := make([]Role, 0)

	resp, err := c.RolesList(ctx)
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

	result = RolesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
