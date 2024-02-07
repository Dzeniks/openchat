import type { Handle } from '@sveltejs/kit';
import { refresh, type AuthResponse, type AuthErrorResponse } from '$lib/auth/auth';
import dotenv from 'dotenv';

dotenv.config();

const authorizedURLs = ['/chat', '/admin'];

const IsDev = process.env.DEV === 'true';

export const handle: Handle = async ({ event, resolve }) => {
    const { pathname } = event.url;
    const accessToken = event.cookies.get('accessToken');
    const refreshToken = event.cookies.get('refreshToken');
    let hasPrivilege = false;
    
    if (IsDev){
        hasPrivilege = true;
    }

    if (authorizedURLs.includes(pathname)) {
        if (accessToken){
            hasPrivilege = await auth(accessToken);
        }
        if (accessToken && refreshToken) {
            const data : AuthResponse | AuthErrorResponse | undefined = await refresh(refreshToken)
            if (data){
                if ('accessToken' in data && 'refreshToken' in data) {
                    event.cookies.set('accessToken', data.accessToken, {
                        httpOnly: true,
                        maxAge: 3600,
                        path: '/',
                    });
                    event.cookies.set('refreshToken', data.refreshToken, {
                        httpOnly: true,
                        maxAge: 3600 * 24 * 30,
                        path: '/',
                    });

                    hasPrivilege = await auth(data.accessToken);
                }
            }
        }
        
        if (!hasPrivilege) {
            return new Response('Unauthorized', {
                status: 302,
                headers: {
                    Location: '/login', // replace with your login page URL
                },
            });    
        }
    }
    const response = await resolve(event);
    return response;
};

async function auth(accessToken: string): Promise<boolean> {
    const response = await fetch('http://localhost:8080/api/auth/', {
        method: 'POST',
        headers: {
            Authorization: `${accessToken}`
        }
    });
    return response.ok;
}

