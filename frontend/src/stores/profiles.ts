import { writable } from "svelte/store";
import { GetProfiles, CreateProfile, GetProfile } from "wailsjs/go/main/App";

const LS_KEY = "cur_profile_id" as const;

export const profiles = writable<Awaited<ReturnType<typeof GetProfiles>>>([]);
export const curProfile =
	writable<Awaited<ReturnType<typeof GetProfile>>>(null);

export const fetchProfiles = async () => {
	profiles.set(await GetProfiles());
	const curProfileId = localStorage.getItem(LS_KEY);
	if (curProfileId && !Number.isNaN(Number(curProfileId))) {
		await setProfile(Number(curProfileId));
	}
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

export const setProfile = async (id: number) => {
	const profile = await GetProfile(id);
	if (profile) {
		curProfile.set(profile);
		localStorage.setItem(LS_KEY, id.toString());
	} else {
		localStorage.removeItem(LS_KEY);
		curProfile.set(null);
	}
};

export const logoutProfile = () => {
	localStorage.removeItem(LS_KEY);
	curProfile.set(null);
};
