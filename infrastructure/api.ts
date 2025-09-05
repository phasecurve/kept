export function createKeptAPI(collector: sst.aws.Service, vpc: sst.aws.Vpc) {
  const api = new sst.aws.ApiGatewayV2("KeptAPI");

  api.route("GET /healthz", {
    handler: "lambda/functions/healthz",
    runtime: "go",
    vpc,
    environment: {
      OTEL_EXPORTER_OTLP_ENDPOINT: collector.service,
    },
    link: [collector],
  });
}
