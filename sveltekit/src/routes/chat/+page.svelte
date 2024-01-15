<script>
// @ts-nocheck

	import Prompt from './../../lib/components/prompt.svelte';
    import { onMount } from 'svelte';
    import { writable } from 'svelte/store';
    import { fade } from "svelte/transition"
    // Create a writable store
    export const refreshToken = writable('');
    export const accessToken = writable('');

    onMount(() => {
        
    });

    let DATA = [
        {
            owner: "Frank",
            date: new Date(),
            prompt: "This is message",
            role: "AI",
            loaded: false
        },
        {
            owner: "AI",
            date: new Date(),
            prompt: "This is message NOW",
            role: "User",
            loaded: false
        },
        {
            owner: "RealAI",
            date: new Date(),
            prompt: "This is message NOW",
            role: "AI",
            loaded: false
        },
    ].map(item => ({ ...item, loaded: false }));

    let newPrompt = '';

    function postPrompt() {
        DATA = [...DATA, {
            owner: "TestMeOut",
            date: new Date(),
            prompt: newPrompt,
            role: "User",
            loaded: false
        }]

        newPrompt = ""
        
        console.log(newPrompt);

    }
</script>

<section>
    <div class="chat-window" style="overflow-y: scroll;">
        {#key DATA}
        <div in:fade={{ duration: 300, delay: item.loaded ? 0 : 200 * index }}>
            {#each DATA as item, index}
                {#if !item.loaded}
                    <Prompt
                        owner={item.owner}
                        date={new Date(item.date)}
                        prompt={item.prompt}
                        role={item.role}
                        loaded={item.loaded}
                    />
                {/if}
            {/each}
        </div>
        {/key}
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
        border: 1px solid #ccc;
        border-radius: 4px;
        resize: vertical; /* Allow vertical resizing of the textarea */
    }

    .input-div {
        padding: 2rem;
    }

    input {
        padding: 2rem;
        display: flex;
        flex-direction: column;
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
        border: 2px red solid   ;

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