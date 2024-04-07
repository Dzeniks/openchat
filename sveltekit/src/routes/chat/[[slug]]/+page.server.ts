import { error, redirect, type NumericRange } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import type { Chat } from '$lib/types';

const PORT = process.env.PORT;

export const load: PageServerLoad = async({ fetch, params, cookies }) => {

    const createChat = async() => {
        const response = await fetch(`http://localhost:${PORT}/api/chat/create`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': cookies.get('accessToken') as string
            }})
        if (response.ok) {
            const data = await response.json();
            if (data.chat_id === undefined || data.chat_id === null) {
                return redirect(302, '/chat');
            }
            return redirect(302, `/chat/${data.chat_id}`);
        }
    }
    
    const getChatByID = async(id: string) => {

		const response = await fetch(`http://localhost:${PORT}/api/chat/get`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				'Authorization': cookies.get('accessToken') as string
			},
			body: JSON.stringify({ chat_id: id })
		})
        if (response.ok) {
            const data = await response.json();
            return data.chat as Chat;
        } else {
            error(response.status as NumericRange<400, 599>,response.statusText);
        }
	};

    const getChatIdsOfUser = async() => {
        const response = await fetch(`http://localhost:${PORT}/api/chats/get`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': cookies.get('accessToken') as string
            }
        })
        if (response.ok) {
            const data = await response.json();
            if (data.chat_ids === undefined || data.chat_ids === null || data.chat_ids.length === 0) {
                return [];
            }
            let chat_ids = [];
            for (let i = 0; i < data.chat_ids.length; i++) {
                chat_ids.push({chat_id: data.chat_ids[i], displayString: `${new Date(data.dates[i]).toLocaleDateString()} - ${new Date(data.dates[i]).toLocaleTimeString()}`});
            }
            return chat_ids;
        }
    }

    if (params.slug === '' || params.slug === undefined || params.slug === null) {
        const ok = await createChat();
        return ok;
    }
    else {
        const DATA = await getChatByID(params.slug as string);
        const chat_ids = await getChatIdsOfUser();
        return {
            chat: DATA,
            history: chat_ids
        };
    }
};