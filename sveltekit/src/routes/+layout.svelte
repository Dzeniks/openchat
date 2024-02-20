<script lang="ts">
    import { onMount, onDestroy } from 'svelte';
    import Header from '$lib/components/Header.svelte';
    import Footer from "$lib/components/Footer.svelte";

    onMount(() => {
        console.log("onMount");
        
    });

</script>


<body>
<Header/>
<main>
    <slot/>
</main>
<!-- <Footer/> -->
</body>

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
    }

    body {
        margin: 0;

    }
</style>