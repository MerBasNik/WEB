import React, { useState } from "react";
import stylesRegister from "./register.module.css";
import stylesReEntry from "./re-entry.module.css";
import { Navigate, useNavigate } from "react-router-dom";
import {
	Button,
	ButtonTextLink,
	ButtonType,
	LocationOnPages,
} from "../../components/button/button";
import {
	LocationInputField,
	TextField,
	TypeInputOnProfile,
} from "../../components/input/input";
import { BackgroundIcon } from "../../components/icons/icons";
import { TypePredicateKind } from "typescript";
import { useLocalStorage } from "../../hooks/useLocalStorage";

export const ReEntryScreen = () => {
	const [InputNameEmail, setInputEmail] = useState("");
	const [InputNamePassword, setInputPassword] = useState("");
	const [InputErrorEmail, setInputErrorEmail] = useState(false);
	const [InputErrorPassword, setInputErrorPassword] = useState(false);
	const [ClassesError, setClassesError] = useState(
		stylesRegister.hidden__block +
			" " +
			stylesRegister.register__window__error,
	);

	const [token, setToken] = useLocalStorage({
		initialValue: {},
		key: "token",
	});

	const navigate = useNavigate();

	return (
		<div>
			<div className={stylesRegister.firstHalf}>{<BackgroundIcon />}</div>
			<div className={stylesRegister.register}>
				<div className={stylesRegister.register__window}>
					<div className={stylesRegister.window__header}>
						<h1 className={stylesRegister.register__header}>
							Вход
						</h1>
					</div>
					<form
						onSubmit={async (event) => {
							event.preventDefault();

							const data = {
								email: InputNameEmail,
								password: InputNamePassword,
							};

							if (data.email.length == 0) {
								setInputErrorEmail(true);
								return;
							}

							if (data.password.length == 0) {
								setInputErrorPassword(true);
								return;
							}

							const response = await fetch(
								"http://localhost:8000/auth/sign-in",
								{
									method: "POST",
									body: JSON.stringify(data),
									headers: {
										"Content-Type": "application/json",
									},
									credentials: "include",
								},
							);
							if (!response.ok) {
								if (response.status == 500) {
									setClassesError(
										stylesRegister.register__window__error,
									);
								}
								return;
							}
							const dataFromResponse = await response.json();

							setToken(dataFromResponse.token);

							setTimeout(() => {
								navigate("/main/home", {
									replace: true,
								});
							}, 1);
						}}
						action="http:://localhost:8000/auth/sign-in"
						method="post"
					>
						<div className={stylesRegister.wrapper__inputs}>
							<div className={stylesRegister.wrapper_input}>
								<TextField
									inputData={InputNameEmail}
									setInput={setInputEmail}
									textLabel={"Email"}
									typeInput={"email"}
									id={"username"}
									error={InputErrorEmail}
									setErrorInput={setInputErrorEmail}
									location={LocationInputField.Authorization}
									typeInputOnProfile={
										TypeInputOnProfile.Double
									}
								/>
							</div>
							<div className={stylesRegister.wrapper_input}>
								<TextField
									inputData={InputNamePassword}
									setInput={setInputPassword}
									textLabel={"Пароль"}
									helpText={"Забыли пароль?"}
									typeInput={"password"}
									id={"password"}
									error={InputErrorPassword}
									setErrorInput={setInputErrorPassword}
									location={LocationInputField.Authorization}
									typeInputOnProfile={
										TypeInputOnProfile.Double
									}
								/>
							</div>
						</div>
						<div
							className={stylesReEntry.wrapper__user__assistance}
						>
							<div
								className={
									stylesReEntry.wrapper__remember__entry__input
								}
							>
								<input
									type={"checkbox"}
									className={
										stylesReEntry.remember__entry__input
									}
								/>
								<label
									className={
										stylesReEntry.label__remember__entry__input
									}
								>
									Запомнить вход
								</label>
							</div>
						</div>
						<div className={stylesRegister.wrapper__input__submit}>
							<Button
								id={"entry"}
								title={"Вход"}
								type={ButtonType.Text}
								typeButton={"submit"}
							/>
						</div>
						<div className={stylesReEntry.wrapper__create__account}>
							<div>
								<h2
									className={
										stylesRegister.inputField__header
									}
								>
									Еще не с нами?
								</h2>
							</div>
							<div
								className={
									stylesRegister.wrapper__buttonTextLink
								}
							>
								<ButtonTextLink
									location={LocationOnPages.Authorization}
									id={"reg"}
									title={"Регистрация"}
									link={"/reg"}
								/>
							</div>
						</div>
					</form>
				</div>
			</div>
		</div>
	);
};
export const foo = () => {
	return <></>;
};
