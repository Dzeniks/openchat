<script lang="ts">
    export var owner: string;
    export var date: Date;
    export var prompt: string;
    export var role: string;
</script>

<div class="prompt {role === 'User' ? 'right' : 'left'}">
    <h3>{owner}</h3>
    <span>
        {prompt}
    </span>
    <p><small>{date}</small></p>
</div>

<style>
    .prompt {
        padding: 1rem;
        width: 40vw;
        background-color: #f1f1f1;
        border-radius: 1rem;
    }

    .right {
        text-align: right;
    }
    
    .left {
        text-align: left;
    }

    

    @media (max-width: 768px) {
        .prompt {
            width: 80vw;
        }
    }
</style>
