import { Context, Request, Response } from "@netlify/functions"

export default async (req: Request, context: Context) => {
	await context.blob.set("some-key", req.body)

	return new Response("Done")
}