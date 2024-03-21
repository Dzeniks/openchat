<script>
    import {onMount} from 'svelte';

    onMount(() => {
        
        
    });

</script>


<section>
    <div id="text-section">
        <h1>
            About OpenChat:
        </h1>
    
        <div>
        <h2>
            Innovative AI Technology:
        </h2>
        <span>
            OpenChat stands at the forefront of innovation, harnessing the power of state-of-the-art AI technology. Our chatbot is not just a conversational tool; it's a dynamic learning entity that adapts and evolves with every interaction.
        </span>
        </div>
        
        <div>
            <h2>
                Open Source and Free:
            </h2>
            <span>
                We believe in the democratization of technology. That's why OpenChat is open source, enabling developers and enthusiasts alike to explore, contribute, and customize the platform. Best of all, it's free to use for everyone!
            </span>    
        </div>
        
        <div>
            <h2>
                Constant Learning:
            </h2>
            <span>
                Our commitment to excellence is reflected in OpenChat's continuous learning process. As users engage with the platform, OpenChat refines its responses, ensuring that it stays ahead of the curve in providing accurate and relevant information.
            </span>
        </div>
    
        <div>
            <h2>
                Students project:
            </h2>
            <span>
                Project from Purkyňka students of information technology
            </span>
        </div>
    </div>
    <div id="img-section">
        <img src="/AI-robot.png" alt="AI">
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

    #text-section {
        padding: 20px;
        width: 50%;
        text-align: center;
    }

    #img-section {
        align-items: center;
        background: var(--secondary);
        height: 100vh;
    }

    img {
        max-width: 100%;
        height: 100vh;
        overflow: hidden;
        object-fit: cover;
    }

    @media (max-width: 800px) {
        img {
            width: 100%;
        }

        #img-section {
            display: flex;
            align-items: center;
            justify-content: center;
            background: var(--secondary);
            height: 50vh;
        }

        #text-section {
            width: 100%;
        }

        section {
            margin-top: 15%;
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