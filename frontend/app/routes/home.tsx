import { VStack, Heading, Text } from "@chakra-ui/react";
import { redirect } from "react-router";
import { authClient } from "~/connect";
import { useAuthStore } from "~/stores/auth";

export async function clientLoader() {
	const { accessToken, user } = useAuthStore.getState();
	if (!accessToken || !user) {
		throw redirect("/login");
	}
	authClient.whoami({}).then((resp) => {
		if (resp.user) {
			useAuthStore.getState().updateUser(resp.user);
		}
	});
	return null;
}

export default function Home() {
	const { user } = useAuthStore();
	return (
		<VStack>
			<Heading>Home</Heading>
			<Text>{user?.username}</Text>
		</VStack>
	);
}
