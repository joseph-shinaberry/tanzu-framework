// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package status contains logic for the status subcommand
package status

import (
	"context"
	"fmt"

	"gopkg.in/yaml.v2"
	v1core "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	runtimeschema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"

	"github.com/vmware-tanzu/tanzu-framework/cli/runtime/component"
	"github.com/vmware-tanzu/tanzu-framework/cmd/cli/plugin/telemetry/kubernetes"
)

type Printer struct {
	Client dynamic.Interface
	Out    component.OutputWriter
}

func (s Printer) PrintStatus() error {
	ceipMsg, err := s.getStatus(kubernetes.TelemetryNamespace, kubernetes.CeipConfigMapName)
	if err != nil {
		return err
	}
	idsMsg, err := s.getStatus(kubernetes.TelemetryNamespace, kubernetes.SharedIdsConfigMapName)
	if err != nil {
		return err
	}
	s.Out.AddRow(ceipMsg, idsMsg)
	s.Out.Render()

	return nil
}

func (s Printer) getStatus(namespace, configMapName string) (string, error) {
	ctx := context.Background()

	configMapUnstructured, err := s.Client.
		Resource(runtimeschema.GroupVersionResource{Version: "v1", Resource: "configmaps"}).
		Namespace(namespace).
		Get(ctx, configMapName, metav1.GetOptions{})
	if apierrors.IsNotFound(err) {
		return fmt.Sprintf("%s/%s configmap resource not found", namespace, configMapName), nil
	} else if err != nil {
		return "", err
	} else {
		var configMap v1core.ConfigMap
		err = runtime.DefaultUnstructuredConverter.FromUnstructured(configMapUnstructured.UnstructuredContent(), &configMap)
		if err != nil {
			return "", err
		}

		yamlData, err := yaml.Marshal(configMap.Data)
		if err != nil {
			return "", err
		}

		return string(yamlData), nil
	}
}
