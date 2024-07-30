package siteinformationprotectionthreatassessmentrequestresult

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

type ListSiteInformationProtectionThreatAssessmentRequestResultsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ThreatAssessmentResult
}

type ListSiteInformationProtectionThreatAssessmentRequestResultsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ThreatAssessmentResult
}

type ListSiteInformationProtectionThreatAssessmentRequestResultsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteInformationProtectionThreatAssessmentRequestResultsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteInformationProtectionThreatAssessmentRequestResults ...
func (c SiteInformationProtectionThreatAssessmentRequestResultClient) ListSiteInformationProtectionThreatAssessmentRequestResults(ctx context.Context, id GroupIdSiteIdInformationProtectionThreatAssessmentRequestId) (result ListSiteInformationProtectionThreatAssessmentRequestResultsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSiteInformationProtectionThreatAssessmentRequestResultsCustomPager{},
		Path:       fmt.Sprintf("%s/results", id.ID()),
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
		Values *[]beta.ThreatAssessmentResult `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteInformationProtectionThreatAssessmentRequestResultsComplete retrieves all the results into a single object
func (c SiteInformationProtectionThreatAssessmentRequestResultClient) ListSiteInformationProtectionThreatAssessmentRequestResultsComplete(ctx context.Context, id GroupIdSiteIdInformationProtectionThreatAssessmentRequestId) (ListSiteInformationProtectionThreatAssessmentRequestResultsCompleteResult, error) {
	return c.ListSiteInformationProtectionThreatAssessmentRequestResultsCompleteMatchingPredicate(ctx, id, ThreatAssessmentResultOperationPredicate{})
}

// ListSiteInformationProtectionThreatAssessmentRequestResultsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteInformationProtectionThreatAssessmentRequestResultClient) ListSiteInformationProtectionThreatAssessmentRequestResultsCompleteMatchingPredicate(ctx context.Context, id GroupIdSiteIdInformationProtectionThreatAssessmentRequestId, predicate ThreatAssessmentResultOperationPredicate) (result ListSiteInformationProtectionThreatAssessmentRequestResultsCompleteResult, err error) {
	items := make([]beta.ThreatAssessmentResult, 0)

	resp, err := c.ListSiteInformationProtectionThreatAssessmentRequestResults(ctx, id)
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

	result = ListSiteInformationProtectionThreatAssessmentRequestResultsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
