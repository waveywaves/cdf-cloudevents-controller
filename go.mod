module github.com/waveywaves/cloudevents-controller

go 1.15

require (
	github.com/cloudevents/sdk-go/v2 v2.4.1
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	k8s.io/api v0.19.7
	k8s.io/apimachinery v0.19.7
	k8s.io/client-go v0.19.7
	k8s.io/code-generator v0.19.7
	k8s.io/kube-openapi v0.0.0-20200805222855-6aeccd4b50c6
	knative.dev/hack v0.0.0-20210428122153-93ad9129c268
	knative.dev/hack/schema v0.0.0-20210428122153-93ad9129c268
	knative.dev/pkg v0.0.0-20210510175900-4564797bf3b7
)
