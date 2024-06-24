<script lang="ts">
	import SelectProfile from "@/components/SelectProfile.svelte";
	import SelectVideoDirectory from "@/components/SelectVideoDirectory.svelte";
	import { config } from "@/stores/config";
	import { loadingStores } from "@/stores/loadingStores";
	import { curProfile, logoutProfile } from "@/stores/profiles";
	import { onMount } from "svelte";
	import { SetConfigDirectory } from "wails/go/main/App";

	let visible:
		| "loading"
		| "select-directory"
		| "select-profile"
		| "video-home" = "loading";
	onMount(() => {
		loadingStores.subscribe((l) => {
			if (l) return;
			if (!$config?.Directory) {
				visible = "select-directory";
			} else if (!$curProfile) {
				visible = "select-profile";
			} else {
				visible = "video-home";
			}
		});
	});

	const setVisible = (v: typeof visible) => {
		visible = v;
	};

	const clear = async () => {
		await SetConfigDirectory("");
		logoutProfile();
		setVisible("loading");
	};
</script>

{#if visible === "loading"}
	<div></div>
{:else if visible === "select-directory"}
	<SelectVideoDirectory next={() => setVisible("select-profile")} />
{:else if visible === "select-profile"}
	<h1 class="text-xl">Select your profile</h1>
	<SelectProfile next={() => setVisible("video-home")} />
{:else}
	<div>
		<button class="text-red-500" type="button" on:click={clear}>clear</button>
	</div>
{/if}
