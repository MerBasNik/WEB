import React, { useState } from "react";
import stylesRegister from "./register.module.css";
import stylesReEntry from "./re-entry.module.css";
import { InputField } from "./register";
<<<<<<< HEAD

export const ReEntryScreen = () => {
	const classesForHeader =
		stylesRegister.text__fontFamily +
		" " +
		stylesRegister.register__header__fontSize;

	const classesForInputSubmit =
		stylesRegister.input__submit +
		" " +
		stylesRegister.text__fontFamily +
		" " +
		stylesRegister.input__header__fontSize;

	const classesForLabel =
		stylesReEntry.label__remember__entry__input +
		" " +
		stylesRegister.text__fontFamily +
		" " +
		stylesReEntry.label__header__fontSize;

	const classesForForgotPassword =
		stylesReEntry.forgot__password +
		" " +
		stylesRegister.text__fontFamily +
		" " +
		stylesReEntry.label__header__fontSize;
	const [InputNameEmail, setInputEmail] = useState("");
	const [InputNamePassword, setInputPassword] = useState("");
=======
import { Navigate, useNavigate } from "react-router-dom";

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
	const navigate = useNavigate();

>>>>>>> main
	return (
		<div>
			<div className={stylesRegister.firstHalf}></div>
			<div className={stylesRegister.register}>
<<<<<<< HEAD
				<div className={stylesRegister.register__window}>
					<div className={stylesRegister.window__header}>
						<h1 className={classesForHeader}>Вход</h1>
					</div>
					<form>
=======
				<div className={ClassesError}>
					<div className={stylesRegister.wrapper__header__error}>
						<h1
							className={
								stylesRegister.register__window__error__header
							}
						>
							Такого пользователя не существует
						</h1>
					</div>
				</div>
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
								username: InputNameEmail,
								password: InputNamePassword,
							};

							if (data.username.length == 0) {
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
							navigate("/main", { replace: true });
						}}
						action="http:://localhost:8000/auth/sign-in"
						method="post"
					>
>>>>>>> main
						<div>
							<InputField
								InputData={InputNameEmail}
								setInput={setInputEmail}
								title={"Email"}
								type={"email"}
								autoComplete={true}
<<<<<<< HEAD
=======
								id={"username"}
								error={InputErrorEmail}
								setErrorInput={setInputErrorEmail}
								setErrorData={setClassesError}
>>>>>>> main
							/>
							<InputField
								InputData={InputNamePassword}
								setInput={setInputPassword}
								title={"Пароль"}
								type={"password"}
								autoComplete={true}
<<<<<<< HEAD
=======
								id={"password"}
								error={InputErrorPassword}
								setErrorInput={setInputErrorPassword}
								setErrorData={setClassesError}
>>>>>>> main
							/>
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
<<<<<<< HEAD
								<label className={classesForLabel}>
=======
								<label
									className={
										stylesReEntry.label__remember__entry__input
									}
								>
>>>>>>> main
									Запомнить вход
								</label>
							</div>
							<div
								className={
									stylesReEntry.wrapper__forgot__password
								}
							>
								<a
<<<<<<< HEAD
									className={classesForForgotPassword}
=======
									className={stylesReEntry.forgot__password}
>>>>>>> main
									href={"#"}
								>
									Забыли пароль?
								</a>
							</div>
						</div>
						<div className={stylesRegister.wrapper__input__submit}>
							<button
								type={"submit"}
<<<<<<< HEAD
								className={classesForInputSubmit}
=======
								className={stylesRegister.input__submit}
>>>>>>> main
							>
								Вход
							</button>
						</div>
					</form>
				</div>
			</div>
		</div>
	);
};
<<<<<<< HEAD
=======
export const foo = () => {
	return <></>;
};
>>>>>>> main
