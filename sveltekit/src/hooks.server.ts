import { type Handle, redirect } from '@sveltejs/kit';
import { type AuthErrorResponse, type AuthResponse, refresh, auth } from '$lib/auth/auth';
import dotenv from 'dotenv';

dotenv.config();

const authorizedURLs = ['/chat/'];

export const handle: Handle = async ({ event, resolve }) => {
    const { pathname } = event.url;
    const accessToken = event.cookies.get('accessToken');
    let hasPrivilege = false;

    // Check if in pathname is some part of authorizedURL
    if (pathname.includes('/chat')) {
        if (accessToken){
            hasPrivilege = await auth(accessToken);
        }
        if (!hasPrivilege) {
            redirect(302, '/login');
        }
    }
    return resolve(event);
};



