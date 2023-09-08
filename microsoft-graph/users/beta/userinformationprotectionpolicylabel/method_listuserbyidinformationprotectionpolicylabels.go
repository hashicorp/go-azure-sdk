package userinformationprotectionpolicylabel

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

type ListUserByIdInformationProtectionPolicyLabelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.InformationProtectionLabelCollectionResponse
}

type ListUserByIdInformationProtectionPolicyLabelsCompleteResult struct {
	Items []models.InformationProtectionLabelCollectionResponse
}

// ListUserByIdInformationProtectionPolicyLabels ...
func (c UserInformationProtectionPolicyLabelClient) ListUserByIdInformationProtectionPolicyLabels(ctx context.Context, id UserId) (result ListUserByIdInformationProtectionPolicyLabelsOperationResponse, err error) {
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

// ListUserByIdInformationProtectionPolicyLabelsComplete retrieves all the results into a single object
func (c UserInformationProtectionPolicyLabelClient) ListUserByIdInformationProtectionPolicyLabelsComplete(ctx context.Context, id UserId) (ListUserByIdInformationProtectionPolicyLabelsCompleteResult, error) {
	return c.ListUserByIdInformationProtectionPolicyLabelsCompleteMatchingPredicate(ctx, id, models.InformationProtectionLabelCollectionResponseOperationPredicate{})
}

// ListUserByIdInformationProtectionPolicyLabelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInformationProtectionPolicyLabelClient) ListUserByIdInformationProtectionPolicyLabelsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.InformationProtectionLabelCollectionResponseOperationPredicate) (result ListUserByIdInformationProtectionPolicyLabelsCompleteResult, err error) {
	items := make([]models.InformationProtectionLabelCollectionResponse, 0)

	resp, err := c.ListUserByIdInformationProtectionPolicyLabels(ctx, id)
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

	result = ListUserByIdInformationProtectionPolicyLabelsCompleteResult{
		Items: items,
	}
	return
}
