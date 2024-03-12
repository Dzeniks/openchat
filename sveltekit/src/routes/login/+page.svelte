<script>
	import { onMount } from 'svelte';
  import { slide } from 'svelte/transition';

  let email = '';
  let password = '';
  let rePassword = '';
  let IsLogin = true;

	onMount(() => {
		console.log('onMount');
    console.log(IsLogin);
		const urlSearchParams = new URLSearchParams(window.location.search);

		// Access individual parameters
		const error = urlSearchParams.get('error');
		const message = urlSearchParams.get('message');

		if (error && message) {
			alert(error + ': ' + message);
		}
	});


	const change = () => {
		IsLogin = !IsLogin;
	};
</script>

<section>
    <h1>Login/register:</h1>
    <div id="form-window">
            {#if !IsLogin}
                <div transition:slide>
                    <button on:click={change}>Are you already registered ?</button>
                    <form method="POST" action="?/register">
                        <input type="text" placeholder="email" name="email" bind:value={email}>
                        <input type="password" placeholder="Password" name="password" bind:value={password}>
                        <input type="password" placeholder="Confirm Password" bind:value={rePassword}>
                        <button type="submit">Register</button>
                    </form>
                </div>
            {:else}
                <div transition:slide>
                    <button on:click={change}>Create new account</button>
                    <form method="POST" action="?/login">
                        <input type="text" placeholder="email" name="email" bind:value={email}>
                        <input type="password" placeholder="Password" name="password" bind:value={password}>
                        <button type="submit">Login</button>
                    </form>
                </div>
            {/if}
    </div>
</section>


<style>

    form {
        display: flex;
        flex-direction: column;
        gap: 10px;
    }


    #form-window{
        display: flex;
        flex-direction: column;
        text-align: center;
        align-items: center;
    }

    section {
        padding: 10vh 5vw 0 5vw;

        display: flex;
        flex-direction: column;
        text-align: center;
        align-items: center;
        background-color: var(--secondary);
        font-size: 16px;

        width: 60%;
        height: 90vh;
    }

    h1 {
        text-align: center;
        font-size: 3rem;
        font-weight: 900;
    }

    div {
        display: flex;
        flex-direction: column;
        width: 50%;
        justify-content: space-between;
    }

    @media (max-width: 800px) {

        #form-window{
            display: flex;
            text-align: center;
            justify-items: center;
            justify-content: space-between;
            flex-direction: column;
            gap: 20px;
        }

        section {
            padding-top: 35%;
            flex-direction: column;
            align-items: center;

            justify-content: space-between;
            gap: 50px;
        }
    }


</style>