import type { Handler, HandlerEvent, HandlerContext } from "@netlify/functions";
import { Blobs } from "@netlify/blobs";

const handler: Handler = async (event: HandlerEvent, context: HandlerContext) => {
	const rawData = Buffer.from(context.clientContext.custom.blobs, "base64");
	const data = JSON.parse(rawData.toString("ascii"));
	const blobs = new Blobs({
		authentication: {
			contextURL: data.url,
			token: data.token,
		},
		context: `deploy:${event.headers["x-nf-deploy-id"]}`,
		siteID: event.headers["x-nf-site-id"],
	});

	await blobs.set(event.headers["nf-key"], event.headers["nf-value"]);

	return {
		statusCode: 200,
	}
};

export { handler };