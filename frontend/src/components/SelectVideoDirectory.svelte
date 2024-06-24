<script lang="ts">
	import { setConfigDirectory, config } from "@/stores/config";
	import { Button } from "bits-ui";
	import { OpenVideoFolderDialog } from "wailsjs/go/main/App";

	const onClick = async () => {
		const folder = await OpenVideoFolderDialog();
		if (!folder) return;
		await setConfigDirectory(folder);
	};
</script>

<code>{$config?.Directory}</code>
<h1>Select Video Directory</h1>
<p>
	A folder named "Lowky Video" will be created if one with that name isn't
	selected.
</p>
<Button.Root
	on:click={onClick}
	type="button"
	class="inline-flex h-12 items-center justify-center rounded-input bg-dark
					px-[21px] text-[15px] font-semibold text-background shadow-mini
					hover:bg-dark/95 active:scale-98 active:transition-all"
>
	Select Video Directory
</Button.Root>
