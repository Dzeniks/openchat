import type { Actions } from './$types';
import type { AuthResponse, AuthErrorResponse } from '$lib/auth/auth';
import { login, register } from '$lib/auth/auth';
import { redirect } from '@sveltejs/kit';

export const actions = {
    chat: async ({ cookies, request }) => {
        // const { token } = cookies.get('accessToken');
        // console.log

        // Make request to localhost:8080/api/ChatCompletition
        const response = await fetch('http://localhost:8080/api/chat/ChatCompletition', {
            method: 'POST',
            headers: {
                // Authorization: `${token}`,
            },
        });

        // Handle the response here
        console.log(response);
    },
    }