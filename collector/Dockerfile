FROM otel/opentelemetry-collector-contrib:0.98.0

COPY otel-config.yaml /etc/otelcol-contrib/config.yaml

CMD ["--config", "/etc/otelcol-contrib/config.yaml"]