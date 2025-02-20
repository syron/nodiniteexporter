package nodiniteexporter

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

var (
	typeStr = component.MustNewType("nodinite")
)

// NewFactory creates a Datadog exporter factory
func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		typeStr,
		createDefaultConfig,
		exporter.WithLogs(createLogsExporter, component.StabilityLevelDevelopment),
	)
}

func createDefaultConfig() component.Config {

	return &config{}
}

func createLogsExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Logs, error) {

	return exporterhelper.NewLogs(ctx, set, cfg,
		pushLogs,
		//	The parameters below are optional. Uncomment any as you need.
		//	exporterhelper.WithStart(start component.StartFunc),
		//exporterhelper.WithShutdown(shutdown component.ShutdownFunc),
		//exporterhelper.WithTimeout(timeoutSettings TimeoutSettings),
		//exporterhelper.WithRetry(retrySettings RetrySettings),
		//exporterhelper.WithQueue(queueSettings QueueSettings),
		//exporterhelper.WithCapabilities(capabilities consumer.Capabilities)
	)

}
