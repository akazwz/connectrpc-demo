import {
	VStack,
	Heading,
	Field,
	Fieldset,
	Input,
	Stack,
	chakra,
	Button,
} from "@chakra-ui/react";
import { useForm } from "@tanstack/react-form";
import { useMutation } from "@tanstack/react-query";
import { useNavigate } from "react-router";
import { z } from "zod";
import { PasswordInput } from "~/components/ui/password-input";
import { toaster } from "~/components/ui/toaster";
import { authClient } from "~/connect";
import { useAuthStore } from "~/stores/auth";

export default function Register() {
	const navigate = useNavigate();
	const formSchema = z.object({
		email: z.string().email({ message: "邮箱格式错误" }),
		password: z.string().min(6, { message: "密码长度至少为6位" }),
	});
	const form = useForm({
		defaultValues: {
			email: "",
			password: "",
		},
		validators: {
			onSubmit: formSchema,
		},
		onSubmit: (data) => {
			mutation.mutate(data.value);
		},
	});
	const mutation = useMutation({
		mutationFn: async (data: z.infer<typeof formSchema>) => {
			const resp = await authClient.login(data);
			return resp;
		},
		onSettled(data, error) {
			if (error) {
				toaster.error({
					title: "登录失败",
					description: error.message,
				});
				console.error(error);
			}
			if (data?.user) {
				toaster.success({
					title: "登录成功",
					description: "正在跳转",
				});
				useAuthStore
					.getState()
					.signIn(data.accessToken, data.refreshToken, data.user);
				navigate("/", { viewTransition: true });
			}
		},
	});

	return (
		<VStack>
			<Heading>Register</Heading>
			<chakra.form
				w="full"
				maxW="md"
				onSubmit={(e) => {
					e.preventDefault();
					e.stopPropagation();
					form.handleSubmit();
				}}
			>
				<Fieldset.Root as="form" size="lg" maxW="md">
					<Stack>
						<Fieldset.Legend>登录</Fieldset.Legend>
						<Fieldset.HelperText>请输入您的登录信息</Fieldset.HelperText>
					</Stack>
					<Fieldset.Content>
						<form.Field name="email">
							{(field) => (
								<Field.Root invalid={!!field.state.meta.errors.length}>
									<Field.Label>邮箱</Field.Label>
									<Input
										name={field.name}
										value={field.state.value}
										onBlur={field.handleBlur}
										type="email"
										onChange={(e) => field.handleChange(e.target.value)}
									/>
									{field.state.meta.errors.length > 0 && (
										<Field.ErrorText>
											{field.state.meta.errors
												.map((error) => error?.message)
												.join(", ")}
										</Field.ErrorText>
									)}
								</Field.Root>
							)}
						</form.Field>
						<form.Field name="password">
							{(field) => (
								<Field.Root invalid={!!field.state.meta.errors.length}>
									<Field.Label>密码</Field.Label>
									<PasswordInput
										name={field.name}
										value={field.state.value}
										onBlur={field.handleBlur}
										onChange={(e) => field.handleChange(e.target.value)}
									/>
									{field.state.meta.errors.length > 0 && (
										<Field.ErrorText>
											{field.state.meta.errors
												.map((error) => error?.message)
												.join(", ")}
										</Field.ErrorText>
									)}
								</Field.Root>
							)}
						</form.Field>
					</Fieldset.Content>
					<Button type="submit" w="full" loading={mutation.isPending}>
						登录
					</Button>
				</Fieldset.Root>
			</chakra.form>
		</VStack>
	);
}
