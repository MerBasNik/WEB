import React, { useState } from "react";
import styles from "./register.module.css";
import {
	Button,
	ButtonTextLink,
	ButtonType,
} from "../../components/button/button";
import { useNavigate } from "react-router-dom";
import {
	LocationInputField,
	TextField,
	TypeInputOnProfile,
} from "../../components/input/input";
import { BackgroundIcon } from "../../components/icons/icons";
import { useLocalStorage } from "../../hooks/useLocalStorage";

export const RegistrationScreen = () => {
	const [InputNameEmail, setInputEmail] = useState("");
	const [InputNamePassword, setInputPassword] = useState("");
	const [InputNameRepeatPassword, setInputRepeatPassword] = useState("");
	const [InputErrorEmail, setInputErrorEmail] = useState(false);
	const [InputErrorPassword, setInputErrorPassword] = useState(false);
	const [InputErrorRepeatPassword, setInputErrorRepeatPassword] =
		useState(false);
	const [ClassesError, setClassesError] = useState(
		styles.hidden__block + " " + styles.register__window__error,
	);
	const [token, setToken] = useLocalStorage({
		initialValue: {},
		key: "token",
	});
	const [id_user, setIdUser] = useLocalStorage({
		initialValue: -1,
		key: "id_user",
	});
	const navigate = useNavigate();

	return (
		<div>
			<div className={styles.firstHalf}>{<BackgroundIcon />}</div>
			<div className={styles.register}>
				<div className={ClassesError}>
					<div className={styles.wrapper__header__error}>
						<h1 className={styles.register__window__error__header}>
							Такой пользователь уже существует
						</h1>
					</div>
				</div>
				<div className={styles.register__window}>
					<div className={styles.window__header}>
						<h1 className={styles.register__header}>Регистрация</h1>
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

							if (InputNameRepeatPassword.length == 0) {
								setInputErrorRepeatPassword(true);
								return;
							}

							if (InputNamePassword !== InputNameRepeatPassword) {
								setInputErrorRepeatPassword(true);
								return;
							}

							const response = await fetch(
								"http://localhost:8000/auth/sign-up",
								{
									method: "POST",
									body: JSON.stringify(data),
									headers: {
										"Content-Type":
											"application/json;charset=utf-8",
									},
								},
							);
							if (!response.ok) {
								if (response.status == 500) {
									setClassesError(
										styles.register__window__error,
									);
								}
								return;
							}
							const dataFromResponse = await response.json();

							setToken(dataFromResponse.token);
							setIdUser(dataFromResponse.id);

							setTimeout(() => {
								navigate("/auth_profile/contacts", {
									replace: true,
								});
							}, 1);
						}}
						action="http:://localhost:8000/auth/sign-up"
						method="post"
					>
						<div className={styles.wrapper__inputs}>
							<div className={styles.wrapper_input}>
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
							<div className={styles.wrapper_input}>
								<TextField
									inputData={InputNamePassword}
									setInput={setInputPassword}
									textLabel={"Пароль"}
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
							<div className={styles.wrapper_input}>
								<TextField
									inputData={InputNameRepeatPassword}
									setInput={setInputRepeatPassword}
									textLabel={"Повторите пароль"}
									typeInput={"password"}
									id={"repeat__password"}
									error={InputErrorRepeatPassword}
									setErrorInput={setInputErrorRepeatPassword}
									location={LocationInputField.Authorization}
									typeInputOnProfile={
										TypeInputOnProfile.Double
									}
								/>
							</div>
						</div>
						<div className={styles.wrapper__input__submit}>
							<Button
								id={"auth"}
								title={"Регистрация"}
								type={ButtonType.Text}
								typeButton={"submit"}
							/>
						</div>
					</form>
				</div>
			</div>
		</div>
	);
};
