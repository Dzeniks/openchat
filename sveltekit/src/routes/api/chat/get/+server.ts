import { json } from '@sveltejs/kit';
import type { RequestHandler } from '@sveltejs/kit';


export const POST: RequestHandler = async ({ request, cookies }) => {
	try {
    const accessToken = cookies.get('accessToken');
    if (!accessToken) {
			return json({ error: 'Unauthorized' });
		}
		const url = `${process.env.BACKEND_URL}/api/chat/CreateChat`;
		const response = await fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				'Authorization': `${accessToken}`
			},
		});
		if (!response.ok) {
			return json({ error: response.statusText});
		}
		const result = await response.json();
		return json(result);
	} catch (error) {
		console.error(error);
		// throw error;
		return json({ error: 'Failed to create chat' });
	}
};