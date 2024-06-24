<script lang="ts">
	import { setConfigDirectory, config } from "@/stores/config";
	import { Button } from "bits-ui";
	import { OpenVideoFolderDialog } from "wailsjs/go/main/App";

	export let next: () => void;

	let error: string | null = null;

	const onClick = async () => {
		const folder = await OpenVideoFolderDialog();
		if (!folder) return;
		try {
			await setConfigDirectory(folder);
		} catch (err) {
			console.error(err);
			error = "Failed to set directory. Please try again.";
		}
	};
</script>

<div
	class="max-w-md prose dark:prose-invert mx-auto px-4 flex flex-col gap-2 w-full"
>
	<h1>Select Video Directory</h1>

	<p>
		A folder named "Lowky Video" will be created if one with that name isn't
		selected.
	</p>

	<div class="flex flex-col w-full gap-1">
		<p class="!m-0">Current Directory:</p>
		<p class="!m-0 font-mono text-dark">{$config?.Directory}</p>
	</div>
	<Button.Root
		on:click={onClick}
		type="button"
		class="inline-flex h-8 items-center justify-center rounded-input border border-dark
						px-[21px] text-[15px] font-semibold text-dark shadow-mini
						hover:bg-dark/15 active:scale-98 active:transition-all"
	>
		Select Video Directory
	</Button.Root>
	<div class="flex w-full justify-end py-12">
		<Button.Root
			on:click={next}
			disabled={!$config?.Directory}
			type="button"
			class="inline-flex h-12 items-center justify-center rounded-input bg-dark
							px-[21px] text-[15px] font-semibold text-background shadow-mini
							hover:bg-dark/95 active:scale-98 active:transition-all disabled:opacity-50 disabled:cursor-not-allowed"
		>
			Next
		</Button.Root>
	</div>
</div>
