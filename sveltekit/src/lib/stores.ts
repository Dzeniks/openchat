import { writable } from "svelte/store";

export const _accessTokenRefreshTime = writable(new Date().getTime() + 1000 * 60 * 30);
export const _refreshTokenRefreshTime = writable(new Date().getTime() + 1000 * 60 * 60 * 24 * 30);