<script lang="ts">
	import { Button, Label } from "bits-ui";
	import { CreateUser } from "wailsjs/go/main/App";
	let name = "";
	let error: null | string = null;

	const addProfile = async () => {
		try {
			await CreateUser(name);
		} catch (error) {
			console.error(error);
			error = "Failed to create user";
		}
	};
</script>

<form on:submit|preventDefault={addProfile}>
	<div class="flex flex-col items-start gap-1 pb-11 pt-7">
		<Label.Root for="profile-name" class="text-sm font-medium"
			>Profile Name</Label.Root
		>
		<div class="relative w-full">
			<input
				id="profile-name"
				class="inline-flex h-input w-full items-center rounded-card-sm border border-border-input bg-background px-4 text-sm placeholder:text-foreground-alt/50 hover:border-dark-40 focus:outline-none focus:ring-2 focus:ring-foreground focus:ring-offset-2 focus:ring-offset-background"
				placeholder="Your Name"
				autocomplete="off"
			/>
		</div>
	</div>
	<Button.Root
		class="inline-flex h-12 items-center justify-center rounded-input bg-dark
  px-[21px] text-[15px] font-semibold text-background shadow-mini
  hover:bg-dark/95 active:scale-98 active:transition-all"
	>
		Add Profile
	</Button.Root>
</form>
