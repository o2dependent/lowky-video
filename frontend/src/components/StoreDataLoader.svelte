<script lang="ts">
	import { loadingStores } from "./../stores/loadingStores.ts";
	import { fetchConfig } from "./../stores/config.ts";
	import { fetchProfiles } from "@/stores/profiles";
	import { onMount } from "svelte";

	onMount(async () => {
		try {
			await Promise.all([fetchProfiles(), fetchConfig()]);
		} catch (error) {}
		loadingStores.set(false);
	});
</script>
