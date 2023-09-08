package groupsiteinformationprotectionpolicylabel

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

type ListGroupByIdSiteByIdInformationProtectionPolicyLabelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.InformationProtectionLabelCollectionResponse
}

type ListGroupByIdSiteByIdInformationProtectionPolicyLabelsCompleteResult struct {
	Items []models.InformationProtectionLabelCollectionResponse
}

// ListGroupByIdSiteByIdInformationProtectionPolicyLabels ...
func (c GroupSiteInformationProtectionPolicyLabelClient) ListGroupByIdSiteByIdInformationProtectionPolicyLabels(ctx context.Context, id GroupSiteId) (result ListGroupByIdSiteByIdInformationProtectionPolicyLabelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/informationProtection/policy/labels", id.ID()),
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
		Values *[]models.InformationProtectionLabelCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdSiteByIdInformationProtectionPolicyLabelsComplete retrieves all the results into a single object
func (c GroupSiteInformationProtectionPolicyLabelClient) ListGroupByIdSiteByIdInformationProtectionPolicyLabelsComplete(ctx context.Context, id GroupSiteId) (ListGroupByIdSiteByIdInformationProtectionPolicyLabelsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdInformationProtectionPolicyLabelsCompleteMatchingPredicate(ctx, id, models.InformationProtectionLabelCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdInformationProtectionPolicyLabelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteInformationProtectionPolicyLabelClient) ListGroupByIdSiteByIdInformationProtectionPolicyLabelsCompleteMatchingPredicate(ctx context.Context, id GroupSiteId, predicate models.InformationProtectionLabelCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdInformationProtectionPolicyLabelsCompleteResult, err error) {
	items := make([]models.InformationProtectionLabelCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdInformationProtectionPolicyLabels(ctx, id)
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

	result = ListGroupByIdSiteByIdInformationProtectionPolicyLabelsCompleteResult{
		Items: items,
	}
	return
}
