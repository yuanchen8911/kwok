/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package k8s

// Don't edit this file directly.  It is generated by feature_gates_data.sh.
var rawData = []FeatureSpec{

	// APIListChunking
	{"APIListChunking", Alpha, 8, 8},
	{"APIListChunking", Beta, 9, 28},
	{"APIListChunking", GA, 29, -1},

	// APIPriorityAndFairness
	{"APIPriorityAndFairness", Alpha, 17, 19},
	{"APIPriorityAndFairness", Beta, 20, 28},
	{"APIPriorityAndFairness", GA, 29, -1},

	// APIResponseCompression
	{"APIResponseCompression", Alpha, 8, 15},
	{"APIResponseCompression", Beta, 16, -1},

	// APISelfSubjectReview
	{"APISelfSubjectReview", Alpha, 26, 26},
	{"APISelfSubjectReview", Beta, 27, 27},
	{"APISelfSubjectReview", GA, 28, -1},

	// APIServerIdentity
	{"APIServerIdentity", Alpha, 20, 25},
	{"APIServerIdentity", Beta, 26, -1},

	// APIServerTracing
	{"APIServerTracing", Alpha, 22, 26},
	{"APIServerTracing", Beta, 27, -1},

	// Accelerators
	{"Accelerators", Alpha, 6, 10},

	// AdmissionWebhookMatchConditions
	{"AdmissionWebhookMatchConditions", Alpha, 27, 27},
	{"AdmissionWebhookMatchConditions", Beta, 28, -1},

	// AdvancedAuditing
	{"AdvancedAuditing", Alpha, 7, 7},
	{"AdvancedAuditing", Beta, 8, 11},
	{"AdvancedAuditing", GA, 12, 27},

	// AffinityInAnnotations
	{"AffinityInAnnotations", Alpha, 6, 7},

	// AggregatedDiscoveryEndpoint
	{"AggregatedDiscoveryEndpoint", Alpha, 26, 26},
	{"AggregatedDiscoveryEndpoint", Beta, 27, -1},

	// AllowInsecureBackendProxy
	{"AllowInsecureBackendProxy", Beta, 17, 20},
	{"AllowInsecureBackendProxy", GA, 21, 22},

	// AllowServiceLBStatusOnNonLB
	{"AllowServiceLBStatusOnNonLB", Deprecated, 29, -1},

	// AnyVolumeDataSource
	{"AnyVolumeDataSource", Alpha, 18, 23},
	{"AnyVolumeDataSource", Beta, 24, -1},

	// AppArmor
	{"AppArmor", Beta, 6, -1},

	// AttachVolumeLimit
	{"AttachVolumeLimit", Alpha, 11, 11},
	{"AttachVolumeLimit", Beta, 12, 16},
	{"AttachVolumeLimit", GA, 17, 20},

	// BalanceAttachedNodeVolumes
	{"BalanceAttachedNodeVolumes", Alpha, 11, 21},

	// BlockVolume
	{"BlockVolume", Alpha, 9, 12},
	{"BlockVolume", Beta, 13, 17},
	{"BlockVolume", GA, 18, 20},

	// BoundServiceAccountTokenVolume
	{"BoundServiceAccountTokenVolume", Alpha, 13, 20},
	{"BoundServiceAccountTokenVolume", Beta, 21, 21},
	{"BoundServiceAccountTokenVolume", GA, 22, 22},

	// CPUCFSQuotaPeriod
	{"CPUCFSQuotaPeriod", Alpha, 12, -1},

	// CPUManager
	{"CPUManager", Alpha, 8, 9},
	{"CPUManager", Beta, 10, 25},
	{"CPUManager", GA, 26, -1},

	// CPUManagerPolicyAlphaOptions
	{"CPUManagerPolicyAlphaOptions", Alpha, 23, -1},

	// CPUManagerPolicyBetaOptions
	{"CPUManagerPolicyBetaOptions", Beta, 23, -1},

	// CPUManagerPolicyOptions
	{"CPUManagerPolicyOptions", Alpha, 22, 22},
	{"CPUManagerPolicyOptions", Beta, 23, -1},

	// CRDValidationRatcheting
	{"CRDValidationRatcheting", Alpha, 28, -1},

	// CRIContainerLogRotation
	{"CRIContainerLogRotation", Alpha, 10, 10},
	{"CRIContainerLogRotation", Beta, 11, 20},
	{"CRIContainerLogRotation", GA, 21, 21},

	// CSIBlockVolume
	{"CSIBlockVolume", Alpha, 11, 13},
	{"CSIBlockVolume", Beta, 14, 17},
	{"CSIBlockVolume", GA, 18, 20},

	// CSIDriverRegistry
	{"CSIDriverRegistry", Alpha, 12, 13},
	{"CSIDriverRegistry", Beta, 14, 17},
	{"CSIDriverRegistry", GA, 18, 20},

	// CSIInlineVolume
	{"CSIInlineVolume", Alpha, 14, 15},
	{"CSIInlineVolume", Beta, 16, 24},
	{"CSIInlineVolume", GA, 25, 26},

	// CSIMigration
	{"CSIMigration", Alpha, 14, 16},
	{"CSIMigration", Beta, 17, 24},
	{"CSIMigration", GA, 25, 26},

	// CSIMigrationAWS
	{"CSIMigrationAWS", Alpha, 14, 16},
	{"CSIMigrationAWS", Beta, 17, 24},
	{"CSIMigrationAWS", GA, 25, 26},

	// CSIMigrationAWSComplete
	{"CSIMigrationAWSComplete", Alpha, 17, 20},

	// CSIMigrationAzureDisk
	{"CSIMigrationAzureDisk", Alpha, 15, 18},
	{"CSIMigrationAzureDisk", Beta, 19, 23},
	{"CSIMigrationAzureDisk", GA, 24, 26},

	// CSIMigrationAzureDiskComplete
	{"CSIMigrationAzureDiskComplete", Alpha, 17, 20},

	// CSIMigrationAzureFile
	{"CSIMigrationAzureFile", Alpha, 15, 20},
	{"CSIMigrationAzureFile", Beta, 21, 25},
	{"CSIMigrationAzureFile", GA, 26, -1},

	// CSIMigrationAzureFileComplete
	{"CSIMigrationAzureFileComplete", Alpha, 17, 20},

	// CSIMigrationGCE
	{"CSIMigrationGCE", Alpha, 14, 16},
	{"CSIMigrationGCE", Beta, 17, 24},
	{"CSIMigrationGCE", GA, 25, 27},

	// CSIMigrationGCEComplete
	{"CSIMigrationGCEComplete", Alpha, 17, 20},

	// CSIMigrationOpenStack
	{"CSIMigrationOpenStack", Alpha, 14, 17},
	{"CSIMigrationOpenStack", Beta, 18, 23},
	{"CSIMigrationOpenStack", GA, 24, 25},

	// CSIMigrationOpenStackComplete
	{"CSIMigrationOpenStackComplete", Alpha, 17, 20},

	// CSIMigrationPortworx
	{"CSIMigrationPortworx", Alpha, 23, 24},
	{"CSIMigrationPortworx", Beta, 25, -1},

	// CSIMigrationRBD
	{"CSIMigrationRBD", Alpha, 23, 27},
	{"CSIMigrationRBD", Deprecated, 28, -1},

	// CSIMigrationvSphere
	{"CSIMigrationvSphere", Beta, 19, 25},
	{"CSIMigrationvSphere", GA, 26, 28},

	// CSIMigrationvSphereComplete
	{"CSIMigrationvSphereComplete", Beta, 19, 21},

	// CSINodeExpandSecret
	{"CSINodeExpandSecret", Alpha, 25, 26},
	{"CSINodeExpandSecret", Beta, 27, 28},
	{"CSINodeExpandSecret", GA, 29, -1},

	// CSINodeInfo
	{"CSINodeInfo", Alpha, 12, 13},
	{"CSINodeInfo", Beta, 14, 16},
	{"CSINodeInfo", GA, 17, 20},

	// CSIPersistentVolume
	{"CSIPersistentVolume", Alpha, 9, 9},
	{"CSIPersistentVolume", Beta, 10, 12},
	{"CSIPersistentVolume", GA, 13, 15},

	// CSIServiceAccountToken
	{"CSIServiceAccountToken", Alpha, 20, 20},
	{"CSIServiceAccountToken", Beta, 21, 21},
	{"CSIServiceAccountToken", GA, 22, 24},

	// CSIStorageCapacity
	{"CSIStorageCapacity", Alpha, 19, 20},
	{"CSIStorageCapacity", Beta, 21, 23},
	{"CSIStorageCapacity", GA, 24, 27},

	// CSIVolumeFSGroupPolicy
	{"CSIVolumeFSGroupPolicy", Alpha, 19, 19},
	{"CSIVolumeFSGroupPolicy", Beta, 20, 22},
	{"CSIVolumeFSGroupPolicy", GA, 23, 24},

	// CSIVolumeHealth
	{"CSIVolumeHealth", Alpha, 21, -1},

	// CSRDuration
	{"CSRDuration", Beta, 22, 23},
	{"CSRDuration", GA, 24, 25},

	// CloudControllerManagerWebhook
	{"CloudControllerManagerWebhook", Alpha, 27, -1},

	// CloudDualStackNodeIPs
	{"CloudDualStackNodeIPs", Alpha, 27, 28},
	{"CloudDualStackNodeIPs", Beta, 29, -1},

	// ClusterTrustBundle
	{"ClusterTrustBundle", Alpha, 27, -1},

	// ClusterTrustBundleProjection
	{"ClusterTrustBundleProjection", Alpha, 29, -1},

	// ComponentSLIs
	{"ComponentSLIs", Alpha, 26, 26},
	{"ComponentSLIs", Beta, 27, -1},

	// ConfigurableFSGroupPolicy
	{"ConfigurableFSGroupPolicy", Alpha, 18, 19},
	{"ConfigurableFSGroupPolicy", Beta, 20, 22},
	{"ConfigurableFSGroupPolicy", GA, 23, 24},

	// ConsistentHTTPGetHandlers
	{"ConsistentHTTPGetHandlers", GA, 26, -1},

	// ConsistentListFromCache
	{"ConsistentListFromCache", Alpha, 28, -1},

	// ContainerCheckpoint
	{"ContainerCheckpoint", Alpha, 25, -1},

	// ContextualLogging
	{"ContextualLogging", Alpha, 25, -1},

	// ControllerManagerLeaderMigration
	{"ControllerManagerLeaderMigration", Alpha, 21, 21},
	{"ControllerManagerLeaderMigration", Beta, 22, 23},
	{"ControllerManagerLeaderMigration", GA, 24, 26},

	// CronJobControllerV2
	{"CronJobControllerV2", Alpha, 20, 20},
	{"CronJobControllerV2", Beta, 21, 21},
	{"CronJobControllerV2", GA, 22, 22},

	// CronJobTimeZone
	{"CronJobTimeZone", Alpha, 24, 24},
	{"CronJobTimeZone", Beta, 25, 26},
	{"CronJobTimeZone", GA, 27, 28},

	// CronJobsScheduledAnnotation
	{"CronJobsScheduledAnnotation", Beta, 28, -1},

	// CrossNamespaceVolumeDataSource
	{"CrossNamespaceVolumeDataSource", Alpha, 26, -1},

	// CustomPodDNS
	{"CustomPodDNS", Alpha, 9, 9},
	{"CustomPodDNS", Beta, 10, 13},
	{"CustomPodDNS", GA, 14, 15},

	// CustomResourceDefaulting
	{"CustomResourceDefaulting", Alpha, 15, 15},
	{"CustomResourceDefaulting", Beta, 16, 16},
	{"CustomResourceDefaulting", GA, 17, 17},

	// CustomResourcePublishOpenAPI
	{"CustomResourcePublishOpenAPI", Alpha, 14, 14},
	{"CustomResourcePublishOpenAPI", Beta, 15, 15},
	{"CustomResourcePublishOpenAPI", GA, 16, 17},

	// CustomResourceSubresources
	{"CustomResourceSubresources", Alpha, 10, 10},
	{"CustomResourceSubresources", Beta, 11, 15},
	{"CustomResourceSubresources", GA, 16, 17},

	// CustomResourceValidation
	{"CustomResourceValidation", Alpha, 8, 8},
	{"CustomResourceValidation", Beta, 9, 15},
	{"CustomResourceValidation", GA, 16, 17},

	// CustomResourceValidationExpressions
	{"CustomResourceValidationExpressions", Alpha, 23, 24},
	{"CustomResourceValidationExpressions", Beta, 25, 28},
	{"CustomResourceValidationExpressions", GA, 29, -1},

	// CustomResourceWebhookConversion
	{"CustomResourceWebhookConversion", Alpha, 13, 14},
	{"CustomResourceWebhookConversion", Beta, 15, 15},
	{"CustomResourceWebhookConversion", GA, 16, 17},

	// DaemonSetUpdateSurge
	{"DaemonSetUpdateSurge", Alpha, 21, 21},
	{"DaemonSetUpdateSurge", Beta, 22, 24},
	{"DaemonSetUpdateSurge", GA, 25, 26},

	// DebugContainers
	{"DebugContainers", Alpha, 8, 15},

	// DefaultHostNetworkHostPortsInPodTemplates
	{"DefaultHostNetworkHostPortsInPodTemplates", Deprecated, 28, -1},

	// DefaultIngressClass
	{"DefaultIngressClass", Beta, 18, 18},
	{"DefaultIngressClass", GA, 19, 19},

	// DefaultPodTopologySpread
	{"DefaultPodTopologySpread", Alpha, 19, 19},
	{"DefaultPodTopologySpread", Beta, 20, 23},
	{"DefaultPodTopologySpread", GA, 24, 25},

	// DelegateFSGroupToCSIDriver
	{"DelegateFSGroupToCSIDriver", Alpha, 22, 22},
	{"DelegateFSGroupToCSIDriver", Beta, 23, 25},
	{"DelegateFSGroupToCSIDriver", GA, 26, 27},

	// DevicePluginCDIDevices
	{"DevicePluginCDIDevices", Alpha, 28, 28},
	{"DevicePluginCDIDevices", Beta, 29, -1},

	// DevicePlugins
	{"DevicePlugins", Alpha, 8, 9},
	{"DevicePlugins", Beta, 10, 25},
	{"DevicePlugins", GA, 26, 27},

	// DisableAcceleratorUsageMetrics
	{"DisableAcceleratorUsageMetrics", Alpha, 19, 19},
	{"DisableAcceleratorUsageMetrics", Beta, 20, 24},
	{"DisableAcceleratorUsageMetrics", GA, 25, 27},

	// DisableCloudProviders
	{"DisableCloudProviders", Alpha, 22, 28},
	{"DisableCloudProviders", Beta, 29, -1},

	// DisableKubeletCloudCredentialProviders
	{"DisableKubeletCloudCredentialProviders", Alpha, 23, 28},
	{"DisableKubeletCloudCredentialProviders", Beta, 29, -1},

	// DisableNodeKubeProxyVersion
	{"DisableNodeKubeProxyVersion", Alpha, 29, -1},

	// DownwardAPIHugePages
	{"DownwardAPIHugePages", Alpha, 20, 20},
	{"DownwardAPIHugePages", Beta, 21, 26},
	{"DownwardAPIHugePages", GA, 27, 28},

	// DryRun
	{"DryRun", Alpha, 12, 12},
	{"DryRun", Beta, 13, 18},
	{"DryRun", GA, 19, 27},

	// DynamicAuditing
	{"DynamicAuditing", Alpha, 13, 18},

	// DynamicKubeletConfig
	{"DynamicKubeletConfig", Alpha, 6, 10},
	{"DynamicKubeletConfig", Beta, 11, 21},
	{"DynamicKubeletConfig", Deprecated, 22, 25},

	// DynamicProvisioningScheduling
	{"DynamicProvisioningScheduling", Alpha, 11, 11},

	// DynamicResourceAllocation
	{"DynamicResourceAllocation", Alpha, 26, -1},

	// DynamicVolumeProvisioning
	{"DynamicVolumeProvisioning", Alpha, 6, 7},

	// EfficientWatchResumption
	{"EfficientWatchResumption", Alpha, 20, 20},
	{"EfficientWatchResumption", Beta, 21, 23},
	{"EfficientWatchResumption", GA, 24, -1},

	// ElasticIndexedJob
	{"ElasticIndexedJob", Beta, 27, -1},

	// EnableAggregatedDiscoveryTimeout
	{"EnableAggregatedDiscoveryTimeout", Deprecated, 16, 16},

	// EnableEquivalenceClassCache
	{"EnableEquivalenceClassCache", Alpha, 8, 13},

	// EndpointSlice
	{"EndpointSlice", Alpha, 16, 16},
	{"EndpointSlice", Beta, 17, 20},
	{"EndpointSlice", GA, 21, 24},

	// EndpointSliceNodeName
	{"EndpointSliceNodeName", Alpha, 20, 20},
	{"EndpointSliceNodeName", GA, 21, 24},

	// EndpointSliceProxying
	{"EndpointSliceProxying", Alpha, 18, 18},
	{"EndpointSliceProxying", Beta, 19, 21},
	{"EndpointSliceProxying", GA, 22, 24},

	// EndpointSliceTerminatingCondition
	{"EndpointSliceTerminatingCondition", Alpha, 20, 21},
	{"EndpointSliceTerminatingCondition", Beta, 22, 25},
	{"EndpointSliceTerminatingCondition", GA, 26, 27},

	// EphemeralContainers
	{"EphemeralContainers", Alpha, 16, 22},
	{"EphemeralContainers", Beta, 23, 24},
	{"EphemeralContainers", GA, 25, 26},

	// EvenPodsSpread
	{"EvenPodsSpread", Alpha, 16, 17},
	{"EvenPodsSpread", Beta, 18, 18},
	{"EvenPodsSpread", GA, 19, 20},

	// EventedPLEG
	{"EventedPLEG", Alpha, 26, 26},
	{"EventedPLEG", Beta, 27, -1},

	// ExecProbeTimeout
	{"ExecProbeTimeout", GA, 20, -1},

	// ExpandCSIVolumes
	{"ExpandCSIVolumes", Alpha, 14, 15},
	{"ExpandCSIVolumes", Beta, 16, 23},
	{"ExpandCSIVolumes", GA, 24, 26},

	// ExpandInUsePersistentVolumes
	{"ExpandInUsePersistentVolumes", Alpha, 11, 14},
	{"ExpandInUsePersistentVolumes", Beta, 15, 23},
	{"ExpandInUsePersistentVolumes", GA, 24, 26},

	// ExpandPersistentVolumes
	{"ExpandPersistentVolumes", Alpha, 8, 10},
	{"ExpandPersistentVolumes", Beta, 11, 23},
	{"ExpandPersistentVolumes", GA, 24, 26},

	// ExpandedDNSConfig
	{"ExpandedDNSConfig", Alpha, 22, 25},
	{"ExpandedDNSConfig", Beta, 26, 27},
	{"ExpandedDNSConfig", GA, 28, -1},

	// ExperimentalCriticalPodAnnotation
	{"ExperimentalCriticalPodAnnotation", Alpha, 6, 15},

	// ExperimentalHostUserNamespaceDefaulting
	{"ExperimentalHostUserNamespaceDefaulting", Beta, 6, 27},
	{"ExperimentalHostUserNamespaceDefaulting", Deprecated, 28, -1},

	// ExternalPolicyForExternalIP
	{"ExternalPolicyForExternalIP", GA, 18, 21},

	// ExternalTrafficLocalOnly
	{"ExternalTrafficLocalOnly", Beta, 6, 6},
	{"ExternalTrafficLocalOnly", GA, 7, 9},

	// GCERegionalPersistentDisk
	{"GCERegionalPersistentDisk", Beta, 10, 12},
	{"GCERegionalPersistentDisk", GA, 13, 14},

	// GRPCContainerProbe
	{"GRPCContainerProbe", Alpha, 23, 23},
	{"GRPCContainerProbe", Beta, 24, 26},
	{"GRPCContainerProbe", GA, 27, 28},

	// GenericEphemeralVolume
	{"GenericEphemeralVolume", Alpha, 19, 20},
	{"GenericEphemeralVolume", Beta, 21, 22},
	{"GenericEphemeralVolume", GA, 23, 24},

	// GracefulNodeShutdown
	{"GracefulNodeShutdown", Alpha, 20, 20},
	{"GracefulNodeShutdown", Beta, 21, -1},

	// GracefulNodeShutdownBasedOnPodPriority
	{"GracefulNodeShutdownBasedOnPodPriority", Alpha, 23, 23},
	{"GracefulNodeShutdownBasedOnPodPriority", Beta, 24, -1},

	// HPAContainerMetrics
	{"HPAContainerMetrics", Alpha, 20, 26},
	{"HPAContainerMetrics", Beta, 27, -1},

	// HPAScaleToZero
	{"HPAScaleToZero", Alpha, 16, -1},

	// HonorPVReclaimPolicy
	{"HonorPVReclaimPolicy", Alpha, 23, -1},

	// HugePageStorageMediumSize
	{"HugePageStorageMediumSize", Alpha, 18, 18},
	{"HugePageStorageMediumSize", Beta, 19, 21},
	{"HugePageStorageMediumSize", GA, 22, 23},

	// HugePages
	{"HugePages", Alpha, 8, 9},
	{"HugePages", Beta, 10, 13},
	{"HugePages", GA, 14, 15},

	// HyperVContainer
	{"HyperVContainer", Alpha, 10, 19},
	{"HyperVContainer", Deprecated, 20, 20},

	// IPTablesOwnershipCleanup
	{"IPTablesOwnershipCleanup", Alpha, 25, 26},
	{"IPTablesOwnershipCleanup", Beta, 27, 27},
	{"IPTablesOwnershipCleanup", GA, 28, -1},

	// IPv6DualStack
	{"IPv6DualStack", Alpha, 16, 20},
	{"IPv6DualStack", Beta, 21, 22},
	{"IPv6DualStack", GA, 23, 26},

	// IdentifyPodOS
	{"IdentifyPodOS", Alpha, 23, 23},
	{"IdentifyPodOS", Beta, 24, 24},
	{"IdentifyPodOS", GA, 25, 26},

	// ImageMaximumGCAge
	{"ImageMaximumGCAge", Alpha, 29, -1},

	// ImmutableEphemeralVolumes
	{"ImmutableEphemeralVolumes", Alpha, 18, 18},
	{"ImmutableEphemeralVolumes", Beta, 19, 20},
	{"ImmutableEphemeralVolumes", GA, 21, 23},

	// InPlacePodVerticalScaling
	{"InPlacePodVerticalScaling", Alpha, 27, -1},

	// InTreePluginAWSUnregister
	{"InTreePluginAWSUnregister", Alpha, 21, -1},

	// InTreePluginAzureDiskUnregister
	{"InTreePluginAzureDiskUnregister", Alpha, 21, -1},

	// InTreePluginAzureFileUnregister
	{"InTreePluginAzureFileUnregister", Alpha, 21, -1},

	// InTreePluginGCEUnregister
	{"InTreePluginGCEUnregister", Alpha, 21, -1},

	// InTreePluginOpenStackUnregister
	{"InTreePluginOpenStackUnregister", Alpha, 21, -1},

	// InTreePluginPortworxUnregister
	{"InTreePluginPortworxUnregister", Alpha, 23, -1},

	// InTreePluginRBDUnregister
	{"InTreePluginRBDUnregister", Alpha, 23, 27},
	{"InTreePluginRBDUnregister", Deprecated, 28, -1},

	// InTreePluginvSphereUnregister
	{"InTreePluginvSphereUnregister", Alpha, 21, -1},

	// IndexedJob
	{"IndexedJob", Alpha, 21, 21},
	{"IndexedJob", Beta, 22, 23},
	{"IndexedJob", GA, 24, 25},

	// IngressClassNamespacedParams
	{"IngressClassNamespacedParams", Alpha, 21, 21},
	{"IngressClassNamespacedParams", Beta, 22, 22},
	{"IngressClassNamespacedParams", GA, 23, 24},

	// Initializers
	{"Initializers", Alpha, 8, 13},

	// JobBackoffLimitPerIndex
	{"JobBackoffLimitPerIndex", Alpha, 28, 28},
	{"JobBackoffLimitPerIndex", Beta, 29, -1},

	// JobMutableNodeSchedulingDirectives
	{"JobMutableNodeSchedulingDirectives", Beta, 23, 26},
	{"JobMutableNodeSchedulingDirectives", GA, 27, 28},

	// JobPodFailurePolicy
	{"JobPodFailurePolicy", Alpha, 25, 25},
	{"JobPodFailurePolicy", Beta, 26, -1},

	// JobPodReplacementPolicy
	{"JobPodReplacementPolicy", Alpha, 28, 28},
	{"JobPodReplacementPolicy", Beta, 29, -1},

	// JobReadyPods
	{"JobReadyPods", Alpha, 23, 23},
	{"JobReadyPods", Beta, 24, 28},
	{"JobReadyPods", GA, 29, -1},

	// JobTrackingWithFinalizers
	{"JobTrackingWithFinalizers", Alpha, 22, 22},
	{"JobTrackingWithFinalizers", Beta, 23, 25},
	{"JobTrackingWithFinalizers", GA, 26, 28},

	// KMSv1
	{"KMSv1", Deprecated, 28, -1},

	// KMSv2
	{"KMSv2", Alpha, 25, 26},
	{"KMSv2", Beta, 27, 28},
	{"KMSv2", GA, 29, -1},

	// KMSv2KDF
	{"KMSv2KDF", Beta, 28, 28},
	{"KMSv2KDF", GA, 29, -1},

	// KubeProxyDrainingTerminatingNodes
	{"KubeProxyDrainingTerminatingNodes", Alpha, 28, -1},

	// KubeletCgroupDriverFromCRI
	{"KubeletCgroupDriverFromCRI", Alpha, 28, -1},

	// KubeletConfigFile
	{"KubeletConfigFile", Alpha, 8, 9},

	// KubeletCredentialProviders
	{"KubeletCredentialProviders", Alpha, 20, 23},
	{"KubeletCredentialProviders", Beta, 24, 25},
	{"KubeletCredentialProviders", GA, 26, 27},

	// KubeletInUserNamespace
	{"KubeletInUserNamespace", Alpha, 22, -1},

	// KubeletPluginsWatcher
	{"KubeletPluginsWatcher", Alpha, 11, 11},
	{"KubeletPluginsWatcher", Beta, 12, 12},
	{"KubeletPluginsWatcher", GA, 13, 15},

	// KubeletPodResources
	{"KubeletPodResources", Alpha, 13, 14},
	{"KubeletPodResources", Beta, 15, 27},
	{"KubeletPodResources", GA, 28, -1},

	// KubeletPodResourcesDynamicResources
	{"KubeletPodResourcesDynamicResources", Alpha, 27, -1},

	// KubeletPodResourcesGet
	{"KubeletPodResourcesGet", Alpha, 27, -1},

	// KubeletPodResourcesGetAllocatable
	{"KubeletPodResourcesGetAllocatable", Alpha, 21, 22},
	{"KubeletPodResourcesGetAllocatable", Beta, 23, 27},
	{"KubeletPodResourcesGetAllocatable", GA, 28, -1},

	// KubeletSeparateDiskGC
	{"KubeletSeparateDiskGC", Alpha, 29, -1},

	// KubeletTracing
	{"KubeletTracing", Alpha, 25, 26},
	{"KubeletTracing", Beta, 27, -1},

	// LegacyNodeRoleBehavior
	{"LegacyNodeRoleBehavior", Alpha, 16, 18},
	{"LegacyNodeRoleBehavior", Beta, 19, 20},
	{"LegacyNodeRoleBehavior", GA, 21, 21},

	// LegacyServiceAccountTokenCleanUp
	{"LegacyServiceAccountTokenCleanUp", Alpha, 28, 28},
	{"LegacyServiceAccountTokenCleanUp", Beta, 29, -1},

	// LegacyServiceAccountTokenNoAutoGeneration
	{"LegacyServiceAccountTokenNoAutoGeneration", Beta, 24, 25},
	{"LegacyServiceAccountTokenNoAutoGeneration", GA, 26, 28},

	// LegacyServiceAccountTokenTracking
	{"LegacyServiceAccountTokenTracking", Alpha, 26, 26},
	{"LegacyServiceAccountTokenTracking", Beta, 27, 27},
	{"LegacyServiceAccountTokenTracking", GA, 28, -1},

	// LoadBalancerIPMode
	{"LoadBalancerIPMode", Alpha, 29, -1},

	// LocalStorageCapacityIsolation
	{"LocalStorageCapacityIsolation", Alpha, 7, 9},
	{"LocalStorageCapacityIsolation", Beta, 10, 24},
	{"LocalStorageCapacityIsolation", GA, 25, 26},

	// LocalStorageCapacityIsolationFSQuotaMonitoring
	{"LocalStorageCapacityIsolationFSQuotaMonitoring", Alpha, 15, -1},

	// LogarithmicScaleDown
	{"LogarithmicScaleDown", Alpha, 21, 21},
	{"LogarithmicScaleDown", Beta, 22, -1},

	// LoggingAlphaOptions
	{"LoggingAlphaOptions", Alpha, 25, -1},

	// LoggingBetaOptions
	{"LoggingBetaOptions", Beta, 25, -1},

	// MatchLabelKeysInPodAffinity
	{"MatchLabelKeysInPodAffinity", Alpha, 29, -1},

	// MatchLabelKeysInPodTopologySpread
	{"MatchLabelKeysInPodTopologySpread", Alpha, 25, 26},
	{"MatchLabelKeysInPodTopologySpread", Beta, 27, -1},

	// MaxUnavailableStatefulSet
	{"MaxUnavailableStatefulSet", Alpha, 24, -1},

	// MemoryManager
	{"MemoryManager", Alpha, 21, 21},
	{"MemoryManager", Beta, 22, -1},

	// MemoryQoS
	{"MemoryQoS", Alpha, 22, -1},

	// MinDomainsInPodTopologySpread
	{"MinDomainsInPodTopologySpread", Alpha, 24, 24},
	{"MinDomainsInPodTopologySpread", Beta, 25, -1},

	// MinimizeIPTablesRestore
	{"MinimizeIPTablesRestore", Alpha, 26, 26},
	{"MinimizeIPTablesRestore", Beta, 27, 27},
	{"MinimizeIPTablesRestore", GA, 28, -1},

	// MixedProtocolLBService
	{"MixedProtocolLBService", Alpha, 20, 23},
	{"MixedProtocolLBService", Beta, 24, 25},
	{"MixedProtocolLBService", GA, 26, 27},

	// MountContainers
	{"MountContainers", Alpha, 9, 16},

	// MountPropagation
	{"MountPropagation", Alpha, 8, 9},
	{"MountPropagation", Beta, 10, 11},
	{"MountPropagation", GA, 12, 13},

	// MultiCIDRRangeAllocator
	{"MultiCIDRRangeAllocator", Alpha, 25, 28},

	// MultiCIDRServiceAllocator
	{"MultiCIDRServiceAllocator", Alpha, 27, -1},

	// NFTablesProxyMode
	{"NFTablesProxyMode", Alpha, 29, -1},

	// NamespaceDefaultLabelName
	{"NamespaceDefaultLabelName", Beta, 21, 21},
	{"NamespaceDefaultLabelName", GA, 22, 23},

	// NetworkPolicyEndPort
	{"NetworkPolicyEndPort", Alpha, 21, 21},
	{"NetworkPolicyEndPort", Beta, 22, 24},
	{"NetworkPolicyEndPort", GA, 25, 26},

	// NetworkPolicyStatus
	{"NetworkPolicyStatus", Alpha, 24, 27},

	// NewVolumeManagerReconstruction
	{"NewVolumeManagerReconstruction", Beta, 27, -1},

	// NodeDisruptionExclusion
	{"NodeDisruptionExclusion", Alpha, 16, 18},
	{"NodeDisruptionExclusion", Beta, 19, 20},
	{"NodeDisruptionExclusion", GA, 21, 21},

	// NodeInclusionPolicyInPodTopologySpread
	{"NodeInclusionPolicyInPodTopologySpread", Alpha, 25, 25},
	{"NodeInclusionPolicyInPodTopologySpread", Beta, 26, -1},

	// NodeLease
	{"NodeLease", Alpha, 12, 13},
	{"NodeLease", Beta, 14, 16},
	{"NodeLease", GA, 17, 22},

	// NodeLogQuery
	{"NodeLogQuery", Alpha, 27, -1},

	// NodeOutOfServiceVolumeDetach
	{"NodeOutOfServiceVolumeDetach", Alpha, 24, 25},
	{"NodeOutOfServiceVolumeDetach", Beta, 26, 27},
	{"NodeOutOfServiceVolumeDetach", GA, 28, -1},

	// NodeSwap
	{"NodeSwap", Alpha, 22, 27},
	{"NodeSwap", Beta, 28, -1},

	// NonPreemptingPriority
	{"NonPreemptingPriority", Alpha, 15, 18},
	{"NonPreemptingPriority", Beta, 19, 23},
	{"NonPreemptingPriority", GA, 24, 25},

	// OpenAPIEnums
	{"OpenAPIEnums", Alpha, 23, 23},
	{"OpenAPIEnums", Beta, 24, -1},

	// OpenAPIV3
	{"OpenAPIV3", Alpha, 23, 23},
	{"OpenAPIV3", Beta, 24, 26},
	{"OpenAPIV3", GA, 27, 28},

	// PDBUnhealthyPodEvictionPolicy
	{"PDBUnhealthyPodEvictionPolicy", Alpha, 26, 26},
	{"PDBUnhealthyPodEvictionPolicy", Beta, 27, -1},

	// PVCProtection
	{"PVCProtection", Alpha, 9, 9},

	// PersistentLocalVolumes
	{"PersistentLocalVolumes", Alpha, 7, 9},
	{"PersistentLocalVolumes", Beta, 10, 13},
	{"PersistentLocalVolumes", GA, 14, 16},

	// PersistentVolumeLastPhaseTransitionTime
	{"PersistentVolumeLastPhaseTransitionTime", Alpha, 28, 28},
	{"PersistentVolumeLastPhaseTransitionTime", Beta, 29, -1},

	// PodAffinityNamespaceSelector
	{"PodAffinityNamespaceSelector", Alpha, 21, 21},
	{"PodAffinityNamespaceSelector", Beta, 22, 23},
	{"PodAffinityNamespaceSelector", GA, 24, 25},

	// PodAndContainerStatsFromCRI
	{"PodAndContainerStatsFromCRI", Alpha, 23, -1},

	// PodDeletionCost
	{"PodDeletionCost", Alpha, 21, 21},
	{"PodDeletionCost", Beta, 22, -1},

	// PodDisruptionBudget
	{"PodDisruptionBudget", Beta, 17, 20},
	{"PodDisruptionBudget", GA, 21, 24},

	// PodDisruptionConditions
	{"PodDisruptionConditions", Alpha, 25, 25},
	{"PodDisruptionConditions", Beta, 26, -1},

	// PodHasNetworkCondition
	{"PodHasNetworkCondition", Alpha, 25, 27},

	// PodHostIPs
	{"PodHostIPs", Alpha, 28, 28},
	{"PodHostIPs", Beta, 29, -1},

	// PodIndexLabel
	{"PodIndexLabel", Beta, 28, -1},

	// PodLifecycleSleepAction
	{"PodLifecycleSleepAction", Alpha, 29, -1},

	// PodOverhead
	{"PodOverhead", Alpha, 16, 17},
	{"PodOverhead", Beta, 18, 23},
	{"PodOverhead", GA, 24, 25},

	// PodPriority
	{"PodPriority", Alpha, 8, 10},
	{"PodPriority", Beta, 11, 13},
	{"PodPriority", GA, 14, 17},

	// PodReadinessGates
	{"PodReadinessGates", Beta, 11, 13},
	{"PodReadinessGates", GA, 14, 15},

	// PodReadyToStartContainersCondition
	{"PodReadyToStartContainersCondition", Alpha, 28, 28},
	{"PodReadyToStartContainersCondition", Beta, 29, -1},

	// PodSchedulingReadiness
	{"PodSchedulingReadiness", Alpha, 26, 26},
	{"PodSchedulingReadiness", Beta, 27, -1},

	// PodSecurity
	{"PodSecurity", Alpha, 22, 22},
	{"PodSecurity", Beta, 23, 24},
	{"PodSecurity", GA, 25, 27},

	// PodShareProcessNamespace
	{"PodShareProcessNamespace", Alpha, 10, 11},
	{"PodShareProcessNamespace", Beta, 12, 16},
	{"PodShareProcessNamespace", GA, 17, 18},

	// PreferNominatedNode
	{"PreferNominatedNode", Alpha, 21, 21},
	{"PreferNominatedNode", Beta, 22, 23},
	{"PreferNominatedNode", GA, 24, 25},

	// ProbeTerminationGracePeriod
	{"ProbeTerminationGracePeriod", Alpha, 21, 21},
	{"ProbeTerminationGracePeriod", Beta, 22, 27},
	{"ProbeTerminationGracePeriod", GA, 28, 28},

	// ProcMountType
	{"ProcMountType", Alpha, 12, -1},

	// ProxyTerminatingEndpoints
	{"ProxyTerminatingEndpoints", Alpha, 22, 25},
	{"ProxyTerminatingEndpoints", Beta, 26, 27},
	{"ProxyTerminatingEndpoints", GA, 28, -1},

	// QOSReserved
	{"QOSReserved", Alpha, 11, -1},

	// ReadOnlyAPIDataVolumes
	{"ReadOnlyAPIDataVolumes", Deprecated, 7, 11},

	// ReadWriteOncePod
	{"ReadWriteOncePod", Alpha, 22, 26},
	{"ReadWriteOncePod", Beta, 27, 28},
	{"ReadWriteOncePod", GA, 29, -1},

	// RecoverVolumeExpansionFailure
	{"RecoverVolumeExpansionFailure", Alpha, 23, -1},

	// RemainingItemCount
	{"RemainingItemCount", Alpha, 15, 15},
	{"RemainingItemCount", Beta, 16, 28},
	{"RemainingItemCount", GA, 29, -1},

	// RemoveSelfLink
	{"RemoveSelfLink", Alpha, 16, 19},
	{"RemoveSelfLink", Beta, 20, 23},
	{"RemoveSelfLink", GA, 24, -1},

	// RequestManagement
	{"RequestManagement", Alpha, 15, 16},

	// ResourceLimitsPriorityFunction
	{"ResourceLimitsPriorityFunction", Alpha, 9, 18},

	// ResourceQuotaScopeSelectors
	{"ResourceQuotaScopeSelectors", Alpha, 11, 11},
	{"ResourceQuotaScopeSelectors", Beta, 12, 16},
	{"ResourceQuotaScopeSelectors", GA, 17, 17},

	// RetroactiveDefaultStorageClass
	{"RetroactiveDefaultStorageClass", Alpha, 25, 25},
	{"RetroactiveDefaultStorageClass", Beta, 26, 27},
	{"RetroactiveDefaultStorageClass", GA, 28, 28},

	// RootCAConfigMap
	{"RootCAConfigMap", Beta, 20, 20},
	{"RootCAConfigMap", GA, 21, 21},

	// RotateKubeletClientCertificate
	{"RotateKubeletClientCertificate", Alpha, 7, 7},
	{"RotateKubeletClientCertificate", Beta, 8, 18},
	{"RotateKubeletClientCertificate", GA, 19, 20},

	// RotateKubeletServerCertificate
	{"RotateKubeletServerCertificate", Alpha, 7, 11},
	{"RotateKubeletServerCertificate", Beta, 12, -1},

	// RunAsGroup
	{"RunAsGroup", Alpha, 10, 13},
	{"RunAsGroup", Beta, 14, 20},
	{"RunAsGroup", GA, 21, 21},

	// RuntimeClass
	{"RuntimeClass", Alpha, 12, 13},
	{"RuntimeClass", Beta, 14, 19},
	{"RuntimeClass", GA, 20, 23},

	// RuntimeClassInImageCriAPI
	{"RuntimeClassInImageCriAPI", Alpha, 29, -1},

	// SCTPSupport
	{"SCTPSupport", Alpha, 12, 18},
	{"SCTPSupport", Beta, 19, 19},
	{"SCTPSupport", GA, 20, 21},

	// SELinuxMountReadWriteOncePod
	{"SELinuxMountReadWriteOncePod", Alpha, 25, 26},
	{"SELinuxMountReadWriteOncePod", Beta, 27, -1},

	// ScheduleDaemonSetPods
	{"ScheduleDaemonSetPods", Alpha, 10, 11},
	{"ScheduleDaemonSetPods", Beta, 12, 16},
	{"ScheduleDaemonSetPods", GA, 17, 17},

	// SchedulerQueueingHints
	{"SchedulerQueueingHints", Beta, 28, -1},

	// SeccompDefault
	{"SeccompDefault", Alpha, 22, 24},
	{"SeccompDefault", Beta, 25, 26},
	{"SeccompDefault", GA, 27, 28},

	// SecurityContextDeny
	{"SecurityContextDeny", Alpha, 27, -1},

	// SelectorIndex
	{"SelectorIndex", Alpha, 18, 18},
	{"SelectorIndex", Beta, 19, 19},
	{"SelectorIndex", GA, 20, 24},

	// SeparateTaintEvictionController
	{"SeparateTaintEvictionController", Beta, 29, -1},

	// ServerSideApply
	{"ServerSideApply", Alpha, 14, 15},
	{"ServerSideApply", Beta, 16, 21},
	{"ServerSideApply", GA, 22, -1},

	// ServerSideFieldValidation
	{"ServerSideFieldValidation", Alpha, 23, 24},
	{"ServerSideFieldValidation", Beta, 25, 26},
	{"ServerSideFieldValidation", GA, 27, -1},

	// ServiceAccountIssuerDiscovery
	{"ServiceAccountIssuerDiscovery", Alpha, 18, 19},
	{"ServiceAccountIssuerDiscovery", Beta, 20, 20},
	{"ServiceAccountIssuerDiscovery", GA, 21, 22},

	// ServiceAccountTokenJTI
	{"ServiceAccountTokenJTI", Alpha, 29, -1},

	// ServiceAccountTokenNodeBinding
	{"ServiceAccountTokenNodeBinding", Alpha, 29, -1},

	// ServiceAccountTokenNodeBindingValidation
	{"ServiceAccountTokenNodeBindingValidation", Alpha, 29, -1},

	// ServiceAccountTokenPodNodeInfo
	{"ServiceAccountTokenPodNodeInfo", Alpha, 29, -1},

	// ServiceAppProtocol
	{"ServiceAppProtocol", Alpha, 18, 18},
	{"ServiceAppProtocol", Beta, 19, 19},
	{"ServiceAppProtocol", GA, 20, 21},

	// ServiceIPStaticSubrange
	{"ServiceIPStaticSubrange", Alpha, 24, 24},
	{"ServiceIPStaticSubrange", Beta, 25, 25},
	{"ServiceIPStaticSubrange", GA, 26, 27},

	// ServiceInternalTrafficPolicy
	{"ServiceInternalTrafficPolicy", Alpha, 21, 21},
	{"ServiceInternalTrafficPolicy", Beta, 22, 25},
	{"ServiceInternalTrafficPolicy", GA, 26, 27},

	// ServiceLBNodePortControl
	{"ServiceLBNodePortControl", Alpha, 20, 21},
	{"ServiceLBNodePortControl", Beta, 22, 23},
	{"ServiceLBNodePortControl", GA, 24, 25},

	// ServiceLoadBalancerClass
	{"ServiceLoadBalancerClass", Alpha, 21, 21},
	{"ServiceLoadBalancerClass", Beta, 22, 23},
	{"ServiceLoadBalancerClass", GA, 24, 25},

	// ServiceLoadBalancerFinalizer
	{"ServiceLoadBalancerFinalizer", Alpha, 15, 15},
	{"ServiceLoadBalancerFinalizer", Beta, 16, 16},
	{"ServiceLoadBalancerFinalizer", GA, 17, 19},

	// ServiceNodeExclusion
	{"ServiceNodeExclusion", Alpha, 8, 18},
	{"ServiceNodeExclusion", Beta, 19, 20},
	{"ServiceNodeExclusion", GA, 21, 21},

	// ServiceNodePortStaticSubrange
	{"ServiceNodePortStaticSubrange", Alpha, 27, 27},
	{"ServiceNodePortStaticSubrange", Beta, 28, 28},
	{"ServiceNodePortStaticSubrange", GA, 29, -1},

	// ServiceProxyAllowExternalIPs
	{"ServiceProxyAllowExternalIPs", Deprecated, 7, 11},

	// ServiceTopology
	{"ServiceTopology", Alpha, 17, 21},

	// SetHostnameAsFQDN
	{"SetHostnameAsFQDN", Alpha, 19, 19},
	{"SetHostnameAsFQDN", Beta, 20, 21},
	{"SetHostnameAsFQDN", GA, 22, 23},

	// SidecarContainers
	{"SidecarContainers", Alpha, 28, 28},
	{"SidecarContainers", Beta, 29, -1},

	// SizeMemoryBackedVolumes
	{"SizeMemoryBackedVolumes", Alpha, 20, 21},
	{"SizeMemoryBackedVolumes", Beta, 22, -1},

	// SkipReadOnlyValidationGCE
	{"SkipReadOnlyValidationGCE", Alpha, 28, 28},
	{"SkipReadOnlyValidationGCE", Deprecated, 29, -1},

	// StableLoadBalancerNodeSet
	{"StableLoadBalancerNodeSet", Beta, 27, -1},

	// StartupProbe
	{"StartupProbe", Alpha, 16, 17},
	{"StartupProbe", Beta, 18, 19},
	{"StartupProbe", GA, 20, 22},

	// StatefulSetAutoDeletePVC
	{"StatefulSetAutoDeletePVC", Alpha, 23, 26},
	{"StatefulSetAutoDeletePVC", Beta, 27, -1},

	// StatefulSetMinReadySeconds
	{"StatefulSetMinReadySeconds", Alpha, 22, 22},
	{"StatefulSetMinReadySeconds", Beta, 23, 24},
	{"StatefulSetMinReadySeconds", GA, 25, 26},

	// StatefulSetStartOrdinal
	{"StatefulSetStartOrdinal", Alpha, 26, 26},
	{"StatefulSetStartOrdinal", Beta, 27, -1},

	// StorageObjectInUseProtection
	{"StorageObjectInUseProtection", Beta, 10, 10},
	{"StorageObjectInUseProtection", GA, 11, 24},

	// StorageVersionAPI
	{"StorageVersionAPI", Alpha, 20, -1},

	// StorageVersionHash
	{"StorageVersionHash", Alpha, 14, 14},
	{"StorageVersionHash", Beta, 15, -1},

	// StreamingProxyRedirects
	{"StreamingProxyRedirects", Beta, 6, 17},
	{"StreamingProxyRedirects", Deprecated, 18, 23},

	// StructuredAuthenticationConfiguration
	{"StructuredAuthenticationConfiguration", Alpha, 29, -1},

	// StructuredAuthorizationConfiguration
	{"StructuredAuthorizationConfiguration", Alpha, 29, -1},

	// SupportIPVSProxyMode
	{"SupportIPVSProxyMode", Alpha, 8, 8},
	{"SupportIPVSProxyMode", Beta, 9, 10},
	{"SupportIPVSProxyMode", GA, 11, 19},

	// SupportNodePidsLimit
	{"SupportNodePidsLimit", Alpha, 14, 14},
	{"SupportNodePidsLimit", Beta, 15, 19},
	{"SupportNodePidsLimit", GA, 20, 22},

	// SupportPodPidsLimit
	{"SupportPodPidsLimit", Alpha, 10, 13},
	{"SupportPodPidsLimit", Beta, 14, 19},
	{"SupportPodPidsLimit", GA, 20, 22},

	// SuspendJob
	{"SuspendJob", Alpha, 21, 21},
	{"SuspendJob", Beta, 22, 23},
	{"SuspendJob", GA, 24, 25},

	// Sysctls
	{"Sysctls", Beta, 11, 20},
	{"Sysctls", GA, 21, 22},

	// TTLAfterFinished
	{"TTLAfterFinished", Alpha, 12, 20},
	{"TTLAfterFinished", Beta, 21, 22},
	{"TTLAfterFinished", GA, 23, 24},

	// TaintBasedEvictions
	{"TaintBasedEvictions", Alpha, 6, 12},
	{"TaintBasedEvictions", Beta, 13, 17},
	{"TaintBasedEvictions", GA, 18, 19},

	// TaintNodesByCondition
	{"TaintNodesByCondition", Alpha, 8, 11},
	{"TaintNodesByCondition", Beta, 12, 16},
	{"TaintNodesByCondition", GA, 17, 17},

	// TokenRequest
	{"TokenRequest", Alpha, 10, 11},
	{"TokenRequest", Beta, 12, 19},
	{"TokenRequest", GA, 20, 20},

	// TokenRequestProjection
	{"TokenRequestProjection", Alpha, 11, 11},
	{"TokenRequestProjection", Beta, 12, 19},
	{"TokenRequestProjection", GA, 20, 20},

	// TopologyAwareHints
	{"TopologyAwareHints", Alpha, 21, 22},
	{"TopologyAwareHints", Beta, 23, -1},

	// TopologyManager
	{"TopologyManager", Alpha, 16, 17},
	{"TopologyManager", Beta, 18, 26},
	{"TopologyManager", GA, 27, 28},

	// TopologyManagerPolicyAlphaOptions
	{"TopologyManagerPolicyAlphaOptions", Alpha, 26, -1},

	// TopologyManagerPolicyBetaOptions
	{"TopologyManagerPolicyBetaOptions", Beta, 26, -1},

	// TopologyManagerPolicyOptions
	{"TopologyManagerPolicyOptions", Alpha, 26, 27},
	{"TopologyManagerPolicyOptions", Beta, 28, -1},

	// TranslateStreamCloseWebsocketRequests
	{"TranslateStreamCloseWebsocketRequests", Alpha, 29, -1},

	// UnauthenticatedHTTP2DOSMitigation
	{"UnauthenticatedHTTP2DOSMitigation", Beta, 25, -1},

	// UnknownVersionInteroperabilityProxy
	{"UnknownVersionInteroperabilityProxy", Alpha, 28, -1},

	// UserNamespacesPodSecurityStandards
	{"UserNamespacesPodSecurityStandards", Alpha, 29, -1},

	// UserNamespacesStatelessPodsSupport
	{"UserNamespacesStatelessPodsSupport", Alpha, 25, 27},

	// UserNamespacesSupport
	{"UserNamespacesSupport", Alpha, 28, -1},

	// ValidateProxyRedirects
	{"ValidateProxyRedirects", Alpha, 12, 13},
	{"ValidateProxyRedirects", Beta, 14, 21},
	{"ValidateProxyRedirects", Deprecated, 22, 23},

	// ValidatingAdmissionPolicy
	{"ValidatingAdmissionPolicy", Alpha, 26, 27},
	{"ValidatingAdmissionPolicy", Beta, 28, -1},

	// VolumeAttributesClass
	{"VolumeAttributesClass", Alpha, 29, -1},

	// VolumeCapacityPriority
	{"VolumeCapacityPriority", Alpha, 21, -1},

	// VolumePVCDataSource
	{"VolumePVCDataSource", Alpha, 15, 15},
	{"VolumePVCDataSource", Beta, 16, 17},
	{"VolumePVCDataSource", GA, 18, 20},

	// VolumeScheduling
	{"VolumeScheduling", Alpha, 9, 9},
	{"VolumeScheduling", Beta, 10, 12},
	{"VolumeScheduling", GA, 13, 15},

	// VolumeSnapshotDataSource
	{"VolumeSnapshotDataSource", Alpha, 12, 16},
	{"VolumeSnapshotDataSource", Beta, 17, 19},
	{"VolumeSnapshotDataSource", GA, 20, 21},

	// VolumeSubpath
	{"VolumeSubpath", GA, 7, 24},

	// VolumeSubpathEnvExpansion
	{"VolumeSubpathEnvExpansion", Alpha, 11, 14},
	{"VolumeSubpathEnvExpansion", Beta, 15, 16},
	{"VolumeSubpathEnvExpansion", GA, 17, 18},

	// WarningHeaders
	{"WarningHeaders", Beta, 19, 21},
	{"WarningHeaders", GA, 22, 23},

	// WatchBookmark
	{"WatchBookmark", Alpha, 15, 15},
	{"WatchBookmark", Beta, 16, 16},
	{"WatchBookmark", GA, 17, -1},

	// WatchList
	{"WatchList", Alpha, 27, -1},

	// WinDSR
	{"WinDSR", Alpha, 14, -1},

	// WinOverlay
	{"WinOverlay", Alpha, 14, 19},
	{"WinOverlay", Beta, 20, -1},

	// WindowsEndpointSliceProxying
	{"WindowsEndpointSliceProxying", Alpha, 19, 20},
	{"WindowsEndpointSliceProxying", Beta, 21, 21},
	{"WindowsEndpointSliceProxying", GA, 22, 24},

	// WindowsGMSA
	{"WindowsGMSA", Alpha, 14, 15},
	{"WindowsGMSA", Beta, 16, 17},
	{"WindowsGMSA", GA, 18, 20},

	// WindowsHostNetwork
	{"WindowsHostNetwork", Alpha, 26, -1},

	// WindowsHostProcessContainers
	{"WindowsHostProcessContainers", Alpha, 22, 22},
	{"WindowsHostProcessContainers", Beta, 23, 25},
	{"WindowsHostProcessContainers", GA, 26, 27},

	// WindowsRunAsUserName
	{"WindowsRunAsUserName", Alpha, 16, 16},
	{"WindowsRunAsUserName", Beta, 17, 17},
	{"WindowsRunAsUserName", GA, 18, 20},

	// ZeroLimitedNominalConcurrencyShares
	{"ZeroLimitedNominalConcurrencyShares", Beta, 29, -1},

	// deprecatedGCERegionalPersistentDisk
	{"deprecatedGCERegionalPersistentDisk", GA, 15, 16},
}
