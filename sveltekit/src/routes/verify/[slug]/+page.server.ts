import { error, redirect, type NumericRange } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import { activate } from '$lib/auth/auth';
import type { AuthResponse, AuthErrorResponse } from '$lib/auth/auth';
import { _accessTokenRefreshTime, _refreshTokenRefreshTime } from '$lib/stores';


// TODO: unite times with index/+layout,server.ts 
const accessTokenMaxAge = 3600;
const refreshTokenMaxAge = 3600 * 24 * 30;

export const load: PageServerLoad = async({ params, cookies }) => {
    console.log('verify/[slug]/+page.server.ts load');

    const refreshToken = params.slug;
    if (!refreshToken) {
        return error(400, 'Invalid refresh token');
    }

    const data: AuthResponse | AuthErrorResponse | undefined = await activate(refreshToken);
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
};