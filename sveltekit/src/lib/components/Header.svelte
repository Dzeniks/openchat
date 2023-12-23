<nav>
    <a href="/">Home</a>
    <ul class="list">
        <li><a href="/about">About</a></li>
        <li><a href="/projects">Projects</a></li>
        <li><a href="/contacts">Contact</a></li>
    </ul>
    <!-- <button class="menu">Menu</button> -->
</nav>

<style>
    nav {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        padding: 0px 32px;
        display: flex;
        align-items: center;
        transition: 0.3s ease-out;
        backdrop-filter: blur(12px) brightness(1.2);
        -webkit-backdrop-filter: blur(12px) brightness(1.2);
        text-shadow: 0 0 5px rgba(0, 0, 0, 0.5);
        color: white;
        font-size: 24px;
        font-weight: 500;
        color: black;
    }

    @media (min-width: 640px) {
        nav {
            padding: 0px 16x;
        }
    }

    a {
        color: inherit;
        text-decoration: none;
    }
    a:hover, a:focus {
        text-decoration: underline;
    }

    .list {
        list-style-type: none;
    }
    @media (min-width: 640px) {
        .list {
            display: flex;
        }
    }
    .list li {
        margin-left: 20px;
    }

</style>


<script>
    let isListVisible = false;

    function toggleList() {
        isListVisible = !isListVisible;
    }
</script>