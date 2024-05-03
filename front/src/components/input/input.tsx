import React, { HTMLInputTypeAttribute, useEffect } from "react";
import styles from "./input.module.css";
import { ErrorIcon } from "../icons/icons";

export enum LocationInputField {
	Profile,
	Authorization,
}

export enum TypeInputOnProfile {
	Single,
	Double,
}

type InputData = {
	inputData: string;
	setInput: (data: string) => void;
	setErrorInput?: (change: boolean) => void;
	error?: boolean;
};

type TextFieldProps = InputData & {
	id: string;
	textLabel: string;
	helpText?: string;
	typeInput: string;
	location: LocationInputField;
	typeInputOnProfile: TypeInputOnProfile;
};

export const TextField = (props: TextFieldProps) => {
	const {
		id,
		textLabel,
		helpText,
		typeInput,
		inputData,
		setInput,
		setErrorInput,
		error,
		location,
		typeInputOnProfile,
	} = props;

	const profileFontSize = {
		mainLabel: "var(--label-fontSize-Auth)",
		helpLabel: "var(--label-fontSize-Auth)",
		input: "var(--inputField-fonstSize-Auth)",
		errorLabel: "var(--inputField-fonstSize-Auth)",
	};

	if (location == LocationInputField.Profile) {
		profileFontSize.mainLabel = "var(--label-fontSize-Prof)";
		profileFontSize.helpLabel = "var(--label-fontSize-Prof)";
		profileFontSize.input = "var(--inputField-fonstSize-Prof)";
		profileFontSize.errorLabel = "var(--inputField-fonstSize-Prof)";
	}

	const mainLabelFontSize = {
		fontSize: profileFontSize.mainLabel,
	};
	const helpLabelFontSize = {
		fontSize: profileFontSize.helpLabel,
	};
	const inputFontSize = {
		fontSize: profileFontSize.input,
	};
	const errorLabelFontSize = {
		fontSize: profileFontSize.errorLabel,
	};

	let classesHelpLabel =
		styles.wrapper__helpLabel + " " + styles.wrapper__helpLabel__hidden;
	if (helpText) {
		classesHelpLabel = styles.wrapper__helpLabel;
	}

	let classesInput = styles.input + " " + inputFontSize;
	if (typeInputOnProfile === TypeInputOnProfile.Single) {
		classesInput = classesInput + " " + styles.inputSingle__padding;
	}
	if (typeInputOnProfile === TypeInputOnProfile.Double) {
		classesInput = classesInput + " " + styles.inputDouble__padding;
	}

	let classesErrorLabel =
		styles.wrapper__hidden + " " + styles.wrapper__errorLabel;

	let additionalText = "";
	let claassesAdditionalLabel =
		styles.wrapper__hidden + " " + styles.wrapper__additionalText;

	if (setErrorInput !== undefined && error !== undefined) {
		if (error) {
			classesInput =
				classesInput +
				" " +
				styles.input__error +
				" " +
				styles.no__hover +
				" " +
				inputFontSize;
			classesErrorLabel =
				styles.wrapper__error__display +
				" " +
				styles.wrapper__errorLabel +
				document.addEventListener("click", () => {
					setErrorInput(false);
				});
		} else {
			document.removeEventListener("click", () => {
				setErrorInput(false);
			});
		}
	} else {
		additionalText = "Необезательно";
		claassesAdditionalLabel = styles.wrapper__additionalText;
	}

	let errorText = "";
	if (typeInput == "email") {
		errorText = "Неверный email!";
	} else if (typeInput == "password") {
		errorText = "Неверный пароль!";
	} else if (typeInput == "text") {
		errorText = "Пустое поле!";
	}

	return (
		<div className={styles.wrapper__textField}>
			<div className={styles.wrapper__labels}>
				<div className={styles.wrapper__mainLabel}>
					<label
						className={styles.mainLabel}
						style={mainLabelFontSize}
					>
						{textLabel}
					</label>
				</div>
				<div className={classesHelpLabel}>
					<label
						className={styles.helpLabel}
						style={helpLabelFontSize}
					>
						{helpText}
					</label>
				</div>
			</div>
			<div className={styles.wrapper__input}>
				<input
					name={id}
					id={id}
					type={typeInput}
					className={classesInput}
					value={inputData}
					onChange={(event) => {
						setInput(event.target.value);
					}}
				/>
			</div>
			<div className={classesErrorLabel}>
				<div className={styles.wrapper__errorIcon}>{<ErrorIcon />}</div>
				<div className={styles.wrapper__label}>
					<label
						className={styles.errorLabel}
						style={errorLabelFontSize}
					>
						{errorText}
					</label>
				</div>
			</div>
			<div className={claassesAdditionalLabel}>
				<label className={styles.additionalText}>
					{additionalText}
				</label>
			</div>
		</div>
	);
};
