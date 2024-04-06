import { json } from '@sveltejs/kit';
import type { RequestHandler } from '@sveltejs/kit';

const IsDev = process.env.DEV === 'true';

export const POST: RequestHandler = async ({ request, cookies }) => {
	try {
		const requestJson = await request.json();
		const prompt = requestJson.prompt;
		const chatID = requestJson.chat_id;
		const accessToken = cookies.get('accessToken');

		if (!chatID) {
			return json({ error: 'Chat ID is required' });
		}

		if (!accessToken) {
				return json({ error: 'Unauthorized' });
		}
		if (!prompt) {
			return json({ error: 'Prompt is required' });
		}

		let url = `${process.env.BACKEND_URL}/api/chat/SentPrompt`;
		if (IsDev) {
			url = `${process.env.BACKEND_URL}/api/chat/SentPromptTest`;
		}
		const response = await fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				'Authorization': `${accessToken}`
			},
			body: JSON.stringify({chat_id: chatID , prompt: prompt })
		});
		if (!response.ok) {
			return json({ error: response.statusText });
		}
		const result = await response.json();
		return json(result);
	} catch (error) {
		console.error(error);
		return json({ error: 'Internal server error' });
	}
};