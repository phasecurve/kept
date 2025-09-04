export function createKeptAPI() {
  const api = new sst.aws.ApiGatewayV2("KeptAPI", {
    transform: {
      route: {
        handler: (args, opts) => {
          args.memory ??= "128 MB";
        },
      },
    },
  });

  api.route("GET /healthz", {
    handler: "lambda/functions/healthz",
    runtime: "go",
  });
}