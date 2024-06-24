import { writable } from "svelte/store";
import { GetProfiles, CreateProfile } from "wailsjs/go/main/App";

export const profiles = writable<Awaited<ReturnType<typeof GetProfiles>>>([]);

export const fetchProfiles = async () => {
	profiles.set(await GetProfiles());
};

export const addProfile = async (name: string) => {
	try {
		console.log(name);
		await CreateProfile(name);
	} catch (error) {
		console.error(error);
		error = "Failed to create profile";
	}
	fetchProfiles();
};
