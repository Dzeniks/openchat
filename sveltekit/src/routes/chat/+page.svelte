<script>
    import {onMount} from 'svelte';
    import { writable } from 'svelte/store';

    // Create a writable store
    export const refreshToken = writable('');
    export const accessToken = writable('');

    onMount(() => {
        
    });

    let IsLogin = true;

    const login = () => {
        console.log("login");
        IsLogin = !IsLogin;
    }

    const register = () => {
        console.log("register");
        IsLogin = !IsLogin;
    }

</script>


<section>
    <h1>
        Chat
    </h1>
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