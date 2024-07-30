package termsandconditionassignment

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

type ListTermsAndConditionAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.TermsAndConditionsAssignment
}

type ListTermsAndConditionAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.TermsAndConditionsAssignment
}

type ListTermsAndConditionAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTermsAndConditionAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTermsAndConditionAssignments ...
func (c TermsAndConditionAssignmentClient) ListTermsAndConditionAssignments(ctx context.Context, id DeviceManagementTermsAndConditionId) (result ListTermsAndConditionAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTermsAndConditionAssignmentsCustomPager{},
		Path:       fmt.Sprintf("%s/assignments", id.ID()),
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
		Values *[]beta.TermsAndConditionsAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTermsAndConditionAssignmentsComplete retrieves all the results into a single object
func (c TermsAndConditionAssignmentClient) ListTermsAndConditionAssignmentsComplete(ctx context.Context, id DeviceManagementTermsAndConditionId) (ListTermsAndConditionAssignmentsCompleteResult, error) {
	return c.ListTermsAndConditionAssignmentsCompleteMatchingPredicate(ctx, id, TermsAndConditionsAssignmentOperationPredicate{})
}

// ListTermsAndConditionAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TermsAndConditionAssignmentClient) ListTermsAndConditionAssignmentsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementTermsAndConditionId, predicate TermsAndConditionsAssignmentOperationPredicate) (result ListTermsAndConditionAssignmentsCompleteResult, err error) {
	items := make([]beta.TermsAndConditionsAssignment, 0)

	resp, err := c.ListTermsAndConditionAssignments(ctx, id)
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

	result = ListTermsAndConditionAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
