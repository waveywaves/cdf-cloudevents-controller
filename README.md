# Cloudevents POC Controller

Create cloudevent sinks and conversion sources which map to those particular sinks.

## Terminology

`CloudeventSink` : Sink where cloudevents can be dumped

`ConversionBroker` : Source which can be used to send converted cloudevents from a Sink. 

### CloudeventSink 

`CloudeventSink` creates a `Deployment` (which backs the sink) and a corresponding `Service` 
which can be used wherever we can give a cloudevents sink url.

```yaml
kind: CloudeventSink
metadata:
  name: tekton-sink
  namespace: tekton-pipelines
spec: 
  type: "http" # type of cloudevents which the sink can accept
```

### ConversionBroker

`ConversionBroker` converts incoming cloudevent from the sink to the required type.

```yaml
kind: ConversionBroker
metadata:
  name: tekton-event-broker
  namespace: jenkins
spec:
  sink: "http://sink-url:6666"
```

If you are interested in contributing, see [CONTRIBUTING.md](./CONTRIBUTING.md)
and [DEVELOPMENT.md](./DEVELOPMENT.md).
