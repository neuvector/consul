package telemetry

import (
	"context"
	"net/url"

	metricpb "go.opentelemetry.io/proto/otlp/metrics/v1"
)

// MetricsClient exports Consul metrics in OTLP format to the desired endpoint.
type MetricsClient interface {
	ExportMetrics(ctx context.Context, protoMetrics *metricpb.ResourceMetrics, endpoint string) error
}

// EndpointProvider provides the endpoint where metrics are exported to by the OTELExporter.
// EndpointProvider exposes the GetEndpoint() interface method to fetch the endpoint.
// This abstraction layer offers flexibility, in particular for dynamic configuration or changes to the endpoint.
type EndpointProvider interface {
	GetEndpoint() *url.URL
}
