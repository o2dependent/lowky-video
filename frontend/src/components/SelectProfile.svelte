<script lang="ts">
	import { profiles } from "./../stores/profiles.ts";
	import { Avatar, Button } from "bits-ui";
	import { Plus } from "phosphor-svelte";
	import { onMount } from "svelte";
	import { GetProfiles } from "wailsjs/go/main/App";
	import AddProfileForm from "./AddProfileDialog.svelte";
	import AddProfileDialog from "./AddProfileDialog.svelte";

	const getAbbr = (profile: { Name: string }) =>
		profile?.Name?.split(" ")
			?.slice(0, 2)
			?.reduce((acc, v) => acc + v?.[0]?.toUpperCase(), "");
</script>

<div class="flex gap-2">
	{#each $profiles as profile}
		<Button.Root
			class="inline-flex items-center justify-center rounded-input bg-dark/0
	p-2 text-[15px] text-dark shadow-none hover:shadow-mini
	hover:bg-dark/15 active:scale-98 active:transition-all w-24"
		>
			<div class="flex flex-col gap-2 items-center justify-center">
				<Avatar.Root
					class="h-12 w-12 rounded-full border border-foreground bg-muted text-[17px] font-medium uppercase text-muted-foreground"
				>
					<div
						class="flex h-full w-full items-center justify-center overflow-hidden rounded-full border-2 border-transparent"
					>
						<Avatar.Fallback class="border border-muted"
							>{getAbbr(profile)}</Avatar.Fallback
						>
					</div>
				</Avatar.Root>
				<p class="text-center line-clamp-2 break-words overflow-ellipsis">
					{profile?.Name}
				</p>
			</div>
		</Button.Root>
	{/each}

	<AddProfileDialog>
		<Button.Root
			class="inline-flex items-center justify-center rounded-input bg-dark/0
	p-2 text-[15px] text-dark shadow-none hover:shadow-mini
	hover:bg-dark/15 active:scale-98 active:transition-all w-24"
		>
			<div class="flex flex-col gap-2 items-center justify-center">
				<Avatar.Root
					class="h-12 w-12 rounded-full border border-foreground bg-muted text-[17px] font-medium uppercase text-muted-foreground"
				>
					<div
						class="flex h-full w-full items-center justify-center overflow-hidden rounded-full border-2 border-transparent"
					>
						<Avatar.Fallback class="border border-muted"
							><Plus /></Avatar.Fallback
						>
					</div>
				</Avatar.Root>
				<p class="text-center line-clamp-2 break-words overflow-ellipsis">
					Add Profile
				</p>
			</div>
		</Button.Root>
	</AddProfileDialog>
</div>
