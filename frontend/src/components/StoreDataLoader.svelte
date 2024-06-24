<script lang="ts">
	import { fetchConfig } from "./../stores/config.ts";
	import { fetchProfiles } from "@/stores/profiles";
	import { onMount } from "svelte";

	onMount(() => {
		fetchProfiles();
		fetchConfig();
	});
</script>
