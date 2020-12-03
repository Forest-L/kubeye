package validator

import (
	corev1 "k8s.io/api/core/v1"
	"kubeye/pkg/config"
	"time"
)

type AuditData struct {
	//AuditTime                    string                    `yaml:"auditTime" json:"auditTime,omitempty"`
	//AuditAddress                 string                    `yaml:"auditAddress" json:"auditAddress,omitempty"`
	//BasicClusterInformation     BasicClusterInformation `yaml:"basicClusterInformation" json:"basicClusterInformation,omitempty"`
	BasicComponentStatus        []BasicComponentStatus `yaml:"basicComponentStatus" json:"basicComponentStatus,omitempty"`
	ClusterCheckResults         []ClusterCheckResults  `yaml:"clusterCheckResults" json:"clusterCheckResults,omitempty"`
	ClusterConfigurationResults []PodResult            `yaml:"clusterConfigurationResults" json:"clusterConfigurationResults,omitempty"`
	AllNodeStatusResults        []AllNodeStatusResults `yaml:"allNodeStatusResults" json:"allNodeStatusResults,omitempty"`
}

type ClusterCheckResults struct {
	Namespace string    `yaml:"namespace" json:"namespace,omitempty"`
	Name      string    `yaml:"name" json:"name,omitempty"`
	EventTime time.Time `yaml:"eventTime" json:"eventTime,omitempty"`
	Reason    string    `yaml:"reason" json:"reason,omitempty"`
	Message   string    `yaml:"message" json:"message,omitempty"`
}

type BasicComponentStatus struct {
	Time     string          `yaml:"time" json:"time,omitempty"`
	Name     string          `yaml:"name" json:"name,omitempty"`
	Message  string          `yaml:"message" json:"message,omitempty"`
	Severity config.Severity `yaml:"severity" json:"severity,omitempty"`
}

type AllNodeStatusResults struct {
	Name          string                 `yaml:"name" json:"name,omitempty"`
	Status        corev1.ConditionStatus `yaml:"status" json:"status,omitempty"`
	HeartbeatTime time.Time              `yaml:"heartbeatTime" json:"heartbeatTime,omitempty"`
	Reason        string                 `yaml:"reason" json:"reason,omitempty"`
	Message       string                 `yaml:"message" json:"message,omitempty"`
	Severity      config.Severity        `yaml:"severity" json:"severity,omitempty"`
}

type BasicClusterInformation struct {
	K8sVersion   string `yaml:"k8sVersion" json:"k8sVersion,omitempty"`
	NodeNum      int    `yaml:"nodeNum" json:"nodeNum,omitempty"`
	PodNum       int    `yaml:"podNum" json:"podNum,omitempty"`
	NamespaceNum int    `yaml:"namespaceNum" json:"namespaceNum,omitempty"`
}

type PodResult struct {
	CreatedTime      string            `yaml:"createdTime" json:"createdTime,omitempty"`
	Namespace        string            `yaml:"namespace" json:"namespace,omitempty"`
	Kind             string            `yaml:"kind" json:"kind,omitempty"`
	Name             string            `yaml:"name" json:"name,omitempty"`
	Message          []string          `yaml:"message" json:"message,omitempty"`
	ContainerResults []ContainerResult `yaml:"containerResults" json:"containerResults,omitempty"`
}

type ContainerResult struct {
	Results ResultSet `yaml:"results" json:"results,omitempty"`
}

type ResultSet map[string]ResultMessage

type ResultMessage struct {
	ID       string          `yaml:"id" json:"id,omitempty"`
	Message  string          `yaml:"message" json:"message,omitempty"`
	Success  bool            `yaml:"success" json:"success,omitempty"`
	Severity config.Severity `yaml:"severity" json:"severity,omitempty"`
	Category string          `yaml:"category" json:"category,omitempty"`
}
