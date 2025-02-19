package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// WebSphereLibertyDumpSpec defines the desired state of WebSphereLibertyDump
type WebSphereLibertyDumpSpec struct {
	// The name of the Pod, which must be in the same namespace as the WebSphereLibertyDump CR.
	PodName string `json:"podName"`
	// Optional. List of memory dump types to request: thread, heap, system.
	// +listType=set
	Include []WebSphereLibertyDumpInclude `json:"include,omitempty"`
}

// Defines the possible values for dump types
// +kubebuilder:validation:Enum=thread;heap;system
type WebSphereLibertyDumpInclude string

const (
	//WebSphereLibertyDumpIncludeHeap heap dump
	WebSphereLibertyDumpIncludeHeap WebSphereLibertyDumpInclude = "heap"
	//WebSphereLibertyDumpIncludeThread thread dump
	WebSphereLibertyDumpIncludeThread WebSphereLibertyDumpInclude = "thread"
	//WebSphereLibertyDumpIncludeSystem system (core) dump
	WebSphereLibertyDumpIncludeSystem WebSphereLibertyDumpInclude = "system"
)

// Defines the observed state of WebSphereLibertyDump
type WebSphereLibertyDumpStatus struct {
	// +listType=atomic
	Conditions []OperationStatusCondition `json:"conditions,omitempty"`
	DumpFile   string                     `json:"dumpFile,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=webspherelibertydumps,scope=Namespaced,shortName=wldump;wldumps
// +kubebuilder:printcolumn:name="Started",type="string",JSONPath=".status.conditions[?(@.type=='Started')].status",priority=0,description="Indicates if dump operation has started"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Started')].reason",priority=1,description="Reason for dump operation failing to start"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Started')].message",priority=1,description="Message for dump operation failing to start"
// +kubebuilder:printcolumn:name="Completed",type="string",JSONPath=".status.conditions[?(@.type=='Completed')].status",priority=0,description="Indicates if dump operation has completed"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Completed')].reason",priority=1,description="Reason for dump operation failing to complete"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Completed')].message",priority=1,description="Message for dump operation failing to complete"
// +kubebuilder:printcolumn:name="Dump file",type="string",JSONPath=".status.dumpFile",priority=0,description="Indicates filename of the server dump"
//+operator-sdk:csv:customresourcedefinitions:displayName="WebSphereLibertyDump"
// Day-2 operation for generating server dumps
type WebSphereLibertyDump struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WebSphereLibertyDumpSpec   `json:"spec,omitempty"`
	Status WebSphereLibertyDumpStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// WebSphereLibertyDumpList contains a list of WebSphereLibertyDump
type WebSphereLibertyDumpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebSphereLibertyDump `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WebSphereLibertyDump{}, &WebSphereLibertyDumpList{})
}
