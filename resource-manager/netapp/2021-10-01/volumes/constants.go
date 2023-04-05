package volumes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AvsDataStore string

const (
	AvsDataStoreDisabled AvsDataStore = "Disabled"
	AvsDataStoreEnabled  AvsDataStore = "Enabled"
)

func PossibleValuesForAvsDataStore() []string {
	return []string{
		string(AvsDataStoreDisabled),
		string(AvsDataStoreEnabled),
	}
}

type ChownMode string

const (
	ChownModeRestricted   ChownMode = "Restricted"
	ChownModeUnrestricted ChownMode = "Unrestricted"
)

func PossibleValuesForChownMode() []string {
	return []string{
		string(ChownModeRestricted),
		string(ChownModeUnrestricted),
	}
}

type EnableSubvolumes string

const (
	EnableSubvolumesDisabled EnableSubvolumes = "Disabled"
	EnableSubvolumesEnabled  EnableSubvolumes = "Enabled"
)

func PossibleValuesForEnableSubvolumes() []string {
	return []string{
		string(EnableSubvolumesDisabled),
		string(EnableSubvolumesEnabled),
	}
}

type EndpointType string

const (
	EndpointTypeDst EndpointType = "dst"
	EndpointTypeSrc EndpointType = "src"
)

func PossibleValuesForEndpointType() []string {
	return []string{
		string(EndpointTypeDst),
		string(EndpointTypeSrc),
	}
}

type NetworkFeatures string

const (
	NetworkFeaturesBasic    NetworkFeatures = "Basic"
	NetworkFeaturesStandard NetworkFeatures = "Standard"
)

func PossibleValuesForNetworkFeatures() []string {
	return []string{
		string(NetworkFeaturesBasic),
		string(NetworkFeaturesStandard),
	}
}

type ReplicationSchedule string

const (
	ReplicationScheduleDaily           ReplicationSchedule = "daily"
	ReplicationScheduleHourly          ReplicationSchedule = "hourly"
	ReplicationScheduleOneZerominutely ReplicationSchedule = "_10minutely"
)

func PossibleValuesForReplicationSchedule() []string {
	return []string{
		string(ReplicationScheduleDaily),
		string(ReplicationScheduleHourly),
		string(ReplicationScheduleOneZerominutely),
	}
}

type SecurityStyle string

const (
	SecurityStyleNtfs SecurityStyle = "ntfs"
	SecurityStyleUnix SecurityStyle = "unix"
)

func PossibleValuesForSecurityStyle() []string {
	return []string{
		string(SecurityStyleNtfs),
		string(SecurityStyleUnix),
	}
}

type ServiceLevel string

const (
	ServiceLevelPremium     ServiceLevel = "Premium"
	ServiceLevelStandard    ServiceLevel = "Standard"
	ServiceLevelStandardZRS ServiceLevel = "StandardZRS"
	ServiceLevelUltra       ServiceLevel = "Ultra"
)

func PossibleValuesForServiceLevel() []string {
	return []string{
		string(ServiceLevelPremium),
		string(ServiceLevelStandard),
		string(ServiceLevelStandardZRS),
		string(ServiceLevelUltra),
	}
}

type VolumeStorageToNetworkProximity string

const (
	VolumeStorageToNetworkProximityDefault VolumeStorageToNetworkProximity = "Default"
	VolumeStorageToNetworkProximityTOne    VolumeStorageToNetworkProximity = "T1"
	VolumeStorageToNetworkProximityTTwo    VolumeStorageToNetworkProximity = "T2"
)

func PossibleValuesForVolumeStorageToNetworkProximity() []string {
	return []string{
		string(VolumeStorageToNetworkProximityDefault),
		string(VolumeStorageToNetworkProximityTOne),
		string(VolumeStorageToNetworkProximityTTwo),
	}
}
