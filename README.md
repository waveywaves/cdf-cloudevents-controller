# Cloudevents Controller

Create sinks to ingest cloudevents and conversion brokers which convert cloudevents from on type to the other. 
Converted cloudevents can be forwarded to other sinks. 

## Terminology

`CloudeventSink` : Sink where cloudevents can be dumped

`ConversionBroker` : Source which can be used to send converted cloudevents from a Sink. 

### CloudeventSink 

`CloudeventSink` creates a `Pod` (which backs the sink) and a corresponding `Service` 
which can be used wherever we can give a cloudevents sink url.

```yaml
kind: CloudeventSink
metadata:
  name: tekton-sink
  namespace: tekton-pipelines
spec: 
  type: "http" # type of cloudevents which the sink can accept
```

### ConversionBroker (TODO)

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
