import type { Actions } from './$types';
import { login, register } from '$lib/auth/auth';
import { redirect } from '@sveltejs/kit';

export const actions = {
    	login: async ({cookies, request}) => {	
            const form = await request.formData();
            const email = form.get('email') as string;
            const password = form.get('password') as string;
            if (!email || !password) {
                return;
            }
            const data = await login(email, password);
            if (data) {
                cookies.set('accessToken', data.accessToken, {
                    httpOnly: true,
                    maxAge: 1,
                    path: '/',
                });

                cookies.set('refreshToken', data.refreshToken, {
                    httpOnly: true,
                    maxAge: 3600 * 24 * 30,
                    path: '/',
                });

            }
            throw redirect(302, '/chat')
    },
    register : async ({cookies, request}) => {		
        const form = await request.formData();
            const email = form.get('email') as string;
            const password = form.get('password') as string;
            if (!email || !password) {
                return;
            }
            const data = await register(email, password);
            if (data) {
                cookies.set('accessToken', data.accessToken, {
                    httpOnly: true,
                    maxAge: 1,
                    path: '/',
                });

                cookies.set('refreshToken', data.refreshToken, {
                    httpOnly: true,
                    maxAge: 3600 * 24 * 30,
                    path: '/',
                });

            }
            throw redirect(302, '/chat')
        },
} satisfies Actions;