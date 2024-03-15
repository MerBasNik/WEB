import React, { useState } from "react";
import stylesRegister from "./register.module.css";
import stylesReEntry from "./re-entry.module.css";
import { InputField } from "./register";

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
	return (
		<div>
			<div className={stylesRegister.firstHalf}></div>
			<div className={stylesRegister.register}>
				<div className={stylesRegister.register__window}>
					<div className={stylesRegister.window__header}>
						<h1 className={classesForHeader}>Вход</h1>
					</div>
					<form>
						<div>
							<InputField
								InputData={InputNameEmail}
								setInput={setInputEmail}
								title={"Email"}
								type={"email"}
								autoComplete={true}
							/>
							<InputField
								InputData={InputNamePassword}
								setInput={setInputPassword}
								title={"Пароль"}
								type={"password"}
								autoComplete={true}
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
								<label className={classesForLabel}>
									Запомнить вход
								</label>
							</div>
							<div
								className={
									stylesReEntry.wrapper__forgot__password
								}
							>
								<a
									className={classesForForgotPassword}
									href={"#"}
								>
									Забыли пароль?
								</a>
							</div>
						</div>
						<div className={stylesRegister.wrapper__input__submit}>
							<button
								type={"submit"}
								className={classesForInputSubmit}
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
