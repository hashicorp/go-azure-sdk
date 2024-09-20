package addons

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByRoleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Addon
}

type ListByRoleCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Addon
}

type ListByRoleCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByRoleCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByRole ...
func (c AddonsClient) ListByRole(ctx context.Context, id RoleId) (result ListByRoleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByRoleCustomPager{},
		Path:       fmt.Sprintf("%s/addons", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]Addon, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := UnmarshalAddonImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for Addon (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListByRoleComplete retrieves all the results into a single object
func (c AddonsClient) ListByRoleComplete(ctx context.Context, id RoleId) (ListByRoleCompleteResult, error) {
	return c.ListByRoleCompleteMatchingPredicate(ctx, id, AddonOperationPredicate{})
}

// ListByRoleCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AddonsClient) ListByRoleCompleteMatchingPredicate(ctx context.Context, id RoleId, predicate AddonOperationPredicate) (result ListByRoleCompleteResult, err error) {
	items := make([]Addon, 0)

	resp, err := c.ListByRole(ctx, id)
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

	result = ListByRoleCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
