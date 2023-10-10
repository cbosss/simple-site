import { Context, Request, Response } from "@netlify/functions"

export default async (req: Request, context: Context) => {
	const file = await context.blob.get("some-key")
	console.log(file)
	return new Response("Done")
}