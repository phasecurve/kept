export function createSecrets() {
  const grafanaApiKey = new sst.Secret("GrafanaCloudAPIKey");

  return { grafanaApiKey };
}