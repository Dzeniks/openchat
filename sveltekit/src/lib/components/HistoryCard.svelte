<script lang="ts">
	import { goto } from "$app/navigation";
    import { fade } from "svelte/transition";
    export var displayString: string;
    export var id: string;

    function redirect_to_chat() {
        goto(`/chat/${id}`)
	}

</script>

<button on:click={redirect_to_chat} transition:fade="{{ duration: 2000 }}">
    {displayString}
</button>


<style>
    button {
        padding: 1rem;
        width: 100%;
        background-color: #f1f1f1;
        text-align: left;
    }

    button:hover {
        background-color: var(--tertiary);
    }
</style>
