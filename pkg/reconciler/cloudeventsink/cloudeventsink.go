/*
Copyright 2020 waveywaves

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

package cloudeventsink

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/util/intstr"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	corev1listers "k8s.io/client-go/listers/core/v1"

	"github.com/waveywaves/cloudevents-controller/pkg/apis/samples"
	samplesv1alpha1 "github.com/waveywaves/cloudevents-controller/pkg/apis/samples/v1alpha1"
	cloudeventsinkreconciler "github.com/waveywaves/cloudevents-controller/pkg/client/injection/reconciler/samples/v1alpha1/cloudeventsink"
	"knative.dev/pkg/kmeta"
	"knative.dev/pkg/logging"
	"knative.dev/pkg/reconciler"
)

// podOwnerLabelKey is the key to a label that points to the owner (creator) of the
// pod, allowing us to easily list all pods a single CloudeventSink created.
const podOwnerLabelKey = samples.GroupName + "/podOwner"
const svcOwnerLabelKey = samples.GroupName + "/svcOwner"

// Reconciler implements cloudeventsinkreconciler.Interface for
// CloudeventSink resources.
type Reconciler struct {
	kubeclient kubernetes.Interface
	podLister  corev1listers.PodLister
	svcLister corev1listers.ServiceLister
	SinkImages samples.SinkImages
}

// Check that our Reconciler implements Interface
var _ cloudeventsinkreconciler.Interface = (*Reconciler)(nil)

// ReconcileKind implements Interface.ReconcileKind.
func (r *Reconciler) ReconcileKind(ctx context.Context, ces *samplesv1alpha1.CloudeventSink) reconciler.Event {
	// This logger has all the context necessary to identify which resource is being reconciled.
	logger := logging.FromContext(ctx)

	// Get all the pods created by the current CloudeventSink. The result is read from
	// cache (via the lister).
	podSelector := labels.SelectorFromSet(labels.Set{
		podOwnerLabelKey: ces.Name,
	})

	existingPods, err := r.podLister.Pods(ces.Namespace).List(podSelector)
	if err != nil {
		return fmt.Errorf("failed to list existing pods: %w", err)
	}
	logger.Infof("Found %ces pods in total", len(existingPods))

	var pod *corev1.Pod
	if len(existingPods) == 0 {
		//sinkType := ces.Spec.SinkType
		//switch sinkType {
		//case "http":
		pod = r.makeHTTPSinkPod(ces)
		if _, err := r.kubeclient.CoreV1().Pods(pod.Namespace).Create(ctx, pod, metav1.CreateOptions{}); err != nil {
			return fmt.Errorf("failed to create pod: %w", err)
		}

		svc := r.makeService(ces)
		if _, err := r.kubeclient.CoreV1().Services(svc.Namespace).Create(ctx, svc, metav1.CreateOptions{}); err != nil {
			return fmt.Errorf("failed to create service: %w", err)
		}

		//default:
		//	logger.Infof("Unknown sink: Cannot create sink '"+sinkType+"'")
		//}
	}

	return nil
}

// makeHTTPSinkPod generates a simple pod to be created in the given namespace with the given
// image.
func (r *Reconciler) makeHTTPSinkPod(ces *samplesv1alpha1.CloudeventSink) *corev1.Pod {
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:    ces.Namespace,
			GenerateName: ces.Name + "-cesink-",
			Labels: map[string]string{
				// The label allows for easy querying of all the pods created.
				podOwnerLabelKey: ces.Name,
			},
			// The OwnerReference makes sure the pods get removed automatically once the
			// CloudeventSink is removed.
			OwnerReferences: []metav1.OwnerReference{*kmeta.NewControllerRef(ces)},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{{
				Name:  "http-sink",
				Image: r.SinkImages.HTTP,
			}},
		},
	}
}

// makeHTTPSinkPod generates a simple pod to be created in the given namespace with the given
// image.
func (r *Reconciler) makeService(ces *samplesv1alpha1.CloudeventSink) *corev1.Service {
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:    ces.Namespace,
			Name: 		  ces.Name+"-cesink",
			Labels: map[string]string{
				// The label allows for easy querying of all the pods created.
				svcOwnerLabelKey: ces.Name,
			},
			// The OwnerReference makes sure the pods get removed automatically once the
			// CloudeventSink is removed.
			OwnerReferences: []metav1.OwnerReference{*kmeta.NewControllerRef(ces)},
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				// The label allows for easy querying of all the pods created.
				podOwnerLabelKey: ces.Name,
			},
			Ports: []corev1.ServicePort{
				{
					Port: 80,
					TargetPort: intstr.IntOrString{IntVal: 8080, StrVal: "8080"},
				},
			},
		},
	}
}

// isPodReady returns whether or not the given pod is ready.
func isPodReady(p *corev1.Pod) bool {
	if p.Status.Phase == corev1.PodRunning && p.DeletionTimestamp == nil {
		for _, cond := range p.Status.Conditions {
			if cond.Type == corev1.PodReady && cond.Status == corev1.ConditionTrue {
				return true
			}
		}
	}
	return false
}
