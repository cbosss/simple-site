import type { Handler, HandlerEvent, HandlerContext } from "@netlify/functions";
import { Blobs } from "@netlify/blobs";

const handler: Handler = async (event: HandlerEvent, context: HandlerContext) => {
	console.log("ENVIRONMENT VARIABLES\n" + JSON.stringify(process.env, null, 2))
	fetch("www.netlify.com")
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



	return {
		statusCode: 200,
		body: JSON.stringify({ message: "Hello World" }),
	};
};

export { handler };
