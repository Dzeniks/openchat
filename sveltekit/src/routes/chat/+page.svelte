<script lang="ts">
	import { fade } from 'svelte/transition';
	import Message from '../../lib/components/Message.svelte';
	import { onMount } from 'svelte';

	let chatID = ""

	const getChatID = () => {
		fetch('http://localhost:3000/api/chat/get', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				'Authorization': localStorage.getItem('accessToken') as string
			},
		}).then(response => response.json()).then(data => {
			if (data.chat_id != undefined) {
				console.log(data.chat_id);
				chatID = data.chat_id
			} else {
				if (data.error) {
					alert(data.error);
				}
			}
		}).catch(error => {
			console.error('Error:', error);
		});


	};
	let loaded = false;

	// OnMount make request to server to get chat history
	onMount(() => {
		getChatID()
		loaded = true
		// You can perform any initialization logic here
		// For example, fetching data from an API
		// Updating component state, etc.
	});



	type DataItem = {
		owner: string,
		date: Date,
		prompt: string,
		role: string,
	};

	let DATA: DataItem[] = [];

	let newPrompt = '';
	let isDisabled = false;
	const postPrompt = () => {
		isDisabled = true;



		DATA = [
			...DATA,
			{
				owner: 'User',
				date: new Date(),
				prompt: newPrompt,
				role: 'user'
			}
		];

		fetch('http://localhost:3000/api/chat/sent', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				'Authorization': localStorage.getItem('accessToken') as string
			},
			body: JSON.stringify({chat_id: chatID , prompt: newPrompt })
		}).then(response => response.json()).then(data => {
			isDisabled = false;
			if (data.output != undefined) {
				DATA = [
					...DATA,
					{
						owner: 'OpenAI',
						date: new Date(),
						prompt: data.output,
						role: 'openai'
					}
				];
			} else {
				DATA = DATA.slice(0, -1);
				if (data.error) {
					alert(data.error);
				}
			}
		}).catch(error => {
			isDisabled = false;

			console.error('Error:', error);
		});
	};

</script>

<section>
	<div class="chat-window" style="overflow-y: scroll;">
		{#key loaded}
					<h1 in:fade="{{ duration: 2000 }}" style="color: var(--primary)">OpenChat-Beta</h1>
		{/key}
		{#each DATA as item}
			<Message
				owner={item.owner}
				prompt={item.prompt}
				role={item.role}
			/>
		{/each}
	</div>
	<div class="input-div">
		<textarea id="prompt-input" placeholder="Write me prompt :)" bind:value={newPrompt} />
		<button on:click={postPrompt} disabled={isDisabled}>Post Data</button>
	</div>
</section>

<style>
    textarea {
        width: 100%;
        padding: 8px;
        margin-bottom: 16px;
        box-sizing: border-box;
        border-radius: 4px;
        resize: vertical;
    }

    .input-div {
        padding: 2rem;
    }

    section {
        padding-top: 10vh;
        display: flex;
        flex-direction: column;
        text-align: center;
        align-items: center;
        width: 80vw;
        height: 90vh;
        background-color: var(--secondary);

        justify-content: space-between;
        gap: 50px;

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
    }

    div {
        display: flex;
        flex-direction: column;
        width: 50%;
        justify-content: space-between;
    }

    @media (max-width: 800px) {

        section {
            /*padding-top: 15%;*/
            flex-direction: column;
            align-items: center;

            justify-content: space-between;
            gap: 50px;
        }
    }

    .chat-window {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: flex-start;
        gap: 5px;

        width: 100%;
        height: 90%;
        margin: 0;
    }

</style>
