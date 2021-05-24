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
	"github.com/waveywaves/cloudevents-controller/pkg/apis/samples"

	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"

	cloudeventsinkinformer "github.com/waveywaves/cloudevents-controller/pkg/client/injection/informers/samples/v1alpha1/cloudeventsink"
	cloudeventsinkreconciler "github.com/waveywaves/cloudevents-controller/pkg/client/injection/reconciler/samples/v1alpha1/cloudeventsink"
	kubeclient "knative.dev/pkg/client/injection/kube/client"
	podinformer "knative.dev/pkg/client/injection/kube/informers/core/v1/pod"
	svcinformer "knative.dev/pkg/client/injection/kube/informers/core/v1/service"
)

// NewController creates a Reconciler and returns the result of NewImpl.
func NewController(images samples.SinkImages) func(ctx context.Context, cmw configmap.Watcher) *controller.Impl {
	return func(ctx context.Context, cmw configmap.Watcher) *controller.Impl {
		logger := logging.FromContext(ctx)

		// Obtain an informer to both the main and child resources. These will be started by
		// the injection framework automatically. They'll keep a cached representation of the
		// cluster's state of the respective resource at all times.
		cloudeventsinkInformer := cloudeventsinkinformer.Get(ctx)
		podInformer := podinformer.Get(ctx)
		svcInformer := svcinformer.Get(ctx)

		r := &Reconciler{
			SinkImages: images,
			kubeclient: kubeclient.Get(ctx),
			podLister:  podInformer.Lister(),
			svcLister: svcInformer.Lister(),
		}
		impl := cloudeventsinkreconciler.NewImpl(ctx, r)

		logger.Info("Setting up event handlers.")

		// Listen for events on the main resource and enqueue themselves.
		cloudeventsinkInformer.Informer().AddEventHandler(controller.HandleAll(impl.Enqueue))

		// Listen for events on the child resources and enqueue the owner of them.
		podInformer.Informer().AddEventHandler(controller.HandleAll(impl.EnqueueControllerOf))

		return impl
	}
}
