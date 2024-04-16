<script lang="ts">
    import { fade } from "svelte/transition";
	import { onMount } from "svelte";

    export var owner: string;
    export var prompt: string; // As Markdown   
</script>

<div class="prompt" transition:fade="{{ duration: 500 }}">
    <h3>{owner === "AI" ? owner : "User"}</h3>
    <span class="message-field">
        {@html prompt}
    </span>
</div>

<style>
    .prompt {
        padding: 1rem;
        width: 90%;
        background-color: #f1f1f1;
        border-radius: 1rem;
    }

    .message-field {
        display: block;
        text-align: left;
    }
    
    h3 {
        text-align: left;
    }



    @media (max-width: 768px) {
        .prompt {
            width: 80vw;
        }
    }
</style>
