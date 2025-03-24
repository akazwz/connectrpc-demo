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

export default function Register() {
	const navigate = useNavigate();
	const formSchema = z.object({
		email: z.string().email({ message: "邮箱格式错误" }),
		username: z.string().min(1, { message: "用户名不能为空" }),
		password: z.string().min(6, { message: "密码长度至少为6位" }),
	});
	const form = useForm({
		defaultValues: {
			email: "",
			username: "",
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
			const resp = await authClient.register(data);
			return resp;
		},
		onSettled(data, error) {
			if (error) {
				toaster.error({
					title: "注册失败",
					description: error.message,
				});
				console.error(error);
			}
			if (data) {
				toaster.success({
					title: "注册成功",
					description: "正在跳转",
				});
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
						<Fieldset.Legend>注册</Fieldset.Legend>
						<Fieldset.HelperText>请输入您的注册信息</Fieldset.HelperText>
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
						<form.Field name="username">
							{(field) => (
								<Field.Root invalid={!!field.state.meta.errors.length}>
									<Field.Label>用户名</Field.Label>
									<Input
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
						注册
					</Button>
				</Fieldset.Root>
			</chakra.form>
		</VStack>
	);
}
