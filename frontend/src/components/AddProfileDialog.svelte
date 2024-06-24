<script lang="ts">
	import { Button, Dialog, Label } from "bits-ui";
	import { fade, fly } from "svelte/transition";
	import { X } from "phosphor-svelte";
	import { addProfile } from "@/stores/profiles";

	let name = "";
	let error: null | string = null;
	let open = false;
</script>

<Dialog.Root {open} onOpenChange={(o) => (open = o)}>
	<Dialog.Trigger>
		<slot />
	</Dialog.Trigger>
	<Dialog.Portal>
		<Dialog.Overlay
			transition={fade}
			transitionConfig={{ duration: 150 }}
			class="fixed inset-0 z-50 bg-black/25 backdrop-blur-md"
		/>
		<Dialog.Content
			transition={fly}
			transitionConfig={{ duration: 250, x: 0, y: 32 }}
			class="fixed left-[50%] top-[50%] z-50 w-full max-w-[94%] translate-x-[-50%] translate-y-[-50%] rounded-card-lg border bg-background p-5 shadow-popover outline-none sm:max-w-[490px] md:w-full"
		>
			<Dialog.Title class="flex w-full text-lg font-semibold tracking-tight"
				>Add new profile</Dialog.Title
			>
			<form on:submit|preventDefault={() => addProfile(name)}>
				<div class="flex flex-col items-start gap-1 pb-6 pt-7">
					<Label.Root for="profile-name" class="text-sm font-medium"
						>Profile Name</Label.Root
					>
					<div class="relative w-full">
						<input
							bind:value={name}
							id="profile-name"
							class="inline-flex h-input w-full items-center rounded-card-sm border border-border-input bg-background px-4 text-sm placeholder:text-foreground-alt/50 hover:border-dark-40 focus:outline-none focus:ring-2 focus:ring-foreground focus:ring-offset-2 focus:ring-offset-background"
							placeholder="Your Name"
							autocomplete="off"
						/>
					</div>
				</div>
				<div class="flex w-full justify-end">
					<Button.Root
						type="submit"
						class="inline-flex h-12 items-center justify-center rounded-input bg-dark
					px-[21px] text-[15px] font-semibold text-background shadow-mini
					hover:bg-dark/95 active:scale-98 active:transition-all"
					>
						Add Profile
					</Button.Root>
				</div>
			</form>
			<Dialog.Close
				class="absolute right-5 top-5 rounded-md focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-foreground focus-visible:ring-offset-2 focus-visible:ring-offset-background active:scale-98"
			>
				<div>
					<X class="size-5 text-foreground" />
					<span class="sr-only">Close</span>
				</div>
			</Dialog.Close>
		</Dialog.Content>
	</Dialog.Portal>
</Dialog.Root>
