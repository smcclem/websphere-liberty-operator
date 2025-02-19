package v1

import (
	"time"

	"github.com/application-stacks/runtime-component-operator/common"
	prometheusv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	routev1 "github.com/openshift/api/route/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Defines the desired state of WebSphereLibertyApplication.
type WebSphereLibertyApplicationSpec struct {
	// The name of the application this resource is part of. If not specified, it defaults to the name of the CR.
	// +operator-sdk:csv:customresourcedefinitions:order=1,type=spec,displayName="Application Name",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	ApplicationName string `json:"applicationName,omitempty"`

	// Application image to be installed.
	// +operator-sdk:csv:customresourcedefinitions:order=2,type=spec,displayName="Application Image",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	ApplicationImage string `json:"applicationImage"`

	// Version of the application.
	// +operator-sdk:csv:customresourcedefinitions:order=3,type=spec,displayName="Application Version",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	ApplicationVersion string `json:"applicationVersion,omitempty"`

	// Policy for pulling container images. Defaults to IfNotPresent.
	// +operator-sdk:csv:customresourcedefinitions:order=4,type=spec,displayName="Pull Policy",xDescriptors="urn:alm:descriptor:com.tectonic.ui:imagePullPolicy"
	PullPolicy *corev1.PullPolicy `json:"pullPolicy,omitempty"`

	// Number of pods to create.
	// +operator-sdk:csv:customresourcedefinitions:order=5,type=spec,displayName="Replicas",xDescriptors="urn:alm:descriptor:com.tectonic.ui:podCount"
	Replicas *int32 `json:"replicas,omitempty"`

	// A boolean that toggles the external exposure of this deployment via a Route or a Knative Route resource.
	// +operator-sdk:csv:customresourcedefinitions:order=6,type=spec,displayName="Expose",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	Expose *bool `json:"expose,omitempty"`

	// Limits the amount of required resources.
	// +operator-sdk:csv:customresourcedefinitions:order=7,type=spec,displayName="Resource Requirements",xDescriptors="urn:alm:descriptor:com.tectonic.ui:resourceRequirements"
	ResourceConstraints *corev1.ResourceRequirements `json:"resourceConstraints,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:order=8,type=spec,displayName="Service"
	Service *WebSphereLibertyApplicationService `json:"service,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:order=16,type=spec,displayName="Auto Scaling"
	Autoscaling *WebSphereLibertyApplicationAutoScaling `json:"autoscaling,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:order=20,type=spec,displayName="Deployment"
	Deployment *WebSphereLibertyApplicationDeployment `json:"deployment,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:order=22,type=spec,displayName="StatefulSet"
	StatefulSet *WebSphereLibertyApplicationStatefulSet `json:"statefulSet,omitempty"`

	// The name of the OpenShift service account to be used during deployment.
	// +operator-sdk:csv:customresourcedefinitions:order=28,type=spec,displayName="Service Account Name",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	ServiceAccountName *string `json:"serviceAccountName,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:order=33,type=spec,displayName="Monitoring"
	Monitoring *WebSphereLibertyApplicationMonitoring `json:"monitoring,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:order=36,type=spec,displayName="Affinity"
	Affinity *WebSphereLibertyApplicationAffinity `json:"affinity,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:order=41,type=spec,displayName="Route"
	Route *WebSphereLibertyApplicationRoute `json:"route,omitempty"`

	// A boolean to toggle the creation of Knative resources and usage of Knative serving.
	// +operator-sdk:csv:customresourcedefinitions:order=52,type=spec,displayName="Create Knative Service",xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	CreateKnativeService *bool `json:"createKnativeService,omitempty"`

	// Detects if the services need to be restarted.
	// +operator-sdk:csv:customresourcedefinitions:order=53,type=spec,displayName="Liveness Probe"
	LivenessProbe *corev1.Probe `json:"livenessProbe,omitempty"`

	// Detects if the services are ready to serve.
	// +operator-sdk:csv:customresourcedefinitions:order=54,type=spec,displayName="Readiness Probe"
	ReadinessProbe *corev1.Probe `json:"readinessProbe,omitempty"`

	// Protects slow starting containers from livenessProbe.
	// +operator-sdk:csv:customresourcedefinitions:order=55,type=spec,displayName="Startup Probe"
	StartupProbe *corev1.Probe `json:"startupProbe,omitempty"`

	// Name of the Secret to use to pull images from the specified repository. It is not required if the cluster is configured with a global image pull secret.
	// +operator-sdk:csv:customresourcedefinitions:order=56,type=spec,displayName="Pull Secret",xDescriptors="urn:alm:descriptor:io.kubernetes:Secret"
	PullSecret *string `json:"pullSecret,omitempty"`

	// Represents a pod volume with data that is accessible to the containers.
	// +listType=map
	// +listMapKey=name
	// +operator-sdk:csv:customresourcedefinitions:order=57,type=spec,displayName="Volume"
	Volumes []corev1.Volume `json:"volumes,omitempty"`

	// Represents where to mount the volumes into containers.
	// +listType=atomic
	// +operator-sdk:csv:customresourcedefinitions:order=58,type=spec,displayName="Volume Mounts"
	VolumeMounts []corev1.VolumeMount `json:"volumeMounts,omitempty"`

	// An array of environment variables following the format of {name, value}, where value is a simple string.
	// +listType=map
	// +listMapKey=name
	// +operator-sdk:csv:customresourcedefinitions:order=59,type=spec,displayName="Env Var"
	Env []corev1.EnvVar `json:"env,omitempty"`

	// An array of references to ConfigMap or Secret resources containing environment variables.
	// +listType=atomic
	// +operator-sdk:csv:customresourcedefinitions:order=60,type=spec,displayName="Env From"
	EnvFrom []corev1.EnvFromSource `json:"envFrom,omitempty"`

	// List of containers that run before other containers in a pod.
	// +listType=map
	// +listMapKey=name
	// +operator-sdk:csv:customresourcedefinitions:order=62,type=spec,displayName="Init Containers"
	InitContainers []corev1.Container `json:"initContainers,omitempty"`

	// The list of sidecar containers. These are additional containers to be added to the pods.
	// +listType=map
	// +listMapKey=name
	// +operator-sdk:csv:customresourcedefinitions:order=63,type=spec,displayName="Sidecar Containers"
	SidecarContainers []corev1.Container `json:"sidecarContainers,omitempty"`

	// WebSphere Liberty specific capabilities

	// +operator-sdk:csv:customresourcedefinitions:order=29,type=spec,displayName="Serviceability"
	Serviceability *WebSphereLibertyApplicationServiceability `json:"serviceability,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:order=30,type=spec,displayName="Single sign-on"
	SSO *WebSphereLibertyApplicationSSO `json:"sso,omitempty"`
}

// Configures a Pod to run on particular Nodes.
type WebSphereLibertyApplicationAffinity struct {
	// Controls which nodes the pod are scheduled to run on, based on labels on the node.
	// +operator-sdk:csv:customresourcedefinitions:order=37,type=spec,displayName="Node Affinity",xDescriptors="urn:alm:descriptor:com.tectonic.ui:nodeAffinity"
	NodeAffinity *corev1.NodeAffinity `json:"nodeAffinity,omitempty"`

	// Controls the nodes the pod are scheduled to run on, based on labels on the pods that are already running on the node.
	// +operator-sdk:csv:customresourcedefinitions:order=38,type=spec,displayName="Pod Affinity",xDescriptors="urn:alm:descriptor:com.tectonic.ui:podAffinity"
	PodAffinity *corev1.PodAffinity `json:"podAffinity,omitempty"`

	// Enables the ability to prevent running a pod on the same node as another pod.
	// +operator-sdk:csv:customresourcedefinitions:order=39,type=spec,displayName="Pod Anti Affinity",xDescriptors="urn:alm:descriptor:com.tectonic.ui:podAntiAffinity"
	PodAntiAffinity *corev1.PodAntiAffinity `json:"podAntiAffinity,omitempty"`

	// A YAML object that contains a set of required labels and their values.
	// +operator-sdk:csv:customresourcedefinitions:order=40,type=spec,displayName="Node Affinity Labels",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	NodeAffinityLabels map[string]string `json:"nodeAffinityLabels,omitempty"`

	// An array of architectures to be considered for deployment. Their position in the array indicates preference.
	// +listType=set
	Architecture []string `json:"architecture,omitempty"`
}

// Configures the desired resource consumption of pods.
type WebSphereLibertyApplicationAutoScaling struct {
	// Required field for autoscaling. Upper limit for the number of pods that can be set by the autoscaler. Parameter spec.resourceConstraints.requests.cpu must also be specified.
	// +kubebuilder:validation:Minimum=1
	// +operator-sdk:csv:customresourcedefinitions:order=17,type=spec,displayName="Max Replicas",xDescriptors="urn:alm:descriptor:com.tectonic.ui:number"
	MaxReplicas int32 `json:"maxReplicas,omitempty"`

	// Lower limit for the number of pods that can be set by the autoscaler.
	// +operator-sdk:csv:customresourcedefinitions:order=18,type=spec,displayName="Min Replicas",xDescriptors="urn:alm:descriptor:com.tectonic.ui:number"
	MinReplicas *int32 `json:"minReplicas,omitempty"`

	// Target average CPU utilization (represented as a percentage of requested CPU) over all the pods.
	// +operator-sdk:csv:customresourcedefinitions:order=19,type=spec,displayName="Target CPU Utilization Percentage",xDescriptors="urn:alm:descriptor:com.tectonic.ui:number"
	TargetCPUUtilizationPercentage *int32 `json:"targetCPUUtilizationPercentage,omitempty"`
}

// Configures parameters for the network service of pods.
type WebSphereLibertyApplicationService struct {
	// The port exposed by the container.
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:validation:Minimum=1
	// +operator-sdk:csv:customresourcedefinitions:order=9,type=spec,displayName="Service Port",xDescriptors="urn:alm:descriptor:com.tectonic.ui:number"
	Port int32 `json:"port,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:order=10,type=spec,displayName="Service Type",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	Type *corev1.ServiceType `json:"type,omitempty"`

	// Node proxies this port into your service.
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:validation:Minimum=0
	// +operator-sdk:csv:customresourcedefinitions:order=11,type=spec,displayName="Node Port",xDescriptors="urn:alm:descriptor:com.tectonic.ui:number"
	NodePort *int32 `json:"nodePort,omitempty"`

	// The name for the port exposed by the container.
	// +operator-sdk:csv:customresourcedefinitions:order=12,type=spec,displayName="Port Name",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	PortName string `json:"portName,omitempty"`

	// Annotations to be added to the service.
	// +operator-sdk:csv:customresourcedefinitions:order=13,type=spec,displayName="Service Annotations",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	Annotations map[string]string `json:"annotations,omitempty"`

	// The port that the operator assigns to containers inside pods. Defaults to the value of spec.service.port.
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:validation:Minimum=1
	// +operator-sdk:csv:customresourcedefinitions:order=14,type=spec,displayName="Target Port",xDescriptors="urn:alm:descriptor:com.tectonic.ui:number"
	TargetPort *int32 `json:"targetPort,omitempty"`

	// A name of a secret that already contains TLS key, certificate and CA to be mounted in the pod.
	// +operator-sdk:csv:customresourcedefinitions:order=15,type=spec,displayName="Certificate Secret Reference",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	CertificateSecretRef *string `json:"certificateSecretRef,omitempty"`

	// An array consisting of service ports.
	Ports []corev1.ServicePort `json:"ports,omitempty"`

	// Expose the application as a bindable service. Defaults to false.
	Bindable *bool `json:"bindable,omitempty"`
}

// Defines the desired state and cycle of applications.
type WebSphereLibertyApplicationDeployment struct {
	// Specifies the strategy to replace old deployment pods with new pods.
	// +operator-sdk:csv:customresourcedefinitions:order=21,type=spec,displayName="Deployment Update Strategy",xDescriptors="urn:alm:descriptor:com.tectonic.ui:updateStrategy"
	UpdateStrategy *appsv1.DeploymentStrategy `json:"updateStrategy,omitempty"`

	// Annotations to be added only to the Deployment and resources owned by the Deployment
	Annotations map[string]string `json:"annotations,omitempty"`
}

// Defines the desired state and cycle of stateful applications.
type WebSphereLibertyApplicationStatefulSet struct {
	// Specifies the strategy to replace old StatefulSet pods with new pods.
	// +operator-sdk:csv:customresourcedefinitions:order=23,type=spec,displayName="StatefulSet Update Strategy",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	UpdateStrategy *appsv1.StatefulSetUpdateStrategy `json:"updateStrategy,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:order=24,type=spec,displayName="Storage"
	Storage *WebSphereLibertyApplicationStorage `json:"storage,omitempty"`

	// Annotations to be added only to the StatefulSet and resources owned by the StatefulSet.
	Annotations map[string]string `json:"annotations,omitempty"`
}

// Defines settings of persisted storage for StatefulSets.
type WebSphereLibertyApplicationStorage struct {
	// A convenient field to set the size of the persisted storage.
	// +kubebuilder:validation:Pattern=^([+-]?[0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$
	// +operator-sdk:csv:customresourcedefinitions:order=25,type=spec,displayName="Storage Size",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	Size string `json:"size,omitempty"`

	// The directory inside the container where this persisted storage will be bound to.
	// +operator-sdk:csv:customresourcedefinitions:order=26,type=spec,displayName="Storage Mount Path",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	MountPath string `json:"mountPath,omitempty"`

	// A YAML object that represents a volumeClaimTemplate component of a StatefulSet.
	// +operator-sdk:csv:customresourcedefinitions:order=27,type=spec,displayName="Storage Volume Claim Template",xDescriptors="urn:alm:descriptor:com.tectonic.ui:PersistentVolumeClaim"
	VolumeClaimTemplate *corev1.PersistentVolumeClaim `json:"volumeClaimTemplate,omitempty"`
}

// Specifies parameters for Service Monitor.
type WebSphereLibertyApplicationMonitoring struct {
	// Labels to set on ServiceMonitor.
	// +operator-sdk:csv:customresourcedefinitions:order=34,type=spec,displayName="Monitoring Labels",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	Labels map[string]string `json:"labels,omitempty"`

	// A YAML snippet representing an array of Endpoint component from ServiceMonitor.
	// +listType=atomic
	// +operator-sdk:csv:customresourcedefinitions:order=35,type=spec,displayName="Monitoring Endpoints",xDescriptors="urn:alm:descriptor:com.tectonic.ui:endpointList"
	Endpoints []prometheusv1.Endpoint `json:"endpoints,omitempty"`
}

// Specifies serviceability-related operations, such as gathering server memory dumps and server traces.
type WebSphereLibertyApplicationServiceability struct {
	// A convenient field to request the size of the persisted storage to use for serviceability.
	// +kubebuilder:validation:Pattern=^([+-]?[0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$
	Size string `json:"size,omitempty"`

	// The name of the PersistentVolumeClaim resource you created to be used for serviceability.
	// +kubebuilder:validation:Pattern=.+
	VolumeClaimName string `json:"volumeClaimName,omitempty"`

	// A convenient field to request the StorageClassName of the persisted storage to use for serviceability.
	// +kubebuilder:validation:Pattern=.+
	StorageClassName string `json:"storageClassName,omitempty"`
}

// Configures the ingress resource.
type WebSphereLibertyApplicationRoute struct {

	// Annotations to be added to the Route.
	// +operator-sdk:csv:customresourcedefinitions:order=42,type=spec,displayName="Route Annotations",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	Annotations map[string]string `json:"annotations,omitempty"`

	// Hostname to be used for the Route.
	// +operator-sdk:csv:customresourcedefinitions:order=43,type=spec,displayName="Route Host",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	Host string `json:"host,omitempty"`

	// Path to be used for Route.
	// +operator-sdk:csv:customresourcedefinitions:order=44,type=spec,displayName="Route Path",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	Path string `json:"path,omitempty"`

	// A name of a secret that already contains TLS key, certificate and CA to be used in the route. Also can contain destination CA certificate.
	// +operator-sdk:csv:customresourcedefinitions:order=45,type=spec,displayName="Certificate Secret Reference",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	CertificateSecretRef *string `json:"certificateSecretRef,omitempty"`

	// TLS termination policy. Can be one of edge, reencrypt and passthrough.
	// +operator-sdk:csv:customresourcedefinitions:order=46,type=spec,displayName="Termination",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	Termination *routev1.TLSTerminationType `json:"termination,omitempty"`

	// HTTP traffic policy with TLS enabled. Can be one of Allow, Redirect and None.
	// +operator-sdk:csv:customresourcedefinitions:order=47,type=spec,displayName="Insecure Edge Termination Policy",xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	InsecureEdgeTerminationPolicy *routev1.InsecureEdgeTerminationPolicyType `json:"insecureEdgeTerminationPolicy,omitempty"`
}

// Defines the observed state of WebSphereLibertyApplication.
type WebSphereLibertyApplicationStatus struct {
	// +listType=atomic
	// +operator-sdk:csv:customresourcedefinitions:type=status,displayName="Status Conditions",xDescriptors="urn:alm:descriptor:io.kubernetes.conditions"
	Conditions     []StatusCondition `json:"conditions,omitempty"`
	RouteAvailable *bool             `json:"routeAvailable,omitempty"`
	ImageReference string            `json:"imageReference,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:type=status,displayName="Service Binding Secret",xDescriptors="urn:alm:descriptor:io.kubernetes:Secret"
	Binding *corev1.LocalObjectReference `json:"binding,omitempty"`
}

// Defines possible status conditions.
type StatusCondition struct {
	LastTransitionTime *metav1.Time           `json:"lastTransitionTime,omitempty"`
	Reason             string                 `json:"reason,omitempty"`
	Message            string                 `json:"message,omitempty"`
	Status             corev1.ConditionStatus `json:"status,omitempty"`
	Type               StatusConditionType    `json:"type,omitempty"`
}

// Defines the type of status condition.
type StatusConditionType string

const (
	// StatusConditionTypeReconciled ...
	StatusConditionTypeReconciled StatusConditionType = "Reconciled"
)

// +kubebuilder:resource:path=webspherelibertyapplications,scope=Namespaced,shortName=wlapp;wlapps
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Image",type="string",JSONPath=".spec.applicationImage",priority=0,description="Absolute name of the deployed image containing registry and tag"
// +kubebuilder:printcolumn:name="Exposed",type="boolean",JSONPath=".spec.expose",priority=0,description="Specifies whether deployment is exposed externally via default Route"
// +kubebuilder:printcolumn:name="Reconciled",type="string",JSONPath=".status.conditions[?(@.type=='Reconciled')].status",priority=0,description="Status of the reconcile condition"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Reconciled')].reason",priority=1,description="Reason for the failure of reconcile condition"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Reconciled')].message",priority=1,description="Failure message from reconcile condition"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",priority=0,description="Age of the resource"
//+operator-sdk:csv:customresourcedefinitions:displayName="WebSphereLibertyApplication",resources={{Deployment,v1},{Service,v1},{StatefulSet,v1},{Route,v1},{HorizontalPodAutoscaler,v1},{ServiceAccount,v1},{Secret,v1}}

// Represents the deployment of an WebSphere Liberty application
type WebSphereLibertyApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WebSphereLibertyApplicationSpec   `json:"spec,omitempty"`
	Status WebSphereLibertyApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebSphereLibertyApplicationList contains a list of WebSphereLibertyApplication
type WebSphereLibertyApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebSphereLibertyApplication `json:"items"`
}

// Specifies the configuration for Single sign-on (SSO) providers to authenticate with.
type WebSphereLibertyApplicationSSO struct {
	// +listType=atomic
	OIDC []OidcClient `json:"oidc,omitempty"`

	// +listType=atomic
	Oauth2 []OAuth2Client `json:"oauth2,omitempty"`

	Github *GithubLogin `json:"github,omitempty"`

	// Common parameters for all SSO providers

	// Specifies a callback protocol, host and port number.
	RedirectToRPHostAndPort string `json:"redirectToRPHostAndPort,omitempty"`

	// Specifies whether to map a user identifier to a registry user. This parameter applies to all providers.
	MapToUserRegistry *bool `json:"mapToUserRegistry,omitempty"`
}

// Represents configuration for an OpenID Connect (OIDC) client.
type OidcClient struct {
	// The unique ID for the provider. Default value is oidc.
	ID string `json:"id,omitempty"`

	// Specifies a discovery endpoint URL for the OpenID Connect provider. Required field.
	DiscoveryEndpoint string `json:"discoveryEndpoint"`

	// Specifies the name of the claim. Use its value as the user group membership.
	GroupNameAttribute string `json:"groupNameAttribute,omitempty"`

	// Specifies the name of the claim. Use its value as the authenticated user principal.
	UserNameAttribute string `json:"userNameAttribute,omitempty"`

	// The name of the social login configuration for display.
	DisplayName string `json:"displayName,omitempty"`

	// Specifies whether the UserInfo endpoint is contacted.
	UserInfoEndpointEnabled *bool `json:"userInfoEndpointEnabled,omitempty"`

	// Specifies the name of the claim. Use its value as the subject realm.
	RealmNameAttribute string `json:"realmNameAttribute,omitempty"`

	// Specifies one or more scopes to request.
	Scope string `json:"scope,omitempty"`

	// Specifies the required authentication method.
	TokenEndpointAuthMethod string `json:"tokenEndpointAuthMethod,omitempty"`

	// Specifies whether to enable host name verification when the client contacts the provider.
	HostNameVerificationEnabled *bool `json:"hostNameVerificationEnabled,omitempty"`
}

// Represents configuration for an OAuth2 client.
type OAuth2Client struct {
	// Specifies the unique ID for the provider. The default value is oauth2.
	ID string `json:"id,omitempty"`

	// Specifies a token endpoint URL for the OAuth 2.0 provider. Required field.
	TokenEndpoint string `json:"tokenEndpoint"`

	// Specifies an authorization endpoint URL for the OAuth 2.0 provider. Required field.
	AuthorizationEndpoint string `json:"authorizationEndpoint"`

	// Specifies the name of the claim. Use its value as the user group membership
	GroupNameAttribute string `json:"groupNameAttribute,omitempty"`

	// Specifies the name of the claim. Use its value as the authenticated user principal.
	UserNameAttribute string `json:"userNameAttribute,omitempty"`

	// The name of the social login configuration for display.
	DisplayName string `json:"displayName,omitempty"`

	// Specifies the name of the claim. Use its value as the subject realm.
	RealmNameAttribute string `json:"realmNameAttribute,omitempty"`

	// Specifies the realm name for this social media.
	RealmName string `json:"realmName,omitempty"`

	// Specifies one or more scopes to request.
	Scope string `json:"scope,omitempty"`

	// Specifies the required authentication method.
	TokenEndpointAuthMethod string `json:"tokenEndpointAuthMethod,omitempty"`

	// Name of the header to use when an OAuth access token is forwarded.
	AccessTokenHeaderName string `json:"accessTokenHeaderName,omitempty"`

	// Determines whether the access token that is provided in the request is used for authentication.
	AccessTokenRequired *bool `json:"accessTokenRequired,omitempty"`

	// Determines whether to support access token authentication if an access token is provided in the request.
	AccessTokenSupported *bool `json:"accessTokenSupported,omitempty"`

	// Indicates which specification to use for the user API.
	UserApiType string `json:"userApiType,omitempty"`

	// The URL for retrieving the user information.
	UserApi string `json:"userApi,omitempty"`
}

// Represents configuration for social login using GitHub.
type GithubLogin struct {
	// Specifies the host name of your enterprise GitHub.
	Hostname string `json:"hostname,omitempty"`
}

func init() {
	SchemeBuilder.Register(&WebSphereLibertyApplication{}, &WebSphereLibertyApplicationList{})
}

// GetApplicationImage returns application image
func (cr *WebSphereLibertyApplication) GetApplicationImage() string {
	return cr.Spec.ApplicationImage
}

// GetPullPolicy returns image pull policy
func (cr *WebSphereLibertyApplication) GetPullPolicy() *corev1.PullPolicy {
	return cr.Spec.PullPolicy
}

// GetPullSecret returns secret name for docker registry credentials
func (cr *WebSphereLibertyApplication) GetPullSecret() *string {
	return cr.Spec.PullSecret
}

// GetServiceAccountName returns service account name
func (cr *WebSphereLibertyApplication) GetServiceAccountName() *string {
	return cr.Spec.ServiceAccountName
}

// GetReplicas returns number of replicas
func (cr *WebSphereLibertyApplication) GetReplicas() *int32 {
	return cr.Spec.Replicas
}

// GetLivenessProbe returns liveness probe
func (cr *WebSphereLibertyApplication) GetLivenessProbe() *corev1.Probe {
	return cr.Spec.LivenessProbe
}

// GetReadinessProbe returns readiness probe
func (cr *WebSphereLibertyApplication) GetReadinessProbe() *corev1.Probe {
	return cr.Spec.ReadinessProbe
}

// GetStartupProbe returns startup probe
func (cr *WebSphereLibertyApplication) GetStartupProbe() *corev1.Probe {
	return cr.Spec.StartupProbe
}

// GetVolumes returns volumes slice
func (cr *WebSphereLibertyApplication) GetVolumes() []corev1.Volume {
	return cr.Spec.Volumes
}

// GetVolumeMounts returns volume mounts slice
func (cr *WebSphereLibertyApplication) GetVolumeMounts() []corev1.VolumeMount {
	return cr.Spec.VolumeMounts
}

// GetResourceConstraints returns resource constraints
func (cr *WebSphereLibertyApplication) GetResourceConstraints() *corev1.ResourceRequirements {
	return cr.Spec.ResourceConstraints
}

// GetExpose returns expose flag
func (cr *WebSphereLibertyApplication) GetExpose() *bool {
	return cr.Spec.Expose
}

// GetEnv returns slice of environment variables
func (cr *WebSphereLibertyApplication) GetEnv() []corev1.EnvVar {
	return cr.Spec.Env
}

// GetEnvFrom returns slice of environment variables from source
func (cr *WebSphereLibertyApplication) GetEnvFrom() []corev1.EnvFromSource {
	return cr.Spec.EnvFrom
}

// GetCreateKnativeService returns flag that toggles Knative service
func (cr *WebSphereLibertyApplication) GetCreateKnativeService() *bool {
	return cr.Spec.CreateKnativeService
}

// GetAutoscaling returns autoscaling settings
func (cr *WebSphereLibertyApplication) GetAutoscaling() common.BaseComponentAutoscaling {
	if cr.Spec.Autoscaling == nil {
		return nil
	}
	return cr.Spec.Autoscaling
}

// GetStorage returns storage settings
func (ss *WebSphereLibertyApplicationStatefulSet) GetStorage() common.BaseComponentStorage {
	if ss.Storage == nil {
		return nil
	}
	return ss.Storage
}

// GetService returns service settings
func (cr *WebSphereLibertyApplication) GetService() common.BaseComponentService {
	if cr.Spec.Service == nil {
		return nil
	}
	return cr.Spec.Service
}

// GetApplicationVersion returns application version
func (cr *WebSphereLibertyApplication) GetApplicationVersion() string {
	return cr.Spec.ApplicationVersion
}

// GetApplicationName returns Application name
func (cr *WebSphereLibertyApplication) GetApplicationName() string {
	return cr.Spec.ApplicationName
}

// GetMonitoring returns monitoring settings
func (cr *WebSphereLibertyApplication) GetMonitoring() common.BaseComponentMonitoring {
	if cr.Spec.Monitoring == nil {
		return nil
	}
	return cr.Spec.Monitoring
}

// GetStatus returns WebSphereLibertyApplication status
func (cr *WebSphereLibertyApplication) GetStatus() common.BaseComponentStatus {
	return &cr.Status
}

// GetInitContainers returns list of init containers
func (cr *WebSphereLibertyApplication) GetInitContainers() []corev1.Container {
	return cr.Spec.InitContainers
}

// GetSidecarContainers returns list of sidecar containers
func (cr *WebSphereLibertyApplication) GetSidecarContainers() []corev1.Container {
	return cr.Spec.SidecarContainers
}

// GetGroupName returns group name to be used in labels and annotation
func (cr *WebSphereLibertyApplication) GetGroupName() string {
	return "apps.WebSphereLiberty.io"
}

// GetRoute returns route
func (cr *WebSphereLibertyApplication) GetRoute() common.BaseComponentRoute {
	if cr.Spec.Route == nil {
		return nil
	}
	return cr.Spec.Route
}

// GetAffinity returns deployment's node and pod affinity settings
func (cr *WebSphereLibertyApplication) GetAffinity() common.BaseComponentAffinity {
	if cr.Spec.Affinity == nil {
		return nil
	}
	return cr.Spec.Affinity
}

// GetDeployment returns deployment settings
func (cr *WebSphereLibertyApplication) GetDeployment() common.BaseComponentDeployment {
	if cr.Spec.Deployment == nil {
		return nil
	}
	return cr.Spec.Deployment
}

// GetDeploymentStrategy returns deployment strategy struct
func (cr *WebSphereLibertyApplicationDeployment) GetDeploymentUpdateStrategy() *appsv1.DeploymentStrategy {
	return cr.UpdateStrategy
}

// GetAnnotations returns annotations to be added only to the Deployment and its child resources
func (rcd *WebSphereLibertyApplicationDeployment) GetAnnotations() map[string]string {
	return rcd.Annotations
}

// GetStatefulSet returns statefulSet settings
func (cr *WebSphereLibertyApplication) GetStatefulSet() common.BaseComponentStatefulSet {
	if cr.Spec.StatefulSet == nil {
		return nil
	}
	return cr.Spec.StatefulSet
}

// GetStatefulSetUpdateStrategy returns statefulSet strategy struct
func (cr *WebSphereLibertyApplicationStatefulSet) GetStatefulSetUpdateStrategy() *appsv1.StatefulSetUpdateStrategy {
	return cr.UpdateStrategy
}

// GetAnnotations returns annotations to be added only to the StatefulSet and its child resources
func (rcss *WebSphereLibertyApplicationStatefulSet) GetAnnotations() map[string]string {
	return rcss.Annotations
}

// GetImageReference returns Docker image reference to be deployed by the CR
func (s *WebSphereLibertyApplicationStatus) GetImageReference() string {
	return s.ImageReference
}

// SetImageReference sets Docker image reference on the status portion of the CR
func (s *WebSphereLibertyApplicationStatus) SetImageReference(imageReference string) {
	s.ImageReference = imageReference
}

// GetBinding returns BindingStatus representing binding status
func (s *WebSphereLibertyApplicationStatus) GetBinding() *corev1.LocalObjectReference {
	return s.Binding
}

// SetBinding sets BindingStatus representing binding status
func (s *WebSphereLibertyApplicationStatus) SetBinding(r *corev1.LocalObjectReference) {
	s.Binding = r
}

// GetMinReplicas returns minimum replicas
func (a *WebSphereLibertyApplicationAutoScaling) GetMinReplicas() *int32 {
	return a.MinReplicas
}

// GetMaxReplicas returns maximum replicas
func (a *WebSphereLibertyApplicationAutoScaling) GetMaxReplicas() int32 {
	return a.MaxReplicas
}

// GetTargetCPUUtilizationPercentage returns target cpu usage
func (a *WebSphereLibertyApplicationAutoScaling) GetTargetCPUUtilizationPercentage() *int32 {
	return a.TargetCPUUtilizationPercentage
}

// GetSize returns pesistent volume size
func (s *WebSphereLibertyApplicationStorage) GetSize() string {
	return s.Size
}

// GetMountPath returns mount path for persistent volume
func (s *WebSphereLibertyApplicationStorage) GetMountPath() string {
	return s.MountPath
}

// GetVolumeClaimTemplate returns a template representing requested persistent volume
func (s *WebSphereLibertyApplicationStorage) GetVolumeClaimTemplate() *corev1.PersistentVolumeClaim {
	return s.VolumeClaimTemplate
}

// GetAnnotations returns a set of annotations to be added to the service
func (s *WebSphereLibertyApplicationService) GetAnnotations() map[string]string {
	return s.Annotations
}

// GetServiceability returns serviceability
func (cr *WebSphereLibertyApplication) GetServiceability() *WebSphereLibertyApplicationServiceability {
	return cr.Spec.Serviceability
}

// GetSize returns pesistent volume size for Serviceability
func (s *WebSphereLibertyApplicationServiceability) GetSize() string {
	return s.Size
}

// GetVolumeClaimName returns the name of custom PersistentVolumeClaim (PVC) for Serviceability. Must be in the same namespace as the WebSphereLibertyApplication.
func (s *WebSphereLibertyApplicationServiceability) GetVolumeClaimName() string {
	return s.VolumeClaimName
}

// GetPort returns service port
func (s *WebSphereLibertyApplicationService) GetPort() int32 {
	if s != nil && s.Port != 0 {
		return s.Port
	}
	return 9080
}

// GetNodePort returns service nodePort
func (s *WebSphereLibertyApplicationService) GetNodePort() *int32 {
	if s.NodePort == nil {
		return nil
	}
	return s.NodePort
}

// GetTargetPort returns the internal target port for containers
func (s *WebSphereLibertyApplicationService) GetTargetPort() *int32 {
	return s.TargetPort
}

// GetPortName returns name of service port
func (s *WebSphereLibertyApplicationService) GetPortName() string {
	return s.PortName
}

// GetType returns service type
func (s *WebSphereLibertyApplicationService) GetType() *corev1.ServiceType {
	return s.Type
}

// GetPorts returns a list of service ports
func (s *WebSphereLibertyApplicationService) GetPorts() []corev1.ServicePort {
	return s.Ports
}

// GetCertificateSecretRef returns a secret reference with a certificate
func (s *WebSphereLibertyApplicationService) GetCertificateSecretRef() *string {
	return s.CertificateSecretRef
}

// GetBindable returns whether the application should be exposable as a service
func (s *WebSphereLibertyApplicationService) GetBindable() *bool {
	return s.Bindable
}

// GetLabels returns labels to be added on ServiceMonitor
func (m *WebSphereLibertyApplicationMonitoring) GetLabels() map[string]string {
	return m.Labels
}

// GetEndpoints returns endpoints to be added to ServiceMonitor
func (m *WebSphereLibertyApplicationMonitoring) GetEndpoints() []prometheusv1.Endpoint {
	return m.Endpoints
}

// GetAnnotations returns route annotations
func (r *WebSphereLibertyApplicationRoute) GetAnnotations() map[string]string {
	return r.Annotations
}

// GetCertificateSecretRef returns a secret reference with a certificate
func (r *WebSphereLibertyApplicationRoute) GetCertificateSecretRef() *string {
	return r.CertificateSecretRef
}

// GetTermination returns terminatation of the route's TLS
func (r *WebSphereLibertyApplicationRoute) GetTermination() *routev1.TLSTerminationType {
	return r.Termination
}

// GetInsecureEdgeTerminationPolicy returns terminatation of the route's TLS
func (r *WebSphereLibertyApplicationRoute) GetInsecureEdgeTerminationPolicy() *routev1.InsecureEdgeTerminationPolicyType {
	return r.InsecureEdgeTerminationPolicy
}

// GetHost returns hostname to be used by the route
func (r *WebSphereLibertyApplicationRoute) GetHost() string {
	return r.Host
}

// GetPath returns path to use for the route
func (r *WebSphereLibertyApplicationRoute) GetPath() string {
	return r.Path
}

// GetNodeAffinity returns node affinity
func (a *WebSphereLibertyApplicationAffinity) GetNodeAffinity() *corev1.NodeAffinity {
	return a.NodeAffinity
}

// GetPodAffinity returns pod affinity
func (a *WebSphereLibertyApplicationAffinity) GetPodAffinity() *corev1.PodAffinity {
	return a.PodAffinity
}

// GetPodAntiAffinity returns pod anti-affinity
func (a *WebSphereLibertyApplicationAffinity) GetPodAntiAffinity() *corev1.PodAntiAffinity {
	return a.PodAntiAffinity
}

// GetArchitecture returns list of architecture names
func (a *WebSphereLibertyApplicationAffinity) GetArchitecture() []string {
	return a.Architecture
}

// GetNodeAffinityLabels returns list of architecture names
func (a *WebSphereLibertyApplicationAffinity) GetNodeAffinityLabels() map[string]string {
	return a.NodeAffinityLabels
}

// Initialize sets default values
func (cr *WebSphereLibertyApplication) Initialize() {
	if cr.Spec.PullPolicy == nil {
		pp := corev1.PullIfNotPresent
		cr.Spec.PullPolicy = &pp
	}

	if cr.Spec.ResourceConstraints == nil {
		cr.Spec.ResourceConstraints = &corev1.ResourceRequirements{}
	}

	// Default applicationName to cr.Name, if a user sets createAppDefinition to true but doesn't set applicationName
	if cr.Spec.ApplicationName == "" {
		if cr.Labels != nil && cr.Labels["app.kubernetes.io/part-of"] != "" {
			cr.Spec.ApplicationName = cr.Labels["app.kubernetes.io/part-of"]
		} else {
			cr.Spec.ApplicationName = cr.Name
		}
	}

	if cr.Labels != nil {
		cr.Labels["app.kubernetes.io/part-of"] = cr.Spec.ApplicationName
	}

	// This is to handle when there is no service in the CR
	if cr.Spec.Service == nil {
		cr.Spec.Service = &WebSphereLibertyApplicationService{}
	}

	if cr.Spec.Service.Type == nil {
		st := corev1.ServiceTypeClusterIP
		cr.Spec.Service.Type = &st
	}

	if cr.Spec.Service.Port == 0 {
		cr.Spec.Service.Port = 9080
	}

}

// GetLabels returns set of labels to be added to all resources
func (cr *WebSphereLibertyApplication) GetLabels() map[string]string {
	labels := map[string]string{
		"app.kubernetes.io/instance":   cr.Name,
		"app.kubernetes.io/name":       cr.Name,
		"app.kubernetes.io/managed-by": "websphere-liberty-operator",
		"app.kubernetes.io/component":  "backend",
		"app.kubernetes.io/part-of":    cr.Spec.ApplicationName,
	}

	if cr.Spec.ApplicationVersion != "" {
		labels["app.kubernetes.io/version"] = cr.Spec.ApplicationVersion
	}

	for key, value := range cr.Labels {
		if key != "app.kubernetes.io/instance" {
			labels[key] = value
		}
	}

	return labels
}

// GetAnnotations returns set of annotations to be added to all resources
func (cr *WebSphereLibertyApplication) GetAnnotations() map[string]string {
	annotations := map[string]string{}
	for k, v := range cr.Annotations {
		annotations[k] = v
	}
	delete(annotations, "kubectl.kubernetes.io/last-applied-configuration")
	return annotations
}

// GetType returns status condition type
func (c *StatusCondition) GetType() common.StatusConditionType {
	return convertToCommonStatusConditionType(c.Type)
}

// SetType returns status condition type
func (c *StatusCondition) SetType(ct common.StatusConditionType) {
	c.Type = convertFromCommonStatusConditionType(ct)
}

// GetLastTransitionTime return time of last status change
func (c *StatusCondition) GetLastTransitionTime() *metav1.Time {
	return c.LastTransitionTime
}

// SetLastTransitionTime sets time of last status change
func (c *StatusCondition) SetLastTransitionTime(t *metav1.Time) {
	c.LastTransitionTime = t
}

// GetMessage return condition's message
func (c *StatusCondition) GetMessage() string {
	return c.Message
}

// SetMessage sets condition's message
func (c *StatusCondition) SetMessage(m string) {
	c.Message = m
}

// GetReason return condition's message
func (c *StatusCondition) GetReason() string {
	return c.Reason
}

// SetReason sets condition's reason
func (c *StatusCondition) SetReason(r string) {
	c.Reason = r
}

// GetStatus return condition's status
func (c *StatusCondition) GetStatus() corev1.ConditionStatus {
	return c.Status
}

// SetStatus sets condition's status
func (c *StatusCondition) SetStatus(s corev1.ConditionStatus) {
	c.Status = s
}

// NewCondition returns new condition
func (s *WebSphereLibertyApplicationStatus) NewCondition() common.StatusCondition {
	return &StatusCondition{}
}

// GetConditions returns slice of conditions
func (s *WebSphereLibertyApplicationStatus) GetConditions() []common.StatusCondition {
	var conditions = make([]common.StatusCondition, len(s.Conditions))
	for i := range s.Conditions {
		conditions[i] = &s.Conditions[i]
	}
	return conditions
}

// GetCondition ...
func (s *WebSphereLibertyApplicationStatus) GetCondition(t common.StatusConditionType) common.StatusCondition {
	for i := range s.Conditions {
		if s.Conditions[i].GetType() == t {
			return &s.Conditions[i]
		}
	}
	return nil
}

// SetCondition ...
func (s *WebSphereLibertyApplicationStatus) SetCondition(c common.StatusCondition) {
	condition := &StatusCondition{}
	found := false
	for i := range s.Conditions {
		if s.Conditions[i].GetType() == c.GetType() {
			condition = &s.Conditions[i]
			found = true
		}
	}

	if condition.GetStatus() != c.GetStatus() {
		condition.SetLastTransitionTime(&metav1.Time{Time: time.Now()})
	}

	condition.SetReason(c.GetReason())
	condition.SetMessage(c.GetMessage())
	condition.SetStatus(c.GetStatus())
	condition.SetType(c.GetType())
	if !found {
		s.Conditions = append(s.Conditions, *condition)
	}
}

func convertToCommonStatusConditionType(c StatusConditionType) common.StatusConditionType {
	switch c {
	case StatusConditionTypeReconciled:
		return common.StatusConditionTypeReconciled
	default:
		panic(c)
	}
}

func convertFromCommonStatusConditionType(c common.StatusConditionType) StatusConditionType {
	switch c {
	case common.StatusConditionTypeReconciled:
		return StatusConditionTypeReconciled
	default:
		panic(c)
	}
}
