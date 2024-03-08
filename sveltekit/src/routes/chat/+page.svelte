<script lang="ts">
    import { onMount } from 'svelte';
    import { writable } from 'svelte/store';
    import { fade, fly } from "svelte/transition"
	import Message from '../../lib/components/Message.svelte';
    import { error } from '@sveltejs/kit';
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


     let apologies = [
        "I'm sorry, I didn't understand that. Could you please rephrase?",
        "Apologies, I'm having trouble grasping that. Can you try rephrasing?",
        "I apologize, but I'm having difficulty understanding. Can you provide more context?",
        "Sorry, it seems I'm not quite getting what you're saying. Could you try saying it differently?",
        "I'm sorry for the confusion. Could you clarify or phrase it differently?",
        "My apologies, it seems I'm not following. Can you try explaining it differently?",
        "I'm sorry if I'm not understanding correctly. Could you try saying it in a different way?",
        "Apologies, it appears I'm not grasping your message. Can you simplify or rephrase it?",
        "Sorry, I'm having trouble understanding. Could you reword that for me?",
        "I apologize, it seems I'm not quite on the same page. Can you try explaining it differently?",
        "Sorry, I had some servers down. What did you want again?",
        "Oops! Looks like I was on a coffee break. Can you repeat that?",
        "My circuits must be on vacation. Mind giving that another shot?",
        "Apologies, it seems my wires got crossed. Can you try explaining it differently?",
        "Whoops! Looks like I tripped over a virtual cable. Could you repeat your request?",
        "Sorry, I think I just hit a digital pothole. Can you say that again?",
        "My virtual hamster took a nap. Can you repeat that?",
        "I think there was a hiccup in my programming. Can you repeat what you said?",
        "Oops! Looks like my circuits got tangled. Can you try saying that again?",
        "Sorry, my virtual dog ate your message. Can you resend it?",
        "It appears my virtual parrot didn't quite catch that. Can you repeat?",
        "I think my digital ducks are out of alignment. Can you try that again?",
        "Sorry, I seem to have misplaced that byte. Can you repeat your request?"
    ]

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
        
        fetch('http://localhost:3000/api/chat/', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ prompt: newPrompt })
        }).then(response => response.json()).then(data => {

            if (data.output != undefined) {
                DATA = [
                    ...DATA,
                    {
                        owner: "OpenAI",
                        date: new Date(),
                        prompt: data.output,
                        role: "openai",
                    }
                ]
            }
            else {
                // Remove last element from DATA
                DATA.pop();

                if (data.error){
                    // Alert error message
                    alert(data.error);
                }
            }
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
