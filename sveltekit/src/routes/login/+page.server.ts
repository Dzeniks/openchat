import type { Actions } from './$types';
import type { AuthResponse, AuthErrorResponse } from '$lib/auth/auth';
import { login, register } from '$lib/auth/auth';
import { redirect } from '@sveltejs/kit';

export const actions = {
    	login: async ({cookies, request}) => {
            const form = await request.formData();
            const email = form.get('email') as string;
            const password = form.get('password') as string;
            if (!email || !password) {
                return {status: 400, error: 'No email or password', message: "Failed to register"} as AuthErrorResponse;
            }
            const data: AuthResponse | AuthErrorResponse | undefined = await login(email, password);
            if (data) {
                if ('accessToken' in data && 'refreshToken' in data){
                    cookies.set('accessToken', data.accessToken, {
                        httpOnly: true,
                        maxAge: 3600,
                        path: '/',
                    });

                    cookies.set('refreshToken', data.refreshToken, {
                        httpOnly: true,
                        maxAge: 3600 * 24 * 30,
                        path: '/',
                    });
                }
                console.log(data);
                if ('error' in data){
                    throw redirect(302, `/login?error=${data.error}&message=${data.message}`)
                }
            }
            throw redirect(302, '/chat')
    },
    register : async ({cookies, request}) => {
          const form = await request.formData();
          const email = form.get('email') as string;
          const password = form.get('password') as string;
          if (!email || !password) {
              console.log('No email or password');
              return {status: 400, error: 'No email or password', message: "Failed to register"} as AuthErrorResponse;
          };
          const data : AuthResponse | AuthErrorResponse | undefined  = await register(email, password);
          if (data) {
                if ('accessToken' in data && 'refreshToken' in data){
                    cookies.set('accessToken', data.accessToken, {
                        httpOnly: true,
                        maxAge: 3600,
                        path: '/',
                    });
                    cookies.set('refreshToken', data.refreshToken, {
                        httpOnly: true,
                        maxAge: 3600 * 24 * 30,
                        path: '/',
                    });
                }
                if ('error' in data){
                    throw redirect(301, `/login?error=${data.error}&message=${data.message}`)
                }
            }
            throw redirect(302, '/chat')
        },
} satisfies Actions;