import type { Actions } from './$types';
import type { AuthResponse, AuthErrorResponse } from '$lib/auth/auth';
import { login, register } from '$lib/auth/auth';
import { redirect } from '@sveltejs/kit';
import { _accessTokenRefreshTime, _refreshTokenRefreshTime } from '$lib/stores';

// TODO: unite times with index/+layout,server.ts 
const accessTokenMaxAge = 3600;
const refreshTokenMaxAge = 3600 * 24 * 30;


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
                    maxAge: accessTokenMaxAge,
                    secure: false,
                    path: '/',
                });

                _accessTokenRefreshTime.set(new Date().getTime() + Math.round(accessTokenMaxAge/2) * 1000);
        
                cookies.set('refreshToken', data.refreshToken, {
                    httpOnly: true,
                    maxAge: refreshTokenMaxAge,
                    secure: false,
                    path: '/',
                });
        
                _refreshTokenRefreshTime.set(new Date().getTime() + Math.round(refreshTokenMaxAge/2) * 1000);
            }
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
            return {status: 400, error: 'No email or password', message: "Failed to register"} as AuthErrorResponse;
        };
        const data : AuthResponse | AuthErrorResponse | undefined  = await register(email, password);
        if (data) {
                if ('accessToken' in data && 'refreshToken' in data){
                    cookies.set('accessToken', data.accessToken, {
                        httpOnly: true,
                        maxAge: accessTokenMaxAge,
                        secure: false,
                        path: '/',
                    });
                    cookies.set('refreshToken', data.refreshToken, {
                        httpOnly: true,
                        maxAge: refreshTokenMaxAge,
                        secure: false,
                        path: '/',
                    });
                }
                if ('error' in data){
                    throw redirect(301, `/login?error=${data.error}&message=${data.message}`)
                }
            }
            throw redirect(302, '/login')
        },
} satisfies Actions;