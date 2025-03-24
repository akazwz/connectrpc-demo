import {
	Code,
	ConnectError,
	createClient,
	type Interceptor,
} from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { redirect } from "react-router";
import { PublicAuthService } from "~/gen/auth/v1/auth_public_pb";
import { useAuthStore } from "~/stores/auth";

// 添加刷新token锁
let isRefreshingToken = false;
let refreshPromise: Promise<string | null> | null = null;

const authInterceptor: Interceptor = (next) => async (req) => {
	const accessToken = useAuthStore.getState().accessToken;
	if (accessToken) {
		req.header.append("Authorization", `Bearer ${accessToken}`);
	}
	return await next(req);
};

const refreshTokenInterceptor: Interceptor = (next) => async (req) => {
	try {
		return await next(req);
	} catch (error) {
		if (error instanceof ConnectError && error.code === Code.Unauthenticated) {
			const refreshToken = useAuthStore.getState().refreshToken;
			if (!refreshToken) {
				useAuthStore.getState().signOut();
				throw redirect("/login");
			}

			// 如果已经在刷新token，等待该操作完成
			if (isRefreshingToken) {
				const newToken = await refreshPromise;
				if (newToken) {
					req.header.set("Authorization", `Bearer ${newToken}`);
					try {
						return await next(req);
					} catch (error) {
						useAuthStore.getState().signOut();
						throw redirect("/login");
					}
				} else {
					useAuthStore.getState().signOut();
					throw redirect("/login");
				}
			}

			// 设置锁并创建刷新token的promise
			isRefreshingToken = true;
			refreshPromise = (async () => {
				try {
					const resp = await authClient.refreshToken({ refreshToken });
					if (resp.accessToken) {
						useAuthStore.getState().setAccessToken(resp.accessToken);
						return resp.accessToken;
					}
					return null;
				} catch (error) {
					return null;
				} finally {
					isRefreshingToken = false;
					refreshPromise = null;
				}
			})();

			// 等待刷新token的结果
			const newToken = await refreshPromise;
			if (newToken) {
				req.header.set("Authorization", `Bearer ${newToken}`);
				try {
					return await next(req);
				} catch (error) {
					useAuthStore.getState().signOut();
					throw redirect("/login");
				}
			} else {
				useAuthStore.getState().signOut();
				throw redirect("/login");
			}
		}
		throw error;
	}
};

export const transport = createConnectTransport({
	baseUrl: import.meta.env.VITE_API_URL,
	interceptors: [authInterceptor, refreshTokenInterceptor],
});

export const authClient = createClient(PublicAuthService, transport);
