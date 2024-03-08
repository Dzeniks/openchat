import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';

const IsDev = process.env.DEV === 'true';

export const POST: RequestHandler = async ({ request }) => {
	try {

        const requestJson = await request.json();
        const prompt = requestJson.prompt;

        if (!prompt) {
            return json({ error: 'Prompt is required' });
        }
        let url = `${process.env.BACKEND_URL}/api/chat/SentPrompt`;
        if(IsDev){
            url = `${process.env.BACKEND_URL}/api/chat/SentPromptTest`
        }
        const response = await fetch(url, {

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