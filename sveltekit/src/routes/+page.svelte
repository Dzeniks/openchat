<script>
    import {onMount} from 'svelte';
</script>


<section>
    <div class="text-container">
        <h1>Welcome to OpenChat</h1>
        <div>
            <p>OpenChat is a modern chatbot that uses the latest in AI technology to answer your questions.</p>
            <p>It is open source and free to use.</p>
            <p>It is built with the latest in AI technology and is constantly learning.</p>
        </div>
        <div id="chat-div-btn">
            <a id="chat-btn" href="/chat">
                Chat
            </a>
        </div>
    </div>

    
        <div class="image-container">
            <img src="/AI-hand-1024.png" alt="AI hand">
        </div>
</section>


<style>

    section {
        display: flex;
        flex-direction: row;
        justify-content: flex-end;
        width: 100vw;
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
    }

    .text-container {
        width: 50%;
        text-align: center;
    }

    .image-container {
        align-items: center;
        background: var(--secondary);
        height: 100vh;
    }

    img {
        max-width: 100%;
        height: auto;
    }

    #chat-div-btn{
        margin-top: 20px;
        display: block;
    }

    #chat-btn {
        padding: 16px;
        color: var(--primary);
        background-color: var(--secondary);
        border-radius: 5px;
        transition: background-color 0.3s, color 0.3s, border 0.3s;
    }

    #chat-btn:hover{
        color: var(--secondary);
        background-color: var(--primary);
        border: var(--secondary) solid 2px;
    }

    @media (max-width: 800px) {
        img {
            width: 100%;
        }

        section {
            padding-top: 10vh;
            flex-direction: column;
            align-items: center;

            justify-content: space-between;
            gap: 50px;
        }
    }

    @media (max-width: 1200px) {
        img {
            max-width: 90%;
        }
    }

</style>