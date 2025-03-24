import type { User as UserProto } from "~/gen/auth/v1/auth_public_pb";

export type RemoveProtoMetadata<T> = T extends object
	? T extends Array<infer U>
		? Array<RemoveProtoMetadata<U>>
		: {
				[K in keyof T as K extends "$typeName" | "$unknown"
					? never
					: K]: RemoveProtoMetadata<T[K]>;
			}
	: T;

export type User = RemoveProtoMetadata<UserProto>;
