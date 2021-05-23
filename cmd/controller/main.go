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
	"flag"
	"github.com/waveywaves/cloudevents-controller/pkg/apis/samples"
	"log"

	// The set of controllers this controller process runs.
	"github.com/waveywaves/cloudevents-controller/pkg/reconciler/addressableservice"
	"github.com/waveywaves/cloudevents-controller/pkg/reconciler/cloudeventsink"

	// This defines the shared main for injected controllers.
	"knative.dev/pkg/injection/sharedmain"
)

var (
	sinkHTTPImage = flag.String("sink-http-image", "", "The container image used for creating a HTTP based cloudevent sink")
)

func main() {

	sinkImages := samples.SinkImages{
		HTTP: *sinkHTTPImage,
	}
	if err := sinkImages.Validate(); err != nil {
		log.Fatal(err)
	}

	sharedmain.Main("controller",
		addressableservice.NewController,
		cloudeventsink.NewController(sinkImages),
	)
}
