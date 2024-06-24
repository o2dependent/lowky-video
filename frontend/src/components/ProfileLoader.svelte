<script lang="ts">
	import { fetchProfiles } from "@/stores/profiles";
	import { onMount } from "svelte";

	onMount(() => {
		fetchProfiles();
	});
</script>
