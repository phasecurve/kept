/// <reference path="./.sst/platform/config.d.ts" />

import { transform } from "./.sst/platform/src/components/component";

export default $config({
  app(input) {
    return {
      name: "kept",
      removal: input?.stage === "production" ? "retain" : "remove",
      home: "aws",
      providers: {
        aws: {
          region: "eu-west-2",
        },
      },
    };
  },
  async run() {
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
      handler: "functions/healthz",
      runtime: "go",
    });

    return {
      ApiUrl: api.url,
    };
  },
});
