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
	const [InputNamePassword, setInputPassword] = useState("");
	const [InputNamePasswordRepaet, setInputPasswordRepeat] = useState("");
	const [InputErrorPassword, setInputErrorPassword] = useState(false);
	const [InputErrorPasswordRepeat, setInputErrorPasswordRepeat] = useState(false);
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
							Отправить
						</h1>
					</div>
					<form
						onSubmit={async (event) => {
							event.preventDefault();

							const data = {
								password: InputNamePassword,
								passwordRepeat: InputNamePasswordRepaet,
							};

							if (data.password.length == 0) {
								setInputErrorPassword(true);
								return;
							}

							if (data.passwordRepeat.length == 0) {
								setInputErrorPasswordRepeat(true);
								return;
							}

							const response = await fetch(
								`http://localhost:8000/auth/reset-password/${token}`,
								{
									method: "POST",
									body: JSON.stringify(data),
									headers: {
										"Content-Type": "application/json",
										token: token,
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
						action="http:://localhost:8000/auth/reset-password/{token}"
						method="post"
					>
						<div className={stylesRegister.wrapper__inputs}>
							<div className={stylesRegister.wrapper_input}>
								<TextField
									inputData={InputNamePassword}
									setInput={setInputPassword}
									textLabel={"Email"}
									typeInput={"email"}
									id={"username"}
									error={InputErrorPasswordRepeat}
									setErrorInput={setInputErrorPasswordRepeat}
									location={LocationInputField.Authorization}
									typeInputOnProfile={
										TypeInputOnProfile.Double
									}
								/>
							</div>
						</div>
						<div className={stylesRegister.wrapper__input__submit}>
							<Button
								id={"reset"}
								title={"Сбросить"}
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
									Вернуться на экран входа?
								</h2>
							</div>
							<div
								className={
									stylesRegister.wrapper__buttonTextLink
								}
							>
								<ButtonTextLink
									location={LocationOnPages.Authorization}
									id={"back"}
									title={"Назад"}
									link={"/back"}
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
