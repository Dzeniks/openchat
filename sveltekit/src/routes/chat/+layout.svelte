<script lang="ts">
    import ChatSideBar from '$lib/components/chatSideBar.svelte';
</script>

<main>
    <ChatSideBar/>
    <div>
        <slot/>
    </div>
</main>

<style>
    main {
        padding: 0;
        margin: 0;
        display: flex;
        flex-direction: row;
        text-align: center;
        width: 100vw;
        height: 100%;
    }
</style>
