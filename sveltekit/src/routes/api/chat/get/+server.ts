import { json } from '@sveltejs/kit';
import type { RequestHandler } from '@sveltejs/kit';


export const POST: RequestHandler = async ({ request, cookies }) => {
	try {
    const accessToken = cookies.get('accessToken');

	const requestJson = await request.json();
	const chatID = requestJson.chat_id;

    if (!accessToken) {
			return json({ error: 'Unauthorized' });
		}
		const url = `${process.env.BACKEND_URL}/api/chat/GetChat`;
		const response = await fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				'Authorization': `${accessToken}`
			},
			body: JSON.stringify({chat_id: chatID})
		});
		// console.log('response', response.status);
		if (!response.ok) {
			return json({ error: response.statusText});
		}
		const result = await response.json();
		return json(result);
	} catch (error) {
		throw error;
	}
};