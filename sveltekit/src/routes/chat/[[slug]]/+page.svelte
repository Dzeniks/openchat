<script lang="ts">
	import type { Chat } from '$lib/types';
	import type { PageData } from './$types';
	import MessageCard from '$lib/components/MessageCard.svelte';
	import { fade } from 'svelte/transition';
	import { onMount } from 'svelte';
	import { error } from '@sveltejs/kit';
	import { afterUpdate, tick } from 'svelte';
	import HistoryCard from '$lib/components/HistoryCard.svelte';
    import { marked } from 'marked';
	
	export let data: PageData;

	let loaded = false;

	let element: Element;

	let DATA: Chat = data.chat as Chat;
	let HISTORY: { chat_id: string; displayString: string; }[] = data.history || [];
	let chatWindow: Element | null = null;

	// Add create new chat to history
	HISTORY = [
		...HISTORY,
		{
			chat_id: "",
			displayString: "Create new chat"
		}
	];

	// Reverse the history
	HISTORY = HISTORY.reverse();

	let newPrompt = '';
	let isDisabled = false;

	let hostname: string;
	let port: string;


	onMount(() => {
		loaded = true

		hostname = window.location.hostname;
		port = window.location.port;

		chatWindow = document.querySelector('.chat-window')
	});

	
	if (data.chat === undefined) {
		error(404, 'Chat not found');
	}


	const scrollToBottom = async (node: Element) => {
    	node.scroll({ top: node.scrollHeight, behavior: 'smooth' });
  	}; 

	afterUpdate(() => {
		if(DATA) scrollToBottom(element);
  	});

	$: if(DATA && element && chatWindow) {
		scrollToBottom(chatWindow);
	}

	const postPrompt = () => {
		newPrompt = newPrompt.trim().replace(/\n/g, "<br>").replace(/\t/g, "&emsp;");
		isDisabled = true;
		DATA.Messages = [
			...DATA.Messages,
			{
				SenderID: "User",
				Content: newPrompt,
				SentAt: new Date()
			} 
		];

		// Loading the prompt
		DATA.Messages = [
			...DATA.Messages,
			{
				SenderID: "AI",
				Content: "Thinking...",
				SentAt: new Date(0)
			} 
		];

		fetch(`http://${hostname}:${port}/api/chat/sent`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				'Authorization': localStorage.getItem('accessToken') as string
			},
			body: JSON.stringify({chat_id: DATA.ChatID , prompt: newPrompt })
		}).then(response => response.json()).then(async data => {
			isDisabled = false;
			if (data.output != undefined) {
				DATA.Messages[DATA.Messages.length - 1] = {
					SenderID: "AI",
					Content: await marked(data.output),
					SentAt: new Date()
				};
			} else if (data.error) {
					alert(data.error);
					DATA.Messages = DATA.Messages.slice(0, DATA.Messages.length - 2);
			}
			newPrompt = '';
		}).catch(error => {
			DATA.Messages = DATA.Messages.slice(0, DATA.Messages.length - 2);
			isDisabled = false;
			console.error('Error:', error);
		});
	};
</script>

<main>
	<aside>
    	<div id="history" style="overflow-y: scroll;">
			{#each HISTORY as item}
				<HistoryCard displayString={item.displayString} id={item.chat_id} />
			{/each}
    	</div>
	</aside>
	<section>
	{#key loaded }
		<h1 in:fade="{{ delay:100 ,duration: 2000 }}" style="color: var(--primary)">OpenChat</h1>
		<div bind:this={element} id="chat-window" style="overflow-y: scroll;" in:fade="{{ delay:100 ,duration: 2000 }}">
			{#each DATA.Messages as item}
				<MessageCard 
					owner={item.SenderID}
					prompt={item.Content}
				/>
			{/each}
		</div>
		<div id="input-div">
			<textarea id="prompt-input" placeholder="Write me prompt :)" bind:value={newPrompt} />
			<button id="send-btn" on:click={postPrompt} disabled={isDisabled}>Send prompt</button>
		</div>
	{/key}

	</section>
</main>

<style>

    textarea {
        width: 100%;
        padding: 8px;
        margin-bottom: 16px;
        box-sizing: border-box;
        border-radius: 4px;
        resize: vertical;
    }

	#prompt-input {
		height: 50%;
		width: 100%;
	}

    #input-div {
		display: flex;
		flex-direction: column;
		align-items: center;
        padding: 2rem;
		width: 80%;
		height: 15vh;
    }

	button {
		background-color: var(--primary);
		padding: 16px 20px;
		border: none;
		border-radius: 4px;
		cursor: pointer;
		width: 20%;
	}

	button:hover {
		background-color: var(--tertiary);
	}

    section {
        padding-top: 8vh;
        display: flex;
        flex-direction: column;
        text-align: center;
        align-items: center;
        width: 80vw;
        height: 92vh;
        background-color: var(--secondary);

        justify-content: space-between;
        font-size: 16px;
    }


    @media (max-width: 768px) {
        section {
            width: 100vw;
        }
    }

    h1 {
        text-align: center;
        font-size: 3rem;
        font-weight: 900;
		padding: 0;
		margin: 1rem;
    }

    div {
        display: flex;
        flex-direction: column;
        width: 50%;
        justify-content: space-between;
    }

    @media (max-width: 800px) {
        section {
            flex-direction: column;
            align-items: center;

            justify-content: space-between;
        }
    }

    #chat-window {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: flex-start;
        gap: 5px;

        width: 100%;
		height: 100%;
    }

	aside {
        padding-top: 8vh;
        display: flex;
        flex-direction: column;
        text-align: center;
        align-items: center;
        width: 20vw;
        height: 92vh;
        background-color: var(--primary);

        font-size: 16px;

		overflow: auto;
        direction: rtl; 
	}

    /* Track */
    aside::-webkit-scrollbar-track {
        background: #f1f1f1;
    }

    /* Handle */
    aside::-webkit-scrollbar-thumb {
        background: #888;
    }

    /* Handle on hover */
    aside::-webkit-scrollbar-thumb:hover {
        background: #555;
    }

	#history {
		justify-content: flex-start;

		overflow-y: scroll;
		width: 100%;
		height: 100%;
	}

	@media (max-width: 768px) {
        aside {
			display: none;
		}

		section {
			padding-top: 15vh;
			height: 85vh;
			width: 100vw;
		}

		#input-div {
			width: 90%;
			height: 20vh;
			padding: 1rem;

			justify-content: space-between;
			gap: 0px;
		}

		#prompt-input {
			height: 50%;
			width: 100%;
		}

		button {
			width: 30%;
			padding: 0.5rem;
		}
    }

    main {
        padding: 0;
        margin: 0;
        display: flex;
        flex-direction: row;
		justify-content: space-between;


		width: 100vw;
		height: 100%;
    }

</style>
