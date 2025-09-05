/// <reference path="./.sst/platform/config.d.ts" />

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
    const { createKeptAPI } = await import("./infrastructure/api");
    const { createCollector } = await import("./infrastructure/collector");
    const { collector, vpc } = await createCollector();
    createKeptAPI(collector, vpc);
  },
});
