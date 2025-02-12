/*
Copyright 2018 The Kubernetes Authors.

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

// TODO: Refactor common part of functions in this file for generic object kinds.

package utils

import (
	"context"
	"fmt"
	"time"

	apps "k8s.io/api/apps/v1"
	batch "k8s.io/api/batch/v1"
	storage "k8s.io/api/storage/v1"

	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	clientset "k8s.io/client-go/kubernetes"
)

const (
	// Parameters for retrying with exponential backoff.
	retryBackoffInitialDuration = 100 * time.Millisecond
	retryBackoffFactor          = 3
	retryBackoffJitter          = 0
	retryBackoffSteps           = 60
)

// Utility for retrying the given function with exponential backoff.
func RetryWithExponentialBackOff(fn wait.ConditionFunc) error {
	backoff := wait.Backoff{
		Duration: retryBackoffInitialDuration,
		Factor:   retryBackoffFactor,
		Jitter:   retryBackoffJitter,
		Steps:    retryBackoffSteps,
	}
	return wait.ExponentialBackoff(backoff, fn)
}

func CreatePodWithRetries(c clientset.Interface, namespace string, obj *v1.Pod) error {
	if obj == nil {
		return fmt.Errorf("object provided to create is empty")
	}
	createFunc := func() (bool, error) {
		_, err := c.CoreV1().Pods(namespace).Create(context.TODO(), obj, metav1.CreateOptions{})
		if isGenerateNameConflict(obj.ObjectMeta, err) {
			return false, nil
		}
		if err == nil || apierrors.IsAlreadyExists(err) {
			return true, nil
		}
		return false, fmt.Errorf("failed to create object with non-retriable error: %v ", err)
	}
	return RetryWithExponentialBackOff(createFunc)
}

func CreateRCWithRetries(c clientset.Interface, namespace string, obj *v1.ReplicationController) error {
	if obj == nil {
		return fmt.Errorf("object provided to create is empty")
	}
	createFunc := func() (bool, error) {
		_, err := c.CoreV1().ReplicationControllers(namespace).Create(context.TODO(), obj, metav1.CreateOptions{})
		if isGenerateNameConflict(obj.ObjectMeta, err) {
			return false, nil
		}
		if err == nil || apierrors.IsAlreadyExists(err) {
			return true, nil
		}
		return false, fmt.Errorf("failed to create object with non-retriable error: %v", err)
	}
	return RetryWithExponentialBackOff(createFunc)
}

func CreateReplicaSetWithRetries(c clientset.Interface, namespace string, obj *apps.ReplicaSet) error {
	if obj == nil {
		return fmt.Errorf("object provided to create is empty")
	}
	createFunc := func() (bool, error) {
		_, err := c.AppsV1().ReplicaSets(namespace).Create(context.TODO(), obj, metav1.CreateOptions{})
		if isGenerateNameConflict(obj.ObjectMeta, err) {
			return false, nil
		}
		if err == nil || apierrors.IsAlreadyExists(err) {
			return true, nil
		}
		return false, fmt.Errorf("failed to create object with non-retriable error: %v", err)
	}
	return RetryWithExponentialBackOff(createFunc)
}

func CreateDeploymentWithRetries(c clientset.Interface, namespace string, obj *apps.Deployment) error {
	if obj == nil {
		return fmt.Errorf("object provided to create is empty")
	}
	createFunc := func() (bool, error) {
		_, err := c.AppsV1().Deployments(namespace).Create(context.TODO(), obj, metav1.CreateOptions{})
		if isGenerateNameConflict(obj.ObjectMeta, err) {
			return false, nil
		}
		if err == nil || apierrors.IsAlreadyExists(err) {
			return true, nil
		}
		return false, fmt.Errorf("failed to create object with non-retriable error: %v", err)
	}
	return RetryWithExponentialBackOff(createFunc)
}

func CreateDaemonSetWithRetries(c clientset.Interface, namespace string, obj *apps.DaemonSet) error {
	if obj == nil {
		return fmt.Errorf("object provided to create is empty")
	}
	createFunc := func() (bool, error) {
		_, err := c.AppsV1().DaemonSets(namespace).Create(context.TODO(), obj, metav1.CreateOptions{})
		if isGenerateNameConflict(obj.ObjectMeta, err) {
			return false, nil
		}
		if err == nil || apierrors.IsAlreadyExists(err) {
			return true, nil
		}
		return false, fmt.Errorf("failed to create object with non-retriable error: %v", err)
	}
	return RetryWithExponentialBackOff(createFunc)
}

func CreateJobWithRetries(c clientset.Interface, namespace string, obj *batch.Job) error {
	if obj == nil {
		return fmt.Errorf("object provided to create is empty")
	}
	createFunc := func() (bool, error) {
		_, err := c.BatchV1().Jobs(namespace).Create(context.TODO(), obj, metav1.CreateOptions{})
		if isGenerateNameConflict(obj.ObjectMeta, err) {
			return false, nil
		}
		if err == nil || apierrors.IsAlreadyExists(err) {
			return true, nil
		}
		return false, fmt.Errorf("failed to create object with non-retriable error: %v", err)
	}
	return RetryWithExponentialBackOff(createFunc)
}

func CreateSecretWithRetries(c clientset.Interface, namespace string, obj *v1.Secret) error {
	if obj == nil {
		return fmt.Errorf("object provided to create is empty")
	}
	createFunc := func() (bool, error) {
		_, err := c.CoreV1().Secrets(namespace).Create(context.TODO(), obj, metav1.CreateOptions{})
		if isGenerateNameConflict(obj.ObjectMeta, err) {
			return false, nil
		}
		if err == nil || apierrors.IsAlreadyExists(err) {
			return true, nil
		}
		return false, fmt.Errorf("failed to create object with non-retriable error: %v", err)
	}
	return RetryWithExponentialBackOff(createFunc)
}

func CreateConfigMapWithRetries(c clientset.Interface, namespace string, obj *v1.ConfigMap) error {
	if obj == nil {
		return fmt.Errorf("object provided to create is empty")
	}
	createFunc := func() (bool, error) {
		_, err := c.CoreV1().ConfigMaps(namespace).Create(context.TODO(), obj, metav1.CreateOptions{})
		if isGenerateNameConflict(obj.ObjectMeta, err) {
			return false, nil
		}
		if err == nil || apierrors.IsAlreadyExists(err) {
			return true, nil
		}
		return false, fmt.Errorf("failed to create object with non-retriable error: %v", err)
	}
	return RetryWithExponentialBackOff(createFunc)
}

func CreateServiceWithRetries(c clientset.Interface, namespace string, obj *v1.Service) error {
	if obj == nil {
		return fmt.Errorf("object provided to create is empty")
	}
	createFunc := func() (bool, error) {
		_, err := c.CoreV1().Services(namespace).Create(context.TODO(), obj, metav1.CreateOptions{})
		if isGenerateNameConflict(obj.ObjectMeta, err) {
			return false, nil
		}
		if err == nil || apierrors.IsAlreadyExists(err) {
			return true, nil
		}
		return false, fmt.Errorf("failed to create object with non-retriable error: %v", err)
	}
	return RetryWithExponentialBackOff(createFunc)
}

func CreateStorageClassWithRetries(c clientset.Interface, obj *storage.StorageClass) error {
	if obj == nil {
		return fmt.Errorf("object provided to create is empty")
	}
	createFunc := func() (bool, error) {
		_, err := c.StorageV1().StorageClasses().Create(context.TODO(), obj, metav1.CreateOptions{})
		if isGenerateNameConflict(obj.ObjectMeta, err) {
			return false, nil
		}
		if err == nil || apierrors.IsAlreadyExists(err) {
			return true, nil
		}
		return false, fmt.Errorf("failed to create object with non-retriable error: %v", err)
	}
	return RetryWithExponentialBackOff(createFunc)
}

func CreateResourceQuotaWithRetries(c clientset.Interface, namespace string, obj *v1.ResourceQuota) error {
	if obj == nil {
		return fmt.Errorf("object provided to create is empty")
	}
	createFunc := func() (bool, error) {
		_, err := c.CoreV1().ResourceQuotas(namespace).Create(context.TODO(), obj, metav1.CreateOptions{})
		if isGenerateNameConflict(obj.ObjectMeta, err) {
			return false, nil
		}
		if err == nil || apierrors.IsAlreadyExists(err) {
			return true, nil
		}
		return false, fmt.Errorf("failed to create object with non-retriable error: %v", err)
	}
	return RetryWithExponentialBackOff(createFunc)
}

func CreatePersistentVolumeWithRetries(c clientset.Interface, obj *v1.PersistentVolume) error {
	if obj == nil {
		return fmt.Errorf("object provided to create is empty")
	}
	createFunc := func() (bool, error) {
		_, err := c.CoreV1().PersistentVolumes().Create(context.TODO(), obj, metav1.CreateOptions{})
		if isGenerateNameConflict(obj.ObjectMeta, err) {
			return false, nil
		}
		if err == nil || apierrors.IsAlreadyExists(err) {
			return true, nil
		}
		return false, fmt.Errorf("failed to create object with non-retriable error: %v", err)
	}
	return RetryWithExponentialBackOff(createFunc)
}

func CreatePersistentVolumeClaimWithRetries(c clientset.Interface, namespace string, obj *v1.PersistentVolumeClaim) error {
	if obj == nil {
		return fmt.Errorf("object provided to create is empty")
	}
	createFunc := func() (bool, error) {
		_, err := c.CoreV1().PersistentVolumeClaims(namespace).Create(context.TODO(), obj, metav1.CreateOptions{})
		if isGenerateNameConflict(obj.ObjectMeta, err) {
			return false, nil
		}
		if err == nil || apierrors.IsAlreadyExists(err) {
			return true, nil
		}
		return false, fmt.Errorf("failed to create object with non-retriable error: %v", err)
	}
	return RetryWithExponentialBackOff(createFunc)
}

// isGenerateNameConflict returns whether the error is generateName conflict or not.
func isGenerateNameConflict(meta metav1.ObjectMeta, err error) bool {
	if apierrors.IsAlreadyExists(err) && meta.Name == "" {
		return true
	}
	return false
}
