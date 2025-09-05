export async function createCollector() {
  const { GRAFANA_CLOUD_INSTANCE_ID, GRAFANA_CLOUD_OTLP_ENDPOINT } =
    await import("./env");
  const { createSecrets } = await import("./secret-manager");
  const { grafanaApiKey } = createSecrets();

  const vpc = new sst.aws.Vpc("KeptVpc");

  const cluster = new sst.aws.Cluster("KeptCluster", { vpc });

  const collector = new sst.aws.Service("GrafanaCollector", {
    cluster,
    image: {
      context: "./collector",
      dockerfile: "Dockerfile",
    },
    environment: {
      GRAFANA_CLOUD_INSTANCE_ID,
      GRAFANA_CLOUD_OTLP_ENDPOINT,
      GRAFANA_CLOUD_API_KEY: grafanaApiKey.value,
    },
    permissions: [
      {
        actions: [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents",
        ],
        resources: ["*"],
      },
    ],
  });

  return { collector, vpc };
}
