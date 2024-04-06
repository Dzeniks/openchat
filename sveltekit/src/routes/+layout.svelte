<script lang="ts">
    import Header from '$lib/components/Header.svelte';
	import { fade } from 'svelte/transition';
	import type { LayoutData } from './$types';
	import { onMount } from 'svelte';

    export let data: LayoutData;

    onMount(() => {
        console.log('onMount Layout');
        const cookies = document.cookie;
        console.log(cookies);
    });
</script>


<Header/>
{#key data.pathname}
		<main transition:fade="{{duration: 200}}">
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