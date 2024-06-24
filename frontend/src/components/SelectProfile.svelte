<script lang="ts">
	import { Button } from "bits-ui";
	import { onMount } from "svelte";
	import { GetUsers } from "wailsjs/go/main/App";

	let users: Awaited<ReturnType<typeof GetUsers>> = [];

	onMount(() => {
		GetUsers()
			.then((data) => {
				if (data) users = data;
			})
			.catch((err) => {
				console.error(err);
			});
	});
</script>

<Button.Root
	on:click={() => {
		console.log("clicked");
	}}
	class="inline-flex h-12 items-center justify-center rounded-input bg-dark
px-[21px] text-[15px] font-semibold text-background shadow-mini
hover:bg-dark/95 active:scale-98 active:transition-all"
>
	FUCK
</Button.Root>
{#each users as user}
	<Button.Root
		class="inline-flex h-12 items-center justify-center rounded-input bg-dark
px-[21px] text-[15px] font-semibold text-background shadow-mini
hover:bg-dark/95 active:scale-98 active:transition-all"
	>
		{user?.name}
	</Button.Root>
{/each}
