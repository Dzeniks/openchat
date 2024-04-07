import { _accessTokenRefreshTime, _refreshTokenRefreshTime } from '$lib/stores.js';
import { refresh, type AuthErrorResponse, type AuthResponse } from '$lib/auth/auth';

export const trailingSlash = 'always';
export const ssr = false;
// Load from global storage variable LastTime

// TODO: unite times with login/+page.server.ts 
const accessTokenMaxAge = 3600;
const refreshTokenMaxAge = 3600 * 24 * 30;


export const load = async({ url, cookies }) => {
	const { pathname } = url;

	const accessToken = cookies.get('accessToken');
	const refreshToken = cookies.get('refreshToken');

	let accessTokenRefreshTime;

	let unsubscribe = _accessTokenRefreshTime.subscribe(value => {
		accessTokenRefreshTime = value;
	});
	unsubscribe();

	if (accessTokenRefreshTime !== undefined && refreshToken !== undefined) {
		if (new Date().getTime() > accessTokenRefreshTime) {

			// Refresh the access token
			const authResponse: AuthResponse | AuthErrorResponse= await refresh(refreshToken);

			if (!('accessToken' in authResponse) || !('refreshToken' in authResponse)) {
					return {
						status: 302,
						redirect: '/login'
					};
				}
			else if ('accessToken' in authResponse && 'refreshToken' in authResponse) {
				const data = authResponse;

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
		}
	}

	return {
		pathname,
		accessToken,
		refreshToken
	};
};
