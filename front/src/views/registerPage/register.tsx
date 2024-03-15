import React, { useState } from "react";
import styles from "./register.module.css";

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
					type={type}
					value={InputData}
					onChange={(event) => setInput(event.target.value)}
					autoComplete={autoComplete ? "on" : "off"}
				/>
			</div>
		</div>
	);
};

export const RegistrationScreen = () => {
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
	return (
		<div>
			<div className={styles.firstHalf}></div>
			<div className={styles.register}>
				<div className={styles.register__window}>
					<div className={styles.window__header}>
						<h1 className={classesForHeader}>Регистрация</h1>
					</div>
					<form>
						<div>
							<InputField
								InputData={InputNameEmail}
								setInput={setInputEmail}
								title={"Email"}
								type={"email"}
								autoComplete={false}
							/>
							<InputField
								InputData={InputNamePassword}
								setInput={setInputPassword}
								title={"Пароль"}
								type={"password"}
								autoComplete={false}
							/>
							<InputField
								InputData={InputNameRepeatPassword}
								setInput={setInputRepeatPassword}
								title={"Повторите пароль"}
								type={"password"}
								autoComplete={false}
							/>
						</div>
						<div className={styles.wrapper__input__submit}>
							<button
								type={"submit"}
								className={classesForInputSubmit}
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
