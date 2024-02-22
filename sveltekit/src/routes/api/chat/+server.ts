import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';


export const POST: RequestHandler = async ({ request }) => {
	try {

        const requestJson = await request.json();
        const prompt = requestJson.prompt;

        if (!prompt) {
            return json({ error: 'Prompt is required' });
        }
        const response = await fetch(`${process.env.BACKEND_URL}/api/chat/ChatCompletetion`, {

            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({"prompt": prompt})
        });

        if (!response.ok) {
            return json({ error: 'Failed to fetch completions' });
        }

        const result = await response.json();
        return json(result);
    } catch (error) {
        console.error(error);
        throw error;
    }
};