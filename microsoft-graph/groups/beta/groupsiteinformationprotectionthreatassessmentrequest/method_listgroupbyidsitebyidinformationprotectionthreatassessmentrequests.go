package groupsiteinformationprotectionthreatassessmentrequest

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

type ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ThreatAssessmentRequestCollectionResponse
}

type ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequestsCompleteResult struct {
	Items []models.ThreatAssessmentRequestCollectionResponse
}

// ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequests ...
func (c GroupSiteInformationProtectionThreatAssessmentRequestClient) ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequests(ctx context.Context, id GroupSiteId) (result ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
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
		Values *[]models.ThreatAssessmentRequestCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequestsComplete retrieves all the results into a single object
func (c GroupSiteInformationProtectionThreatAssessmentRequestClient) ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequestsComplete(ctx context.Context, id GroupSiteId) (ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequestsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequestsCompleteMatchingPredicate(ctx, id, models.ThreatAssessmentRequestCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteInformationProtectionThreatAssessmentRequestClient) ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequestsCompleteMatchingPredicate(ctx context.Context, id GroupSiteId, predicate models.ThreatAssessmentRequestCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequestsCompleteResult, err error) {
	items := make([]models.ThreatAssessmentRequestCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequests(ctx, id)
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

	result = ListGroupByIdSiteByIdInformationProtectionThreatAssessmentRequestsCompleteResult{
		Items: items,
	}
	return
}
