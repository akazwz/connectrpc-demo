import { create } from "zustand";
import { persist } from "zustand/middleware";
import type { User } from "~/types";

interface AuthState {
	user: User | null;
	accessToken: string | null;
	refreshToken: string | null;
	signIn: (accessToken: string, refreshToken: string, user: User) => void;
	updateUser: (user: User) => void;
	setAccessToken: (accessToken: string) => void;
	signOut: () => void;
}

export const useAuthStore = create<AuthState>()(
	persist(
		(set) => ({
			user: null,
			accessToken: null,
			refreshToken: null,
			signIn: (accessToken: string, refreshToken: string, user: User) => {
				set({ accessToken, refreshToken, user });
			},
			updateUser: (user: User) => {
				set({ user });
			},
			setAccessToken: (accessToken: string) => {
				set({ accessToken });
			},
			signOut: () => {
				set({ accessToken: null, refreshToken: null, user: null });
			},
		}),
		{
			name: "auth-storage",
		},
	),
);
