/*
Copyright 2019 waveywaves

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

package main

import (
	"context"
	"flag"
	"github.com/waveywaves/cloudevents-controller/pkg/apis/samples"
	"log"

	// The set of controllers this controller process runs.
	"github.com/waveywaves/cloudevents-controller/pkg/reconciler/addressableservice"
	"github.com/waveywaves/cloudevents-controller/pkg/reconciler/cloudeventsink"
	filteredinformerfactory "knative.dev/pkg/client/injection/kube/informers/factory/filtered"

	// This defines the shared main for injected controllers.
	"knative.dev/pkg/injection/sharedmain"

	"k8s.io/client-go/rest"
)

const (
	// ControllerLogKey is the name of the logger for the controller cmd
	ControllerLogKey = "cloudevents-controller"
)

var (
	sinkHTTPImage = flag.String("sink-http-image", "", "The container image used for creating a HTTP based cloudevent sink")
)

func main() {
	cfg := sharedmain.ParseAndGetConfigOrDie()
	sinkImages := samples.SinkImages{
		HTTP: *sinkHTTPImage,
	}
	if err := sinkImages.Validate(); err != nil {
		log.Fatal(err)
	}
	if cfg.QPS == 0 {
		cfg.QPS = 2 * rest.DefaultQPS
	}
	if cfg.Burst == 0 {
		cfg.Burst = rest.DefaultBurst
	}

	ctx := filteredinformerfactory.WithSelectors(context.Background(), "")
	sharedmain.MainWithConfig(ctx, ControllerLogKey, cfg,
		addressableservice.NewController,
		cloudeventsink.NewController(sinkImages),
	)

	sharedmain.Main("controller",

	)
}
