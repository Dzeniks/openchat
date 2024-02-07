<script lang="ts">
    import { fade } from "svelte/transition";
    export var owner: string;
    export var date: Date;
    export var prompt: string;
    export var role: string;
</script>

<div class="prompt" transition:fade="{{ duration: 600 }}">
    <h3  class="{role === 'user' ? 'right' : 'left'}">{owner}</h3>
    <span>
        {prompt}
    </span>
    <!-- <p class="{role === 'user' ? 'right' : 'left'}"><small>{date}</small></p> -->
</div>


<style>
    
    
    .prompt {
        padding: 1rem;
        width: 40vw;
        background-color: #f1f1f1;
        border-radius: 1rem;
    }

    .right {
        text-align: right;
    }
    
    .left {
        text-align: left;
    }

    span {
        text-align: center;
        
    }

    @media (max-width: 768px) {
        .prompt {
            width: 80vw;
        }
    }
</style>
