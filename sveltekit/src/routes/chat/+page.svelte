<script lang="ts">
    import { onMount } from 'svelte';
    import { writable } from 'svelte/store';
    import { fade, fly } from "svelte/transition"
	import Message from '../../lib/components/Message.svelte';
    // Create a writable store
    export const refreshToken = writable('');
    export const accessToken = writable('');

    type DataItem = {
        owner: string,
        date: Date,
        prompt: string,
        role: string,
    };
    

    let DATA: DataItem[] = [
    ]

    let newPrompt = '';

    const postPrompt = () => {
        DATA = [
            ...DATA,
                {
                    owner: "User",
                    date: new Date(),
                    prompt: newPrompt,
                    role: "user",
                }
            ]
            
        
        // Make request to localhost:8080/api/ChatCompletition
        const response = fetch('http://localhost:5173/api/chat/', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ prompt: newPrompt })
        }).then(response => response.json()).then(data => {
            
            DATA = [
                ...DATA,
                {
                    owner: "AI",
                    date: new Date(),
                    prompt: data?.output,
                    role: "ai",
                }
            ]
        }).catch(error => {
            console.error('Error:', error);

        }).catch(error => {
            console.error('Error:', error);
        })
    }  
        
</script>

<section>
    <div class="chat-window" style="overflow-y: scroll;">
        <h1 in:fade="{{ duration: 1000 }}">OpenChat-Beta</h1>
        {#each DATA as item}
            <Message
                owner={item.owner}
                date={new Date(item.date)}
                prompt={item.prompt}
                role={item.role}
            />
        {/each}
    </div>
    <div class="input-div">
        <textarea id="prompt-input" placeholder="Write me prompt :)" bind:value={newPrompt} />
        <button on:click={postPrompt}>Post Data</button>
    </div>
</section>

<style>
    textarea {
        width: 100%;
        padding: 8px;
        margin-bottom: 16px;
        box-sizing: border-box;
        border-radius: 4px;
        resize: vertical; /* Allow vertical resizing of the textarea */
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

    h1{
        text-align: center;
        font-size: 3rem;
        font-weight: 900;
    }

    div {
        display: flex;
        flex-direction: column;
        justify-content: center;
        
        width: 50%;

        justify-content: space-between;
        grid-area: 20px;
    }

    @media (max-width: 800px) {

        section {
            padding-top: 35%;
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
