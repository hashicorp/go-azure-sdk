package groupsiteinformationprotectionthreatassessmentrequestresult

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequestByIdResultsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ThreatAssessmentResultCollectionResponse
}

type ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequestByIdResultsCompleteResult struct {
	Items []models.ThreatAssessmentResultCollectionResponse
}

// ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequestByIdResults ...
func (c GroupSiteInformationProtectionThreatAssessmentRequestResultClient) ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequestByIdResults(ctx context.Context, id GroupSiteInformationProtectionThreatAssessmentRequestId) (result ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequestByIdResultsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
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
		Values *[]models.ThreatAssessmentResultCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequestByIdResultsComplete retrieves all the results into a single object
func (c GroupSiteInformationProtectionThreatAssessmentRequestResultClient) ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequestByIdResultsComplete(ctx context.Context, id GroupSiteInformationProtectionThreatAssessmentRequestId) (ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequestByIdResultsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequestByIdResultsCompleteMatchingPredicate(ctx, id, models.ThreatAssessmentResultCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequestByIdResultsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteInformationProtectionThreatAssessmentRequestResultClient) ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequestByIdResultsCompleteMatchingPredicate(ctx context.Context, id GroupSiteInformationProtectionThreatAssessmentRequestId, predicate models.ThreatAssessmentResultCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequestByIdResultsCompleteResult, err error) {
	items := make([]models.ThreatAssessmentResultCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequestByIdResults(ctx, id)
	if err != nil {
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

	result = ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequestByIdResultsCompleteResult{
		Items: items,
	}
	return
}
