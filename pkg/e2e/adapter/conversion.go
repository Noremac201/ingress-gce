/*
Copyright 2021 The Kubernetes Authors.

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

package adapter

// **NOTE** These conversion functions have been copied from kubernetes/kubernetes v1.19.0
// The only changes made were to convert between v1beta1.Ingress and v1.Ingress instead of
// internal ingress structure

import (
	networking "k8s.io/api/networking/v1"
	v1beta1 "k8s.io/api/networking/v1beta1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func Convert_v1beta1_IngressBackend_To_networking_IngressBackend(in *v1beta1.IngressBackend, out *networking.IngressBackend, s conversion.Scope) error {
	if err := autoConvert_v1beta1_IngressBackend_To_networking_IngressBackend(in, out, s); err != nil {
		return err
	}
	if len(in.ServiceName) > 0 || in.ServicePort.IntVal != 0 || in.ServicePort.StrVal != "" || in.ServicePort.Type == intstr.String {
		out.Service = &networking.IngressServiceBackend{}
		out.Service.Name = in.ServiceName
		out.Service.Port.Name = in.ServicePort.StrVal
		out.Service.Port.Number = in.ServicePort.IntVal
	}
	return nil
}

func Convert_networking_IngressBackend_To_v1beta1_IngressBackend(in *networking.IngressBackend, out *v1beta1.IngressBackend, s conversion.Scope) error {
	if err := autoConvert_networking_IngressBackend_To_v1beta1_IngressBackend(in, out, s); err != nil {
		return err
	}
	if in.Service != nil {
		out.ServiceName = in.Service.Name
		if len(in.Service.Port.Name) > 0 {
			out.ServicePort = intstr.FromString(in.Service.Port.Name)
		} else {
			out.ServicePort = intstr.FromInt(int(in.Service.Port.Number))
		}
	}
	return nil
}
func Convert_v1beta1_IngressSpec_To_networking_IngressSpec(in *v1beta1.IngressSpec, out *networking.IngressSpec, s conversion.Scope) error {
	if err := autoConvert_v1beta1_IngressSpec_To_networking_IngressSpec(in, out, s); err != nil {
		return nil
	}
	if in.Backend != nil {
		out.DefaultBackend = &networking.IngressBackend{}
		if err := Convert_v1beta1_IngressBackend_To_networking_IngressBackend(in.Backend, out.DefaultBackend, s); err != nil {
			return err
		}
	}
	return nil
}

func Convert_networking_IngressSpec_To_v1beta1_IngressSpec(in *networking.IngressSpec, out *v1beta1.IngressSpec, s conversion.Scope) error {
	if err := autoConvert_networking_IngressSpec_To_v1beta1_IngressSpec(in, out, s); err != nil {
		return nil
	}
	if in.DefaultBackend != nil {
		out.Backend = &v1beta1.IngressBackend{}
		if err := Convert_networking_IngressBackend_To_v1beta1_IngressBackend(in.DefaultBackend, out.Backend, s); err != nil {
			return err
		}
	}
	return nil
}
