<script lang="ts">
    import { fade } from "svelte/transition";
    export var owner: string;
    export var prompt: string;
</script>

<div class="prompt" transition:fade="{{ duration: 500 }}">
    <h3>{owner === "AI" ? owner : "User"}</h3>
    <span>
        {prompt}
    </span>
</div>

<style>
    .prompt {
        padding: 1rem;
        width: 90%;
        background-color: #f1f1f1;
        border-radius: 1rem;
    }

    
    span {
        text-align: right;
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
