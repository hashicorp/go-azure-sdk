package virtualmachinescalesets

type UpgradeOperationHistoricalStatusInfoOperationPredicate struct {
	Location *string
	Type     *string
}

func (p UpgradeOperationHistoricalStatusInfoOperationPredicate) Matches(input UpgradeOperationHistoricalStatusInfo) bool {

	if p.Location != nil && (input.Location == nil && *p.Location != *input.Location) {
		return false
	}

	if p.Type != nil && (input.Type == nil && *p.Type != *input.Type) {
		return false
	}

	return true
}

type VirtualMachineScaleSetOperationPredicate struct {
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p VirtualMachineScaleSetOperationPredicate) Matches(input VirtualMachineScaleSet) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && *p.Location != input.Location {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil && *p.Type != *input.Type) {
		return false
	}

	return true
}

type VirtualMachineScaleSetSkuOperationPredicate struct {
	ResourceType *string
}

func (p VirtualMachineScaleSetSkuOperationPredicate) Matches(input VirtualMachineScaleSetSku) bool {

	if p.ResourceType != nil && (input.ResourceType == nil && *p.ResourceType != *input.ResourceType) {
		return false
	}

	return true
}
