<script>
    import {onMount} from 'svelte';

    // const refreshToken = null;

    onMount(() => {
        console.log("onMount");
        const urlSearchParams = new URLSearchParams(window.location.search);
        
        // Access individual parameters
        const error = urlSearchParams.get('error');
        const message = urlSearchParams.get('message');

        if (error && message) {
            alert(error + ": " +message);
        }
    });
    let email = "";
    let password = "";
    let rePassword = "";


    let IsLogin = true;

    const change = () => {
        IsLogin = !IsLogin;
    }
</script>


<section>
    <h1>
        Login/register:
    </h1>
    <div>
        
        {#if !IsLogin}
        <form method="POST" action="?/register">
            <button on:click={change}>Are you already registered ?</button>
            <input type="text" placeholder="email" name="email" bind:value={email}>
            <input type="password" placeholder="Password" name="password" bind:value={password}>
            <input type="password" placeholder="Password" bind:value={rePassword}>
            <button>Register</button>
        </form>
        {/if}

        {#if IsLogin}
        <form method="POST" action="?/login">
            <button on:click={change}>Create new accounts</button>
            <input type="text" placeholder="email" name="email" bind:value={email}>
            <input type="password" placeholder="Password" name="password" bind:value={password}>
            <button>Login</button>
        </form>
        {/if}
    </div>    
</section>


<style>

    section {
        padding: 10vh 5vw;

        display: flex;
        flex-direction: column;
        text-align: center;
        align-items: center;
        width: 60vw;
        background-color: var(--secondary);

        justify-content: space-between;
        gap: 50px;

        font-size: 16px;
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


</style>