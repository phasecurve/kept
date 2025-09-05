export async function createCollector() {
  const { GRAFANA_CLOUD_INSTANCE_ID, GRAFANA_CLOUD_OTLP_ENDPOINT } =
    await import("./env");
  const vpc = new sst.aws.Vpc("KeptVpc");

  const cluster = new sst.aws.Cluster("KeptCluster", { vpc });

  const collector = new sst.aws.Service("GrafanaCollector", {
    cluster,
    image: {
      context: "./collector",
      dockerfile: "Dockerfile",
    },
    port: 4318,
    environment: {
      GRAFANA_CLOUD_INSTANCE_ID,
      GRAFANA_CLOUD_OTLP_ENDPOINT,
      GRAFANA_CLOUD_API_KEY: new sst.Secret("GRAFANA_CLOUD_API_KEY").value,
    },
    permissions: [
      {
        actions: ["logs:CreateLogGroup", "logs:CreateLogStream", "logs:PutLogEvents"],
        resources: ["*"],
      },
    ],
  });

  return { collector, vpc };
}
