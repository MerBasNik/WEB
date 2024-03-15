import React, { useState } from "react";
import styles from "./register.module.css";
<<<<<<< HEAD

type InputFieldProps = {
	type: string;
	title: string;
	setInput: (data: string) => void;
	InputData: string;
	autoComplete: boolean;
};

export const InputField = (props: InputFieldProps) => {
	const classes =
		styles.text__fontFamily + " " + styles.input__header__fontSize;

	const { title, InputData, setInput, type, autoComplete } = props;
	return (
		<div>
			<div>
				<h2 className={classes}>{title}</h2>
			</div>
			<div>
				<input
					className={styles.window__input}
=======
import { ButtonTextLink } from "../../components/button/button";
import { Navigate, useNavigate } from "react-router-dom";

type InputFieldProps = {
	id: string;
	type: string;
	title: string;
	setInput: (data: string) => void;
	setErrorInput: (change: boolean) => void;
	setErrorData: (classes: string) => void;
	InputData: string;
	autoComplete: boolean;
	error: boolean;
};

export const InputField = (props: InputFieldProps) => {
	const {
		title,
		InputData,
		setInput,
		type,
		autoComplete,
		id,
		error,
		setErrorInput,
		setErrorData,
	} = props;

	let classesInput = styles.window__input;
	if (error) {
		classesInput = classesInput + " " + styles.window__input__error;
	}

	return (
		<div>
			<div>
				<h2 className={styles.inputField__header}>{title}</h2>
			</div>
			<div>
				<input
					name={id}
					id={id}
					className={classesInput}
>>>>>>> main
					type={type}
					value={InputData}
					onChange={(event) => setInput(event.target.value)}
					autoComplete={autoComplete ? "on" : "off"}
<<<<<<< HEAD
=======
					onClick={() => {
						setErrorInput(false);
						setErrorData(
							styles.hidden__block +
								" " +
								styles.register__window__error,
						);
					}}
>>>>>>> main
				/>
			</div>
		</div>
	);
};

export const RegistrationScreen = () => {
<<<<<<< HEAD
	const classesForHeader =
		styles.text__fontFamily + " " + styles.register__header__fontSize;

	const classesForInputSubmit =
		styles.input__submit +
		" " +
		styles.text__fontFamily +
		" " +
		styles.input__header__fontSize;

	const [InputNameEmail, setInputEmail] = useState("");
	const [InputNamePassword, setInputPassword] = useState("");
	const [InputNameRepeatPassword, setInputRepeatPassword] = useState("");
=======
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
	const navigate = useNavigate();

>>>>>>> main
	return (
		<div>
			<div className={styles.firstHalf}></div>
			<div className={styles.register}>
<<<<<<< HEAD
				<div className={styles.register__window}>
					<div className={styles.window__header}>
						<h1 className={classesForHeader}>Регистрация</h1>
					</div>
					<form>
=======
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
								name: "name",
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
										"Content-Type": "application/json",
									},
									credentials: "include",
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
							navigate("/main/profile", { replace: true });
						}}
						action="http:://localhost:8000/auth/sign-up"
						method="post"
					>
>>>>>>> main
						<div>
							<InputField
								InputData={InputNameEmail}
								setInput={setInputEmail}
								title={"Email"}
								type={"email"}
								autoComplete={false}
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
								autoComplete={false}
<<<<<<< HEAD
=======
								id={"password"}
								error={InputErrorPassword}
								setErrorInput={setInputErrorPassword}
								setErrorData={setClassesError}
>>>>>>> main
							/>
							<InputField
								InputData={InputNameRepeatPassword}
								setInput={setInputRepeatPassword}
								title={"Повторите пароль"}
								type={"password"}
								autoComplete={false}
<<<<<<< HEAD
							/>
						</div>
						<div className={styles.wrapper__input__submit}>
							<button
								type={"submit"}
								className={classesForInputSubmit}
=======
								id={"repeat__password"}
								error={InputErrorRepeatPassword}
								setErrorInput={setInputErrorRepeatPassword}
								setErrorData={setClassesError}
							/>
						</div>
						<div className={styles.wrapper__input__submit}>
							<div className={styles.wrapper__buttonTextLink}>
								<ButtonTextLink
									id={"reEntry"}
									title={"Уже имеете аккаунт"}
									link={"/reEnt"}
								/>
							</div>
							<button
								type={"submit"}
								className={styles.input__submit}
>>>>>>> main
							>
								Регистрация
							</button>
						</div>
					</form>
				</div>
			</div>
		</div>
	);
};
