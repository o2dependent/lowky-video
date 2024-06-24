import { writable } from "svelte/store";
import { GetConfig, SetConfigDirectory } from "wailsjs/go/main/App";

export const config = writable<Awaited<ReturnType<typeof GetConfig>>>({});

export const fetchConfig = async () => {
	const configData = await GetConfig();
	console.log(configData);
	config.set(configData ?? {});
};

export const setConfigDirectory = async (directory: string) => {
	try {
		await SetConfigDirectory(directory);
	} catch (error) {
		console.error(error);
		error = "Failed to set config directory";
	}
	await fetchConfig();
};
