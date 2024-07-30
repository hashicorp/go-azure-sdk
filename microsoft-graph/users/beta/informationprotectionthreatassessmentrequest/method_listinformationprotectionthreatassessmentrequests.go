package informationprotectionthreatassessmentrequest

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

type ListInformationProtectionThreatAssessmentRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ThreatAssessmentRequest
}

type ListInformationProtectionThreatAssessmentRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ThreatAssessmentRequest
}

type ListInformationProtectionThreatAssessmentRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListInformationProtectionThreatAssessmentRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListInformationProtectionThreatAssessmentRequests ...
func (c InformationProtectionThreatAssessmentRequestClient) ListInformationProtectionThreatAssessmentRequests(ctx context.Context, id UserId) (result ListInformationProtectionThreatAssessmentRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListInformationProtectionThreatAssessmentRequestsCustomPager{},
		Path:       fmt.Sprintf("%s/informationProtection/threatAssessmentRequests", id.ID()),
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
		Values *[]beta.ThreatAssessmentRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListInformationProtectionThreatAssessmentRequestsComplete retrieves all the results into a single object
func (c InformationProtectionThreatAssessmentRequestClient) ListInformationProtectionThreatAssessmentRequestsComplete(ctx context.Context, id UserId) (ListInformationProtectionThreatAssessmentRequestsCompleteResult, error) {
	return c.ListInformationProtectionThreatAssessmentRequestsCompleteMatchingPredicate(ctx, id, ThreatAssessmentRequestOperationPredicate{})
}

// ListInformationProtectionThreatAssessmentRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c InformationProtectionThreatAssessmentRequestClient) ListInformationProtectionThreatAssessmentRequestsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate ThreatAssessmentRequestOperationPredicate) (result ListInformationProtectionThreatAssessmentRequestsCompleteResult, err error) {
	items := make([]beta.ThreatAssessmentRequest, 0)

	resp, err := c.ListInformationProtectionThreatAssessmentRequests(ctx, id)
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

	result = ListInformationProtectionThreatAssessmentRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
