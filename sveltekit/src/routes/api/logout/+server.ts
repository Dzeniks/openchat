import { json } from '@sveltejs/kit';
import type { RequestHandler } from '@sveltejs/kit';


export const POST: RequestHandler = async ({ request, cookies }) => {
	cookies.set('accessToken', '', {
		httpOnly: true,
		maxAge: 0,
		secure: false,
		path: '/',
	});
	cookies.set('refreshToken', '', {
		httpOnly: true,
		maxAge: 0,
		secure: false,
		path: '/',
	});
	return json({ message: 'Logged out' });
};