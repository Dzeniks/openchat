<script lang="ts">
	import { page } from '$app/stores';
</script>

<section>
	<h1>{$page.status}</h1>
	<h1>{$page.error.message}</h1>
</section>

<style>
	section {
		text-align: center;
		margin-top: 100px;
	}
</style>