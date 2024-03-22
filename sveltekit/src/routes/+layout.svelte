<script lang="ts">
    import Header from '$lib/components/Header.svelte';
	import { fade, scale, slide, fly, draw, blur } from 'svelte/transition';

    export let data;
</script>


<Header/>

{#key data.pathname}
	<main transition:fade="{{delay: 50, duration: 1000}}">
		<slot />
    </main>
{/key}

<style>
    @import url('https://fonts.googleapis.com/css2?family=Oxanium:wght@400;700&display=swap');

    :root {
        font-family: 'Oxanium', sans-serif;

        --primary: #fff;
        --secondary: #484D54;
        --tertiary: #85DDFB;
    }

    main {
            display: flex;
            flex-direction: column;
            justify-content: space-between;
            align-items: center;
            margin: 0;
    }
</style>