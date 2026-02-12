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

type EnrollmentGroupsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]EnrollmentGroup
}

type EnrollmentGroupsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []EnrollmentGroup
}

type EnrollmentGroupsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *EnrollmentGroupsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// EnrollmentGroupsList ...
func (c IotcentralsClient) EnrollmentGroupsList(ctx context.Context) (result EnrollmentGroupsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &EnrollmentGroupsListCustomPager{},
		Path:       "/enrollmentGroups",
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
		Values *[]EnrollmentGroup `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// EnrollmentGroupsListComplete retrieves all the results into a single object
func (c IotcentralsClient) EnrollmentGroupsListComplete(ctx context.Context) (EnrollmentGroupsListCompleteResult, error) {
	return c.EnrollmentGroupsListCompleteMatchingPredicate(ctx, EnrollmentGroupOperationPredicate{})
}

// EnrollmentGroupsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IotcentralsClient) EnrollmentGroupsListCompleteMatchingPredicate(ctx context.Context, predicate EnrollmentGroupOperationPredicate) (result EnrollmentGroupsListCompleteResult, err error) {
	items := make([]EnrollmentGroup, 0)

	resp, err := c.EnrollmentGroupsList(ctx)
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

	result = EnrollmentGroupsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
