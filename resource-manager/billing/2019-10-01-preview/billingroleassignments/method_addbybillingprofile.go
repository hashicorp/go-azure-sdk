package billingroleassignments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddByBillingProfileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BillingRoleAssignment
}

type AddByBillingProfileCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BillingRoleAssignment
}

type AddByBillingProfileCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *AddByBillingProfileCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AddByBillingProfile ...
func (c BillingRoleAssignmentsClient) AddByBillingProfile(ctx context.Context, id BillingProfileId, input BillingRoleAssignmentPayload) (result AddByBillingProfileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod: http.MethodPost,
		Pager:      &AddByBillingProfileCustomPager{},
		Path:       fmt.Sprintf("%s/createBillingRoleAssignment", id.ID()),
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
		Values *[]BillingRoleAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AddByBillingProfileComplete retrieves all the results into a single object
func (c BillingRoleAssignmentsClient) AddByBillingProfileComplete(ctx context.Context, id BillingProfileId, input BillingRoleAssignmentPayload) (AddByBillingProfileCompleteResult, error) {
	return c.AddByBillingProfileCompleteMatchingPredicate(ctx, id, input, BillingRoleAssignmentOperationPredicate{})
}

// AddByBillingProfileCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BillingRoleAssignmentsClient) AddByBillingProfileCompleteMatchingPredicate(ctx context.Context, id BillingProfileId, input BillingRoleAssignmentPayload, predicate BillingRoleAssignmentOperationPredicate) (result AddByBillingProfileCompleteResult, err error) {
	items := make([]BillingRoleAssignment, 0)

	resp, err := c.AddByBillingProfile(ctx, id, input)
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

	result = AddByBillingProfileCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
