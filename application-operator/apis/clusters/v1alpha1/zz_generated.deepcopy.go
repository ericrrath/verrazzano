// +build !ignore_autogenerated

// Copyright (c) 2021, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationConfigurationTemplate) DeepCopyInto(out *ApplicationConfigurationTemplate) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationConfigurationTemplate.
func (in *ApplicationConfigurationTemplate) DeepCopy() *ApplicationConfigurationTemplate {
	if in == nil {
		return nil
	}
	out := new(ApplicationConfigurationTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentTemplate) DeepCopyInto(out *ComponentTemplate) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentTemplate.
func (in *ComponentTemplate) DeepCopy() *ComponentTemplate {
	if in == nil {
		return nil
	}
	out := new(ComponentTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapTemplate) DeepCopyInto(out *ConfigMapTemplate) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	if in.Immutable != nil {
		in, out := &in.Immutable, &out.Immutable
		*out = new(bool)
		**out = **in
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.BinaryData != nil {
		in, out := &in.BinaryData, &out.BinaryData
		*out = make(map[string][]byte, len(*in))
		for key, val := range *in {
			var outVal []byte
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]byte, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapTemplate.
func (in *ConfigMapTemplate) DeepCopy() *ConfigMapTemplate {
	if in == nil {
		return nil
	}
	out := new(ConfigMapTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingScopeTemplate) DeepCopyInto(out *LoggingScopeTemplate) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingScopeTemplate.
func (in *LoggingScopeTemplate) DeepCopy() *LoggingScopeTemplate {
	if in == nil {
		return nil
	}
	out := new(LoggingScopeTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterApplicationConfiguration) DeepCopyInto(out *MultiClusterApplicationConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterApplicationConfiguration.
func (in *MultiClusterApplicationConfiguration) DeepCopy() *MultiClusterApplicationConfiguration {
	if in == nil {
		return nil
	}
	out := new(MultiClusterApplicationConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiClusterApplicationConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterApplicationConfigurationList) DeepCopyInto(out *MultiClusterApplicationConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MultiClusterApplicationConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterApplicationConfigurationList.
func (in *MultiClusterApplicationConfigurationList) DeepCopy() *MultiClusterApplicationConfigurationList {
	if in == nil {
		return nil
	}
	out := new(MultiClusterApplicationConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiClusterApplicationConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterApplicationConfigurationSpec) DeepCopyInto(out *MultiClusterApplicationConfigurationSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	in.Placement.DeepCopyInto(&out.Placement)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterApplicationConfigurationSpec.
func (in *MultiClusterApplicationConfigurationSpec) DeepCopy() *MultiClusterApplicationConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(MultiClusterApplicationConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterApplicationConfigurationStatus) DeepCopyInto(out *MultiClusterApplicationConfigurationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterApplicationConfigurationStatus.
func (in *MultiClusterApplicationConfigurationStatus) DeepCopy() *MultiClusterApplicationConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(MultiClusterApplicationConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterComponent) DeepCopyInto(out *MultiClusterComponent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterComponent.
func (in *MultiClusterComponent) DeepCopy() *MultiClusterComponent {
	if in == nil {
		return nil
	}
	out := new(MultiClusterComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiClusterComponent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterComponentList) DeepCopyInto(out *MultiClusterComponentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MultiClusterComponent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterComponentList.
func (in *MultiClusterComponentList) DeepCopy() *MultiClusterComponentList {
	if in == nil {
		return nil
	}
	out := new(MultiClusterComponentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiClusterComponentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterComponentSpec) DeepCopyInto(out *MultiClusterComponentSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	in.Placement.DeepCopyInto(&out.Placement)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterComponentSpec.
func (in *MultiClusterComponentSpec) DeepCopy() *MultiClusterComponentSpec {
	if in == nil {
		return nil
	}
	out := new(MultiClusterComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterComponentStatus) DeepCopyInto(out *MultiClusterComponentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterComponentStatus.
func (in *MultiClusterComponentStatus) DeepCopy() *MultiClusterComponentStatus {
	if in == nil {
		return nil
	}
	out := new(MultiClusterComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterConfigMap) DeepCopyInto(out *MultiClusterConfigMap) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterConfigMap.
func (in *MultiClusterConfigMap) DeepCopy() *MultiClusterConfigMap {
	if in == nil {
		return nil
	}
	out := new(MultiClusterConfigMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiClusterConfigMap) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterConfigMapList) DeepCopyInto(out *MultiClusterConfigMapList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MultiClusterConfigMap, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterConfigMapList.
func (in *MultiClusterConfigMapList) DeepCopy() *MultiClusterConfigMapList {
	if in == nil {
		return nil
	}
	out := new(MultiClusterConfigMapList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiClusterConfigMapList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterConfigMapSpec) DeepCopyInto(out *MultiClusterConfigMapSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	in.Placement.DeepCopyInto(&out.Placement)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterConfigMapSpec.
func (in *MultiClusterConfigMapSpec) DeepCopy() *MultiClusterConfigMapSpec {
	if in == nil {
		return nil
	}
	out := new(MultiClusterConfigMapSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterConfigMapStatus) DeepCopyInto(out *MultiClusterConfigMapStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterConfigMapStatus.
func (in *MultiClusterConfigMapStatus) DeepCopy() *MultiClusterConfigMapStatus {
	if in == nil {
		return nil
	}
	out := new(MultiClusterConfigMapStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterLoggingScope) DeepCopyInto(out *MultiClusterLoggingScope) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterLoggingScope.
func (in *MultiClusterLoggingScope) DeepCopy() *MultiClusterLoggingScope {
	if in == nil {
		return nil
	}
	out := new(MultiClusterLoggingScope)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiClusterLoggingScope) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterLoggingScopeList) DeepCopyInto(out *MultiClusterLoggingScopeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MultiClusterLoggingScope, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterLoggingScopeList.
func (in *MultiClusterLoggingScopeList) DeepCopy() *MultiClusterLoggingScopeList {
	if in == nil {
		return nil
	}
	out := new(MultiClusterLoggingScopeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiClusterLoggingScopeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterLoggingScopeSpec) DeepCopyInto(out *MultiClusterLoggingScopeSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	in.Placement.DeepCopyInto(&out.Placement)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterLoggingScopeSpec.
func (in *MultiClusterLoggingScopeSpec) DeepCopy() *MultiClusterLoggingScopeSpec {
	if in == nil {
		return nil
	}
	out := new(MultiClusterLoggingScopeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterLoggingScopeStatus) DeepCopyInto(out *MultiClusterLoggingScopeStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterLoggingScopeStatus.
func (in *MultiClusterLoggingScopeStatus) DeepCopy() *MultiClusterLoggingScopeStatus {
	if in == nil {
		return nil
	}
	out := new(MultiClusterLoggingScopeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterSecret) DeepCopyInto(out *MultiClusterSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterSecret.
func (in *MultiClusterSecret) DeepCopy() *MultiClusterSecret {
	if in == nil {
		return nil
	}
	out := new(MultiClusterSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiClusterSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterSecretList) DeepCopyInto(out *MultiClusterSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MultiClusterSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterSecretList.
func (in *MultiClusterSecretList) DeepCopy() *MultiClusterSecretList {
	if in == nil {
		return nil
	}
	out := new(MultiClusterSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiClusterSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterSecretSpec) DeepCopyInto(out *MultiClusterSecretSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	in.Placement.DeepCopyInto(&out.Placement)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterSecretSpec.
func (in *MultiClusterSecretSpec) DeepCopy() *MultiClusterSecretSpec {
	if in == nil {
		return nil
	}
	out := new(MultiClusterSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterSecretStatus) DeepCopyInto(out *MultiClusterSecretStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterSecretStatus.
func (in *MultiClusterSecretStatus) DeepCopy() *MultiClusterSecretStatus {
	if in == nil {
		return nil
	}
	out := new(MultiClusterSecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceTemplate) DeepCopyInto(out *NamespaceTemplate) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceTemplate.
func (in *NamespaceTemplate) DeepCopy() *NamespaceTemplate {
	if in == nil {
		return nil
	}
	out := new(NamespaceTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Placement) DeepCopyInto(out *Placement) {
	*out = *in
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]Cluster, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Placement.
func (in *Placement) DeepCopy() *Placement {
	if in == nil {
		return nil
	}
	out := new(Placement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectTemplate) DeepCopyInto(out *ProjectTemplate) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]NamespaceTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Security.DeepCopyInto(&out.Security)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectTemplate.
func (in *ProjectTemplate) DeepCopy() *ProjectTemplate {
	if in == nil {
		return nil
	}
	out := new(ProjectTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretTemplate) DeepCopyInto(out *SecretTemplate) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make(map[string][]byte, len(*in))
		for key, val := range *in {
			var outVal []byte
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]byte, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.StringData != nil {
		in, out := &in.StringData, &out.StringData
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretTemplate.
func (in *SecretTemplate) DeepCopy() *SecretTemplate {
	if in == nil {
		return nil
	}
	out := new(SecretTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecuritySpec) DeepCopyInto(out *SecuritySpec) {
	*out = *in
	if in.ProjectAdminSubjects != nil {
		in, out := &in.ProjectAdminSubjects, &out.ProjectAdminSubjects
		*out = make([]v1.Subject, len(*in))
		copy(*out, *in)
	}
	if in.ProjectMonitorSubjects != nil {
		in, out := &in.ProjectMonitorSubjects, &out.ProjectMonitorSubjects
		*out = make([]v1.Subject, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecuritySpec.
func (in *SecuritySpec) DeepCopy() *SecuritySpec {
	if in == nil {
		return nil
	}
	out := new(SecuritySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoProject) DeepCopyInto(out *VerrazzanoProject) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoProject.
func (in *VerrazzanoProject) DeepCopy() *VerrazzanoProject {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoProject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VerrazzanoProject) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoProjectList) DeepCopyInto(out *VerrazzanoProjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VerrazzanoProject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoProjectList.
func (in *VerrazzanoProjectList) DeepCopy() *VerrazzanoProjectList {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoProjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VerrazzanoProjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoProjectSpec) DeepCopyInto(out *VerrazzanoProjectSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoProjectSpec.
func (in *VerrazzanoProjectSpec) DeepCopy() *VerrazzanoProjectSpec {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoProjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoProjectStatus) DeepCopyInto(out *VerrazzanoProjectStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoProjectStatus.
func (in *VerrazzanoProjectStatus) DeepCopy() *VerrazzanoProjectStatus {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoProjectStatus)
	in.DeepCopyInto(out)
	return out
}
